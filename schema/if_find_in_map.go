package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-findinmap.html
func validateFindInMap(builtin parse.IntrinsicFunction, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	value, found := builtin.UnderlyingMap["Fn::FindInMap"]
	if !found || value == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::FindInMap\" key")}
	}

	args, ok := value.([]interface{})
	if !ok {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::FindInMap\" key: %T", value)}
	}

	if len(builtin.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(builtin.UnderlyingMap, "Fn::FindInMap"))}
	}

	if len(args) != 3 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Options has wrong number of items, expected: 3, actual: %d", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	mapName := args[0]
	if _, errs := validateFindInMapMapName(builtin, mapName, PropertyContextAdd(ctx, "MapName")); errs != nil {
		reports = append(reports, errs...)
	}

	topLevelKey := args[1]
	if _, errs := validateFindInMapTopLevelKey(builtin, mapName, topLevelKey, PropertyContextAdd(ctx, "TopLevelKey")); errs != nil {
		reports = append(reports, errs...)
	}

	secondLevelKey := args[2]
	if _, errs := validateFindInMapSecondLevelKey(builtin, mapName, topLevelKey, secondLevelKey, PropertyContextAdd(ctx, "SecondLevelKey")); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.ValidateAbort, reporting.Safe(reports)
}

func validateFindInMapMapName(builtin parse.IntrinsicFunction, mapName interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if mapName == nil {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := mapName.(type) {
	case string:
		// TODO: validate actual map exists
		return reporting.ValidateOK, nil
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid MapName: %v", mapName)}
}

func validateFindInMapTopLevelKey(builtin parse.IntrinsicFunction, mapName, topLevelKey interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if topLevelKey == nil {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := topLevelKey.(type) {
	case string:
		// TODO: validate actual map top level key exists, if possible
		return reporting.ValidateOK, nil
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid TopLevelKey: %v", topLevelKey)}
}

func validateFindInMapSecondLevelKey(builtin parse.IntrinsicFunction, mapName, topLevelKey, secondLevelKey interface{}, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	if secondLevelKey == nil {
		return reporting.ValidateOK, reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := secondLevelKey.(type) {
	case string:
		// TODO: validate actual map second level key exists, if possible
		return reporting.ValidateOK, nil
	case parse.IntrinsicFunction:
		return ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
	}

	return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid SecondLevelKey: %v", secondLevelKey)}
}
