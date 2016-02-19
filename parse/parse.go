package parse

import "encoding/json"

func ParseTemplateJSON(data []byte, forgiving bool) (*Template, error) {
	var temp Template

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	return &temp, nil
}
