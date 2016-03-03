package deprecations

import "fmt"

type Deprecation interface {
	Describe(property string) string
}

func Deprecated(description string) Deprecation {
	return simpleDeprecation(description)
}

func ReplacedBy(property, description string) Deprecation {
	return replacement{description, property}
}

type simpleDeprecation string

func (d simpleDeprecation) Describe(property string) string {
	return fmt.Sprintf("%s is deprecated: %s", property, d)
}

type replacement struct {
	description, replacedBy string
}

func (d replacement) Describe(property string) string {
	return fmt.Sprintf("%s is replaced by %s: %s", property, d.replacedBy, d.description)
}
