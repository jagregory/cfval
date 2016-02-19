package parse

import "encoding/json"

type TemplateResource struct {
	Type       string
	properties map[string]interface{}
	Metadata   map[string]interface{}
}

func (d *TemplateResource) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Type                 string
		Properties, Metadata map[string]interface{}
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	d.Type = tmp.Type
	d.properties = tmp.Properties
	d.Metadata = tmp.Metadata

	return nil
}

func (tr TemplateResource) Properties() []string {
	props := make([]string, 0, len(tr.properties))
	for key, _ := range tr.properties {
		props = append(props, key)
	}
	return props
}

func (tr TemplateResource) PropertyValue(name string) (interface{}, bool) {
	val, ok := tr.properties[name]
	return val, ok
}

func NewTemplateResource(awsType string, properties map[string]interface{}) TemplateResource {
	return TemplateResource{
		Type:       awsType,
		properties: properties,
	}
}
