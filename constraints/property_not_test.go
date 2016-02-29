package constraints

import "testing"

func TestConstraintsPropertyNot(t *testing.T) {
	constraint := PropertyNot("MyProp", "Value")

	if !constraint.Pass(TestCurrentResource{map[string]interface{}{"MyProp": "Yay"}, map[string]interface{}{}}) {
		t.Error("Should pass when property exists with different value")
	}

	if !constraint.Pass(TestCurrentResource{}) {
		t.Error("Should pass when property doesn't exist")
	}

	if !constraint.Pass(TestCurrentResource{map[string]interface{}{}, map[string]interface{}{"MyProp": "Yay"}}) {
		t.Error("Should pass if property doesn't exist but has default value not matching expectation")
	}

	if constraint.Pass(TestCurrentResource{map[string]interface{}{"MyProp": "Value"}, map[string]interface{}{}}) {
		t.Error("Should fail if property has specified value")
	}

	if constraint.Pass(TestCurrentResource{map[string]interface{}{}, map[string]interface{}{"MyProp": "Value"}}) {
		t.Error("Should fail if property has specified default value")
	}
}
