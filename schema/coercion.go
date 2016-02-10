package schema

type Coercion int

const (
	CoercionNever Coercion = iota
	CoercionAlways
	CoercionBegrudgingly
)
