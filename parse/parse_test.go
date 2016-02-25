package parse

import "testing"

func TestParsing(t *testing.T) {
	json := `{
  "Parameters": {
    "ParamA": {
      "Type": "String"
    }
  },

  "Resources": {
    "ResourceA": {
      "Type": "AWS::EC2::Instance",
      "Properties": {
        "Name": "TestInstance",
				"RefTarget": { "Ref": "ParamA" },
				"MapTarget": { "Fn::FindInMap": ["a", "b", "c"] }
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
    }
  }
}`

	template, err := ParseTemplateJSON([]byte(json))

	if err != nil {
		t.Error("Failed to parse template", err)
	}

	if len(template.Resources) != 1 {
		t.Error("Incorrect number of resources found, expected 1 got: %d", len(template.Resources))
	} else if len(template.Resources["ResourceA"].properties) != 3 {
		t.Errorf("Incorrect number of properties found, expected 2 got %d", len(template.Resources["ResourceA"].properties))
	} else {
		if template.Resources["ResourceA"].properties["Name"] != "TestInstance" {
			t.Error("Didn't parse Properties of ResourceA")
		}

		if _, ok := template.Resources["ResourceA"].properties["RefTarget"].(Ref); !ok {
			t.Error("Didn't convert Ref")
		}

		if _, ok := template.Resources["ResourceA"].properties["MapTarget"].(FindInMap); !ok {
			t.Error("Didn't convert FindInMap")
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

	if len(template.Outputs) != 3 {
		t.Error("Incorrect number of outputs found, expected 1 got ", len(template.Outputs))
	} else {
		if template.Outputs["OutputA"].Value != "Test" {
			t.Error("Didn't parse OutputA")
		}

		if _, ok := template.Outputs["OutputB"].Value.(Ref); !ok {
			t.Error("Didn't convert Ref")
		}

		if _, ok := template.Outputs["OutputC"].Value.(FindInMap); !ok {
			t.Error("Didn't convert FindInMap")
		}
	}
}
