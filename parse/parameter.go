package parse

type Parameter struct {
	Type string
}

// func (p *Parameter) UnmarshalJSON(b []byte) (err error) {
// 	temp := struct {
// 		Type string
// 	}{}
//
// 	if err = json.Unmarshal(b, &temp); err != nil {
// 		panic("Unexpected type unmarshalling parameter")
// 	}
//
// 	if s, found := parameterTypeSchemas[temp.Type]; found {
// 		p.Schema = s
// 		return nil
// 	}
//
// 	return fmt.Errorf("Unexpected type for Parameter %s", temp.Type)
// }
