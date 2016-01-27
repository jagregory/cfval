package main

import (
	"fmt"
	"encoding/json"
)

type Fn interface {
	Validate(t Template, context []string) (bool, []Failure)
}

func UnmarshalFns(data []byte) (Fn,error) {
	var fns struct {
		Ref *string `json:",omitempty"`
		Join []json.RawMessage `json:"Fn::Join,omitempty"`
		GetAtt []string `json:"Fn::GetAtt,omitempty"`
	}

	if err := json.Unmarshal(data, &fns); err != nil {
		return nil, err
	}

	if fns.Ref != nil {
		return Ref{Target: *fns.Ref}, nil
	}

	if fns.Join != nil {
		var delim string
		var items []StringOrBuiltinFns
		json.Unmarshal(fns.Join[0], &delim)
		json.Unmarshal(fns.Join[1], &items)
		return Join{
			Delim: delim,
			Items: items,
		}, nil
	}

	if fns.GetAtt != nil {
		return GetAtt{
			Attribute: fns.GetAtt[1],
			Resource: fns.GetAtt[0],
		}, nil
	}

	return nil, fmt.Errorf("Unrecognised value: %s", data)
}

type StringOrBuiltinFns struct {
	Value *string
	Fn    Fn
}

func (sob StringOrBuiltinFns) Validate(t Template, context []string) (bool, []Failure) {
	if sob.Fn != nil {
		return sob.Fn.Validate(t, context)
	}

	return true, nil
}

func (sob *StringOrBuiltinFns) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		sob.Value = &s
		return nil
	}

	fns, err := UnmarshalFns(data)
	if err != nil {
		return err
	}

	sob.Fn = fns
	return nil
}

type BoolOrBuiltinFns struct {
	Value *bool
	Fn    Fn
}

func (sob BoolOrBuiltinFns) Validate(t Template, context []string) (bool, []Failure) {
	if sob.Fn != nil {
		return sob.Fn.Validate(t, context)
	}

	return true, nil
}

func (sob *BoolOrBuiltinFns) UnmarshalJSON(data []byte) error {
	var s bool
	if err := json.Unmarshal(data, &s); err == nil {
		sob.Value = &s
		return nil
	}

	fns, err := UnmarshalFns(data)
	if err != nil {
		return err
	}

	sob.Fn = fns
	return nil
}

type Cidr struct {
	StringOrBuiltinFns
}

type VpcId struct {
	StringOrBuiltinFns
}

type ResourceTag struct {
	Key   StringOrBuiltinFns
	Value StringOrBuiltinFns
}

func (tag ResourceTag) Validate(t Template, context []string) (bool, []Failure) {
	errors := make([]Failure, 0, 10)

	if ok, errs := tag.Key.Validate(t, context); !ok {
		errors = append(errors, errs...)
	}

	if ok, errs := tag.Value.Validate(t, context); !ok {
		errors = append(errors, errs...)
	}

	return len(errors) == 0, errors
}

type ResourceTags []ResourceTag

func (ts ResourceTags) Validate(t Template, context []string) (bool, []Failure) {
	errors := make([]Failure, 0, len(ts)*2)

	for _, tag := range ts {
		if ok, errs := tag.Validate(t, context); !ok {
			errors = append(errors, errs...)
		}
	}

	return len(errors) == 0, errors
}
