package parse

type IntrinsicFunction struct {
	Key           IntrinsicFunctionSignature
	UnderlyingMap map[string]interface{}
}

type IntrinsicFunctionSignature string

const (
	FnAnd       IntrinsicFunctionSignature = "Fn::And"
	FnBase64                               = "Fn::Base64"
	FnEquals                               = "Fn::Equals"
	FnFindInMap                            = "Fn::FindInMap"
	FnGetAtt                               = "Fn::GetAtt"
	FnGetAZs                               = "Fn::GetAZs"
	FnIf                                   = "Fn::If"
	FnJoin                                 = "Fn::Join"
	FnNot                                  = "Fn::Not"
	FnOr                                   = "Fn::Or"
	FnRef                                  = "Ref"
	FnSelect                               = "Fn::Select"

	// FnCondition is not really a Intrinsic Function, but it behaves like one for
	// parsing purposes.
	FnCondition = "Condition"
)

type IntrinsicFunctionSignatures []IntrinsicFunctionSignature

func (s IntrinsicFunctionSignatures) Except(keys ...IntrinsicFunctionSignature) IntrinsicFunctionSignatures {
	set := make(map[IntrinsicFunctionSignature]bool, len(keys))
	for _, k := range keys {
		set[k] = true
	}

	arr := make(IntrinsicFunctionSignatures, 0, len(s)-1)
	for _, k := range s {
		if !set[k] {
			arr = append(arr, k)
		}
	}
	return arr
}

var AllIntrinsicFunctions = IntrinsicFunctionSignatures{
	FnAnd,
	FnBase64,
	FnCondition,
	FnEquals,
	FnFindInMap,
	FnGetAtt,
	FnGetAZs,
	FnIf,
	FnJoin,
	FnNot,
	FnOr,
	FnRef,
	FnSelect,
}

func convertAnyIntrinsicFunctions(value interface{}, fns IntrinsicFunctionSignatures) interface{} {
	for _, key := range fns {
		if fn, ok := convertToIntrinsicFunction(key, value); ok {
			return fn
		}
	}

	switch t := value.(type) {
	case map[string]interface{}:
		return convertMapToIntrinsicFunction(t, fns)
	case []interface{}:
		return convertArrayToIntrinsicFunction(t, fns)
	default:
		return value
	}
}

func convertMapToIntrinsicFunction(value map[string]interface{}, fns IntrinsicFunctionSignatures) map[string]interface{} {
	converted := make(map[string]interface{}, len(value))
	for id, prop := range value {
		converted[id] = convertAnyIntrinsicFunctions(prop, fns)
	}
	return converted
}

func convertArrayToIntrinsicFunction(value []interface{}, fns IntrinsicFunctionSignatures) []interface{} {
	arr := make([]interface{}, len(value))
	for i, v := range value {
		arr[i] = convertAnyIntrinsicFunctions(v, fns)
	}
	return arr
}

func convertToIntrinsicFunction(key IntrinsicFunctionSignature, value interface{}) (IntrinsicFunction, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(key)]; found {
			return IntrinsicFunction{key, convertMapToIntrinsicFunction(m, AllIntrinsicFunctions)}, true
		}
	}

	return IntrinsicFunction{key, map[string]interface{}{}}, false
}
