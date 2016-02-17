package constraints

import "testing"

func TestConstraintsPropertyExists(t *testing.T) {
	constraint := PropertyExists("MyProp")

	if !constraint.Pass(map[string]interface{}{"MyProp": "Value"}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(map[string]interface{}{}) {
		t.Error("Should fail if property doesn't exist")
	}
}
