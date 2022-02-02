package slices

import (
	"constraints"
	"math/rand"
)

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
func Insert[T any](slice []T, index int, v T) []T {
	return append(slice[:index], append([]T{v}, slice[index:]...)...)
}

// 指定した位置にスライスの全要素を追加する。
func InsertAll[T any](slice []T, index int, v []T) []T {
	return append(slice[:index], append(v, slice[index:]...)...)
}

// 指定した位置の要素を削除する。
func Remove[T any](slice []T, index int) []T {
	if index < len(slice)-1 {
		copy(slice[index:], slice[index+1:])
	}
	slice[len(slice)-1] = *new(T)
	return slice[:len(slice)-1]
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
func ContainsFunc[T any](slice []T, f func(T) bool) bool {
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
func ContainsAllFunc[T any](slice []T, f func(T) bool) bool {
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
func CountFunc[T any](slice []T, f func(T) bool) int {
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
func IndexFunc[T any](slice []T, f func(T) bool) int {
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
func LastIndexFunc[T any](slice []T, f func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// 値を変換したスライスを返す。
func MapFunc[T1, T2 any](slice []T1, f func(T1) T2) []T2 {
	dst := make([]T2, len(slice))
	for i := range slice {
		dst[i] = f(slice[i])
	}
	return dst
}

// 要素を先頭から順に演算する。
func ReduceFunc[T any](slice []T, f func(T, T) T) T {
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
func ReduceRightFunc[T any](slice []T, f func(T, T) T) T {
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
func FoldFunc[T1, T2 any](slice []T1, v T2) func(func(T2, T1) T2) T2 {
	return func(f func(T2, T1) T2) T2 {
		for i := range slice {
			v = f(v, slice[i])
		}
		return v
	}
}

// 初期値と要素を終端から順に演算する。
func FoldRightFunc[T1, T2 any](slice []T1, v T2) func(func(T2, T1) T2) T2 {
	return func(f func(T2, T1) T2) T2 {
		for i := len(slice) - 1; i >= 0; i-- {
			v = f(v, slice[i])
		}
		return v
	}
}

// 初期値と要素を先頭から順に演算して途中経過のスライスを返す。
func ScanFunc[T1, T2 any](slice []T1, v T2) func(func(T2, T1) T2) []T2 {
	dst := make([]T2, 0, len(slice)+1)
	dst = append(dst, v)
	return func(f func(T2, T1) T2) []T2 {
		for i := range slice {
			v = f(v, slice[i])
			dst = append(dst, v)
		}
		return dst
	}
}

// 初期値と要素を終端から順に演算して途中経過のスライスを返す。
func ScanRightFunc[T1, T2 any](slice []T1, v T2) func(func(T2, T1) T2) []T2 {
	dst := make([]T2, 0, len(slice)+1)
	dst = append(dst, v)
	return func(f func(T2, T1) T2) []T2 {
		for i := len(slice) - 1; i >= 0; i-- {
			v = f(v, slice[i])
			dst = append(dst, v)
		}
		return dst
	}
}

// 値をスライスに変換し、それらを結合したスライスを返す。
func FlatMapFunc[T1, T2 any](slice []T1, f func(T1) []T2) []T2 {
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
func SplitFunc[T any](slice []T, f func(T) bool) [][]T {
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
func SplitAfterFunc[T any](slice []T, f func(T) bool) [][]T {
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
func FindFunc[T any](slice []T, f func(T) bool) (ret T, ok bool) {
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
func SpanFunc[T any](slice []T, f func(T) bool) ([]T, []T) {
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
func TakeWhileFunc[T any](slice []T, f func(T) bool) []T {
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
func DropWhileFunc[T any](slice []T, f func(T) bool) []T {
	for i := range slice {
		if !f(slice[i]) {
			return slice[i:]
		}
	}
	return []T{}
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
func UniqueInPlace[T comparable](slice []T) []T {
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
func FilterFunc[T any](slice []T, f func(T) bool) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致する要素だけのスライスを返す。
func FilterInPlace[T comparable](slice []T, v T) []T {
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
func FilterInPlaceFunc[T any](slice []T, f func(T) bool) []T {
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
func FilterNotFunc[T any](slice []T, f func(T) bool) []T {
	dst := make([]T, 0, len(slice))
	for i := range slice {
		if !f(slice[i]) {
			dst = append(dst, slice[i])
		}
	}
	return dst
}

// 値の一致しない要素だけのスライスを返す。
func FilterNotInPlace[T comparable](slice []T, v T) []T {
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
func FilterNotInPlaceFunc[T any](slice []T, f func(T) bool) []T {
	c := 0
	for i := range slice {
		if !f(slice[i]) {
			slice[c], slice[i] = slice[i], slice[c]
			c++
		}
	}
	return slice[:c]
}

// 条件を満たす要素を変換したスライスを返す。
func CollectFunc[T1, T2 any](slice []T1, f func(T1) (T2, bool)) []T2 {
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

// 値の一致するスライスと一致しないスライスを返す。
func PartitionInPlace[T comparable](slice []T, v T) ([]T, []T) {
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
func PartitionFunc[T any](slice []T, f func(T) bool) ([]T, []T) {
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

// 条件を満たすスライスと満たさないスライスを返す。
func PartitionInPlaceFunc[T any](slice []T, f func(T) bool) ([]T, []T) {
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
func Sum[T constraints.Ordered | constraints.Complex](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[0]
	for i := 1; i < len(slice); i++ {
		v += slice[i]
	}
	return v
}

// 要素をすべて掛ける。
func Product[T constraints.Integer | constraints.Float | constraints.Complex](slice []T) T {
	if len(slice) == 0 {
		return *new(T)
	}

	v := slice[0]
	for i := 1; i < len(slice); i++ {
		v *= slice[i]
	}
	return v
}

// 最大の要素を返す。
func Max[T constraints.Ordered](slice []T) T {
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
func MaxBy[T1 any, T2 constraints.Ordered](slice []T1, f func(T1) T2) T2 {
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
func Min[T constraints.Ordered](slice []T) T {
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
func MinBy[T1 any, T2 constraints.Ordered](slice []T1, f func(T1) T2) T2 {
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

// すべての要素に値を代入する
func Fill[T any](slice []T, v T) {
	for i := range slice {
		slice[i] = v
	}
}

// ふたつのスライスの同じ位置の要素をペアにしたスライスを返す。
func Zip[T1, T2 any](slice1 []T1, slice2 []T2) []Tuple2[T1, T2] {
	size := len(slice1)
	if size > len(slice2) {
		size = len(slice2)
	}
	dst := make([]Tuple2[T1, T2], 0, size)
	for i := 0; i < size; i++ {
		dst = append(dst, NewTuple2(slice1[i], slice2[i]))
	}
	return dst
}

// スライスの要素と位置をペアにしたスライスを返す。
func ZipWithIndex[T any](slice []T) []Tuple2[T, int] {
	dst := make([]Tuple2[T, int], 0, len(slice))
	for i := range slice {
		dst = append(dst, NewTuple2(slice[i], i))
	}
	return dst
}

// 要素を分離してふたつのスライスを返す。
func Unzip[T1, T2 any](slice []Tuple2[T1, T2]) ([]T1, []T2) {
	dst1 := make([]T1, 0, len(slice))
	dst2 := make([]T2, 0, len(slice))
	for i := range slice {
		dst1 = append(dst1, slice[i].V1)
		dst2 = append(dst2, slice[i].V2)
	}
	return dst1, dst2
}
