package constraints

type CurrentResource interface {
	Id() string
	AwsType() string
	Properties() []string
	PropertyDefault(string) (interface{}, bool)
	PropertyValue(string) (interface{}, bool)
	PropertyValueOrDefault(string) (interface{}, bool)
}
