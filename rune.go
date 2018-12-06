package slices

import (
	"math/rand"
)

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func GetOrDefaultRune(slice []rune, index int, defaultValue rune) rune {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}

// 要素を１つランダムに返す。
func SampleRune(slice []rune, r *rand.Rand) rune {
	return slice[r.Intn(len(slice))]
}

// スライスの各要素を組み合わせたスライスを返す。
func CombineRune(slice []rune, slices ...[]rune) [][]rune {
	size := len(slice)
	for _, slice := range slices {
		size *= len(slice)
	}

	out := make([][]rune, 0, size)
	for _, v := range slice {
		out = append(out, []rune{v})
	}

	for _, slice := range slices {
		out = combineRune(out, slice)
	}

	return out
}

func combineRune(out [][]rune, slice []rune) [][]rune {
	length := len(out)
	for i := 0; i < length; i++ {
		for j := 0; j < len(slice); j++ {
			out = append(out, append(out[i], slice[j]))
		}
	}
	return out[length:]
}

// １つでも値と一致する要素が存在したらtrue。
func ContainsRune(slice []rune, v rune) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

// １つでも条件を満たす要素が存在したらtrue。
func ContainsRuneFunc(slice []rune, f func(rune) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}
	return false
}

// 他のスライスのすべての要素を内包していたらtrue。
func ContainsAllRune(slice []rune, subset []rune) bool {
	for i := range subset {
		if !ContainsRune(slice, subset[i]) {
			return false
		}
	}
	return true
}

// すべての要素が条件を満たしたらtrue。
func ContainsAllRuneFunc(slice []rune, f func(rune) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// 他のスライスの要素をひとつでも内包していたらtrue。
func ContainsAnyRune(slice []rune, subset []rune) bool {
	for i := range subset {
		if ContainsRune(slice, subset[i]) {
			return true
		}
	}
	return false
}

// 値と一致する要素の数を返す。
func CountRune(slice []rune, v rune) int {
	c := 0
	for i := range slice {
		if slice[i] == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func CountRuneFunc(slice []rune, f func(rune) bool) int {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			c++
		}
	}
	return c
}

// 値と一致する最初の要素の位置を返す。
func IndexRune(slice []rune, v rune) int {
	for i := range slice {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最初の要素の位置を返す。
func IndexRuneFunc(slice []rune, f func(rune) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値と一致する最後の要素の位置を返す。
func LastIndexRune(slice []rune, v rune) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最後の要素の位置を返す。
func LastIndexRuneFunc(slice []rune, f func(rune) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値をbool型の値に変換したスライスを返す。
func MapRuneToBool(slice []rune, f func(rune) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をbyte型の値に変換したスライスを返す。
func MapRuneToByte(slice []rune, f func(rune) byte) []byte {
	dst := make([]byte, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex64型の値に変換したスライスを返す。
func MapRuneToComplex64(slice []rune, f func(rune) complex64) []complex64 {
	dst := make([]complex64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex128型の値に変換したスライスを返す。
func MapRuneToComplex128(slice []rune, f func(rune) complex128) []complex128 {
	dst := make([]complex128, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat32型の値に変換したスライスを返す。
func MapRuneToFloat32(slice []rune, f func(rune) float32) []float32 {
	dst := make([]float32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat64型の値に変換したスライスを返す。
func MapRuneToFloat64(slice []rune, f func(rune) float64) []float64 {
	dst := make([]float64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint型の値に変換したスライスを返す。
func MapRuneToInt(slice []rune, f func(rune) int) []int {
	dst := make([]int, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint8型の値に変換したスライスを返す。
func MapRuneToInt8(slice []rune, f func(rune) int8) []int8 {
	dst := make([]int8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint16型の値に変換したスライスを返す。
func MapRuneToInt16(slice []rune, f func(rune) int16) []int16 {
	dst := make([]int16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint32型の値に変換したスライスを返す。
func MapRuneToInt32(slice []rune, f func(rune) int32) []int32 {
	dst := make([]int32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint64型の値に変換したスライスを返す。
func MapRuneToInt64(slice []rune, f func(rune) int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をrune型の値に変換したスライスを返す。
func MapRuneToRune(slice []rune, f func(rune) rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をstring型の値に変換したスライスを返す。
func MapRuneToString(slice []rune, f func(rune) string) []string {
	dst := make([]string, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint型の値に変換したスライスを返す。
func MapRuneToUint(slice []rune, f func(rune) uint) []uint {
	dst := make([]uint, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint8型の値に変換したスライスを返す。
func MapRuneToUint8(slice []rune, f func(rune) uint8) []uint8 {
	dst := make([]uint8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint16型の値に変換したスライスを返す。
func MapRuneToUint16(slice []rune, f func(rune) uint16) []uint16 {
	dst := make([]uint16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint32型の値に変換したスライスを返す。
func MapRuneToUint32(slice []rune, f func(rune) uint32) []uint32 {
	dst := make([]uint32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint64型の値に変換したスライスを返す。
func MapRuneToUint64(slice []rune, f func(rune) uint64) []uint64 {
	dst := make([]uint64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 各要素に関数を適用して単一のbool型の値を返す。
func ReduceRuneToBool(slice []rune, v bool, f func(bool, rune) bool) bool {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のbyte型の値を返す。
func ReduceRuneToByte(slice []rune, v byte, f func(byte, rune) byte) byte {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex64型の値を返す。
func ReduceRuneToComplex64(slice []rune, v complex64, f func(complex64, rune) complex64) complex64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex128型の値を返す。
func ReduceRuneToComplex128(slice []rune, v complex128, f func(complex128, rune) complex128) complex128 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat32型の値を返す。
func ReduceRuneToFloat32(slice []rune, v float32, f func(float32, rune) float32) float32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat64型の値を返す。
func ReduceRuneToFloat64(slice []rune, v float64, f func(float64, rune) float64) float64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint型の値を返す。
func ReduceRuneToInt(slice []rune, v int, f func(int, rune) int) int {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint8型の値を返す。
func ReduceRuneToInt8(slice []rune, v int8, f func(int8, rune) int8) int8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint16型の値を返す。
func ReduceRuneToInt16(slice []rune, v int16, f func(int16, rune) int16) int16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint32型の値を返す。
func ReduceRuneToInt32(slice []rune, v int32, f func(int32, rune) int32) int32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint64型の値を返す。
func ReduceRuneToInt64(slice []rune, v int64, f func(int64, rune) int64) int64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のrune型の値を返す。
func ReduceRuneToRune(slice []rune, v rune, f func(rune, rune) rune) rune {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のstring型の値を返す。
func ReduceRuneToString(slice []rune, v string, f func(string, rune) string) string {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint型の値を返す。
func ReduceRuneToUint(slice []rune, v uint, f func(uint, rune) uint) uint {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint8型の値を返す。
func ReduceRuneToUint8(slice []rune, v uint8, f func(uint8, rune) uint8) uint8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint16型の値を返す。
func ReduceRuneToUint16(slice []rune, v uint16, f func(uint16, rune) uint16) uint16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint32型の値を返す。
func ReduceRuneToUint32(slice []rune, v uint32, f func(uint32, rune) uint32) uint32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint64型の値を返す。
func ReduceRuneToUint64(slice []rune, v uint64, f func(uint64, rune) uint64) uint64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 値を繰り返したスライスを返す。
func RepeatRune(v rune, count int) []rune {
	slice := make([]rune, count)
	for i := range slice {
		slice[i] = v
	}
	return slice
}

// old を new で置き換えたスライスを返す。
// n < 0 とき、すべての old を new で置き換える。
func ReplaceRune(slice []rune, old, new rune, n int) []rune {
	dst := make([]rune, len(slice))
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
func SplitRune(slice []rune, v rune) [][]rune {
	ret := [][]rune{[]rune{}}
	for i := range slice {
		if slice[i] == v {
			ret = append(ret, []rune{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 条件を満たす要素で分割したスライスを返す。
func SplitRuneFunc(slice []rune, f func(rune) bool) [][]rune {
	ret := [][]rune{[]rune{}}
	for i := range slice {
		if f(slice[i]) {
			ret = append(ret, []rune{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 値と一致する要素の直後で分割したスライスを返す。
func SplitAfterRune(slice []rune, v rune) [][]rune {
	ret := [][]rune{[]rune{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if slice[i] == v {
			ret = append(ret, []rune{})
		}
	}
	return ret
}

// 条件を満たす要素の直後で分割したスライスを返す。
func SplitAfterRuneFunc(slice []rune, f func(rune) bool) [][]rune {
	ret := [][]rune{[]rune{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if f(slice[i]) {
			ret = append(ret, []rune{})
		}
	}
	return ret
}

// 値と一致する最初の要素を返す。
func FindRune(slice []rune, v rune) (ret rune, ok bool) {
	for _, t := range slice {
		if t == v {
			return t, true
		}
	}
	return
}

// 条件を満たす最初の要素を返す。
func FindRuneFunc(slice []rune, f func(rune) bool) (ret rune, ok bool) {
	for _, t := range slice {
		if f(t) {
			return t, true
		}
	}
	return
}

// 値と一致する先頭部分と一致しない残りの部分を返す。
func SpanRune(slice []rune, v rune) ([]rune, []rune) {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []rune{}
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func SpanRuneFunc(slice []rune, f func(rune) bool) ([]rune, []rune) {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []rune{}
}

// 値と一致する先頭のスライスを返す。
// 値と一致しなかった時点で終了する。
func TakeWhileRune(slice []rune, v rune) []rune {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i]
		}
	}
	return slice
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhileRuneFunc(slice []rune, f func(rune) bool) []rune {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i]
		}
	}
	return slice
}

// 値と一致する先頭の要素を除いていったスライスを返す。
// 値と一致しなかった時点で終了する。
func DropWhileRune(slice []rune, v rune) []rune {
	for i := range slice {
		if slice[i] != v {
			return slice[i:]
		}
	}
	return []rune{}
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhileRuneFunc(slice []rune, f func(rune) bool) []rune {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []rune{}
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func UniqueRune(slice []rune) []rune {
	dst := make([]rune, 0, len(slice))

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
func UniqueInPlaceRune(slice []rune) []rune {
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
func FilterRune(slice []rune, v rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, v)
		}
	}
	return dst
}

// 条件を満たす要素だけのスライスを返す。
func FilterRuneFunc(slice []rune, f func(rune) bool) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInPlaceRune(slice []rune, v rune) []rune {
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
func FilterInPlaceRuneFunc(slice []rune, f func(rune) bool) []rune {
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
func FilterNotRune(slice []rune, v rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		if slice[i] != v {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotRuneFunc(slice []rune, f func(rune) bool) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInPlaceRune(slice []rune, v rune) []rune {
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
func FilterNotInPlaceRuneFunc(slice []rune, f func(rune) bool) []rune {
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
func PartitionRune(slice []rune, v rune) ([]rune, []rune) {
	a := make([]rune, 0, len(slice)/2)
	b := make([]rune, 0, len(slice)/2)
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
func PartitionInPlaceRune(slice []rune, v rune) ([]rune, []rune) {
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
func PartitionRuneFunc(slice []rune, f func(rune) bool) ([]rune, []rune) {
	a := make([]rune, 0, len(slice)/2)
	b := make([]rune, 0, len(slice)/2)
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
func PartitionInPlaceRuneFunc(slice []rune, f func(rune) bool) ([]rune, []rune) {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}