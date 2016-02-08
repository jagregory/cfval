package reporting

type ValidateResult int

const (
	// ValidateOK indicates that validation should continue regardless of
	// whether the current set of tests passed or failed. This is useful for when
	// you are chaining tests together.
	ValidateOK ValidateResult = iota
	// ValidateAbort means validation either passed or failed, and should not
	// execute any more validations; this is used when something like a Ref is
	// validated, no more validations should run from the ValueString suite
	// because they're meaningless against a Ref.
	ValidateAbort
)
