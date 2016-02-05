package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/pseudo-parameter-reference.html
var pseudoParameters = map[string]Schema{
	"AWS::AccountId": Schema{
		Type: TypeString,
	},

	"AWS::NotificationARNs": Schema{
		Type:  TypeString,
		Array: true,
	},

	"AWS::NoValue": Schema{
		Type: TypeString,
	},

	"AWS::Region": Schema{
		Type: TypeString,
	},

	"AWS::StackId": Schema{
		Type: TypeString,
	},

	"AWS::StackName": Schema{
		Type: TypeString,
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

func (ref Ref) InferType(template *Template) ValueType {
	if target := ref.resolveTarget(template); target != nil {
		return target.TargetType()
	}

	return TypeUnknown
}

func (ref Ref) Validate(template *Template, context []string) (bool, []reporting.Failure) {
	target := ref.resolveTarget(template)
	if target == nil {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Ref '%s' is not a resource, parameter, or pseudo-parameter", ref.target), context)}
	}

	// fail if types don't match, except special case TypeUnknown for types with an unspecified Ref
	// TODO: Fix up all resources to have Ref types and remove this special case
	if targetType := target.TargetType(); targetType != ref.source.Type && targetType != TypeUnknown {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Ref value of '%s' is %s but is being assigned to a %s property", ref.target, targetType, ref.source.Type), context)}
	}

	return true, nil
}
