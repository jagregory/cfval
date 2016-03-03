package constraints

import "testing"

func TestConstraintsPropertyExists(t *testing.T) {
	constraint := PropertyExists("MyProp")

	if !constraint.Pass(testCurrentResource{map[string]interface{}{"MyProp": "Value"}, map[string]interface{}{}}) {
		t.Error("Should pass property exists")
	}

	if constraint.Pass(testCurrentResource{}) {
		t.Error("Should fail if property doesn't exist")
	}
}
