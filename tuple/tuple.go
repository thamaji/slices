package tuple

func NewT2[V1, V2 any](v1 V1, v2 V2) T2[V1, V2] {
	return T2[V1, V2]{v1, v2}
}

type T2[V1, V2 any] struct {
	V1 V1
	V2 V2
}

func (t T2[V1, V2]) Values() (V1, V2) {
	return t.V1, t.V2
}

func NewT3[V1, V2, V3 any](v1 V1, v2 V2, v3 V3) T3[V1, V2, V3] {
	return T3[V1, V2, V3]{v1, v2, v3}
}

type T3[V1, V2, V3 any] struct {
	V1 V1
	V2 V2
	V3 V3
}

func (t T3[V1, V2, V3]) Values() (V1, V2, V3) {
	return t.V1, t.V2, t.V3
}

func NewT4[V1, V2, V3, V4 any](v1 V1, v2 V2, v3 V3, v4 V4) T4[V1, V2, V3, V4] {
	return T4[V1, V2, V3, V4]{v1, v2, v3, v4}
}

type T4[V1, V2, V3, V4 any] struct {
	V1 V1
	V2 V2
	V3 V3
	V4 V4
}

func (t T4[V1, V2, V3, V4]) Values() (V1, V2, V3, V4) {
	return t.V1, t.V2, t.V3, t.V4
}
