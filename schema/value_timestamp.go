package schema

var Timestamp = FuncType{
	Description: "Timestamp",

	Fn: RegexpValidate("^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}Z$", "Not a valid Timestamp"),
}
