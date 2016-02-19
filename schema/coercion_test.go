package schema

import "testing"

type testCase struct {
	from, to PropertyType
	result   Coercion
}

func data() []testCase {
	coercions := []testCase{
		testCase{from: ValueString, to: ValueBool, result: CoercionBegrudgingly},
		testCase{from: ValueString, to: ValueNumber, result: CoercionBegrudgingly},
		testCase{from: ValueString, to: ValueString, result: CoercionAlways},
		testCase{from: ValueString, to: ValueUnknown, result: CoercionBegrudgingly},

		testCase{from: ValueNumber, to: ValueBool, result: CoercionNever},
		testCase{from: ValueNumber, to: ValueNumber, result: CoercionAlways},
		testCase{from: ValueNumber, to: ValueString, result: CoercionAlways},
		testCase{from: ValueNumber, to: ValueUnknown, result: CoercionBegrudgingly},

		testCase{from: ValueBool, to: ValueBool, result: CoercionAlways},
		testCase{from: ValueBool, to: ValueNumber, result: CoercionNever},
		testCase{from: ValueBool, to: ValueString, result: CoercionAlways},
		testCase{from: ValueBool, to: ValueUnknown, result: CoercionBegrudgingly},

		testCase{from: ValueUnknown, to: ValueBool, result: CoercionBegrudgingly},
		testCase{from: ValueUnknown, to: ValueNumber, result: CoercionBegrudgingly},
		testCase{from: ValueUnknown, to: ValueString, result: CoercionBegrudgingly},
		testCase{from: ValueUnknown, to: ValueUnknown, result: CoercionBegrudgingly},
	}

	// TODO: add more types here
	for _, enum := range []PropertyType{ARN, AvailabilityZone, CIDR, JSON, KeyName, Period, VpcID} {
		coercions = append(coercions, testCase{from: enum, to: enum, result: CoercionAlways})

		coercions = append(coercions, testCase{from: enum, to: ValueBool, result: CoercionNever})
		coercions = append(coercions, testCase{from: enum, to: ValueNumber, result: CoercionNever})
		coercions = append(coercions, testCase{from: enum, to: ValueString, result: CoercionAlways})
		coercions = append(coercions, testCase{from: enum, to: ValueUnknown, result: CoercionBegrudgingly})

		coercions = append(coercions, testCase{from: ValueBool, to: enum, result: CoercionNever})
		coercions = append(coercions, testCase{from: ValueNumber, to: enum, result: CoercionNever})
		coercions = append(coercions, testCase{from: ValueString, to: enum, result: CoercionBegrudgingly})
		coercions = append(coercions, testCase{from: ValueUnknown, to: enum, result: CoercionBegrudgingly})
	}

	return coercions
}

func TestCoercions(t *testing.T) {
	for _, c := range data() {
		result := c.from.CoercibleTo(c.to)

		if result != c.result {
			t.Errorf("%s should %s be coercible to %s but is %s", c.from.Describe(), coercionString(c.result), c.to.Describe(), coercionString(result))
		}
	}
}

func coercionString(c Coercion) string {
	switch c {
	case CoercionAlways:
		return "always"
	case CoercionNever:
		return "never"
	case CoercionBegrudgingly:
		return "begrudgingly"
	default:
		panic("Unexpected coercion")
	}
}
