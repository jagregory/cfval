package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// validateCondition validates the inline { "Condition": "X" } structure. It
// isn't technically a condition itself, acting more like a Ref.
func validateCondition(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnCondition, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnCondition)]
	condition, ok := value.(string)
	if !ok || condition == "" {
		return reporting.Reports{reporting.NewFailure(ctx, `Invalid type for "Condition" key: %T`, value)}
	}

	if _, found := ctx.Template().Conditions[condition]; !found {
		return reporting.Reports{reporting.NewFailure(ctx, "%s is not defined in the Conditions of the template", condition)}
	}

	return nil
}
