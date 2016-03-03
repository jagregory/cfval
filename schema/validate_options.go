package schema

//go:generate stringer -type=ValidateOption

type ValidationOptions map[ValidateOption]bool

type ValidateOption int

const (
	OptionExperimentDisableObjectArrayCoercion ValidateOption = iota
)
