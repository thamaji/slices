package slices

import (
	"math/rand"
)

// 指定した位置の要素を返す。indexが範囲外のときはdefaultValueを返す。
func GetOrDefault{{.Type|ToUpper}}(slice []{{.Type}}, index int, defaultValue {{.Type}}) {{.Type}} {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}

// 要素を１つランダムに返す。
func Sample{{.Type|ToUpper}}(slice []{{.Type}}, r *rand.Rand) {{.Type}} {
	return slice[r.Intn(len(slice))]
}

// スライスの各要素を組み合わせたスライスを返す。
func Combine{{.Type|ToUpper}}(slice []{{.Type}}, slices ...[]{{.Type}}) [][]{{.Type}} {
	size := len(slice)
	for _, slice := range slices {
		size *= len(slice)
	}

	out := make([][]{{.Type}}, 0, size)
	for _, v := range slice {
		out = append(out, []{{.Type}}{v})
	}

	for _, slice := range slices {
		out = combine{{.Type|ToUpper}}(out, slice)
	}

	return out
}

func combine{{.Type|ToUpper}}(out [][]{{.Type}}, slice []{{.Type}}) [][]{{.Type}} {
	length := len(out)
	for i := 0; i < length; i++ {
		for j := 0; j < len(slice); j++ {
			out = append(out, append(out[i], slice[j]))
		}
	}
	return out[length:]
}

// １つでも値と一致する要素が存在したらtrue。
func Contains{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

// １つでも条件を満たす要素が存在したらtrue。
func Contains{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}
	return false
}

// 他のスライスのすべての要素を内包していたらtrue。
func ContainsAll{{.Type|ToUpper}}(slice []{{.Type}}, subset []{{.Type}}) bool {
	for i := range subset {
		if !Contains{{.Type|ToUpper}}(slice, subset[i]) {
			return false
		}
	}
	return true
}

// すべての要素が条件を満たしたらtrue。
func ContainsAll{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// 他のスライスの要素をひとつでも内包していたらtrue。
func ContainsAny{{.Type|ToUpper}}(slice []{{.Type}}, subset []{{.Type}}) bool {
	for i := range subset {
		if Contains{{.Type|ToUpper}}(slice, subset[i]) {
			return true
		}
	}
	return false
}

// 値と一致する要素の数を返す。
func Count{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) int {
	c := 0
	for i := range slice {
		if slice[i] == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func Count{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) int {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			c++
		}
	}
	return c
}

// 値と一致する最初の要素の位置を返す。
func Index{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) int {
	for i := range slice {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最初の要素の位置を返す。
func Index{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値と一致する最後の要素の位置を返す。
func LastIndex{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最後の要素の位置を返す。
func LastIndex{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値をbool型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToBool(slice []{{.Type}}, f func({{.Type}}) bool) []bool {
	dst := make([]bool, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をbyte型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToByte(slice []{{.Type}}, f func({{.Type}}) byte) []byte {
	dst := make([]byte, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex64型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToComplex64(slice []{{.Type}}, f func({{.Type}}) complex64) []complex64 {
	dst := make([]complex64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をcomplex128型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToComplex128(slice []{{.Type}}, f func({{.Type}}) complex128) []complex128 {
	dst := make([]complex128, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat32型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToFloat32(slice []{{.Type}}, f func({{.Type}}) float32) []float32 {
	dst := make([]float32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をfloat64型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToFloat64(slice []{{.Type}}, f func({{.Type}}) float64) []float64 {
	dst := make([]float64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToInt(slice []{{.Type}}, f func({{.Type}}) int) []int {
	dst := make([]int, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint8型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToInt8(slice []{{.Type}}, f func({{.Type}}) int8) []int8 {
	dst := make([]int8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint16型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToInt16(slice []{{.Type}}, f func({{.Type}}) int16) []int16 {
	dst := make([]int16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint32型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToInt32(slice []{{.Type}}, f func({{.Type}}) int32) []int32 {
	dst := make([]int32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をint64型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToInt64(slice []{{.Type}}, f func({{.Type}}) int64) []int64 {
	dst := make([]int64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をrune型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToRune(slice []{{.Type}}, f func({{.Type}}) rune) []rune {
	dst := make([]rune, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をstring型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToString(slice []{{.Type}}, f func({{.Type}}) string) []string {
	dst := make([]string, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToUint(slice []{{.Type}}, f func({{.Type}}) uint) []uint {
	dst := make([]uint, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint8型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToUint8(slice []{{.Type}}, f func({{.Type}}) uint8) []uint8 {
	dst := make([]uint8, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint16型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToUint16(slice []{{.Type}}, f func({{.Type}}) uint16) []uint16 {
	dst := make([]uint16, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint32型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToUint32(slice []{{.Type}}, f func({{.Type}}) uint32) []uint32 {
	dst := make([]uint32, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 値をuint64型の値に変換したスライスを返す。
func Map{{.Type|ToUpper}}ToUint64(slice []{{.Type}}, f func({{.Type}}) uint64) []uint64 {
	dst := make([]uint64, 0, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 各要素に関数を適用して単一のbool型の値を返す。
func Reduce{{.Type|ToUpper}}ToBool(slice []{{.Type}}, v bool, f func(bool, {{.Type}}) bool) bool {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のbyte型の値を返す。
func Reduce{{.Type|ToUpper}}ToByte(slice []{{.Type}}, v byte, f func(byte, {{.Type}}) byte) byte {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex64型の値を返す。
func Reduce{{.Type|ToUpper}}ToComplex64(slice []{{.Type}}, v complex64, f func(complex64, {{.Type}}) complex64) complex64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のcomplex128型の値を返す。
func Reduce{{.Type|ToUpper}}ToComplex128(slice []{{.Type}}, v complex128, f func(complex128, {{.Type}}) complex128) complex128 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat32型の値を返す。
func Reduce{{.Type|ToUpper}}ToFloat32(slice []{{.Type}}, v float32, f func(float32, {{.Type}}) float32) float32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のfloat64型の値を返す。
func Reduce{{.Type|ToUpper}}ToFloat64(slice []{{.Type}}, v float64, f func(float64, {{.Type}}) float64) float64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint型の値を返す。
func Reduce{{.Type|ToUpper}}ToInt(slice []{{.Type}}, v int, f func(int, {{.Type}}) int) int {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint8型の値を返す。
func Reduce{{.Type|ToUpper}}ToInt8(slice []{{.Type}}, v int8, f func(int8, {{.Type}}) int8) int8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint16型の値を返す。
func Reduce{{.Type|ToUpper}}ToInt16(slice []{{.Type}}, v int16, f func(int16, {{.Type}}) int16) int16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint32型の値を返す。
func Reduce{{.Type|ToUpper}}ToInt32(slice []{{.Type}}, v int32, f func(int32, {{.Type}}) int32) int32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のint64型の値を返す。
func Reduce{{.Type|ToUpper}}ToInt64(slice []{{.Type}}, v int64, f func(int64, {{.Type}}) int64) int64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のrune型の値を返す。
func Reduce{{.Type|ToUpper}}ToRune(slice []{{.Type}}, v rune, f func(rune, {{.Type}}) rune) rune {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のstring型の値を返す。
func Reduce{{.Type|ToUpper}}ToString(slice []{{.Type}}, v string, f func(string, {{.Type}}) string) string {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint型の値を返す。
func Reduce{{.Type|ToUpper}}ToUint(slice []{{.Type}}, v uint, f func(uint, {{.Type}}) uint) uint {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint8型の値を返す。
func Reduce{{.Type|ToUpper}}ToUint8(slice []{{.Type}}, v uint8, f func(uint8, {{.Type}}) uint8) uint8 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint16型の値を返す。
func Reduce{{.Type|ToUpper}}ToUint16(slice []{{.Type}}, v uint16, f func(uint16, {{.Type}}) uint16) uint16 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint32型の値を返す。
func Reduce{{.Type|ToUpper}}ToUint32(slice []{{.Type}}, v uint32, f func(uint32, {{.Type}}) uint32) uint32 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 各要素に関数を適用して単一のuint64型の値を返す。
func Reduce{{.Type|ToUpper}}ToUint64(slice []{{.Type}}, v uint64, f func(uint64, {{.Type}}) uint64) uint64 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 値を繰り返したスライスを返す。
func Repeat{{.Type|ToUpper}}(v {{.Type}}, count int) []{{.Type}} {
	slice := make([]{{.Type}}, count)
	for i := range slice {
		slice[i] = v
	}
	return slice
}

// old を new で置き換えたスライスを返す。
// n < 0 とき、すべての old を new で置き換える。
func Replace{{.Type|ToUpper}}(slice []{{.Type}}, old, new {{.Type}}, n int) []{{.Type}} {
	dst := make([]{{.Type}}, len(slice))
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
func Split{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) [][]{{.Type}} {
	ret := [][]{{.Type}}{[]{{.Type}}{}}
	for i := range slice {
		if slice[i] == v {
			ret = append(ret, []{{.Type}}{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 条件を満たす要素で分割したスライスを返す。
func Split{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) [][]{{.Type}} {
	ret := [][]{{.Type}}{[]{{.Type}}{}}
	for i := range slice {
		if f(slice[i]) {
			ret = append(ret, []{{.Type}}{})
			continue
		}
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
	}
	return ret
}

// 値と一致する要素の直後で分割したスライスを返す。
func SplitAfter{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) [][]{{.Type}} {
	ret := [][]{{.Type}}{[]{{.Type}}{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if slice[i] == v {
			ret = append(ret, []{{.Type}}{})
		}
	}
	return ret
}

// 条件を満たす要素の直後で分割したスライスを返す。
func SplitAfter{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) [][]{{.Type}} {
	ret := [][]{{.Type}}{[]{{.Type}}{}}
	for i := range slice {
		ret[len(ret)-1] = append(ret[len(ret)-1], slice[i])
		if f(slice[i]) {
			ret = append(ret, []{{.Type}}{})
		}
	}
	return ret
}

// 値と一致する最初の要素を返す。
func Find{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) (ret {{.Type}}, ok bool) {
	for _, t := range slice {
		if t == v {
			return t, true
		}
	}
	return
}

// 条件を満たす最初の要素を返す。
func Find{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) (ret {{.Type}}, ok bool) {
	for _, t := range slice {
		if f(t) {
			return t, true
		}
	}
	return
}

// 値と一致する先頭部分と一致しない残りの部分を返す。
func Span{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) ([]{{.Type}}, []{{.Type}}) {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []{{.Type}}{}
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func Span{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) ([]{{.Type}}, []{{.Type}}) {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []{{.Type}}{}
}

// 値と一致する先頭のスライスを返す。
// 値と一致しなかった時点で終了する。
func TakeWhile{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i]
		}
	}
	return slice
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhile{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i]
		}
	}
	return slice
}

// 値と一致する先頭の要素を除いていったスライスを返す。
// 値と一致しなかった時点で終了する。
func DropWhile{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
	for i := range slice {
		if slice[i] != v {
			return slice[i:]
		}
	}
	return []{{.Type}}{}
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhile{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []{{.Type}}{}
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func Unique{{.Type|ToUpper}}(slice []{{.Type}}) []{{.Type}} {
	dst := make([]{{.Type}}, 0, len(slice))

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
func UniqueInPlace{{.Type|ToUpper}}(slice []{{.Type}}) []{{.Type}} {
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
func Filter{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
	dst := make([]{{.Type}}, 0, len(slice))
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, v)
		}
	}
	return dst
}

// 条件を満たす要素だけのスライスを返す。
func Filter{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
	dst := make([]{{.Type}}, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInPlace{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
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
func FilterInPlace{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
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
func FilterNot{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
	dst := make([]{{.Type}}, 0, len(slice))
	for i := range slice {
		if slice[i] != v {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNot{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
	dst := make([]{{.Type}}, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInPlace{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) []{{.Type}} {
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
func FilterNotInPlace{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) []{{.Type}} {
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
func Partition{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) ([]{{.Type}}, []{{.Type}}) {
	a := make([]{{.Type}}, 0, len(slice)/2)
	b := make([]{{.Type}}, 0, len(slice)/2)
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
func PartitionInPlace{{.Type|ToUpper}}(slice []{{.Type}}, v {{.Type}}) ([]{{.Type}}, []{{.Type}}) {
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
func Partition{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) ([]{{.Type}}, []{{.Type}}) {
	a := make([]{{.Type}}, 0, len(slice)/2)
	b := make([]{{.Type}}, 0, len(slice)/2)
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
func PartitionInPlace{{.Type|ToUpper}}Func(slice []{{.Type}}, f func({{.Type}}) bool) ([]{{.Type}}, []{{.Type}}) {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}