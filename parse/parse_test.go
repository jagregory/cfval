package parse

import "testing"

func TestParsing(t *testing.T) {
	json := `{
  "Parameters": {
    "ParamA": {
      "Type": "String"
    }
  },

	"Conditions": {
		"ConditionA": { "Fn::Equals": ["a", "b"] }
	},

  "Resources": {
    "ResourceA": {
      "Type": "AWS::EC2::Instance",
      "Properties": {
        "Name": "TestInstance",
				"RefTarget": { "Ref": "ParamA" },
				"MapTarget": { "Fn::FindInMap": ["a", "b", "c"] },
				"JoinTarget": { "Fn::Join": ["a", ["b", "c"]] },
				"Base64Target": { "Fn::Base64": { "Ref": "ParamA" } },
				"Array": [{ "Ref": "ParamA" }],
				"Map": {
					"Nested": { "Ref": "ParamA" }
				}
      },
      "Metadata": {
        "Some": "JSON"
      }
    }
  },

  "Outputs": {
    "OutputA": {
      "Value": "Test"
    },
		"OutputB": {
      "Value": { "Ref": "ParamA" }
    },
		"OutputC": {
      "Value": { "Fn::FindInMap": ["a", "b", "c"] }
    },
		"OutputD": {
			"Value": { "Fn::Join": ["a", ["b", "c"]] }
		},
		"OutputE": {
			"Value": { "Fn::Not": { "Condition": "ConditionA" } }
		}
  }
}`

	template, err := ParseTemplateJSON([]byte(json))

	if err != nil {
		t.Error("Failed to parse template", err)
	}

	if len(template.Resources) != 1 {
		t.Errorf("Incorrect number of resources found, expected 1 got: %d", len(template.Resources))
	} else if len(template.Resources["ResourceA"].properties) != 7 {
		t.Errorf("Incorrect number of properties found, expected 7 got %d", len(template.Resources["ResourceA"].properties))
	} else {
		if template.Resources["ResourceA"].properties["Name"] != "TestInstance" {
			t.Error("Didn't parse Properties of ResourceA")
		}

		if b, _ := template.Resources["ResourceA"].properties["RefTarget"].(IntrinsicFunction); b.Key != "Ref" {
			t.Error("Didn't convert Ref")
		}

		if b, _ := template.Resources["ResourceA"].properties["MapTarget"].(IntrinsicFunction); b.Key != "Fn::FindInMap" {
			t.Error("Didn't convert FindInMap")
		}

		if b, _ := template.Resources["ResourceA"].properties["JoinTarget"].(IntrinsicFunction); b.Key != "Fn::Join" {
			t.Error("Didn't convert Join")
		}

		if b, _ := template.Resources["ResourceA"].properties["Base64Target"].(IntrinsicFunction); b.Key != "Fn::Base64" {
			t.Error("Didn't convert Base64")
		} else if nb, _ := b.UnderlyingMap["Fn::Base64"].(IntrinsicFunction); nb.Key != "Ref" {
			t.Error("Didn't convert Ref in Base64", b, nb)
		}

		if b, _ := template.Resources["ResourceA"].properties["Array"].([]interface{})[0].(IntrinsicFunction); b.Key != "Ref" {
			t.Error("Didn't convert Array[Ref]")
		}

		if b, _ := template.Resources["ResourceA"].properties["Map"].(map[string]interface{})["Nested"].(IntrinsicFunction); b.Key != "Ref" {
			t.Error("Didn't convert Map[Ref]")
		}

		if template.Resources["ResourceA"].Metadata["Some"] != "JSON" {
			t.Error("Didn't parse Metadata of ResourceA")
		}
	}

	if len(template.Parameters) != 1 {
		t.Error("Incorrect number of parameters found, expected 1 got ", len(template.Parameters))
	} else if template.Parameters["ParamA"].Type != "String" {
		t.Error("Didn't parse ParamA")
	}

	if len(template.Outputs) != 5 {
		t.Error("Incorrect number of outputs found, expected 5 got ", len(template.Outputs))
	} else {
		if template.Outputs["OutputA"].Value != "Test" {
			t.Error("Didn't parse OutputA")
		}

		if b, _ := template.Outputs["OutputB"].Value.(IntrinsicFunction); b.Key != "Ref" {
			t.Error("Didn't convert output Ref")
		}

		if b, _ := template.Outputs["OutputC"].Value.(IntrinsicFunction); b.Key != "Fn::FindInMap" {
			t.Error("Didn't convert output FindInMap")
		}

		if b, _ := template.Outputs["OutputD"].Value.(IntrinsicFunction); b.Key != "Fn::Join" {
			t.Error("Didn't convert output Join")
		}

		if b, _ := template.Outputs["OutputE"].Value.(IntrinsicFunction); b.Key != "Fn::Not" {
			t.Error("Didn't convert output Join")
		} else if bn, _ := b.UnderlyingMap["Fn::Not"].(IntrinsicFunction); bn.Key != "Condition" {
			t.Error("Didn't convert output Condition in Not")
		}
	}

	if len(template.Conditions) != 1 {
		t.Error("Incorrect number of conditions found, expected 1 got ", len(template.Conditions))
	} else {
		if b, _ := template.Conditions["ConditionA"].Fn.(IntrinsicFunction); b.Key != "Fn::Equals" {
			t.Error("Didn't parse ConditionA")
		}
	}
}
