package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateGetAtt(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::GetAtt"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::GetAtt\" key")}
	}

	args, ok := value.([]interface{})
	if !ok || args == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::GetAtt\" key: %s", args)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::GetAtt"))}
	}

	if len(args) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	resourceID := args[0]
	attributeID := args[1]

	if _, errs := validateGetAttResourceID(builtin, resourceID, ctx); errs != nil {
		reports = append(reports, errs...)
	} else if _, errs := validateGetAttAttributeID(builtin, resourceID, attributeID, ctx); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.ValidateOK, reporting.Safe(reports)
	// 	if resourceID, ok := args[0].(string); ok {
	// 		if resource, ok := template.Resources[resourceID]; ok {
	// 			if attributeName, ok := args[1].(string); ok {
	// 				definition := ctx.Definitions().Lookup(resource.Type)
	// 				if attribute, ok := definition.Attributes[attributeName]; ok {
	// 					// TODO: make this common, so GetAtt and others can use it
	// 					targetType := attribute.Type
	// 					switch targetType.CoercibleTo(ctx.Property().Type) {
	// 					case CoercionNever:
	// 						return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe())}
	// 					case CoercionBegrudgingly:
	// 						return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(ctx, "GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe())}
	// 					}
	//
	// 					return reporting.ValidateAbort, nil
	// 				}
	// 			}
	//
	// 			// attribute not found on resource
	// 			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s.%s is not an attribute", resourceID, args[1])}
	// 		}
	//
	// 		// resource not found
	// 		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt '%s' is not a resource", resourceID)}
	// 	}
	//
	// 	// resource not a string
	// 	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt '%s' is not a valid resource name", args[0])}
}

func validateGetAttResourceID(builtin parse.IntrinsicFunction, resourceID interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	switch t := resourceID.(type) {
	case string:
		if _, found := ctx.Template().Resources[t]; found {
			return reporting.ValidateOK, nil
		}

		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s is not a resource", t)}
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s is not a valid resource name", resourceID)}
}

func validateGetAttAttributeID(builtin parse.IntrinsicFunction, resourceID, attributeID interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	resource := ctx.Template().Resources[resourceID.(string)]
	definition := ctx.Definitions().Lookup(resource.Type)

	switch t := attributeID.(type) {
	case string:
		if attribute, ok := definition.Attributes[t]; ok {
			targetType := attribute.Type
			switch targetType.CoercibleTo(ctx.Property().Type) {
			case CoercionNever:
				return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, t, targetType.Describe(), ctx.Property().Type.Describe())}
			case CoercionBegrudgingly:
				return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(ctx, "GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, t, targetType.Describe(), ctx.Property().Type.Describe())}
			}

			return reporting.ValidateAbort, nil
		}
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnRef: true,
		})
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s.%s is not an attribute", resourceID, attributeID)}
}
