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
)

var AllIntrinsicFunctions = []IntrinsicFunctionSignature{
	FnAnd,
	FnBase64,
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

func convertMapToIntrinsicFunction(value map[string]interface{}) map[string]interface{} {
	converted := make(map[string]interface{}, len(value))
	for id, prop := range value {
		converted[id] = convertAnyIntrinsicFunctions(prop)
	}
	return converted
}

func convertArrayToIntrinsicFunction(value []interface{}) []interface{} {
	arr := make([]interface{}, len(value))
	for i, v := range value {
		arr[i] = convertAnyIntrinsicFunctions(v)
	}
	return arr
}

func convertAnyIntrinsicFunctions(value interface{}) interface{} {
	for _, key := range AllIntrinsicFunctions {
		if fn, ok := convertToIntrinsicFunction(key, value); ok {
			return fn
		}
	}

	switch t := value.(type) {
	case map[string]interface{}:
		return convertMapToIntrinsicFunction(t)
	case []interface{}:
		return convertArrayToIntrinsicFunction(t)
	default:
		return value
	}
}

func convertToIntrinsicFunction(key IntrinsicFunctionSignature, value interface{}) (IntrinsicFunction, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(key)]; found {
			return IntrinsicFunction{key, convertMapToIntrinsicFunction(m)}, true
		}
	}

	return IntrinsicFunction{key, map[string]interface{}{}}, false
}
