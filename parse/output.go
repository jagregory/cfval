package parse

import "encoding/json"

type Output struct {
	Description, Value interface{}
}

type tempOutput Output

func (d *Output) UnmarshalJSON(b []byte) error {
	var tmp tempOutput

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	d.Value = convertAnyBuiltins(tmp.Value)

	return nil
}
