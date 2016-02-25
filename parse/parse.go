package parse

import "encoding/json"

func ParseTemplateJSON(data []byte) (*Template, error) {
	var temp Template

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	for _, resource := range temp.Resources {
		for property, value := range resource.properties {
			if builtin, ok := convertToBuiltin(value); ok {
				resource.properties[property] = builtin
			}
		}
	}

	for name, output := range temp.Outputs {
		if builtin, ok := convertToBuiltin(output.Value); ok {
			// TODO: This is a nasty way of modifying the struct
			output.Value = builtin
			temp.Outputs[name] = output
		}
	}

	return &temp, nil
}

func convertToBuiltin(value interface{}) (interface{}, bool) {
	if ref, ok := convertToRef(value); ok {
		return ref, ok
	} else if findInMap, ok := convertToFindInMap(value); ok {
		return findInMap, ok
	} else if join, ok := convertToJoin(value); ok {
		return join, ok
	}

	return nil, false
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

type Ref struct {
	UnderlyingMap map[string]interface{}
}

type FindInMap struct {
	UnderlyingMap map[string]interface{}
}

type Join struct {
	UnderlyingMap map[string]interface{}
}
