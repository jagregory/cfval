package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-builtin.html
func validateJoin(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnJoin, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnJoin)]

	// TODO: this will fail with { "Fn::Join": { "Fn::GetAZs": "" }} and such
	items, ok := value.([]interface{})
	if !ok || items == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::Join\" key: %s", items)}
	}

	if len(items) != 2 {
		return reporting.Reports{reporting.NewFailure(ctx, "Join has incorrect number of arguments (expected: 2, actual: %d)", len(items))}
	}

	reports := make(reporting.Reports, 0, 10)
	delimiter := items[0]
	values := items[1]

	if errs := validateJoinDelimiter(builtin, delimiter, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	if errs := validateJoinList(builtin, values, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateJoinDelimiter(builtin parse.IntrinsicFunction, delimiter interface{}, ctx PropertyContext) reporting.Reports {
	if _, ok := delimiter.(string); !ok {
		return reporting.Reports{reporting.NewFailure(ctx, "\"%s\" is not a valid delimiter", delimiter)}
	}

	return nil
}

func validateJoinList(builtin parse.IntrinsicFunction, values interface{}, ctx PropertyContext) reporting.Reports {
	valuesCtx := PropertyContextAdd(ctx, "Values")

	switch parts := values.(type) {
	case []interface{}:
		reports := make(reporting.Reports, 0, 10)
		valueType := Schema{Type: ValueString}
		for i, part := range parts {
			if errs := validateJoinListValue(part, PropertyContextAdd(NewPropertyContext(valuesCtx, valueType), strconv.Itoa(i))); errs != nil {
				reports = append(reports, errs...)
			}
		}

		return reporting.Safe(reports)
	case parse.IntrinsicFunction:
		listType := Schema{Type: Multiple(ValueString)}
		_, errs := ValidateIntrinsicFunctions(parts, NewPropertyContext(valuesCtx, listType), SupportedFunctions{
			parse.FnBase64:    true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
			parse.FnJoin:      true,
			parse.FnSelect:    true,
			parse.FnRef:       true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Join items are not valid: %s", values)}
}

func validateJoinListValue(value interface{}, ctx PropertyContext) reporting.Reports {
	switch vt := value.(type) {
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(vt, ctx, SupportedFunctions{
			parse.FnBase64:    true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
			parse.FnJoin:      true,
			parse.FnSelect:    true,
			parse.FnRef:       true,
		})
		return errs
	}

	_, errs := ValueString.Validate(value, ctx)
	return errs
}
