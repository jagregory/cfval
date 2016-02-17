package parse

type Template struct {
	Resources  map[string]*TemplateResource
	Parameters map[string]Parameter
	Outputs    map[string]Output
}
