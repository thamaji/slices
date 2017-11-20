package slices

import (
	"reflect"
	"errors"
)

func sliceValue(slice interface{}) reflect.Value {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		panic(errors.New("slices: invalid type"))
	}

	return value
}

func Find(slice interface{}, f func(int) bool) interface{} {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			return src.Index(i).Interface()
		}
	}

	return nil
}

func Split(slice interface{}, f func(int) bool) (interface{}, interface{}) {
	src := sliceValue(slice)

	length := src.Len()
	for i := 0; i < length; i++ {
		if f(i) {
			return src.Slice(0, i + 1).Interface(), src.Slice(i + 1, length).Interface()
		}
	}

	return slice, reflect.MakeSlice(src.Type(), 0, 0)
}

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

func FilterFast(slice interface{}, f func(int) bool) interface{} {
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

func PartitionFast(slice interface{}, f func(int) bool) (interface{}, interface{}) {
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

func Contains(slice interface{}, f func(int) bool) bool {
	return Index(slice, f) >= 0
}

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

