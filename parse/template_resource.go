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

	return nil, false
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
