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
			if ref, ok := convertToRef(value); ok {
				resource.properties[property] = ref
			}
		}
	}

	for name, output := range temp.Outputs {
		if ref, ok := convertToRef(output.Value); ok {
			// TODO: This is a nasty way of modifying the struct
			output.Value = ref
			temp.Outputs[name] = output
		}
	}

	return &temp, nil
}

func convertToRef(value interface{}) (Ref, bool) {
	if m, ok := value.(map[string]interface{}); ok {
		if _, found := m["Ref"]; found {
			return Ref{m}, true
		}
	}

	return Ref{}, false
}

type Ref struct {
	UnderlyingMap map[string]interface{}
}
