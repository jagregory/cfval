package parse

type Template struct {
	Conditions map[string]Condition
	Outputs    map[string]Output
	Parameters map[string]Parameter
	Resources  map[string]TemplateResource
}
