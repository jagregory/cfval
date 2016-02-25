package parse

type Builtin struct {
	Key           BuiltinSignature
	UnderlyingMap map[string]interface{}
}

type BuiltinSignature string

const (
	BuiltinBase64    BuiltinSignature = "Fn::Base64"
	BuiltinFindInMap                  = "Fn::FindInMap"
	BuiltinGetAtt                     = "Fn::GetAtt"
	BuiltinJoin                       = "Fn::Join"
	BuiltinRef                        = "Ref"
)

func convertMapToBuiltin(value map[string]interface{}) map[string]interface{} {
	converted := make(map[string]interface{}, len(value))
	for id, prop := range value {
		converted[id] = convertToBuiltin(prop)
	}
	return converted
}

func convertArrayToBuiltin(value []interface{}) []interface{} {
	arr := make([]interface{}, len(value))
	for i, v := range value {
		arr[i] = convertToBuiltin(v)
	}
	return arr
}

func convertToBuiltin(value interface{}) interface{} {
	if ref, ok := convertToRef(value); ok {
		return ref
	} else if findInMap, ok := convertToFindInMap(value); ok {
		return findInMap
	} else if join, ok := convertToJoin(value); ok {
		return join
	} else if getAtt, ok := convertToGetAtt(value); ok {
		return getAtt
	} else if base64, ok := convertToBase64(value); ok {
		return base64
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

func convertToRef(value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(BuiltinRef)]; found {
			return Builtin{BuiltinRef, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{BuiltinRef, map[string]interface{}{}}, false
}

func convertToFindInMap(value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(BuiltinFindInMap)]; found {
			return Builtin{BuiltinFindInMap, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{BuiltinFindInMap, map[string]interface{}{}}, false
}

func convertToJoin(value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(BuiltinJoin)]; found {
			return Builtin{BuiltinJoin, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{BuiltinJoin, map[string]interface{}{}}, false
}

func convertToGetAtt(value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(BuiltinGetAtt)]; found {
			return Builtin{BuiltinGetAtt, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{BuiltinGetAtt, map[string]interface{}{}}, false
}

func convertToBase64(value interface{}) (Builtin, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m[string(BuiltinBase64)]; found {
			return Builtin{BuiltinBase64, convertMapToBuiltin(m)}, true
		}
	}

	return Builtin{BuiltinBase64, map[string]interface{}{}}, false
}
