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
		testCase{from: ValueString, to: JSON, result: CoercionNever},

		testCase{from: ValueNumber, to: ValueBool, result: CoercionNever},
		testCase{from: ValueNumber, to: ValueString, result: CoercionAlways},
		testCase{from: ValueNumber, to: JSON, result: CoercionNever},

		testCase{from: ValueBool, to: ValueNumber, result: CoercionNever},
		testCase{from: ValueBool, to: ValueString, result: CoercionAlways},
		testCase{from: ValueBool, to: JSON, result: CoercionNever},

		testCase{from: JSON, to: ValueBool, result: CoercionNever},
		testCase{from: JSON, to: ValueNumber, result: CoercionNever},
		testCase{from: JSON, to: ValueString, result: CoercionNever},
	}

	// TODO: add more types here
	for _, enum := range []PropertyType{ARN, AvailabilityZone, CIDR, KeyName, Period, InternetGatewayID, VpcID} {
		coercions = append(coercions, testCase{from: enum, to: ValueBool, result: CoercionNever})
		coercions = append(coercions, testCase{from: enum, to: ValueNumber, result: CoercionNever})
		coercions = append(coercions, testCase{from: enum, to: ValueString, result: CoercionAlways})

		coercions = append(coercions, testCase{from: ValueBool, to: enum, result: CoercionNever})
		coercions = append(coercions, testCase{from: ValueNumber, to: enum, result: CoercionNever})
		coercions = append(coercions, testCase{from: ValueString, to: enum, result: CoercionBegrudgingly})
	}

	return coercions
}

func TestCoercions(t *testing.T) {
	for _, c := range data() {
		if result := c.from.CoercibleTo(c.from); result != CoercionAlways {
			t.Errorf("%s should always be coercible to itself but is %s", c.from.Describe(), coercionString(result))
		}

		if result := c.from.CoercibleTo(c.to); result != c.result {
			t.Errorf("%s should %s be coercible to %s but is %s", c.from.Describe(), coercionString(c.result), c.to.Describe(), coercionString(result))
		}

		if result := c.from.CoercibleTo(ValueUnknown); result != CoercionBegrudgingly {
			t.Errorf("%s should be begrudgingly coercible to Unknown but is %s", c.from.Describe(), coercionString(result))
		}

		if result := ValueUnknown.CoercibleTo(c.to); result != CoercionBegrudgingly {
			t.Errorf("Unknown should be begrudgingly coercible to %s but is %s", c.from.Describe(), coercionString(result))
		}

		mulF := Multiple(c.from)
		mulT := Multiple(c.to)

		if mulF.CoercibleTo(c.to) != CoercionNever {
			t.Errorf("%s should not be coercible to %s", mulF.Describe(), c.to.Describe())
		}

		if mulF.CoercibleTo(mulT) != CoercionNever {
			t.Errorf("%s should not be coercible to %s", mulF.Describe(), mulT.Describe())
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
