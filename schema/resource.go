package schema

import "github.com/jagregory/cfval/reporting"

type Resource struct {
	AwsType      string
	Attributes   map[string]Schema
	Properties   ValidatableProperties
	ReturnValue  Schema
	ValidateFunc func(ResourceContext) (reporting.ValidateResult, reporting.Reports)
}

func (rd Resource) PropertyDefault(name string) (interface{}, bool) {
	return rd.Properties.PropertyDefault(name)
}

func (rd Resource) Validate(ctx ResourceContext) (reporting.ValidateResult, reporting.Reports) {
	if rd.ValidateFunc != nil {
		return rd.ValidateFunc(ctx)
	}

	failures := rd.Properties.Validate(ctx)

	return reporting.ValidateOK, reporting.Safe(failures)
}

func (rd Resource) TargetType() PropertyType {
	return rd.ReturnValue.TargetType()
}

func NewUnrecognisedResource(awsType string) Resource {
	return Resource{
		ValidateFunc: func(ctx ResourceContext) (reporting.ValidateResult, reporting.Reports) {
			return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Unrecognised resource %s", awsType)}
		},
	}
}
