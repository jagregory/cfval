package schema

import (
	"testing"
)

func TestOutputValidation(t *testing.T) {
	template := &Template{
		Resources: map[string]TemplateResource{
			"MyResource": TemplateResource{},
		},
	}
	context := []string{}

	goodResourceOutput := Output{
		Description: "Ref with valid resource",
		Value:       map[string]interface{}{"Ref": "MyResource"},
	}
	badResourceOutput := Output{
		Description: "Ref with invalid resource",
		Value:       map[string]interface{}{"Ref": "Missing"},
	}
	goodAttrOutput := Output{
		Description: "GetAtt with valid resource",
		Value:       map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "Id"}},
	}
	badAttrOutput := Output{
		Description: "GetAtt with invalid resource",
		Value:       map[string]interface{}{"Fn::GetAtt": []interface{}{"Missing", "Id"}},
	}

	if ok, errs := goodResourceOutput.Validate(template, context); !ok {
		t.Error("Resource output should pass if resource exists", errs)
	}

	if ok, errs := badResourceOutput.Validate(template, context); ok {
		t.Error("Resource output should fail if resource doesn't exist", errs)
	}

	if ok, errs := goodAttrOutput.Validate(template, context); !ok {
		t.Error("GetAtt output should pass if resource exists", errs)
	}

	if ok, errs := badAttrOutput.Validate(template, context); ok {
		t.Error("GetAtt output should fail if resource doesn't exist", errs)
	}
}
