package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

func TestBase64(t *testing.T) {
	template := &parse.Template{
		Resources: map[string]parse.TemplateResource{
			"MyResource": parse.TemplateResource{
				Type: "TestResource",
			},
		},
	}
	currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
		"TestResource": Resource{
			Attributes: Properties{
				"InstanceId": Schema{
					Type: InstanceID,
				},

				"Name": Schema{
					Type: ValueString,
				},
			},

			ReturnValue: Schema{
				Type: ValueString,
			},
		},
	}), currentResource, Schema{Type: InstanceID}, ValidationOptions{})

	if _, errs := validateBase64(parse.Base64{map[string]interface{}{"Fn::Base64": 123}}, ctx); errs == nil {
		t.Error("Should fail when invalid type used for args", errs)
	}

	if _, errs := validateBase64(parse.Base64{map[string]interface{}{}}, ctx); errs == nil {
		t.Error("Should fail when no args", errs)
	}

	if _, errs := validateBase64(parse.Base64{map[string]interface{}{"Fn::Base64": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, ctx); errs == nil {
		t.Error("Should fail when valid with extra properties", errs)
	}

	if _, errs := validateBase64(parse.Base64{map[string]interface{}{"Fn::Base64": "blah"}}, ctx); errs != nil {
		t.Error("Should pass when valid types used", errs)
	}

	if _, errs := validateBase64(parse.Base64{map[string]interface{}{"Fn::Base64": parse.Ref{map[string]interface{}{"Ref": "MyResource"}}}}, ctx); errs != nil {
		t.Error("Should short circuit and pass when ref used", errs)
	}
}
