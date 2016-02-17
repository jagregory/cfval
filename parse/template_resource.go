package parse

type TemplateResource struct {
	Tmpl       *Template // haha, get rid of this
	Type       string
	Properties map[string]interface{}
	Metadata   map[string]interface{}
}

func (tr TemplateResource) Template() *Template {
	if tr.Tmpl == nil {
		panic("Nil Template")
	}

	return tr.Tmpl
}

func (tr TemplateResource) Property(name string) (interface{}, bool) {
	val, ok := tr.Properties[name]

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
	if value, found := tr.Properties[name]; found {
		return value == expected
	}

	return false
}

func NewTemplateResource(template *Template) TemplateResource {
	if template == nil {
		panic("Template is nil")
	}

	return TemplateResource{Tmpl: template}
}
