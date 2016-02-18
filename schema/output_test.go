package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestOutputValidation(t *testing.T) {
	res := Resource{
		Attributes: map[string]Schema{
			"Id": Schema{
				Type: ValueString,
			},
		},

		ReturnValue: Schema{
			Type: ValueString,
		},
	}

	definitions := NewResourceDefinitions(map[string]Resource{
		"TestResource": res,
	})

	template := &parse.Template{
		Resources: map[string]*parse.TemplateResource{
			"MyResource": &parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	context := []string{}

	goodResourceOutput := parse.Output{
		Description: "Ref with valid resource",
		Value:       map[string]interface{}{"Ref": "MyResource"},
	}
	badResourceOutput := parse.Output{
		Description: "Ref with invalid resource",
		Value:       map[string]interface{}{"Ref": "Missing"},
	}
	goodAttrOutput := parse.Output{
		Description: "GetAtt with valid resource",
		Value:       map[string]interface{}{"Fn::GetAtt": []interface{}{"MyResource", "Id"}},
	}
	badAttrOutput := parse.Output{
		Description: "GetAtt with invalid resource",
		Value:       map[string]interface{}{"Fn::GetAtt": []interface{}{"Missing", "Id"}},
	}

	if _, errs := outputValidate(goodResourceOutput, template, definitions, context); errs != nil {
		t.Error("Resource output should pass if resource exists", errs)
	}

	if _, errs := outputValidate(badResourceOutput, template, definitions, context); errs == nil {
		t.Error("Resource output should fail if resource doesn't exist", errs)
	}

	if _, errs := outputValidate(goodAttrOutput, template, definitions, context); errs != nil {
		t.Error("GetAtt output should pass if resource exists", errs)
	}

	if _, errs := outputValidate(badAttrOutput, template, definitions, context); errs == nil {
		t.Error("GetAtt output should fail if resource doesn't exist", errs)
	}
}
