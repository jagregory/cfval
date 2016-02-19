package constraints

type CurrentResource interface {
	Properties() []string
	PropertyValueOrDefault(string) (interface{}, bool)
	PropertyValue(string) (interface{}, bool)
	PropertyDefault(string) (interface{}, bool)
}
