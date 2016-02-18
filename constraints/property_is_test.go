package constraints

import "testing"

func TestConstraintsPropertyIs(t *testing.T) {
	constraint := PropertyIs("MyProp", "Value")

	if !constraint.Pass(TestCurrentResource{"MyProp": "Value"}) {
		t.Error("Should pass property exists")
	}

	if !constraint.Pass(TestCurrentResource{}) {
		t.Error("Should pass if property doesn't exist but has a matching default value in the schema")
	}

	if constraint.Pass(TestCurrentResource{}) {
		t.Error("Should fail if property doesn't exist and has a non-matching default value in the schema")
	}

	if constraint.Pass(TestCurrentResource{}) {
		t.Error("Should fail if property doesn't exist")
	}

	if constraint.Pass(TestCurrentResource{"MyProp": "Boo"}) {
		t.Error("Should fail if property has wrong value")
	}
}
