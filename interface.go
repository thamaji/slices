package slices

import (
	"errors"
	"math/rand"
	"reflect"
)

func sliceValue(slice interface{}) reflect.Value {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic(errors.New("slices: invalid type"))
	}

	return value
}

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func GetOrDefault(slice interface{}, index int, defaultValue interface{}) interface{} {
	src := sliceValue(slice)
	if index < src.Len() {
		return src.Index(index).Interface()
	}
	return defaultValue
}

// 要素をランダムに入れ替える。
func Shuffle(slice interface{}, r *rand.Rand) {
	src := sliceValue(slice)

	swap := reflect.Swapper(src.Interface())

	length := src.Len()
	for i := 0; i < length; i++ {
		swap(i, r.Intn(i+1))
	}
}

// 要素を１つランダムに返す。
func Sample(slice interface{}, r *rand.Rand) interface{} {
	src := sliceValue(slice)
	length := src.Len()
	return src.Index(r.Intn(length)).Interface()
}

// １つでも条件を満たす要素が存在したらtrue。
func Contains(slice interface{}, f func(int) bool) bool {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			return true
		}
	}

	return false
}

// 全ての要素が条件を満たすとき、true。
func ContainsAll(slice interface{}, f func(int) bool) bool {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			return false
		}
	}

	return true
}

// 条件を満たす要素の数を返す。
func Count(slice interface{}, f func(int) bool) (count int) {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			count++
		}
	}

	return count
}

// 最初に条件を満たした要素の位置を返す。
// 無い場合は-1が返る。
func Index(slice interface{}, f func(int) bool) int {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			return i
		}
	}

	return -1
}

// 最後に条件を満たした要素の位置を返す。
// 無い場合は-1が返る。
func LastIndex(slice interface{}, f func(int) bool) int {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := length - 1; i >= 0; i-- {
		if f(i) {
			return i
		}
	}

	return -1
}

// 条件を満たす最初の要素を返す。
func Find(slice interface{}, f func(int) bool) (interface{}, bool) {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			return src.Index(i).Interface(), true
		}
	}

	return reflect.Zero(src.Type().Elem()).Interface(), false
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func Span(slice interface{}, f func(int) bool) (interface{}, interface{}) {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			return src.Slice(0, i).Interface(), src.Slice(i, length).Interface()
		}
	}

	return slice, reflect.MakeSlice(src.Type(), 0, 0).Interface()
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhile(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			return src.Slice(0, i).Interface()
		}
	}

	return slice
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhile(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			return src.Slice(i, length).Interface()
		}
	}

	return reflect.MakeSlice(src.Type(), 0, 0).Interface()
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func Unique(slice interface{}, f func(int, int) bool) interface{} {
	src := sliceValue(slice)

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	length := src.Len()

	if length > 0 {
		dst = reflect.Append(dst, src.Index(0))
	}

	for i := 1; i < length; i++ {
		if !f(i, i-1) {
			dst = reflect.Append(dst, src.Index(i))
		}
	}

	return dst.Interface()
}

// 重複を排除したスライスを返す。
func UniqueInPlace(slice interface{}, f func(int, int) bool) interface{} {
	src := sliceValue(slice)

	swap := reflect.Swapper(src.Interface())

	length := src.Len()
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; {
			if f(i, j) {
				length--
				swap(j, length)
			} else {
				j++
			}
		}
	}

	return src.Slice(0, length).Interface()
}

// 条件を満たす要素だけのスライスを返す。
func Filter(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	swap := reflect.Swapper(dst.Interface())

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			swap(count, i)
			count++
		}
	}

	return dst.Slice(0, count).Interface()
}

// 条件を満たす要素だけのスライスを返す。
func FilterInPlace(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	swap := reflect.Swapper(slice)

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			swap(count, i)
			count++
		}
	}

	return src.Slice(0, count).Interface()
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNot(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	swap := reflect.Swapper(dst.Interface())

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			swap(count, i)
			count++
		}
	}

	return dst.Slice(0, count).Interface()
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotInPlace(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	swap := reflect.Swapper(slice)

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			swap(count, i)
			count++
		}
	}

	return src.Slice(0, count).Interface()
}

// 条件を満たすスライスと満たさないスライスを返す。
func Partition(slice interface{}, f func(int) bool) (interface{}, interface{}) {
	src := sliceValue(slice)

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	swap := reflect.Swapper(dst.Interface())

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			swap(count, i)
			count++
		}
	}

	return dst.Slice(0, count).Interface(), dst.Slice(count, dst.Len()).Interface()
}

// 条件を満たすスライスと満たさないスライスを返す。
func PartitionInPlace(slice interface{}, f func(int) bool) (interface{}, interface{}) {
	src := sliceValue(slice)

	swap := reflect.Swapper(slice)

	count := 0
	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			swap(count, i)
			count++
		}
	}

	return src.Slice(0, count).Interface(), src.Slice(count, src.Len()).Interface()
}
