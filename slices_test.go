package slices

import (
	"testing"
	"reflect"
)

func BenchmarkFind_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() int {
		for _, v := range slice {
			if v == n {
				return v
			}
		}
		return 0
	}()
}

func BenchmarkFind_BasicInterface(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	find := func(f func(interface{}) bool) interface{} {
		for _, v := range slice {
			if f(v) {
				return v
			}
		}
		return 0
	}

	b.ResetTimer()

	find(func(v interface{}) bool {
		return v.(int) == n
	})
}

func BenchmarkFind(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Find(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkSplit_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() ([]int, []int) {
		for i, v := range slice {
			if v == n {
				return slice[0:i + 1], slice[i + 1:]
			}
		}
		return slice, []int{}
	}()
}

func BenchmarkSplit(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Split(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkFilter_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	func() ([]int) {
		dst := make([]int, 0, len(slice) / 2)
		for _, v := range slice {
			if v % 2 == 0 {
				dst = append(dst, v)
			}
		}
		return dst
	}()
}

func BenchmarkFilter_BasicReflect(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	filter := func(slice interface{}, f func(int) bool) interface{} {
		src := sliceValue(slice)
		dst := reflect.MakeSlice(src.Type(), 0, src.Len() / 2)

		length := src.Len()
		for i := 0; i < length; i++ {
			if f(i) {
				dst = reflect.Append(dst, src.Index(i))
			}
		}

		return dst.Interface()
	}

	b.ResetTimer()

	filter(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkFilter(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	Filter(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkFilter_Fast(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	FilterFast(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkPartition_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	func() ([]int, []int) {
		dst1 := make([]int, 0, len(slice) / 2)
		dst2 := make([]int, 0, len(slice) / 2)
		for _, v := range slice {
			if v % 2 == 0 {
				dst1 = append(dst1, v)
			} else {
				dst2 = append(dst2, v)
			}
		}
		return dst1, dst2
	}()
}

func BenchmarkPartition_BasicReflect(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	partition := func(slice interface{}, f func(int) bool) (interface{}, interface{}) {
		src := sliceValue(slice)
		dst1 := reflect.MakeSlice(src.Type(), 0, src.Len() / 2)
		dst2 := reflect.MakeSlice(src.Type(), 0, src.Len() / 2)

		length := src.Len()
		for i := 0; i < length; i++ {
			if f(i) {
				dst1 = reflect.Append(dst1, src.Index(i))
			} else {
				dst2 = reflect.Append(dst2, src.Index(i))
			}
		}

		return dst1.Interface(), dst2.Interface()
	}

	b.ResetTimer()

	partition(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkPartition(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	Partition(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkPartition_Fast(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	PartitionFast(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkCount_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	func() int {
		count := 0
		for _, v := range slice {
			if v % 2 == 0 {
				count++
			}
		}
		return count
	}()
}

func BenchmarkCount(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	Count(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkContains_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() bool {
		for _, v := range slice {
			if v == n {
				return true
			}
		}
		return false
	}()
}

func BenchmarkContains(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Contains(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkIndex_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() int {
		for i, v := range slice {
			if v == n {
				return i
			}
		}
		return -1
	}()
}

func BenchmarkIndex(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Index(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkLastIndex_Basic(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() int {
		for i := len(slice) - 1; i >= 0; i-- {
			if slice[i] == n {
				return i
			}
		}
		return -1
	}()
}

func BenchmarkLastIndex(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	LastIndex(slice, func(i int) bool {
		return slice[i] == n
	})
}