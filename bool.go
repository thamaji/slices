package slices

import (
	"math/rand"
)

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func GetOrDefaultBool(slice []bool, index int, defaultValue bool) bool {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}

// 要素を１つランダムに返す。
func SampleBool(slice []bool, r *rand.Rand) bool {
	return slice[r.Intn(len(slice))]
}

// スライスの各要素を組み合わせたスライスを返す。
func CombineBool(slice []bool, slices ...[]bool) [][]bool {
	size := len(slice)
	for _, slice := range slices {
		size *= len(slice)
	}

	out := make([][]bool, 0, size)
	for _, v := range slice {
		out = append(out, []bool{v})
	}

	for _, slice := range slices {
		out = combineBool(out, slice)
	}

	return out
}

func combineBool(out [][]bool, slice []bool) [][]bool {
	length := len(out)
	for i := 0; i < length; i++ {
		for j := 0; j < len(slice); j++ {
			out = append(out, append(out[i], slice[j]))
		}
	}
	return out[length:]
}

// １つでも値と一致する要素が存在したらtrue。
func ContainsBool(slice []bool, v bool) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

// １つでも条件を満たす要素が存在したらtrue。
func ContainsBoolFunc(slice []bool, f func(bool) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}
	return false
}

// 他のスライスのすべての要素を内包していたらtrue。
func ContainsAllBool(slice []bool, subset []bool) bool {
	for i := range subset {
		if !ContainsBool(slice, subset[i]) {
			return false
		}
	}
	return true
}

// すべての要素が条件を満たしたらtrue。
func ContainsAllBoolFunc(slice []bool, f func(bool) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// 他のスライスの要素をひとつでも内包していたらtrue。
func ContainsAnyBool(slice []bool, subset []bool) bool {
	for i := range subset {
		if ContainsBool(slice, subset[i]) {
			return true
		}
	}
	return false
}

// 値と一致する要素の数を返す。
func CountBool(slice []bool, v bool) int {
	c := 0
	for i := range slice {
		if slice[i] == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func CountBoolFunc(slice []bool, f func(bool) bool) int {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			c++
		}
	}
	return c
}

// 値と一致する最初の要素の位置を返す。
func IndexBool(slice []bool, v bool) int {
	for i := range slice {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最初の要素の位置を返す。
func IndexBoolFunc(slice []bool, f func(bool) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値と一致する最後の要素の位置を返す。
func LastIndexBool(slice []bool, v bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最後の要素の位置を返す。
func LastIndexBoolFunc(slice []bool, f func(bool) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値をbool型の値に変換したスライスを返す。
func MapBoolToBool(slice []bool, f func(bool) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をbyte型の値に変換したスライスを返す。
func MapBoolToByte(slice []bool, f func(bool) byte) []byte {
	dst := make([]byte, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex64型の値に変換したスライスを返す。
func MapBoolToComplex64(slice []bool, f func(bool) complex64) []complex64 {
	dst := make([]complex64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex128型の値に変換したスライスを返す。
func MapBoolToComplex128(slice []bool, f func(bool) complex128) []complex128 {
	dst := make([]complex128, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat32型の値に変換したスライスを返す。
func MapBoolToFloat32(slice []bool, f func(bool) float32) []float32 {
	dst := make([]float32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat64型の値に変換したスライスを返す。
func MapBoolToFloat64(slice []bool, f func(bool) float64) []float64 {
	dst := make([]float64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint型の値に変換したスライスを返す。
func MapBoolToInt(slice []bool, f func(bool) int) []int {
	dst := make([]int, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint8型の値に変換したスライスを返す。
func MapBoolToInt8(slice []bool, f func(bool) int8) []int8 {
	dst := make([]int8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint16型の値に変換したスライスを返す。
func MapBoolToInt16(slice []bool, f func(bool) int16) []int16 {
	dst := make([]int16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint32型の値に変換したスライスを返す。
func MapBoolToInt32(slice []bool, f func(bool) int32) []int32 {
	dst := make([]int32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint64型の値に変換したスライスを返す。
func MapBoolToInt64(slice []bool, f func(bool) int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をrune型の値に変換したスライスを返す。
func MapBoolToRune(slice []bool, f func(bool) rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をstring型の値に変換したスライスを返す。
func MapBoolToString(slice []bool, f func(bool) string) []string {
	dst := make([]string, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint型の値に変換したスライスを返す。
func MapBoolToUint(slice []bool, f func(bool) uint) []uint {
	dst := make([]uint, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint8型の値に変換したスライスを返す。
func MapBoolToUint8(slice []bool, f func(bool) uint8) []uint8 {
	dst := make([]uint8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint16型の値に変換したスライスを返す。
func MapBoolToUint16(slice []bool, f func(bool) uint16) []uint16 {
	dst := make([]uint16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint32型の値に変換したスライスを返す。
func MapBoolToUint32(slice []bool, f func(bool) uint32) []uint32 {
	dst := make([]uint32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint64型の値に変換したスライスを返す。
func MapBoolToUint64(slice []bool, f func(bool) uint64) []uint64 {
	dst := make([]uint64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 各要素に関数を適用して単一のbool型の値を返す。
func ReduceBoolToBool(slice []bool, v bool, f func(bool, bool) bool) bool {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のbyte型の値を返す。
func ReduceBoolToByte(slice []bool, v byte, f func(byte, bool) byte) byte {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex64型の値を返す。
func ReduceBoolToComplex64(slice []bool, v complex64, f func(complex64, bool) complex64) complex64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex128型の値を返す。
func ReduceBoolToComplex128(slice []bool, v complex128, f func(complex128, bool) complex128) complex128 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat32型の値を返す。
func ReduceBoolToFloat32(slice []bool, v float32, f func(float32, bool) float32) float32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat64型の値を返す。
func ReduceBoolToFloat64(slice []bool, v float64, f func(float64, bool) float64) float64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint型の値を返す。
func ReduceBoolToInt(slice []bool, v int, f func(int, bool) int) int {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint8型の値を返す。
func ReduceBoolToInt8(slice []bool, v int8, f func(int8, bool) int8) int8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint16型の値を返す。
func ReduceBoolToInt16(slice []bool, v int16, f func(int16, bool) int16) int16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint32型の値を返す。
func ReduceBoolToInt32(slice []bool, v int32, f func(int32, bool) int32) int32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint64型の値を返す。
func ReduceBoolToInt64(slice []bool, v int64, f func(int64, bool) int64) int64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のrune型の値を返す。
func ReduceBoolToRune(slice []bool, v rune, f func(rune, bool) rune) rune {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のstring型の値を返す。
func ReduceBoolToString(slice []bool, v string, f func(string, bool) string) string {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint型の値を返す。
func ReduceBoolToUint(slice []bool, v uint, f func(uint, bool) uint) uint {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint8型の値を返す。
func ReduceBoolToUint8(slice []bool, v uint8, f func(uint8, bool) uint8) uint8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint16型の値を返す。
func ReduceBoolToUint16(slice []bool, v uint16, f func(uint16, bool) uint16) uint16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint32型の値を返す。
func ReduceBoolToUint32(slice []bool, v uint32, f func(uint32, bool) uint32) uint32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint64型の値を返す。
func ReduceBoolToUint64(slice []bool, v uint64, f func(uint64, bool) uint64) uint64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 値を繰り返したスライスを返す。
func RepeatBool(v bool, count int) []bool {
	slice := make([]bool, count)
	for i := range slice {
		slice[i] = v
	}
	return slice
}

// old を new で置き換えたスライスを返す。
// n < 0 とき、すべての old を new で置き換える。
func ReplaceBool(slice []bool, old, new bool, n int) []bool {
	dst := make([]bool, len(slice))
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
func SplitBool(slice []bool, v bool) [][]bool {
	ret := [][]bool{[]bool{}}
	for i := range slice {
		if slice[i] == v {
			ret = append(ret, []bool{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 条件を満たす要素で分割したスライスを返す。
func SplitBoolFunc(slice []bool, f func(bool) bool) [][]bool {
	ret := [][]bool{[]bool{}}
	for i := range slice {
		if f(slice[i]) {
			ret = append(ret, []bool{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 値と一致する要素の直後で分割したスライスを返す。
func SplitAfterBool(slice []bool, v bool) [][]bool {
	ret := [][]bool{[]bool{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if slice[i] == v {
			ret = append(ret, []bool{})
		}
	}
	return ret
}

// 条件を満たす要素の直後で分割したスライスを返す。
func SplitAfterBoolFunc(slice []bool, f func(bool) bool) [][]bool {
	ret := [][]bool{[]bool{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if f(slice[i]) {
			ret = append(ret, []bool{})
		}
	}
	return ret
}

// 値と一致する最初の要素を返す。
func FindBool(slice []bool, v bool) (ret bool, ok bool) {
	for _, t := range slice {
		if t == v {
			return t, true
		}
	}
	return
}

// 条件を満たす最初の要素を返す。
func FindBoolFunc(slice []bool, f func(bool) bool) (ret bool, ok bool) {
	for _, t := range slice {
		if f(t) {
			return t, true
		}
	}
	return
}

// 値と一致する先頭部分と一致しない残りの部分を返す。
func SpanBool(slice []bool, v bool) ([]bool, []bool) {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []bool{}
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func SpanBoolFunc(slice []bool, f func(bool) bool) ([]bool, []bool) {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []bool{}
}

// 値と一致する先頭のスライスを返す。
// 値と一致しなかった時点で終了する。
func TakeWhileBool(slice []bool, v bool) []bool {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i]
		}
	}
	return slice
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhileBoolFunc(slice []bool, f func(bool) bool) []bool {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i]
		}
	}
	return slice
}

// 値と一致する先頭の要素を除いていったスライスを返す。
// 値と一致しなかった時点で終了する。
func DropWhileBool(slice []bool, v bool) []bool {
	for i := range slice {
		if slice[i] != v {
			return slice[i:]
		}
	}
	return []bool{}
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhileBoolFunc(slice []bool, f func(bool) bool) []bool {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []bool{}
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func UniqueBool(slice []bool) []bool {
	dst := make([]bool, 0, len(slice))

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
func UniqueInPlaceBool(slice []bool) []bool {
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
func FilterBool(slice []bool, v bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, v)
		}
	}
	return dst
}

// 条件を満たす要素だけのスライスを返す。
func FilterBoolFunc(slice []bool, f func(bool) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInPlaceBool(slice []bool, v bool) []bool {
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
func FilterInPlaceBoolFunc(slice []bool, f func(bool) bool) []bool {
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
func FilterNotBool(slice []bool, v bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		if slice[i] != v {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotBoolFunc(slice []bool, f func(bool) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInPlaceBool(slice []bool, v bool) []bool {
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
func FilterNotInPlaceBoolFunc(slice []bool, f func(bool) bool) []bool {
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
func PartitionBool(slice []bool, v bool) ([]bool, []bool) {
	a := make([]bool, 0, len(slice)/2)
	b := make([]bool, 0, len(slice)/2)
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
func PartitionInPlaceBool(slice []bool, v bool) ([]bool, []bool) {
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
func PartitionBoolFunc(slice []bool, f func(bool) bool) ([]bool, []bool) {
	a := make([]bool, 0, len(slice)/2)
	b := make([]bool, 0, len(slice)/2)
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
func PartitionInPlaceBoolFunc(slice []bool, f func(bool) bool) ([]bool, []bool) {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}