package schema

import (
	"fmt"

	"github.com/jagregory/cfval/reporting"
)

var pseudoParameters = map[string]bool{
	"AWS::AccountId":        true,
	"AWS::NotificationARNs": true,
	"AWS::NoValue":          true,
	"AWS::Region":           true,
	"AWS::StackId":          true,
	"AWS::StackName":        true,
}

type Ref struct {
	target string
}

func NewRef(target string) Ref {
	return Ref{target}
}

func (ref Ref) InferType(template *Template) interface{} {
	return TypeString
}

func (ref Ref) Validate(template *Template, context []string) (bool, []reporting.Failure) {
	if _, ok := template.Resources[ref.target]; ok {
		// ref is to a resource and we've found it
		// TODO: validate resource ref value is correct type for property
		return true, nil
	} else if _, ok := template.Parameters[ref.target]; ok {
		// ref is to a parameter and we've found it
		// TODO: validate parameter type is correct for property
		return true, nil
	} else if _, ok := pseudoParameters[ref.target]; ok {
		// ref is to a cloudformation pseudo parameter and we've found it
		// TODO: validate parameter type is correct for property
		return true, nil
	}

	return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Ref '%s' is not a resource or parameter", ref), context)}
}
