package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html
var pseudoParameters = map[string]Schema{
	"AWS::AccountId": Schema{
		Type: ValueString,
	},

	"AWS::NotificationARNs": Schema{
		Type: Multiple(ValueString), // TODO: ARN loop here
	},

	"AWS::NoValue": Schema{
		Type: ValueString,
	},

	"AWS::Region": Schema{
		Type: ValueString,
	},

	"AWS::StackId": Schema{
		Type: ValueString,
	},

	"AWS::StackName": Schema{
		Type: ValueString,
	},
}

type RefTarget interface {
	TargetType() PropertyType
}

func keysExcept(m map[string]interface{}, ignore string) []string {
	keys := make([]string, 0, len(m)-1)
	for key := range m {
		if key != ignore {
			keys = append(keys, key)
		}
	}
	return keys
}

func validateRef(ref parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	refValue, found := ref.UnderlyingMap["Ref"]
	if !found || refValue == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Ref\" key")}
	}

	refString, ok := refValue.(string)
	if !ok || refString == "" {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Ref\" key: %s", refString)}
	}

	if len(ref.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(ref.UnderlyingMap, "Ref"))}
	}

	target := resolveTarget(refString, ctx.Definitions(), ctx.Template())
	if target == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Ref '%s' is not a resource, parameter, or pseudo-parameter", refString)}
	}

	targetType := target.TargetType()
	if targetType == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "%s cannot be used in a Ref", refString)}
	}

	switch targetType.CoercibleTo(ctx.Property().Type) {
	case CoercionNever:
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Ref value of '%s' is %s but is being assigned to a %s property", refString, targetType.Describe(), ctx.Property().Type.Describe())}
	case CoercionBegrudgingly:
		return reporting.ValidateAbort, reporting.Reports{reporting.NewWarning(ctx, "Ref value of '%s' is %s but is being dangerously coerced to a %s property", refString, targetType.Describe(), ctx.Property().Type.Describe())}
	}

	return reporting.ValidateAbort, nil
}

func resolveTarget(target string, definitions ResourceDefinitions, template *parse.Template) RefTarget {
	if resource, ok := template.Resources[target]; ok {
		return definitions.Lookup(resource.Type)
	} else if parameter, ok := template.Parameters[target]; ok {
		return definitions.LookupParameter(parameter.Type)
	} else if pseudoParameters, ok := pseudoParameters[target]; ok {
		return pseudoParameters
	}

	return nil
}
