package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-builtin.html
func validateJoin(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::Join"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Join\" key")}
	}

	// TODO: this will fail with { "Fn::Join": { "Fn::GetAZs": "" }} and such
	items, ok := value.([]interface{})
	if !ok || items == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::Join\" key: %s", items)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::Join"))}
	}

	if len(items) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join has incorrect number of arguments (expected: 2, actual: %d)", len(items))}
	}

	reports := make(reporting.Reports, 0, 10)
	delimiter := items[0]
	values := items[1]

	if _, errs := validateJoinDelimiter(builtin, delimiter, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	if _, errs := validateJoinList(builtin, values, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.ValidateAbort, reporting.Safe(reports)
}

func validateJoinDelimiter(builtin parse.IntrinsicFunction, delimiter interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if _, ok := delimiter.(string); !ok {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "\"%s\" is not a valid delimiter", delimiter)}
	}

	return reporting.ValidateOK, nil
}

func validateJoinList(builtin parse.IntrinsicFunction, values interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch parts := values.(type) {
	case []interface{}:
		reports := make(reporting.Reports, 0, 10)
		builtinItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
		for i, part := range parts {
			if _, errs := validateJoinListValue(part, PropertyContextAdd(builtinItemContext, "Values", strconv.Itoa(i))); errs != nil {
				reports = append(reports, errs...)
			}
		}

		return reporting.ValidateOK, reporting.Safe(reports)
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(parts, ctx, SupportedFunctions{
			parse.FnBase64:    true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
			parse.FnJoin:      true,
			parse.FnSelect:    true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Join items are not valid: %s", values)}
}

func validateJoinListValue(value interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch vt := value.(type) {
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(vt, ctx, SupportedFunctions{
			parse.FnBase64:    true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
			parse.FnJoin:      true,
			parse.FnSelect:    true,
			parse.FnRef:       true,
		})
	}

	return ValueString.Validate(value, ctx)
}
