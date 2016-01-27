package main

import "encoding/json"

func parseTemplateJSON(data []byte) (*Template, error) {
	var temp struct {
		Parameters map[string]Parameter
		Resources map[string]struct {
			Type       string
			Properties json.RawMessage
		}
	}

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	template := &Template{
		Resources: make(map[string]Resource),
	}
	template.Parameters = temp.Parameters

	for key, rawResource := range temp.Resources {
		if rawResource.Type == "AWS::EC2::Subnet" {
			var res Aws_Ec2_Subnet
			err = json.Unmarshal(rawResource.Properties, &res)
			if err != nil {
				return nil, err
			}
			template.Resources[key] = res
		}
	}

	return template, nil
}
