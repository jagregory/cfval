package schema

import "github.com/jagregory/cfval/reporting"

type Coercion string

const (
	CoercionNever        Coercion = "never"
	CoercionAlways       Coercion = "always"
	CoercionBegrudgingly Coercion = "begrudgingly"
)

func coerce(from, to PropertyType, ctx PropertyContext) reporting.Reports {
	switch from.CoercibleTo(to) {
	case CoercionNever:
		return reporting.Reports{
			reporting.NewFailure(ctx, "Value of %s used in %s property", from.Describe(), to.Describe()),
		}
	case CoercionBegrudgingly:
		return reporting.Reports{
			reporting.NewWarning(ctx, "%s is dangerously coerced to a %s property", from.Describe(), to.Describe()),
		}
	case CoercionAlways:
		return nil
	}

	return nil
}
