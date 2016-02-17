package constraints

import "testing"

func TestConstraintsPropertyIs(t *testing.T) {
	constraint := PropertyIs("MyProp", "Value")

	if !constraint.Pass(map[string]interface{}{"MyProp": "Value"}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(map[string]interface{}{}) {
		t.Error("Should fail if property doesn't exist")
	}

	if constraint.Pass(map[string]interface{}{"MyProp": "Boo"}) {
		t.Error("Should fail if property has wrong value")
	}
}
