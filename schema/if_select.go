package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-select.html
func validateSelect(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnSelect, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnSelect)]
	switch t := value.(type) {
	case []interface{}:
		return validateSelectParameters(builtin, t, ctx)
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::Select\" key: %s", value)}
}

func validateSelectParameters(builtin parse.IntrinsicFunction, args []interface{}, ctx PropertyContext) reporting.Reports {
	if len(args) != 2 {
		return reporting.Reports{reporting.NewFailure(ctx, "Wrong number of arguments to Fn::Select (expected 2, got %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	index := args[0]
	array := args[1]

	if errs := validateSelectIndex(builtin, index, array, PropertyContextAdd(ctx, "Index")); errs != nil {
		reports = append(reports, errs...)
	}

	if errs := validateSelectArray(builtin, array, PropertyContextAdd(ctx, "Array")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateSelectIndex(builtin parse.IntrinsicFunction, index interface{}, array interface{}, ctx PropertyContext) reporting.Reports {
	if index == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Index cannot be null")}
	}

	switch t := index.(type) {
	case string:
		return reporting.Reports{reporting.NewFailure(ctx, "Wrong type for index %T", index)}
	case float64:
		return validateIndexNumericalValue(t, array, ctx)
	case parse.IntrinsicFunction:
		indexType := Schema{Type: ValueNumber}
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, indexType), SupportedFunctions{
			parse.FnRef:       true,
			parse.FnFindInMap: true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid value for index %#v", index)}
}

func validateIndexNumericalValue(index float64, array interface{}, ctx PropertyContext) reporting.Reports {
	if index < 0 {
		return reporting.Reports{reporting.NewFailure(ctx, "Index cannot less than zero")}
	} else if arr, ok := array.([]interface{}); ok && int(index) >= len(arr) {
		return reporting.Reports{reporting.NewFailure(ctx, "Index cannot greater than array length")}
	}

	return nil
}

func validateSelectArray(builtin parse.IntrinsicFunction, array interface{}, ctx PropertyContext) reporting.Reports {
	if array == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Array cannot be null")}
	}

	switch t := array.(type) {
	case []interface{}:
		reports := make(reporting.Reports, 0, 10)
		for i, item := range t {
			if item == nil {
				reports = append(reports, reporting.NewFailure(PropertyContextAdd(ctx, strconv.Itoa(i)), "Array item cannot be null"))
			}
		}
		return reporting.Safe(reports)
	case parse.IntrinsicFunction:
		arrayType := Schema{Type: Multiple(ValueString)}
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, arrayType), SupportedFunctions{
			parse.FnRef:       true,
			parse.FnFindInMap: true,
			parse.FnGetAtt:    true,
			parse.FnGetAZs:    true,
			parse.FnIf:        true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid value for array %s", array)}
}
