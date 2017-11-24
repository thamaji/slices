package slices

import (
	"testing"
	"errors"
	"strconv"
	"reflect"
)

func assert(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		return
	}

	t.Errorf("expected: <%v>, but was: <%v>", expected, actual)
}

func TestFetch(t *testing.T) {
	input := []int{1, 2, 3}

	output := Fetch(input, 1, 100)

	assert(t, 2, output.(int))

	input = []int{1, 2, 3}

	output = Fetch(input, 3, 100)

	assert(t, 100, output.(int))
}

func TestCycle(t *testing.T) {
	input := []int{1, 2, 3}
	output := []int{}

	e := Cycle(input, 2, func(i int) error {
		output = append(output, input[i])
		return nil
	})

	assert(t, nil, e)
	assert(t, 6, len(output))
	assert(t, 1, output[0])
	assert(t, 2, output[1])
	assert(t, 3, output[2])
	assert(t, 1, output[3])
	assert(t, 2, output[4])
	assert(t, 3, output[5])

	input = []int{1, 2, 3}
	output = []int{}
	dummy := errors.New("")

	e = Cycle(input, 2, func(i int) error {
		if i == 2 {
			return dummy
		}

		output = append(output, input[i])
		return nil
	})

	assert(t, dummy, e)
	assert(t, 2, len(output))
	assert(t, 1, output[0])
	assert(t, 2, output[1])
}

func TestMaxLength(t *testing.T) {
	assert(t, 2, MaxLength([]int{}, []int{1}, []int{1, 2}))
	assert(t, 2, MaxLength([]int{1}, []int{1, 2}, []int{}))
	assert(t, 2, MaxLength([]int{1, 2}, []int{}, []int{1}))
	assert(t, 0, MaxLength([]int{}, []int{}, []int{}))
	assert(t, 1, MaxLength([]int{1}, []int{1}, []int{1}))
	assert(t, 1, MaxLength([]int{1}))
	assert(t, 0, MaxLength())
}

func TestMinLength(t *testing.T) {
	assert(t, 0, MinLength([]int{}, []int{1}, []int{1, 2}))
	assert(t, 0, MinLength([]int{1}, []int{1, 2}, []int{}))
	assert(t, 0, MinLength([]int{1, 2}, []int{}, []int{1}))
	assert(t, 0, MinLength([]int{}, []int{}, []int{}))
	assert(t, 1, MinLength([]int{1}, []int{1}, []int{1}))
	assert(t, 1, MinLength([]int{1}))
	assert(t, 0, MinLength())
}

func TestJoin(t *testing.T) {
	input := []int{1, 2, 3}
	output := Join(input, ", ", func(i int) string {
		return strconv.Itoa(input[i])
	})
	assert(t, "1, 2, 3", output)

	input = []int{1}
	output = Join(input, ", ", func(i int) string {
		return strconv.Itoa(input[i])
	})
	assert(t, "1", output)

	input = []int{}
	output = Join(input, ", ", func(i int) string {
		return strconv.Itoa(input[i])
	})
	assert(t, "", output)
}

func TestClone(t *testing.T) {
	input := []int{1, 2, 3}

	output := Clone(input)

	input[0] = 100;
	input[1] = 200;
	input[2] = 300;

	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])
}

func TestTop(t *testing.T) {
	input := []int{1, 2, 3}

	output := Top(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	assert(t, 3, output.(int))

	input = []int{1, 2, 3}

	output = Top(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	assert(t, 1, output.(int))
}

func TestFind(t *testing.T) {
	input := []int{1, 2, 3}

	output, ok := Find(input, func(i int) bool {
		return input[i] == 1
	})

	assert(t, true, ok)
	assert(t, 1, output.(int))

	input = []int{1, 2, 3}

	output, ok = Find(input, func(i int) bool {
		return input[i] == 3
	})

	assert(t, true, ok)
	assert(t, 3, output.(int))

	output, ok = Find(input, func(i int) bool {
		return input[i] == 4
	})

	assert(t, false, ok)
	assert(t, 0, output.(int))
}

func TestSpan(t *testing.T) {
	input := []int{1, 2, 3}

	output1, output2 := Span(input, func(i int) bool {
		return input[i] <= 2
	})

	assert(t, 2, len(output1.([]int)))
	assert(t, 1, output1.([]int)[0])
	assert(t, 2, output1.([]int)[1])
	assert(t, 1, len(output2.([]int)))
	assert(t, 3, output2.([]int)[0])

	output1, output2 = Span(input, func(i int) bool {
		return input[i] <= 0
	})

	assert(t, 0, len(output1.([]int)))
	assert(t, 3, len(output2.([]int)))
	assert(t, 1, output2.([]int)[0])
	assert(t, 2, output2.([]int)[1])
	assert(t, 3, output2.([]int)[2])

	output1, output2 = Span(input, func(i int) bool {
		return input[i] <= 3
	})

	assert(t, 3, len(output1.([]int)))
	assert(t, 1, output1.([]int)[0])
	assert(t, 2, output1.([]int)[1])
	assert(t, 3, output1.([]int)[2])
	assert(t, 0, len(output2.([]int)))
}

func TestTakeWhile(t *testing.T) {
	input := []int{1, 2, 3}

	output := TakeWhile(input, func(i int) bool {
		return input[i] <= 2
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])

	output = TakeWhile(input, func(i int) bool {
		return input[i] <= 0
	})

	assert(t, 0, len(output.([]int)))

	output = TakeWhile(input, func(i int) bool {
		return input[i] <= 3
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])
}

func TestDropWhile(t *testing.T) {
	input := []int{1, 2, 3}

	output := DropWhile(input, func(i int) bool {
		return input[i] <= 2
	})

	assert(t, 1, len(output.([]int)))
	assert(t, 3, output.([]int)[0])

	output = DropWhile(input, func(i int) bool {
		return input[i] <= 0
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	output = DropWhile(input, func(i int) bool {
		return input[i] <= 3
	})

	assert(t, 0, len(output.([]int)))
}

func TestDistinct(t *testing.T) {
	input := []int{1, 2, 3}

	output := Distinct(input, func(i, j int) bool {
		return input[i] == input[j]
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	input = []int{1, 1, 1}

	output = Distinct(input, func(i, j int) bool {
		return input[i] == input[j]
	})

	assert(t, 1, len(output.([]int)))
	assert(t, 1, output.([]int)[0])

	input = []int{1, 2, 1}

	output = Distinct(input, func(i, j int) bool {
		return input[i] == input[j]
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])

	input = []int{1, 2, 2}

	output = Distinct(input, func(i, j int) bool {
		return input[i] == input[j]
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])

	input = []int{1, 1, 2}

	output = Distinct(input, func(i, j int) bool {
		return input[i] == input[j]
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3}

	output := Filter(input, func(i int) bool {
		return input[i] != 2
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 3, output.([]int)[1])

	input = []int{1, 2, 3}

	output = Filter(input, func(i int) bool {
		return true
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	input = []int{1, 2, 3}

	output = Filter(input, func(i int) bool {
		return false
	})

	assert(t, 0, len(output.([]int)))
}

func TestFastFilter(t *testing.T) {
	input := []int{1, 2, 3}

	output := FastFilter(input, func(i int) bool {
		return input[i] != 2
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 3, output.([]int)[1])

	input = []int{1, 2, 3}

	output = FastFilter(input, func(i int) bool {
		return true
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	input = []int{1, 2, 3}

	output = FastFilter(input, func(i int) bool {
		return false
	})

	assert(t, 0, len(output.([]int)))
}

func TestFilterNot(t *testing.T) {
	input := []int{1, 2, 3}

	output := FilterNot(input, func(i int) bool {
		return input[i] == 2
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 3, output.([]int)[1])

	input = []int{1, 2, 3}

	output = FilterNot(input, func(i int) bool {
		return false
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	input = []int{1, 2, 3}

	output = FilterNot(input, func(i int) bool {
		return true
	})

	assert(t, 0, len(output.([]int)))
}

func TestFastFilterNot(t *testing.T) {
	input := []int{1, 2, 3}

	output := FastFilterNot(input, func(i int) bool {
		return input[i] == 2
	})

	assert(t, 2, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 3, output.([]int)[1])

	input = []int{1, 2, 3}

	output = FastFilterNot(input, func(i int) bool {
		return false
	})

	assert(t, 3, len(output.([]int)))
	assert(t, 1, output.([]int)[0])
	assert(t, 2, output.([]int)[1])
	assert(t, 3, output.([]int)[2])

	input = []int{1, 2, 3}

	output = FastFilterNot(input, func(i int) bool {
		return true
	})

	assert(t, 0, len(output.([]int)))
}

func TestPartition(t *testing.T) {
	input := []int{1, 2, 3}

	output1, output2 := Partition(input, func(i int) bool {
		return input[i] == 2
	})

	assert(t, 1, len(output1.([]int)))
	assert(t, 2, output1.([]int)[0])
	assert(t, 2, len(output2.([]int)))
	assert(t, 1, output2.([]int)[0])
	assert(t, 3, output2.([]int)[1])

	input = []int{1, 2, 3}

	output1, output2 = Partition(input, func(i int) bool {
		return true
	})

	assert(t, 3, len(output1.([]int)))
	assert(t, 1, output1.([]int)[0])
	assert(t, 2, output1.([]int)[1])
	assert(t, 3, output1.([]int)[2])
	assert(t, 0, len(output2.([]int)))

	input = []int{1, 2, 3}

	output1, output2 = Partition(input, func(i int) bool {
		return false
	})

	assert(t, 0, len(output1.([]int)))
	assert(t, 3, len(output2.([]int)))
	assert(t, 1, output2.([]int)[0])
	assert(t, 2, output2.([]int)[1])
	assert(t, 3, output2.([]int)[2])
}

func TestFastPartition(t *testing.T) {
	input := []int{1, 2, 3}

	output1, output2 := FastPartition(input, func(i int) bool {
		return input[i] == 2
	})

	assert(t, 1, len(output1.([]int)))
	assert(t, 2, output1.([]int)[0])
	assert(t, 2, len(output2.([]int)))
	assert(t, 1, output2.([]int)[0])
	assert(t, 3, output2.([]int)[1])

	input = []int{1, 2, 3}

	output1, output2 = FastPartition(input, func(i int) bool {
		return true
	})

	assert(t, 3, len(output1.([]int)))
	assert(t, 1, output1.([]int)[0])
	assert(t, 2, output1.([]int)[1])
	assert(t, 3, output1.([]int)[2])
	assert(t, 0, len(output2.([]int)))

	input = []int{1, 2, 3}

	output1, output2 = FastPartition(input, func(i int) bool {
		return false
	})

	assert(t, 0, len(output1.([]int)))
	assert(t, 3, len(output2.([]int)))
	assert(t, 1, output2.([]int)[0])
	assert(t, 2, output2.([]int)[1])
	assert(t, 3, output2.([]int)[2])
}

func TestCount(t *testing.T) {
	input := []int{1, 2, 3}

	output := Count(input, func(i int) bool {
		return input[i] == 2
	})

	assert(t, 1, output)

	input = []int{1, 2, 3}

	output = Count(input, func(i int) bool {
		return true
	})

	assert(t, 3, output)

	input = []int{1, 2, 3}

	output = Count(input, func(i int) bool {
		return false
	})

	assert(t, 0, output)
}

func TestExists(t *testing.T) {
	input := []int{1, 2, 3}

	output := Exists(input, func(i int) bool {
		return input[i] == 1
	})

	assert(t, true, output)

	input = []int{1, 2, 3}

	output = Exists(input, func(i int) bool {
		return input[i] == 3
	})

	assert(t, true, output)

	input = []int{1, 2, 3}

	output = Exists(input, func(i int) bool {
		return input[i] == 4
	})

	assert(t, false, output)
}

func TestForAll(t *testing.T) {
	input := []int{1, 2, 3}

	output := ForAll(input, func(i int) bool {
		return true
	})

	assert(t, true, output)

	input = []int{1, 2, 3}

	output = ForAll(input, func(i int) bool {
		return input[i] <= 2
	})

	assert(t, false, output)
}

func TestIndexWhere(t *testing.T) {
	input := []int{1, 1, 3}

	output := IndexWhere(input, func(i int) bool {
		return input[i] == 1
	})

	assert(t, 0, output)

	input = []int{1, 2, 3}

	output = IndexWhere(input, func(i int) bool {
		return input[i] == 3
	})

	assert(t, 2, output)

	input = []int{1, 2, 3}

	output = IndexWhere(input, func(i int) bool {
		return input[i] == 4
	})

	assert(t, -1, output)
}

func TestLastIndexWhere(t *testing.T) {
	input := []int{1, 1, 3}

	output := LastIndexWhere(input, func(i int) bool {
		return input[i] == 1
	})

	assert(t, 1, output)

	input = []int{1, 2, 3}

	output = LastIndexWhere(input, func(i int) bool {
		return input[i] == 3
	})

	assert(t, 2, output)

	input = []int{1, 2, 3}

	output = LastIndexWhere(input, func(i int) bool {
		return input[i] == 4
	})

	assert(t, -1, output)
}

func TestTopIndex(t *testing.T) {
	input := []int{1, 2, 3}

	output := TopIndex(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	assert(t, 2, output)

	input = []int{1, 2, 3}

	output = TopIndex(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	assert(t, 0, output)
}

func BenchmarkFind_Muscle(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() (int, bool) {
		for _, v := range slice {
			if v == n {
				return v, true
			}
		}
		return 0, false
	}()
}

func BenchmarkFind_BasicInterface(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	find := func(f func(interface{}) bool) (interface{}, bool) {
		for _, v := range slice {
			if f(v) {
				return v, true
			}
		}
		return 0, false
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

func BenchmarkSpan_Muscle(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	func() ([]int, []int) {
		for i, v := range slice {
			if !(v <= n) {
				return slice[0:i + 1], slice[i + 1:]
			}
		}
		return slice, []int{}
	}()
}

func BenchmarkSpan(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Span(slice, func(i int) bool {
		return slice[i] <= n
	})
}

func BenchmarkFilter_Muscle(b *testing.B) {
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

func BenchmarkFastFilter(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	FastFilter(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkPartition_Muscle(b *testing.B) {
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

func BenchmarkFastPartition(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()

	FastPartition(slice, func(i int) bool {
		return slice[i] % 2 == 0
	})
}

func BenchmarkCount_Muscle(b *testing.B) {
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

func BenchmarkExists_Muscle(b *testing.B) {
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

func BenchmarkExists(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	Exists(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkIndexWhere_Muscle(b *testing.B) {
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

func BenchmarkIndexWhere(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	IndexWhere(slice, func(i int) bool {
		return slice[i] == n
	})
}

func BenchmarkLastIndexWhere_Muscle(b *testing.B) {
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

func BenchmarkLastIndexWhere(b *testing.B) {
	slice := make([]int, b.N)
	for i := range slice {
		slice[i] = i
	}

	n := len(slice) / 2

	b.ResetTimer()

	LastIndexWhere(slice, func(i int) bool {
		return slice[i] == n
	})
}
