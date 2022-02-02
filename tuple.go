package slices

func NewTuple2[T1, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{v1, v2}
}

type Tuple2[T1, T2 any] struct {
	V1 T1
	V2 T2
}

func (t Tuple2[T1, T2]) Values() (T1, T2) {
	return t.V1, t.V2
}

func NewTuple3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) Tuple3[T1, T2, T3] {
	return Tuple3[T1, T2, T3]{v1, v2, v3}
}

type Tuple3[T1, T2, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

func (t Tuple3[T1, T2, T3]) Values() (T1, T2, T3) {
	return t.V1, t.V2, t.V3
}

func NewTuple4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) Tuple4[T1, T2, T3, T4] {
	return Tuple4[T1, T2, T3, T4]{v1, v2, v3, v4}
}

type Tuple4[T1, T2, T3, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
}

func (t Tuple4[T1, T2, T3, T4]) Values() (T1, T2, T3, T4) {
	return t.V1, t.V2, t.V3, t.V4
}
