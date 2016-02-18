package constraints

type BoolConstraint bool

func (b BoolConstraint) Describe(CurrentResource) string {
	if bool(b) {
		return "it's mandatory"
	} else {
		return "never"
	}
}

func (b BoolConstraint) Pass(CurrentResource) bool {
	return bool(b)
}

var Always Constraint = BoolConstraint(true)
var Never Constraint = BoolConstraint(false)
