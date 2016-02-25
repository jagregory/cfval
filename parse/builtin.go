package parse

type Builtin struct {
	Key           BuiltinSignature
	UnderlyingMap map[string]interface{}
}

type BuiltinSignature string

const (
	BuiltinAnd       BuiltinSignature = "Fn::And"
	BuiltinBase64                     = "Fn::Base64"
	BuiltinEquals                     = "Fn::Equals"
	BuiltinFindInMap                  = "Fn::FindInMap"
	BuiltinGetAtt                     = "Fn::GetAtt"
	BuiltinGetAZs                     = "Fn::GetAZs"
	BuiltinIf                         = "Fn::If"
	BuiltinJoin                       = "Fn::Join"
	BuiltinNot                        = "Fn::Not"
	BuiltinOr                         = "Fn::Or"
	BuiltinRef                        = "Ref"
	BuiltinSelect                     = "Fn::Select"
)

var allBuiltins = []BuiltinSignature{
	BuiltinAnd,
	BuiltinBase64,
	BuiltinEquals,
	BuiltinFindInMap,
	BuiltinGetAtt,
	BuiltinGetAZs,
	BuiltinIf,
	BuiltinJoin,
	BuiltinNot,
	BuiltinOr,
	BuiltinRef,
	BuiltinSelect,
}

func convertMapToBuiltin(value map[string]interface{}) map[string]interface{} {
	converted := make(map[string]interface{}, len(value))
	for id, prop := range value {
		converted[id] = convertAnyBuiltins(prop)
	}
	return converted
}

func convertArrayToBuiltin(value []interface{}) []interface{} {
	arr := make([]interface{}, len(value))
	for i, v := range value {
		arr[i] = convertAnyBuiltins(v)
	}
	return arr
}

func convertAnyBuiltins(value interface{}) interface{} {
	for _, key := range allBuiltins {
		if builtin, ok := convertToBuiltin(key, value); ok {
			return builtin
		}
	}

	switch t := value.(type) {
	case map[string]interface{}:
		return convertMapToBuiltin(t)
	case []interface{}:
		return convertArrayToBuiltin(t)
	default:
		return value
	}
}

func convertToBuiltin(key BuiltinSignature, value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(key)]; found {
			return Builtin{key, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{key, map[string]interface{}{}}, false
}
