package schema

import "testing"

func TestPropertyTypeEquality(t *testing.T) {
	types := []PropertyType{
		AvailabilityZone,
		InstanceID,
		JSON,
		ValueBool,
		ValueNumber,
		ValueString,
		ValueUnknown,
		NestedResource{Description: "Something"},

		Multiple(AvailabilityZone),
		Multiple(InstanceID),
		Multiple(JSON),
		Multiple(ValueBool),
		Multiple(ValueNumber),
		Multiple(ValueString),
		Multiple(ValueUnknown),
		Multiple(NestedResource{Description: "Something"}),
	}

	for fi, from := range types {
		for ti, to := range types {
			if fi == ti && !from.Same(to) {
				t.Errorf("Expected %s and %s to be the same", from.Describe(), to.Describe())
			} else if fi != ti && from.Same(to) {
				t.Errorf("Expected %s and %s to not be the same", from.Describe(), to.Describe())
			}
		}
	}
}
