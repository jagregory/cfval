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

func TestConstraintsPropertyNot(t *testing.T) {
	constraint := PropertyNot("MyProp", "Value")

	if !constraint.Pass(map[string]interface{}{"MyProp": "Yay"}) {
		t.Error("Should pass when property exists with different value")
	}

	if !constraint.Pass(map[string]interface{}{}) {
		t.Error("Should pass when property doesn't exist")
	}

	if constraint.Pass(map[string]interface{}{"MyProp": "Value"}) {
		t.Error("Should fail if property has specified value")
	}
}
