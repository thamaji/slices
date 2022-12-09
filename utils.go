package slices

import "github.com/thamaji/slices/tuple"

// 指定した値のポインタを返す。
func GetPtr[T any](v T) *T {
	return &v
}

// ポインタをスライスに変換する。
func FromPtr[T any](v *T) []T {
	if v == nil {
		return []T{}
	}
	return []T{*v}
}

// マップをスライスに変換する。
func FromMap[K ordered, V any](m map[K]V) []tuple.T2[K, V] {
	entries := make([]tuple.T2[K, V], 0, len(m))
	for key, value := range m {
		entries = append(entries, tuple.NewT2(key, value))
	}
	return entries
}

// マップのキーをスライスに変換する。
func FromMapKeys[K ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// マップの値をスライスに変換する。
func FromMapValues[K ordered, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// stringをruneのスライスに変換する
func FromString(v string) []rune {
	return []rune(v)
}

// 値をスライスに変換する。
func FromValue[T any](values ...T) []T {
	return append(make([]T, 0, len(values)), values...)
}

// 関数をn回実行した結果をスライスに変換する。
func FromFunc[T any](n int, f func(int) T) []T {
	slice := make([]T, 0, n)
	for i := 0; i < n; i++ {
		slice = append(slice, f(i))
	}
	return slice
}
