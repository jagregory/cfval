package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateGetAtt(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnGetAtt, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnGetAtt)]
	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::GetAtt\" key: %T", value)}
	}

	if len(args) != 2 {
		return reporting.Reports{reporting.NewFailure(ctx, "GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	resourceID := args[0]
	attributeID := args[1]

	if errs := validateGetAttResourceID(builtin, resourceID, ctx); errs != nil {
		reports = append(reports, errs...)
	} else if errs := validateGetAttAttributeID(builtin, resourceID, attributeID, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateGetAttResourceID(builtin parse.IntrinsicFunction, resourceID interface{}, ctx PropertyContext) reporting.Reports {
	switch t := resourceID.(type) {
	case string:
		if _, found := ctx.Template().Resources[t]; found {
			return nil
		}

		return reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s is not a resource", t)}
	}

	return reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s is not a valid resource name", resourceID)}
}

func validateGetAttAttributeID(builtin parse.IntrinsicFunction, resourceID, attributeID interface{}, ctx PropertyContext) reporting.Reports {
	resource := ctx.Template().Resources[resourceID.(string)]
	definition := ctx.Definitions().Lookup(resource.Type)

	switch t := attributeID.(type) {
	case string:
		if attribute, ok := definition.Attributes[t]; ok {
			targetType := attribute.Type
			switch targetType.CoercibleTo(ctx.Property().Type) {
			case CoercionNever:
				return reporting.Reports{reporting.NewFailure(ctx, "GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, t, targetType.Describe(), ctx.Property().Type.Describe())}
			case CoercionBegrudgingly:
				return reporting.Reports{reporting.NewWarning(ctx, "GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, t, targetType.Describe(), ctx.Property().Type.Describe())}
			}

			return nil
		}
	case parse.IntrinsicFunction:
		getAttAttributeNameType := Schema{Type: ValueString}
		_, errs := ValidateIntrinsicFunctions(t, NewPropertyContext(ctx, getAttAttributeNameType), SupportedFunctions{
			parse.FnRef: true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "%s is not an attribute of %s", attributeID, resource.Type)}
}
