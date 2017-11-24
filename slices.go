package slices

import (
	"reflect"
	"errors"
	"math/rand"
	"bytes"
)

func sliceValue(slice interface{}) reflect.Value {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic(errors.New("slices: invalid type"))
	}

	return value
}

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func Fetch(slice interface{}, index int, defaultValue interface{}) interface{} {
	src := sliceValue(slice)
	if index < src.Len() {
		return src.Index(index).Interface()
	}
	return defaultValue
}

// スライスの要素の数だけ関数fを繰り返し実行し、その繰り返しをn回続ける。
func Cycle(slice interface{}, n int, f func(int) error) error {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < length; j++ {
			if e := f(j); e != nil {
				return e
			}
		}
	}

	return nil
}

// スライスの長さの最大値を返す。
func MaxLength(slices ...interface{}) int {
	max := 0
	for _, slice := range slices {
		if length := sliceValue(slice).Len(); length > max {
			max = length
		}
	}
	return max
}

// スライスの長さの最小値を返す。
func MinLength(slices ...interface{}) int {
	if len(slices) <= 0 {
		return 0
	}

	min := sliceValue(slices[0]).Len()
	for i := 1; i < len(slices); i++ {
		if length := sliceValue(slices[i]).Len(); length < min {
			min = length
		}
	}
	return min
}

// 要素を結合した文字列を返す。
func Join(slice interface{}, sep string, f func(int) string) string {
	length := sliceValue(slice).Len()

	if length <= 0 {
		return ""
	}

	buf := bytes.NewBuffer(make([]byte, 0, (length + len(sep)) * 4 + 1))

	length--
	for i := 0; i < length; i++ {
		buf.WriteString(f(i))
		buf.WriteString(sep)
	}

	buf.WriteString(f(length))

	return buf.String()
}

// スライスを複製する。
func Clone(slice interface{}) interface{} {
	src := sliceValue(slice)

	dst := reflect.MakeSlice(src.Type(), src.Len(), src.Len())
	reflect.Copy(dst, src)

	return dst.Interface()
}

// 要素をランダムに入れ替える。
func Shuffle(slice interface{}, r *rand.Rand) {
	src := sliceValue(slice)

	swap := reflect.Swapper(src.Interface())

	length := src.Len()
	for i := 0; i < length; i++ {
		swap(i, r.Intn(i + 1))
	}
}

// 要素を１つランダムに返す。
func Sample(slice interface{}, r *rand.Rand) interface{} {
	src := sliceValue(slice)
	length := src.Len()
	return src.Index(r.Intn(length)).Interface()
}

// 最もよく条件を満たす要素を返す。
func Top(slice interface{}, f func(int, int) bool) interface{} {
	src := sliceValue(slice)

	length := src.Len()

	if length <= 0 {
		return reflect.Zero(src.Type().Elem()).Interface()
	}

	top := 0;
	for i := 1; i < length; i++ {
		if f(top, i) {
			top = i
		}
	}

	return src.Index(top).Interface()
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

// 条件を満たす先頭の要素を除いていったスライス。
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

// 重複を排除したコレクションを返す。
// 入力スライスの順序に影響を与える。
func Distinct(slice interface{}, f func(int, int) bool) interface{} {
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
// Filter よりも高速だが、入力スライスの順序に影響を与える。
func FastFilter(slice interface{}, f func(int) bool) interface{} {
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
// FilterNot よりも高速だが、入力スライスの順序に影響を与える。
func FastFilterNot(slice interface{}, f func(int) bool) interface{} {
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
// Partition よりも高速だが、入力スライスの順序に影響を与える。
func FastPartition(slice interface{}, f func(int) bool) (interface{}, interface{}) {
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

// １つでも条件を満たす要素が存在したらtrue。
func Exists(slice interface{}, f func(int) bool) bool {
	return IndexWhere(slice, f) >= 0
}

// 全ての要素が条件を満たすとき、true。
func ForAll(slice interface{}, f func(int) bool) bool {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := 0; i < length; i++ {
		if !f(i) {
			return false
		}
	}

	return true
}

// 最初に条件を満たした要素の位置を返す。
// 無い場合は-1が返る。
func IndexWhere(slice interface{}, f func(int) bool) int {
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
func LastIndexWhere(slice interface{}, f func(int) bool) int {
	rs := sliceValue(slice)

	length := rs.Len()
	for i := length - 1; i >= 0; i-- {
		if f(i) {
			return i
		}
	}

	return -1
}

// 最もよく条件を満たす要素の位置を返す。
func TopIndex(slice interface{}, f func(int, int) bool) int {
	src := sliceValue(slice)

	length := src.Len()

	if length <= 0 {
		return -1
	}

	top := 0;
	for i := 1; i < length; i++ {
		if f(top, i) {
			top = i
		}
	}

	return top
}


