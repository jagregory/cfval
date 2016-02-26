package schema

import (
	"testing"

	"github.com/jagregory/cfval/parse"
)

type IFScenario struct {
	fn      parse.IntrinsicFunction
	pass    bool
	message string
}

func TestAnd(t *testing.T) {
	// template := &parse.Template{
	// 	Resources: map[string]parse.TemplateResource{
	// 		"MyResource": parse.TemplateResource{
	// 			Type: "TestResource",
	// 		},
	// 	},
	// }
	// currentResource := ResourceWithDefinition{parse.TemplateResource{}, Resource{}}
	// ctx := NewContextShorthand(template, NewResourceDefinitions(map[string]Resource{
	// 	"TestResource": Resource{
	// 		Attributes: Properties{
	// 			"InstanceId": Schema{
	// 				Type: InstanceID,
	// 			},
	//
	// 			"Name": Schema{
	// 				Type: ValueString,
	// 			},
	// 		},
	//
	// 		ReturnValue: Schema{
	// 			Type: ValueString,
	// 		},
	// 	},
	// }), currentResource, Schema{Type: ValueString}, ValidationOptions{})
	//
	// scenarios := []IFScenario{
	// 	IFScenario{IF(parse.FnAnd)(123), false, "invalid type used for args"},
	// 	IFScenario{IF(parse.FnAnd)(nil), false, "nil used for args"},
	// 	IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{}}, false, "empty map"},
	// 	IFScenario{parse.IntrinsicFunction{"Fn::And", map[string]interface{}{"Fn::And": []interface{}{"a", []interface{}{"b", "c"}}, "blah": "blah"}}, false, "extra properties"},
	// 	IFScenario{IF(parse.FnAnd)([]interface{}{"a", "b", "c"}), false, "too many arguments"},
	// 	IFScenario{IF(parse.FnAnd)([]interface{}{"a"}), false, "too few arguments"},
	// }

	// TODO: Can't do AND until Conditions are implemented...
	// for i, s := range scenarios {
	// 	_, errs := validateAnd(s.fn, ctx)
	// 	if s.pass && errs != nil {
	// 		t.Errorf("Scenario %d: Should pass with %s (errs: %s)", i+1, s.message, errs)
	// 	} else if !s.pass && errs == nil {
	// 		t.Errorf("Scenario %d: Should fail with %s", i+1, s.message)
	// 	}
	// }
}
