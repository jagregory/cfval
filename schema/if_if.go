package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-conditions.html#d0e97711
func validateIf(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnIf, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnIf)]
	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::If\" key: %T", value)}
	}

	if len(args) != 3 {
		return reporting.Reports{reporting.NewFailure(ctx, "Incorrect number of arguments [condition_name, true_value, false_value]", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	conditionName := args[0]
	trueValue := args[1]
	falseValue := args[2]

	if errs := validateIfConditionName(conditionName, PropertyContextAdd(ctx, "ConditionName")); errs != nil {
		reports = append(reports, errs...)
	}

	if errs := validateIfValue(trueValue, PropertyContextAdd(ctx, "TrueValue")); errs != nil {
		reports = append(reports, errs...)
	}

	if errs := validateIfValue(falseValue, PropertyContextAdd(ctx, "FalseValue")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateIfConditionName(value interface{}, ctx PropertyContext) reporting.Reports {
	if name, ok := value.(string); ok {
		if _, found := ctx.Template().Conditions[name]; found {
			return nil
		}

		return reporting.Reports{reporting.NewFailure(ctx, "%s is not defined in the Conditions of the template", name)}
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid condition %s", value)}
}

func validateIfValue(value interface{}, ctx PropertyContext) reporting.Reports {
	// TODO: This is to fix #45, but in the case of Fn::Ifs I think this is
	//       desirable behaviour. Probably will undo this soon and get all IFs
	//			 inline with this.
	valueType := Schema{Type: ValueString}

	switch t := value.(type) {
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, valueType), SupportedFunctions{
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
