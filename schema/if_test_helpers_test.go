package schema

import "github.com/jagregory/cfval/parse"

var IF = func(key parse.IntrinsicFunctionSignature) func(args interface{}) parse.IntrinsicFunction {
	return func(args interface{}) parse.IntrinsicFunction {
		return parse.IntrinsicFunction{key, map[string]interface{}{
			string(key): args,
		}}
	}
}

func conditionFactory() parse.IntrinsicFunction {
	return IF(parse.FnCondition)("Condition")
}

var ExampleValidIFs = map[parse.IntrinsicFunctionSignature]func() parse.IntrinsicFunction{
	parse.FnAnd: func() parse.IntrinsicFunction {
		return IF(parse.FnAnd)([]interface{}{conditionFactory(), conditionFactory()})
	},

	parse.FnBase64: func() parse.IntrinsicFunction {
		return IF(parse.FnBase64)("a string to encode")
	},

	parse.FnCondition: conditionFactory,

	parse.FnEquals: func() parse.IntrinsicFunction {
		return IF(parse.FnEquals)([]interface{}{"a", "b"})
	},

	parse.FnFindInMap: func() parse.IntrinsicFunction {
		return IF(parse.FnFindInMap)([]interface{}{"MyMap", "Key1", "Key2"})
	},

	parse.FnGetAtt: func() parse.IntrinsicFunction {
		return IF(parse.FnGetAtt)([]interface{}{"MyResource", "Name"})
	},

	parse.FnGetAZs: func() parse.IntrinsicFunction {
		return IF(parse.FnGetAZs)("")
	},

	parse.FnIf: func() parse.IntrinsicFunction {
		return IF(parse.FnIf)([]interface{}{"Condition", "a", "b"})
	},

	parse.FnJoin: func() parse.IntrinsicFunction {
		return IF(parse.FnJoin)([]interface{}{"delim", []interface{}{"item-1", "item-2"}})
	},

	parse.FnNot: func() parse.IntrinsicFunction {
		return IF(parse.FnNot)(conditionFactory())
	},

	parse.FnOr: func() parse.IntrinsicFunction {
		return IF(parse.FnOr)([]interface{}{conditionFactory(), conditionFactory()})
	},

	parse.FnRef: func() parse.IntrinsicFunction {
		return IF(parse.FnRef)("MyResource")
	},

	parse.FnSelect: func() parse.IntrinsicFunction {
		return IF(parse.FnSelect)([]interface{}{float64(1), []interface{}{"1", "2", "3"}})
	},
}
