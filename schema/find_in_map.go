package schema

import "github.com/jagregory/cfval/reporting"

type FindInMap struct {
	args []interface{}
}

func (fim FindInMap) Validate(ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	find := fim.args

	if len(find) != 3 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Options has wrong number of items, expected: 3, actual: %d", len(find))}
	}

	findInMapItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})

	mapName := find[0]
	_, mapNameIsString := mapName.(string)
	if _, errs := ValueString.Validate(mapName, PropertyContextAdd(findInMapItemContext, "0")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString {
		// map name is a string, so we can do some further interrogation
		// TODO: lookup whether MapName is a valid Map
	}

	topLevelKey := find[1]
	_, topLevelKeyIsString := topLevelKey.(string)
	if _, errs := ValueString.Validate(topLevelKey, PropertyContextAdd(findInMapItemContext, "1")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString {
		// TODO: lookup whether topLevelKey is in mapName
	}

	secondLevelKey := find[2]
	_, secondLevelKeyIsString := secondLevelKey.(string)
	if _, errs := ValueString.Validate(secondLevelKey, PropertyContextAdd(findInMapItemContext, "2")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString && secondLevelKeyIsString {
		// TODO: lookup whether secondLevelKeyIsString is in topLevelKey
	}

	return reporting.ValidateAbort, nil
}

func NewFindInMap(args []interface{}) FindInMap {
	return FindInMap{args}
}
