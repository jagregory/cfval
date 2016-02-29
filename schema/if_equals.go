package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateEquals(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	value, found := builtin.UnderlyingMap["Fn::Equals"]
	if !found || value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Equals\" key")}
	}

	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::Equals\" key: %T", value)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::Equals"))}
	}

	if len(args) != 2 {
		return reporting.Reports{reporting.NewFailure(ctx, "Incorrect number of arguments (expected: 2, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	left := args[0]
	if errs := validateEqualsItem(left, PropertyContextAdd(ctx, "Value-1")); errs != nil {
		reports = append(reports, errs...)
	}

	right := args[1]
	if errs := validateEqualsItem(right, PropertyContextAdd(ctx, "Value-2")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateEqualsItem(value interface{}, ctx PropertyContext) reporting.Reports {
	if value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Value is null")}
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnFindInMap: true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
			parse.FnRef:       true,
		})
		return errs
	}

	return nil
}
