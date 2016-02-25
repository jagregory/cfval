package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateFindInMap(findInMap parse.FindInMap, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	argsValue, found := findInMap.UnderlyingMap["Fn::FindInMap"]
	if !found || argsValue == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::FindInMap\" key")}
	}

	argsArray, ok := argsValue.([]interface{})
	if !ok {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::FindInMap\" key: %s", argsArray)}
	}

	if len(findInMap.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(findInMap.UnderlyingMap, "Fn::FindInMap"))}
	}

	if len(argsArray) != 3 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Options has wrong number of items, expected: 3, actual: %d", len(argsArray))}
	}

	findInMapItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})

	mapName := argsArray[0]
	_, mapNameIsString := mapName.(string)
	if _, errs := ValueString.Validate(mapName, PropertyContextAdd(findInMapItemContext, "0")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString {
		// map name is a string, so we can do some further interrogation
		// TODO: lookup whether MapName is a valid Map
	}

	topLevelKey := argsArray[1]
	_, topLevelKeyIsString := topLevelKey.(string)
	if _, errs := ValueString.Validate(topLevelKey, PropertyContextAdd(findInMapItemContext, "1")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString {
		// TODO: lookup whether topLevelKey is in mapName
	}

	secondLevelKey := argsArray[2]
	_, secondLevelKeyIsString := secondLevelKey.(string)
	if _, errs := ValueString.Validate(secondLevelKey, PropertyContextAdd(findInMapItemContext, "2")); errs != nil {
		return reporting.ValidateAbort, errs
	}

	if mapNameIsString && topLevelKeyIsString && secondLevelKeyIsString {
		// TODO: lookup whether secondLevelKeyIsString is in topLevelKey
	}

	return reporting.ValidateAbort, nil
}
