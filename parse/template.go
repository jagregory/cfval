package parse

import "encoding/json"

type Template struct {
	Conditions map[string]Condition
	Outputs    map[string]Output
	Parameters map[string]Parameter
	Resources  map[string]TemplateResource
}

type unmarshalTemplate Template

func (d *Template) UnmarshalJSON(b []byte) error {
	var template unmarshalTemplate

	if err := json.Unmarshal(b, &template); err != nil {
		return err
	}

	d.Conditions = template.Conditions
	d.Outputs = template.Outputs
	d.Parameters = template.Parameters

	d.Resources = make(map[string]TemplateResource, len(template.Resources))
	for logicalID, resource := range template.Resources {
		resource.LogicalID = logicalID
		d.Resources[logicalID] = resource
	}

	return nil
}
