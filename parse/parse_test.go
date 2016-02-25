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
				"Target": { "Ref": "ParamA" }
      },
      "Metadata": {
        "Some": "JSON"
      }
    }
  },

  "Outputs": {
    "OutputA": {
      "Value": "Test"
    }
  }
}`

	template, err := ParseTemplateJSON([]byte(json))

	if err != nil {
		t.Error("Failed to parse template", err)
	}

	if len(template.Resources) != 1 {
		t.Error("Incorrect number of resources found, expected 1 got: %d", len(template.Resources))
	} else if len(template.Resources["ResourceA"].properties) != 2 {
		t.Error("Incorrect number of properties found, expected 2 got ", len(template.Resources["ResourceA"].properties))
	} else {
		if template.Resources["ResourceA"].properties["Name"] != "TestInstance" {
			t.Error("Didn't parse Properties of ResourceA")
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

	if len(template.Outputs) != 1 {
		t.Error("Incorrect number of outputs found, expected 1 got ", len(template.Outputs))
	} else if template.Outputs["OutputA"].Value != "Test" {
		t.Error("Didn't parse OutputA")
	}
}
