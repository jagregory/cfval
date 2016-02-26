package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html#d0e97554
func validateAnd(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::And"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::And\" key")}
	}

	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::And\" key: %T", value)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::And"))}
	}

	if len(args) < 2 || len(args) > 10 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Incorrect number of conditions (expected between 2 and 10, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	for i, condition := range args {
		if _, errs := validateAndItem(condition, PropertyContextAdd(ctx, strconv.Itoa(i))); errs != nil {
			reports = append(reports, errs...)
		}
	}

	return reporting.ValidateOK, reporting.Safe(reports)
}

func validateAndItem(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Value is null")}
	}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnAnd:       true,
			parse.FnCondition: true,
			parse.FnEquals:    true,
			parse.FnFindInMap: true,
			parse.FnIf:        true,
			parse.FnNot:       true,
			parse.FnOr:        true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Invalid condition: %s", value)}
}
