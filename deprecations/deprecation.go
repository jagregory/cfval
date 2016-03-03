package deprecations

import "fmt"

type Deprecation interface {
	Describe() string
}

func Deprecated(description string) Deprecation {
	return simpleDeprecation(description)
}

func ReplacedBy(property, description string) Deprecation {
	return replacement{description, property}
}

type simpleDeprecation string

func (d simpleDeprecation) Describe() string {
	return fmt.Sprintf("Deprecated: %s", d)
}

type replacement struct {
	description, replacedBy string
}

func (d replacement) Describe() string {
	return fmt.Sprintf("Replaced by %s: %s", d.replacedBy, d.description)
}
