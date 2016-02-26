package parse

import "encoding/json"

type Condition struct {
	Fn interface{}
}

type tempCondition Condition

func (d *Condition) UnmarshalJSON(b []byte) error {
	var tmp interface{}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	d.Fn = convertAnyIntrinsicFunctions(tmp)

	return nil
}
