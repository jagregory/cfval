package parse

type TemplateResource struct {
	Tmpl       *Template // TODO: haha, get rid of this
	Type       string
	properties map[string]interface{}
	Metadata   map[string]interface{}
}

func (tr TemplateResource) Template() *Template {
	if tr.Tmpl == nil {
		panic("Nil Template")
	}

	return tr.Tmpl
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

	if ok {
		return val, ok
	}

	// TODO: defaults, need a test for this... should return default value if
	// 			 none specified
	// if def := tr.Definition.Properties[name]; def.Default != nil {
	// 	return def.Default, true
	// }

	return nil, false
}

func (tr TemplateResource) HasProperty(name string, expected interface{}) bool {
	if value, found := tr.properties[name]; found {
		return value == expected
	}

	return false
}

func NewTemplateResource(template *Template, awsType string, properties map[string]interface{}) TemplateResource {
	if template == nil {
		panic("Template is nil")
	}

	return TemplateResource{
		Tmpl:       template,
		Type:       awsType,
		properties: properties,
	}
}
