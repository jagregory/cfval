package constraints

import "testing"

func TestConstraintsPropertyNot(t *testing.T) {
	constraint := PropertyNot("MyProp", "Value")

	if !constraint.Pass(TestCurrentResource{"MyProp": "Yay"}) {
		t.Error("Should pass when property exists with different value")
	}

	if !constraint.Pass(TestCurrentResource{}) {
		t.Error("Should pass when property doesn't exist")
	}

	if constraint.Pass(TestCurrentResource{"MyProp": "Value"}) {
		t.Error("Should fail if property has specified value")
	}
}
