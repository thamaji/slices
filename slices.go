package slices

import (
	"math/rand"

	"github.com/thamaji/slices/tuple"
)

// 指定した値をn個複製したスライスを返す。
func Repeat[T any](n int, v T) []T {
	slice := make([]T, n)
	for i := 0; i < n; i++ {
		slice[i] = v
	}
	return slice
}

// 指定した値をn個複製したスライスを返す。
func RepeatBy[T any](n int, f func() T) []T {
	slice := make([]T, n)
	for i := 0; i < n; i++ {
		slice[i] = f()
	}
	return slice
}

// 指定した位置の要素を返す。
func Get[T any](slice []T, index int) (T, bool) {
	if index < len(slice) {
		return slice[index], true
	}
	return *new(T), false
}

// 指定した位置の要素を返す。無い場合はvを返す。
func GetOrElse[T any](slice []T, index int, v T) T {
	if index < len(slice) {
		return slice[index]
	}
	return v
}

// 先頭の要素を返す。
func GetFirst[T any](slice []T) (T, bool) {
	if len(slice) > 0 {
		return slice[0], true
	}
	return *new(T), false
}

// 先頭の要素を返す。無い場合はvを返す。
func GetFirstOrElse[T any](slice []T, v T) T {
	if len(slice) > 0 {
		return slice[0]
	}
	return v
}

// 終端の要素を返す。
func GetLast[T any](slice []T) (T, bool) {
	if len(slice) > 0 {
		return slice[len(slice)-1], true
	}
	return *new(T), false
}

// 終端の要素を返す。無い場合はvを返す。
func GetLastOrElse[T any](slice []T, v T) T {
	if len(slice) > 0 {
		return slice[len(slice)-1]
	}
	return v
}

// 指定した位置の要素を受け取り任意の処理を実行する関数を返す。
// 関数は、要素が範囲内なら true を、範囲外なら false を返す。
func RunWith[T any](slice []T, f func(T)) func(int) bool {
	return func(i int) bool {
		if i >= len(slice) {
			return false
		}
		f(slice[i])
		return true
	}
}

// 指定した位置に要素を追加する。
func Insert[T any](slice []T, index int, v ...T) []T {
	return append(slice[:index], append(v, slice[index:]...)...)
}

// 末尾に要素を追加する。
func Push[T any](slice []T, v ...T) []T {
	return append(slice, v...)
}

// 先頭に要素を追加する。
func PushBack[T any](slice []T, v ...T) []T {
	return append(v, slice...)
}

// 末尾から要素を取り出す。
func Pop[T any](slice []T) (T, []T) {
	if len(slice) == 0 {
		return *new(T), slice
	}
	return slice[len(slice)-1], slice[:len(slice)-1]
}

// 末尾からn個の要素を取り出す。
func PopN[T any](slice []T, n int) ([]T, []T) {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[len(slice)-n:], slice[:len(slice)-n]
}

// 先頭から要素を取り出す。
func PopBack[T any](slice []T) (T, []T) {
	if len(slice) == 0 {
		return *new(T), slice
	}
	return slice[0], slice[1:]
}

// 先頭からn個の要素を取り出す。
func PopBackN[T any](slice []T, n int) ([]T, []T) {
	if n > len(slice) {
		n = len(slice)
	}
	return slice[:n], slice[n:]
}

// 指定した位置の要素を削除する。
func Remove[T any](slice []T, index int) []T {
	return RemoveN(slice, index, 1)
}

// 指定した位置からn個の要素を削除する。
func RemoveN[T any](slice []T, index int, n int) []T {
	if index+n > len(slice) {
		n = len(slice) - index
	}
	if index > len(slice) {
		return slice
	}
	copy(slice[index:], slice[index+n:])
	for i := len(slice) - n; i < len(slice); i++ {
		slice[i] = *new(T)
	}
	return slice[:len(slice)-n]
}

// 要素をすべて削除する。
func Clear[T any](slice []T) []T {
	for i := range slice {
		slice[i] = *new(T)
	}
	return slice[:0]
}

// 要素をすべてコピーしたスライスを返す。
func Clone[T any](slice []T) []T {
	clone := make([]T, len(slice))
	copy(clone, slice)
	return clone
}

// 要素をランダムに入れ替える。
func Shuffle[T any](slice []T, r *rand.Rand) {
	var n int
	for i := 0; i < len(slice); i++ {
		n = r.Intn(i + 1)
		slice[i], slice[n] = slice[n], slice[i]
	}
}

// 要素を１つランダムに返す。
func Sample[T any](slice []T, r *rand.Rand) T {
	return slice[r.Intn(len(slice))]
}

// スライスの各要素を組み合わせたスライスを返す。
func Combine[T any](slice []T, slices ...[]T) [][]T {
	size := len(slice)
	for _, slice := range slices {
		size *= len(slice)
	}

	dst := make([][]T, 0, size)
	for _, v := range slice {
		dst = append(dst, []T{v})
	}

	for _, slice := range slices {
		dst = combine(dst, slice)
	}

	return dst
}

func combine[T any](dst [][]T, slice []T) [][]T {
	size := len(dst)
	var j int
	for i := 0; i < size; i++ {
		for j = 0; j < len(slice); j++ {
			dst = append(dst, append(dst[i], slice[j]))
		}
	}
	return dst[size:]
}

// スライスをn個ずつ分割したスライスを返す。
func Grouped[T any](slice []T, n int) [][]T {
	if len(slice) == 0 {
		return [][]T{}
	}
	grouped := make([][]T, len(slice)/n)
	for i := range slice {
		grouped[i/n] = append(grouped[i/n], slice[i])
	}
	return grouped
}

// １つでも値と一致する要素が存在したらtrue。
func Contains[T comparable](slice []T, v T) bool {
	for i := range slice {
		if slice[i] == v {
			return true
		}
	}
	return false
}

// １つでも条件を満たす要素が存在したらtrue。
func ContainsBy[T any](slice []T, f func(T) bool) bool {
	for i := range slice {
		if f(slice[i]) {
			return true
		}
	}
	return false
}

// 他のスライスのすべての要素を内包していたらtrue。
func ContainsAll[T comparable](slice []T, subset []T) bool {
	for i := range subset {
		if !Contains(slice, subset[i]) {
			return false
		}
	}
	return true
}

// すべての要素が条件を満たしたらtrue。
func ContainsAllBy[T any](slice []T, f func(T) bool) bool {
	for i := range slice {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// 他のスライスの要素をひとつでも内包していたらtrue。
func ContainsAny[T comparable](slice []T, subset []T) bool {
	for i := range subset {
		if Contains(slice, subset[i]) {
			return true
		}
	}
	return false
}

// 他のスライスを内包していたらtrue。
func ContainsSlice[T comparable](slice []T, subset []T) bool {
OUTER:
	for i := range slice {
		if len(slice)-i > len(subset) {
			break
		}
		for j := range subset {
			if slice[i] != subset[j] {
				continue OUTER
			}
		}
		return true
	}
	return false
}

// 値と一致する要素の数を返す。
func Count[T comparable](slice []T, v T) int {
	c := 0
	for i := range slice {
		if slice[i] == v {
			c++
		}
	}
	return c
}

// 条件を満たす要素の数を返す。
func CountBy[T any](slice []T, f func(T) bool) int {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			c++
		}
	}
	return c
}

// 値と一致する最初の要素の位置を返す。
func Index[T comparable](slice []T, v T) int {
	for i := range slice {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最初の要素の位置を返す。
func IndexBy[T any](slice []T, f func(T) bool) int {
	for i := range slice {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値と一致する最後の要素の位置を返す。
func LastIndex[T comparable](slice []T, v T) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == v {
			return i
		}
	}
	return -1
}

// 条件を満たす最後の要素の位置を返す。
func LastIndexBy[T any](slice []T, f func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 逆順にしたスライスを返す。
func Reverse[T any](slice []T) []T {
	dst := make([]T, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		dst = append(dst, slice[i])
	}
	return dst
}

// 逆順にしたスライスを返す。
func ReverseInplace[T any](slice []T) []T {
	for i := 0; i < len(slice)/2; i++ {
		j := len(slice) - i - 1
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// 連番のスライスを返す。
func Range[T ordered](start T, stop T, step T) []T {
	slice := []T{}
	for i := start; i < stop; i += step {
		slice = append(slice, i)
	}
	return slice
}

// スライスのインデックスのスライスを返す。
func Indices[T any](slice []T) []int {
	indices := make([]int, len(slice))
	for i := range indices {
		indices = append(indices, i)
	}
	return indices
}

// 先頭からひとつ目のoldをnewで置き換えたスライスを返す。
func Replace[T comparable](slice []T, old T, new T) []T {
	flag := true
	dst := make([]T, len(slice))
	for i := range slice {
		if flag && slice[i] == old {
			dst[i] = new
			flag = false
		} else {
			dst[i] = slice[i]
		}
	}
	return dst
}

// すべてのoldをnewで置き換えたスライスを返す。
func ReplaceAll[T comparable](slice []T, old T, new T) []T {
	dst := make([]T, len(slice))
	for i := range slice {
		if slice[i] == old {
			dst[i] = new
		} else {
			dst[i] = slice[i]
		}
	}
	return dst
}

// ゼロ値の要素を除いたスライスを返す。
func Clean[T comparable](slice []T) []T {
	zero := *new(T)
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if slice[i] != zero {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値を変換したスライスを返す。
func Map[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	dst := make([]T2, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 要素を先頭から順に演算する。
func Reduce[T any](slice []T, f func(T, T) T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[0]
	for i := 1; i < len(slice); i++ {
		v = f(v, slice[i])
	}
	return v
}

// 要素を終端から順に演算する。
func ReduceRight[T any](slice []T, f func(T, T) T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[len(slice)-1]
	for i := len(slice) - 2; i >= 0; i-- {
		v = f(v, slice[i])
	}
	return v
}

// 初期値と要素を先頭から順に演算する。
func Fold[T1, T2 any](slice []T1, v T2, f func(T2, T1) T2) T2 {
	for i := range slice {
		v = f(v, slice[i])
	}
	return v
}

// 初期値と要素を終端から順に演算する。
func FoldRight[T1, T2 any](slice []T1, v T2, f func(T2, T1) T2) T2 {
	for i := len(slice) - 1; i >= 0; i-- {
		v = f(v, slice[i])
	}
	return v
}

// 初期値と要素を先頭から順に演算して途中経過のスライスを返す。
func Scan[T1, T2 any](slice []T1, v T2, f func(T2, T1) T2) []T2 {
	dst := make([]T2, 0, len(slice)+1)
	dst = append(dst, v)
	for i := range slice {
		v = f(v, slice[i])
		dst = append(dst, v)
	}
	return dst
}

// 初期値と要素を終端から順に演算して途中経過のスライスを返す。
func ScanRight[T1, T2 any](slice []T1, v T2, f func(T2, T1) T2) []T2 {
	dst := make([]T2, 0, len(slice)+1)
	dst = append(dst, v)
	for i := len(slice) - 1; i >= 0; i-- {
		v = f(v, slice[i])
		dst = append(dst, v)
	}
	return dst
}

// スライスを平坦化する。
func Flatten[T any](slice [][]T) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		dst = append(dst, slice[i]...)
	}
	return dst
}

// 値をスライスに変換し、それらを結合したスライスを返す。
func FlatMap[T1, T2 any](slice []T1, f func(T1) []T2) []T2 {
	dst := make([]T2, 0, len(slice))
	for i := range slice {
		dst = append(dst, f(slice[i])...)
	}
	return dst
}

// 値と一致する要素で分割したスライスを返す。
func Split[T comparable](slice []T, v T) [][]T {
	dst := [][]T{{}}
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, []T{})
			continue
		}
		dst[len(dst)-1] = append(dst[len(dst)-1], slice[i])
	}
	return dst
}

// 条件を満たす要素で分割したスライスを返す。
func SplitBy[T any](slice []T, f func(T) bool) [][]T {
	dst := [][]T{{}}
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, []T{})
			continue
		}
		dst[len(dst)-1] = append(dst[len(dst)-1], slice[i])
	}
	return dst
}

// 値と一致する要素の直後で分割したスライスを返す。
func SplitAfter[T comparable](slice []T, v T) [][]T {
	dst := [][]T{{}}
	for i := range slice {
		dst[len(dst)-1] = append(dst[len(dst)-1], slice[i])
		if slice[i] == v {
			dst = append(dst, []T{})
		}
	}
	return dst
}

// 条件を満たす要素の直後で分割したスライスを返す。
func SplitAfterBy[T any](slice []T, f func(T) bool) [][]T {
	dst := [][]T{{}}
	for i := range slice {
		dst[len(dst)-1] = append(dst[len(dst)-1], slice[i])
		if f(slice[i]) {
			dst = append(dst, []T{})
		}
	}
	return dst
}

// 値と一致する最初の要素を返す。
func Find[T comparable](slice []T, v T) (ret T, ok bool) {
	for _, t := range slice {
		if t == v {
			return t, true
		}
	}
	return
}

// 条件を満たす最初の要素を返す。
func FindBy[T any](slice []T, f func(T) bool) (ret T, ok bool) {
	for _, t := range slice {
		if f(t) {
			return t, true
		}
	}
	return
}

// 値と一致する先頭部分と一致しない残りの部分を返す。
func Span[T comparable](slice []T, v T) ([]T, []T) {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []T{}
}

// 条件を満たす先頭部分と満たさない残りの部分を返す。
func SpanBy[T any](slice []T, f func(T) bool) ([]T, []T) {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i], slice[i:]
		}
	}
	return slice, []T{}
}

// 先頭n個の要素を返す。
func Take[T any](slice []T, n int) []T {
	if n > len(slice) {
		return slice
	}
	return slice[:n]
}

// 値と一致する先頭のスライスを返す。
// 値と一致しなかった時点で終了する。
func TakeWhile[T comparable](slice []T, v T) []T {
	for i := range slice {
		if slice[i] != v {
			return slice[0:i]
		}
	}
	return slice
}

// 条件を満たす先頭のスライスを返す。
// 条件を満たさなかった時点で終了する。
func TakeWhileBy[T any](slice []T, f func(T) bool) []T {
	for i := range slice {
		if !f(slice[i]) {
			return slice[0:i]
		}
	}
	return slice
}

// 先頭n個の要素を除いたスライスを返す。
func Drop[T any](slice []T, n int) []T {
	if n > len(slice) {
		return []T{}
	}
	return slice[n:]
}

// 値と一致する先頭の要素を除いていったスライスを返す。
// 値と一致しなかった時点で終了する。
func DropWhile[T comparable](slice []T, v T) []T {
	for i := range slice {
		if slice[i] != v {
			return slice[i:]
		}
	}
	return []T{}
}

// 条件を満たす先頭の要素を除いていったスライスを返す。
// 条件を満たさなかった時点で終了する。
func DropWhileBy[T any](slice []T, f func(T) bool) []T {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []T{}
}

// スライスが一致していたらtrue。
func Equal[T comparable](slices1 []T, slices2 []T) bool {
	if len(slices1) != len(slices2) {
		return false
	}
	for i := 0; i < len(slices1); i++ {
		if slices1[i] != slices2[i] {
			return false
		}
	}
	return true
}

// スライスが一致していたらtrue。
func EqualBy[T any](slices1 []T, slices2 []T, f func(T, T) bool) bool {
	if len(slices1) != len(slices2) {
		return false
	}
	for i := 0; i < len(slices1); i++ {
		if !f(slices1[i], slices2[i]) {
			return false
		}
	}
	return true
}

// スライスの先頭がスライスと一致していたら true を返す。
func StartWith[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) < len(slice2) {
		return false
	}

	for i := range slice2 {
		if i >= len(slice1) {
			return false
		}

		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

// スライスの先頭がスライスと一致していたら true を返す。
func StartWithBy[T any](slice1 []T, slice2 []T, f func(T, T) bool) bool {
	if len(slice1) < len(slice2) {
		return false
	}

	for i := range slice2 {
		if i >= len(slice1) {
			return false
		}

		if !f(slice1[i], slice2[i]) {
			return false
		}
	}
	return true
}

// スライスの終端がスライスと一致していたら true を返す。
func EndWith[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) < len(slice2) {
		return false
	}

	n := len(slice1) - len(slice2)

	for i := len(slice2) - 1; i >= 0; i-- {
		if i+n >= len(slice1) {
			return false
		}
		if slice1[i+n] != slice2[i] {
			return false
		}
	}

	return true
}

// スライスの終端がスライスと一致していたら true を返す。
func EndWithBy[T any](slice1 []T, slice2 []T, f func(T, T) bool) bool {
	if len(slice1) < len(slice2) {
		return false
	}

	n := len(slice1) - len(slice2)

	for i := len(slice2) - 1; i >= 0; i-- {
		if i+n >= len(slice1) {
			return false
		}
		if !f(slice1[i+n], slice2[i]) {
			return false
		}
	}

	return true
}

// 重複を排除したスライスを返す。
// 入力スライスはソートされている必要がある。
func Unique[T comparable](slice []T) []T {
	dst := make([]T, 0, len(slice))

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
// 入力スライスはソートされている必要がある。
func UniqueBy[T any](slice []T, f func(T, T) bool) []T {
	dst := make([]T, 0, len(slice))

	if len(slice) > 0 {
		dst = append(dst, slice[0])
	}

	for i := 1; i < len(slice); i++ {
		if !f(slice[i], slice[i-1]) {
			dst = append(dst, slice[i])
		}
	}

	return dst
}

// 重複を排除したスライスを返す。
func UniqueInplace[T comparable](slice []T) []T {
	size := len(slice)
	var j int
	for i := 0; i < size; i++ {
		for j = i + 1; j < size; {
			if slice[i] == slice[j] {
				size--
				slice[j], slice[size] = slice[size], slice[j]
			} else {
				j++
			}
		}
	}
	return slice[:size]
}

// 重複を排除したスライスを返す。
func UniqueByInplace[T any](slice []T, f func(T, T) bool) []T {
	size := len(slice)
	var j int
	for i := 0; i < size; i++ {
		for j = i + 1; j < size; {
			if f(slice[i], slice[j]) {
				size--
				slice[j], slice[size] = slice[size], slice[j]
			} else {
				j++
			}
		}
	}
	return slice[:size]
}

// 値の一致する要素だけのスライスを返す。
func Filter[T comparable](slice []T, v T) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if slice[i] == v {
			dst = append(dst, v)
		}
	}
	return dst
}

// 条件を満たす要素だけのスライスを返す。
func FilterBy[T any](slice []T, f func(T) bool) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInplace[T comparable](slice []T, v T) []T {
	c := 0
	for i := range slice {
		if slice[i] == v {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	for i := c; i < len(slice); i++ {
		slice[i] = *new(T)
	}
	return slice[:c]
}

// 条件を満たす要素だけのスライスを返す。
func FilterByInplace[T any](slice []T, f func(T) bool) []T {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	for i := c; i < len(slice); i++ {
		slice[i] = *new(T)
	}
	return slice[:c]
}

// 値の一致しない要素だけのスライスを返す。
func FilterNot[T comparable](slice []T, v T) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if slice[i] != v {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotBy[T any](slice []T, f func(T) bool) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInplace[T comparable](slice []T, v T) []T {
	c := 0
	for i := range slice {
		if slice[i] != v {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	for i := c; i < len(slice); i++ {
		slice[i] = *new(T)
	}
	return slice[:c]
}

// 条件を満たさない要素だけのスライスを返す。
func FilterNotByInplace[T any](slice []T, f func(T) bool) []T {
	c := 0
	for i := range slice {
		if !f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	for i := c; i < len(slice); i++ {
		slice[i] = *new(T)
	}
	return slice[:c]
}

// 条件を満たす要素を変換したスライスを返す。
func Collect[T1, T2 any](slice []T1, f func(T1) (T2, bool)) []T2 {
	dst := make([]T2, 0, len(slice))
	for i := range slice {
		if v, ok := f(slice[i]); ok {
			dst = append(dst, v)
		}
	}
	return dst
}

// 要素ごとに関数の返すキーでグルーピングしたマップを返す。
func GroupBy[T1 any, T2 comparable](slice []T1, f func(T1) T2) map[T2][]T1 {
	m := map[T2][]T1{}
	var k T2
	for i := range slice {
		k = f(slice[i])
		if v, ok := m[k]; !ok {
			m[k] = []T1{slice[i]}
		} else {
			m[k] = append(v, slice[i])
		}
	}
	return m
}

// 値の一致するスライスと一致しないスライスを返す。
func Partition[T comparable](slice []T, v T) ([]T, []T) {
	dst1 := make([]T, 0, len(slice)/2)
	dst2 := make([]T, 0, len(slice)/2)
	for i := range slice {
		if slice[i] == v {
			dst1 = append(dst1, slice[i])
		} else {
			dst2 = append(dst2, slice[i])
		}
	}
	return dst1, dst2
}

// 条件を満たすスライスと満たさないスライスを返す。
func PartitionBy[T any](slice []T, f func(T) bool) ([]T, []T) {
	dst1 := make([]T, 0, len(slice)/2)
	dst2 := make([]T, 0, len(slice)/2)
	for i := range slice {
		if f(slice[i]) {
			dst1 = append(dst1, slice[i])
		} else {
			dst2 = append(dst2, slice[i])
		}
	}
	return dst1, dst2
}

// 値の一致するスライスと一致しないスライスを返す。
func PartitionInplace[T comparable](slice []T, v T) ([]T, []T) {
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
func PartitionByInplace[T any](slice []T, f func(T) bool) ([]T, []T) {
	c := 0
	for i := range slice {
		if f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c], slice[c:]
}

// 要素の合計を返す。
func Sum[T ordered | complex](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[0]
	for i := 1; i < len(slice); i++ {
		v += slice[i]
	}
	return v
}

// 要素を変換して合計を返す。
func SumBy[T1 any, T2 ordered](slice []T1, f func(T1) T2) T2 {
	if len(slice) == 0 {
		return *new(T2)
	}

	v := *new(T2)
	for i := 0; i < len(slice); i++ {
		v += f(slice[i])
	}
	return v
}

// 要素をすべて掛ける。
func Product[T integer | float | complex](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[0]
	for i := 1; i < len(slice); i++ {
		v *= slice[i]
	}
	return v
}

// 要素を変換してすべて掛ける。
func ProductBy[T1 any, T2 integer | float](slice []T1, f func(T1) T2) T2 {
	if len(slice) == 0 {
		return *new(T2)
	}

	v := *new(T2)
	for i := 0; i < len(slice); i++ {
		v *= f(slice[i])
	}
	return v
}

// 最大の要素を返す。
func Max[T ordered](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	max := slice[0]
	for i := 1; i < len(slice); i++ {
		if max < slice[i] {
			max = slice[i]
		}
	}
	return max
}

// 要素を変換して最大の要素を返す。
func MaxBy[T1 any, T2 ordered](slice []T1, f func(T1) T2) T2 {
	if len(slice) == 0 {
		return *new(T2)
	}

	max := f(slice[0])
	for i := 1; i < len(slice); i++ {
		v := f(slice[i])
		if max < v {
			max = v
		}
	}
	return max
}

// 最小の要素を返す。
func Min[T ordered](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	min := slice[0]
	for i := 1; i < len(slice); i++ {
		if min > slice[i] {
			min = slice[i]
		}
	}
	return min
}

// 要素を変換して最小の要素を返す。
func MinBy[T1 any, T2 ordered](slice []T1, f func(T1) T2) T2 {
	if len(slice) == 0 {
		return *new(T2)
	}

	min := f(slice[0])
	for i := 1; i < len(slice); i++ {
		v := f(slice[i])
		if min > v {
			min = v
		}
	}
	return min
}

// すべての要素に値を代入する。
func Fill[T any](slice []T, v T) {
	for i := range slice {
		slice[i] = v
	}
}

// すべての要素にゼロ値を代入する。
func FillZero[T any](slice []T) {
	for i := range slice {
		slice[i] = *new(T)
	}
}

// すべての要素に関数の実行結果を代入する。
func FillBy[T any](slice []T, f func(int) T) {
	for i := range slice {
		slice[i] = f(i)
	}
}

// 要素がn個になるまで先頭にvを挿入する。
func Pad[T any](slice []T, n int, v T) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = v
	}
	return append(t, slice...)
}

// 要素がn個になるまで先頭にゼロ値を挿入する。
func PadZero[T any](slice []T, n int) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = *new(T)
	}
	return append(t, slice...)
}

// 要素がn個になるまで先頭に関数の実行結果を挿入する。
func PadBy[T any](slice []T, n int, f func(int) T) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = f(i)
	}
	return append(t, slice...)
}

// 要素がn個になるまで末尾にvを挿入する。
func PadRight[T any](slice []T, n int, v T) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = v
	}
	return append(slice, t...)
}

// 要素がn個になるまで末尾にゼロ値を挿入する。
func PadZeroRight[T any](slice []T, n int) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = *new(T)
	}
	return append(slice, t...)
}

// 要素がn個になるまで末尾に関数の実行結果を挿入する。
func PadRightBy[T any](slice []T, n int, f func(int) T) []T {
	if len(slice) >= n {
		return slice
	}
	c := n - len(slice)
	t := make([]T, c)
	for i := 0; i < len(t); i++ {
		t[i] = f(len(slice) + i)
	}
	return append(slice, t...)
}

// ふたつのスライスの同じ位置の要素をペアにしたスライスを返す。
// ふたつのスライスの要素数が異なる場合、小さいほうに合わせる。
func Zip[T1, T2 any](slice1 []T1, slice2 []T2) []tuple.T2[T1, T2] {
	size := len(slice1)
	if size > len(slice2) {
		size = len(slice2)
	}
	dst := make([]tuple.T2[T1, T2], 0, size)
	for i := 0; i < size; i++ {
		dst = append(dst, tuple.NewT2(slice1[i], slice2[i]))
	}
	return dst
}

// スライスの要素と位置をペアにしたスライスを返す。
func ZipWithIndex[T any](slice []T) []tuple.T2[T, int] {
	dst := make([]tuple.T2[T, int], 0, len(slice))
	for i := range slice {
		dst = append(dst, tuple.NewT2(slice[i], i))
	}
	return dst
}

// 要素を分離してふたつのスライスを返す。
func Unzip[T1, T2 any](slice []tuple.T2[T1, T2]) ([]T1, []T2) {
	dst1 := make([]T1, 0, len(slice))
	dst2 := make([]T2, 0, len(slice))
	for i := range slice {
		dst1 = append(dst1, slice[i].V1)
		dst2 = append(dst2, slice[i].V2)
	}
	return dst1, dst2
}

// 値に区切り要素を挟んでスライスにする。
func Join[T any](separator T, values ...T) []T {
	if len(values) == 0 {
		return []T{}
	}
	slice := make([]T, 0, len(values)*2-1)
	n := len(values) - 1
	for i := 0; i < n; i++ {
		slice = append(slice, values[i], separator)
	}
	slice = append(slice, values[n])
	return slice
}
