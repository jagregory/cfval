package schema

import "github.com/jagregory/cfval/reporting"

type GetAtt struct {
	definition []interface{}
}

func NewGetAtt(definition []interface{}) GetAtt {
	return GetAtt{definition}
}

func (ga GetAtt) Validate(ctx PropertyContext) (result reporting.ValidateResult, reports reporting.Reports) {
	result = reporting.ValidateAbort

	if len(ga.definition) != 2 {
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(ga.definition)))
		return
	}
	reports = append(reports, reporting.NewSuccess(ctx, "GetAtt has correct number of arguments"))

	template := ctx.Template()

	resourceID, ok := ga.definition[0].(string)
	if !ok {
		// resource not a string
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt %s is not a valid resource name", ga.definition[0]))
		return
	}
	reports = append(reports, reporting.NewSuccess(ctx, "GetAtt %s is a valid resource name", ga.definition[0]))

	resource, ok := template.Resources[resourceID]
	if !ok {
		// resource not found
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt %s is not a resource", resourceID))
		return
	}
	reports = append(reports, reporting.NewSuccess(ctx, "GetAtt %s is a resource", resourceID))

	attributeName, ok := ga.definition[1].(string)
	if !ok {
		// attribute not found on resource
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt %s.%s is not an attribute name", resourceID, attributeName))
		return
	}
	reports = append(reports, reporting.NewSuccess(ctx, "GetAtt %s.%s is an attribute name", resourceID, attributeName))

	definition := ctx.Definitions().Lookup(resource.Type)
	attribute, ok := definition.Attributes[attributeName]
	if !ok {
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt %s.%s is not an attribute", resourceID, attributeName))
		return
	}
	reports = append(reports, reporting.NewSuccess(ctx, "GetAtt %s.%s is an attribute", resourceID, attributeName))

	targetType := attribute.Type
	switch targetType.CoercibleTo(ctx.Property().Type) {
	case CoercionNever:
		reports = append(reports, reporting.NewFailure(ctx, "GetAtt value of %s.%s is %s but is being assigned to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe()))
		return
	case CoercionBegrudgingly:
		reports = append(reports, reporting.NewWarning(ctx, "GetAtt value of %s.%s is %s but is being dangerously coerced to a %s property", resourceID, attributeName, targetType.Describe(), ctx.Property().Type.Describe()))
		return
	case CoercionAlways:
		reports = append(reports, reporting.NewSuccess(ctx, "GetAtt value of %s.%s is being assigned to a property with a compatible type", resourceID, attributeName))
		return
	}

	return
}
