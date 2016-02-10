package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html
var pseudoParameters = map[string]Schema{
	"AWS::AccountId": Schema{
		Type: ValueString,
	},

	"AWS::NotificationARNs": Schema{
		Type:  ValueString,
		Array: true,
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
	TargetType() ValueType
}

type Ref struct {
	source Schema
	target string
}

func NewRef(source Schema, target string) Ref {
	return Ref{source, target}
}

func (ref Ref) Validate(template *Template, context []string) (reporting.ValidateResult, reporting.Failures) {
	if template == nil {
		panic("Template is nil")
	}

	target := ref.resolveTarget(template)
	if target == nil {
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Ref '%s' is not a resource, parameter, or pseudo-parameter", ref.target), context)}
	}

	// TODO: make this common, so GetAtt and others can use it
	targetType := target.TargetType()
	switch targetType.CoercibleTo(ref.source.Type) {
	case CoercionNever:
		return reporting.ValidateAbort, reporting.Failures{reporting.NewFailure(fmt.Sprintf("Ref value of '%s' is %s but is being assigned to a %s property", ref.target, targetType.Describe(), ref.source.Type.Describe()), context)}
	case CoercionBegrudgingly:
		return reporting.ValidateAbort, reporting.Failures{reporting.NewWarning(fmt.Sprintf("Ref value of '%s' is %s but is being dangerously coerced to a %s property", ref.target, targetType.Describe(), ref.source.Type.Describe()), context)}
	}

	return reporting.ValidateAbort, nil
}

func (ref Ref) InferType(template *Template) ValueType {
	if target := ref.resolveTarget(template); target != nil {
		return target.TargetType()
	}

	return ValueUnknown
}

func (ref Ref) resolveTarget(template *Template) RefTarget {
	if resource, ok := template.Resources[ref.target]; ok {
		return resource.Definition.ReturnValue
	} else if parameter, ok := template.Parameters[ref.target]; ok {
		return parameter
	} else if pseudoParameters, ok := pseudoParameters[ref.target]; ok {
		return pseudoParameters
	}

	return nil
}
