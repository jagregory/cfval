package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateAndOr(key parse.IntrinsicFunctionSignature, builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	value, found := builtin.UnderlyingMap[string(key)]
	if !found || value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Missing \"%s\" key", key)}
	}

	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"%s\" key: %T", key, value)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, string(key)))}
	}

	if len(args) < 2 || len(args) > 10 {
		return reporting.Reports{reporting.NewFailure(ctx, "Incorrect number of conditions (expected between 2 and 10, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	for i, condition := range args {
		if errs := validateAndOrItem(condition, PropertyContextAdd(ctx, strconv.Itoa(i))); errs != nil {
			reports = append(reports, errs...)
		}
	}

	return reporting.Safe(reports)
}

func validateAndOrItem(value interface{}, ctx PropertyContext) reporting.Reports {
	if value == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Value is null")}
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid condition: %#v", value)}
}
