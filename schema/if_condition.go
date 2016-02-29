package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// validateCondition validates the inline { "Condition": "X" } structure. It
// isn't technically a condition itself, acting more like a Ref.
func validateCondition(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	value, found := builtin.UnderlyingMap["Condition"]
	if !found || value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Missing \"Condition\" key")}
	}

	condition, ok := value.(string)
	if !ok || condition == "" {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Condition\" key: %T", value)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Condition"))}
	}

	if _, found := ctx.Template().Conditions[condition]; !found {
		return reporting.Reports{reporting.NewFailure(ctx, "%s is not defined in the Conditions of the template", condition)}
	}

	return nil
}
