package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateGetAtt(getAtt parse.GetAtt, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	getAttValue, found := getAtt.UnderlyingMap["Fn::GetAtt"]
	if !found || getAttValue == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::GetAtt\" key")}
	}

	getAttArray, ok := getAttValue.([]interface{})
	if !ok || getAttArray == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::GetAtt\" key: %s", getAttArray)}
	}

	if len(getAtt.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(getAtt.UnderlyingMap, "Fn::GetAtt"))}
	}

	if len(getAttArray) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(getAttArray))}
	}

	template := ctx.Template()
	if resourceID, ok := getAttArray[0].(string); ok {
		if resource, ok := template.Resources[resourceID]; ok {
			if attributeName, ok := getAttArray[1].(string); ok {
				definition := ctx.Definitions().Lookup(resource.Type)
				if attribute, ok := definition.Attributes[attributeName]; ok {
					// TODO: make this common, so GetAtt and others can use it
					targetType := attribute.Type
					switch targetType.CoercibleTo(ctx.Property().Type) {
					case CoercionNever:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe())}
					case CoercionBegrudgingly:
						return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(ctx, "GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe())}
					}

					return reporting.ValidateAbort, nil
				}
			}

			// attribute not found on resource
			return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt %s.%s is not an attribute", resourceID, getAttArray[1])}
		}

		// resource not found
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt '%s' is not a resource", resourceID)}
	}

	// resource not a string
	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "GetAtt '%s' is not a valid resource name", getAttArray[0])}
}
