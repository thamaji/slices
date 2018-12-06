package slices

import (
	"math/rand"
)

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func GetOrDefaultInt64(slice []int64, index int, defaultValue int64) int64 {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}

// 要素を１つランダムに返す。
func SampleInt64(slice []int64, r *rand.Rand) int64 {
	return slice[r.Intn(len(slice))]
}

// スライスの各要素を組み合わせたスライスを返す。
func CombineInt64(slice []int64, slices ...[]int64) [][]int64 {
	size := len(slice)
	for _, slice := range slices {
		size *= len(slice)
	}

	out := make([][]int64, 0, size)
	for _, v := range slice {
		out = append(out, []int64{v})
	}

	for _, slice := range slices {
		out = combineInt64(out, slice)
	}

	return out
}

func combineInt64(out [][]int64, slice []int64) [][]int64 {
	length := len(out)
	for i := 0; i < length; i++ {
		for j := 0; j < len(slice); j++ {
			out = append(out, append(out[i], slice[j]))
		}
	}
	return out[length:]
}

// １つでも値と一致する要素が存在したらtrue。
func ContainsInt64(slice []int64, v int64) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

// １つでも条件を満たす要素が存在したらtrue。
func ContainsInt64Func(slice []int64, f func(int64) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}
	return false
}

// 他のスライスのすべての要素を内包していたらtrue。
func ContainsAllInt64(slice []int64, subset []int64) bool {
	for i := range subset {
		if !ContainsInt64(slice, subset[i]) {
			return false
		}
	}
	return true
}

// すべての要素が条件を満たしたらtrue。
func ContainsAllInt64Func(slice []int64, f func(int64) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// 他のスライスの要素をひとつでも内包していたらtrue。
func ContainsAnyInt64(slice []int64, subset []int64) bool {
	for i := range subset {
		if ContainsInt64(slice, subset[i]) {
			return true
		}
	}
	return false
}

// 値と一致する要素の数を返す。
func CountInt64(slice []int64, v int64) int {
	c := 0
	for i := range slice {
		if slice[i] == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func CountInt64Func(slice []int64, f func(int64) bool) int {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			c++
		}
	}
	return c
}

// 値と一致する最初の要素の位置を返す。
func IndexInt64(slice []int64, v int64) int {
	for i := range slice {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最初の要素の位置を返す。
func IndexInt64Func(slice []int64, f func(int64) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値と一致する最後の要素の位置を返す。
func LastIndexInt64(slice []int64, v int64) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最後の要素の位置を返す。
func LastIndexInt64Func(slice []int64, f func(int64) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値をbool型の値に変換したスライスを返す。
func MapInt64ToBool(slice []int64, f func(int64) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をbyte型の値に変換したスライスを返す。
func MapInt64ToByte(slice []int64, f func(int64) byte) []byte {
	dst := make([]byte, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex64型の値に変換したスライスを返す。
func MapInt64ToComplex64(slice []int64, f func(int64) complex64) []complex64 {
	dst := make([]complex64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex128型の値に変換したスライスを返す。
func MapInt64ToComplex128(slice []int64, f func(int64) complex128) []complex128 {
	dst := make([]complex128, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat32型の値に変換したスライスを返す。
func MapInt64ToFloat32(slice []int64, f func(int64) float32) []float32 {
	dst := make([]float32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat64型の値に変換したスライスを返す。
func MapInt64ToFloat64(slice []int64, f func(int64) float64) []float64 {
	dst := make([]float64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint型の値に変換したスライスを返す。
func MapInt64ToInt(slice []int64, f func(int64) int) []int {
	dst := make([]int, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint8型の値に変換したスライスを返す。
func MapInt64ToInt8(slice []int64, f func(int64) int8) []int8 {
	dst := make([]int8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint16型の値に変換したスライスを返す。
func MapInt64ToInt16(slice []int64, f func(int64) int16) []int16 {
	dst := make([]int16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint32型の値に変換したスライスを返す。
func MapInt64ToInt32(slice []int64, f func(int64) int32) []int32 {
	dst := make([]int32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint64型の値に変換したスライスを返す。
func MapInt64ToInt64(slice []int64, f func(int64) int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をrune型の値に変換したスライスを返す。
func MapInt64ToRune(slice []int64, f func(int64) rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をstring型の値に変換したスライスを返す。
func MapInt64ToString(slice []int64, f func(int64) string) []string {
	dst := make([]string, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint型の値に変換したスライスを返す。
func MapInt64ToUint(slice []int64, f func(int64) uint) []uint {
	dst := make([]uint, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint8型の値に変換したスライスを返す。
func MapInt64ToUint8(slice []int64, f func(int64) uint8) []uint8 {
	dst := make([]uint8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint16型の値に変換したスライスを返す。
func MapInt64ToUint16(slice []int64, f func(int64) uint16) []uint16 {
	dst := make([]uint16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint32型の値に変換したスライスを返す。
func MapInt64ToUint32(slice []int64, f func(int64) uint32) []uint32 {
	dst := make([]uint32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint64型の値に変換したスライスを返す。
func MapInt64ToUint64(slice []int64, f func(int64) uint64) []uint64 {
	dst := make([]uint64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 各要素に関数を適用して単一のbool型の値を返す。
func ReduceInt64ToBool(slice []int64, v bool, f func(bool, int64) bool) bool {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のbyte型の値を返す。
func ReduceInt64ToByte(slice []int64, v byte, f func(byte, int64) byte) byte {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex64型の値を返す。
func ReduceInt64ToComplex64(slice []int64, v complex64, f func(complex64, int64) complex64) complex64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex128型の値を返す。
func ReduceInt64ToComplex128(slice []int64, v complex128, f func(complex128, int64) complex128) complex128 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat32型の値を返す。
func ReduceInt64ToFloat32(slice []int64, v float32, f func(float32, int64) float32) float32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat64型の値を返す。
func ReduceInt64ToFloat64(slice []int64, v float64, f func(float64, int64) float64) float64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint型の値を返す。
func ReduceInt64ToInt(slice []int64, v int, f func(int, int64) int) int {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint8型の値を返す。
func ReduceInt64ToInt8(slice []int64, v int8, f func(int8, int64) int8) int8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint16型の値を返す。
func ReduceInt64ToInt16(slice []int64, v int16, f func(int16, int64) int16) int16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint32型の値を返す。
func ReduceInt64ToInt32(slice []int64, v int32, f func(int32, int64) int32) int32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint64型の値を返す。
func ReduceInt64ToInt64(slice []int64, v int64, f func(int64, int64) int64) int64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のrune型の値を返す。
func ReduceInt64ToRune(slice []int64, v rune, f func(rune, int64) rune) rune {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のstring型の値を返す。
func ReduceInt64ToString(slice []int64, v string, f func(string, int64) string) string {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint型の値を返す。
func ReduceInt64ToUint(slice []int64, v uint, f func(uint, int64) uint) uint {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint8型の値を返す。
func ReduceInt64ToUint8(slice []int64, v uint8, f func(uint8, int64) uint8) uint8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint16型の値を返す。
func ReduceInt64ToUint16(slice []int64, v uint16, f func(uint16, int64) uint16) uint16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint32型の値を返す。
func ReduceInt64ToUint32(slice []int64, v uint32, f func(uint32, int64) uint32) uint32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint64型の値を返す。
func ReduceInt64ToUint64(slice []int64, v uint64, f func(uint64, int64) uint64) uint64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 値を繰り返したスライスを返す。
func RepeatInt64(v int64, count int) []int64 {
	slice := make([]int64, count)
	for i := range slice {
		slice[i] = v
	}
	return slice
}

// old を new で置き換えたスライスを返す。
// n < 0 とき、すべての old を new で置き換える。
func ReplaceInt64(slice []int64, old, new int64, n int) []int64 {
	dst := make([]int64, len(slice))
	c := 0
	for i := range slice {
		if n > 0 && c <= n && slice[i] == old {
			dst = append(dst, new)
			c++
		} else {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値と一致する要素で分割したスライスを返す。
func SplitInt64(slice []int64, v int64) [][]int64 {
	ret := [][]int64{[]int64{}}
	for i := range slice {
		if slice[i] == v {
			ret = append(ret, []int64{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 条件を満たす要素で分割したスライスを返す。
func SplitInt64Func(slice []int64, f func(int64) bool) [][]int64 {
	ret := [][]int64{[]int64{}}
	for i := range slice {
		if f(slice[i]) {
			ret = append(ret, []int64{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 値と一致する要素の直後で分割したスライスを返す。
func SplitAfterInt64(slice []int64, v int64) [][]int64 {
	ret := [][]int64{[]int64{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if slice[i] == v {
			ret = append(ret, []int64{})
		}
	}
	return ret
}

// 条件を満たす要素の直後で分割したスライスを返す。
func SplitAfterInt64Func(slice []int64, f func(int64) bool) [][]int64 {
	ret := [][]int64{[]int64{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if f(slice[i]) {
			ret = append(ret, []int64{})
		}
	}
	return ret
}

// 値と一致する最初の要素を返す。
func FindInt64(slice []int64, v int64) (ret int64, ok bool) {
	for _, t := range slice {
		if t == v {
			return t, true
		}
	}
	return
}

// 条件を満たす最初の要素を返す。
func FindInt64Func(slice []int64, f func(int64) bool) (ret int64, ok bool) {
	for _, t := range slice {
		if f(t) {
			return t, true
		}
	}
	return
}

// 値と一致する先頭部分と一致しない残りの部分を返す。
func SpanInt64(slice []int64, v int64) ([]int64, []int64) {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []int64{}
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func SpanInt64Func(slice []int64, f func(int64) bool) ([]int64, []int64) {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []int64{}
}

// 値と一致する先頭のスライスを返す。
// 値と一致しなかった時点で終了する。
func TakeWhileInt64(slice []int64, v int64) []int64 {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i]
		}
	}
	return slice
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhileInt64Func(slice []int64, f func(int64) bool) []int64 {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i]
		}
	}
	return slice
}

// 値と一致する先頭の要素を除いていったスライスを返す。
// 値と一致しなかった時点で終了する。
func DropWhileInt64(slice []int64, v int64) []int64 {
	for i := range slice {
		if slice[i] != v {
			return slice[i:]
		}
	}
	return []int64{}
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhileInt64Func(slice []int64, f func(int64) bool) []int64 {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []int64{}
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func UniqueInt64(slice []int64) []int64 {
	dst := make([]int64, 0, len(slice))

	if len(slice) > 0 {
		dst = append(dst, slice[0])
	}

	for i := 1; i < len(slice); i++ {
		if slice[i] != slice[i-1] {
			dst = append(dst, slice[i])
		}
	}

	return dst
}

// 重複を排除したスライスを返す。
func UniqueInPlaceInt64(slice []int64) []int64 {
	length := len(slice)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; {
			if slice[i] == slice[j] {
				length--
				slice[j], slice[length] = slice[length], slice[j]
			} else {
				j++
			}
		}
	}
	return slice[:length]
}

// 値の一致する要素だけのスライスを返す。
func FilterInt64(slice []int64, v int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, v)
		}
	}
	return dst
}

// 条件を満たす要素だけのスライスを返す。
func FilterInt64Func(slice []int64, f func(int64) bool) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInPlaceInt64(slice []int64, v int64) []int64 {
	c := 0
	for i := range slice {
		if slice[i] == v {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c]
}

// 条件を満たす要素だけのスライスを返す。
func FilterInPlaceInt64Func(slice []int64, f func(int64) bool) []int64 {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c]
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInt64(slice []int64, v int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		if slice[i] != v {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotInt64Func(slice []int64, f func(int64) bool) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInPlaceInt64(slice []int64, v int64) []int64 {
	c := 0
	for i := range slice {
		if slice[i] != v {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c]
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotInPlaceInt64Func(slice []int64, f func(int64) bool) []int64 {
	c := 0
	for i := range slice {
		if !f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c]
}

// 値の一致するスライスと一致しないスライスを返す。
func PartitionInt64(slice []int64, v int64) ([]int64, []int64) {
	a := make([]int64, 0, len(slice)/2)
	b := make([]int64, 0, len(slice)/2)
	for i := range slice {
		if slice[i] == v {
			a = append(a, slice[i])
		} else {
			b = append(b, slice[i])
		}
	}
	return a, b
}

// 値の一致するスライスと一致しないスライスを返す。
func PartitionInPlaceInt64(slice []int64, v int64) ([]int64, []int64) {
	c := 0
	for i := range slice {
		if slice[i] == v {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}

// 条件を満たすスライスと満たさないスライスを返す。
func PartitionInt64Func(slice []int64, f func(int64) bool) ([]int64, []int64) {
	a := make([]int64, 0, len(slice)/2)
	b := make([]int64, 0, len(slice)/2)
	for i := range slice {
		if f(slice[i]) {
			a = append(a, slice[i])
		} else {
			b = append(b, slice[i])
		}
	}
	return a, b
}

// 条件を満たすスライスと満たさないスライスを返す。
func PartitionInPlaceInt64Func(slice []int64, f func(int64) bool) ([]int64, []int64) {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}