package constraints

type BoolConstraint bool

func (b BoolConstraint) Describe(map[string]interface{}) string {
	if bool(b) {
		return "it's mandatory"
	} else {
		return "never"
	}
}

func (b BoolConstraint) Pass(map[string]interface{}) bool {
	return bool(b)
}

var Always Constraint = BoolConstraint(true)
var Never Constraint = BoolConstraint(false)
