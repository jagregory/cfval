package parse

import "encoding/json"

func ParseTemplateJSON(data []byte) (*Template, error) {
	var temp Template

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	return &temp, nil
}

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

func convertToRef(value interface{}) (Ref, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m["Ref"]; found {
			return Ref{m}, true
		}
	}

	return Ref{}, false
}

func convertToFindInMap(value interface{}) (FindInMap, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m["Fn::FindInMap"]; found {
			return FindInMap{m}, true
		}
	}

	return FindInMap{}, false
}

func convertToJoin(value interface{}) (Join, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m["Fn::Join"]; found {
			return Join{m}, true
		}
	}

	return Join{}, false
}

func convertToGetAtt(value interface{}) (GetAtt, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m["Fn::GetAtt"]; found {
			return GetAtt{m}, true
		}
	}

	return GetAtt{}, false
}

type Ref struct {
	UnderlyingMap map[string]interface{}
}

type FindInMap struct {
	UnderlyingMap map[string]interface{}
}

type Join struct {
	UnderlyingMap map[string]interface{}
}

type GetAtt struct {
	UnderlyingMap map[string]interface{}
}
