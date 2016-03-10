package schema

import (
	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-findinmap.html
func validateFindInMap(builtin parse.IntrinsicFunction, ctx PropertyContext) reporting.Reports {
	if errs := validateIntrinsicFunctionBasicCriteria(parse.FnFindInMap, builtin, ctx); errs != nil {
		return errs
	}

	value := builtin.UnderlyingMap[string(parse.FnFindInMap)]
	args, ok := value.([]interface{})
	if !ok {
		return reporting.Reports{reporting.NewFailure(ctx, "Invalid type for \"Fn::FindInMap\" key: %T", value)}
	}

	if len(args) != 3 {
		return reporting.Reports{reporting.NewFailure(ctx, "Options has wrong number of items, expected: 3, actual: %d", len(args))}
	}

	reports := make(reporting.Reports, 0, 10)

	mapName := args[0]
	mapNameCtx := PropertyContextAdd(NewPropertyContext(ctx, Schema{Type: ValueString}), "MapName")
	if errs := validateFindInMapMapName(builtin, mapName, mapNameCtx); errs != nil {
		reports = append(reports, errs...)
	}

	topLevelKey := args[1]
	topLevelKeyCtx := PropertyContextAdd(NewPropertyContext(ctx, Schema{Type: ValueString}), "TopLevelKey")
	if errs := validateFindInMapTopLevelKey(builtin, mapName, topLevelKey, topLevelKeyCtx); errs != nil {
		reports = append(reports, errs...)
	}

	secondLevelKey := args[2]
	secondLevelKeyCtx := PropertyContextAdd(NewPropertyContext(ctx, Schema{Type: ValueString}), "SecondLevelKey")
	if errs := validateFindInMapSecondLevelKey(builtin, mapName, topLevelKey, secondLevelKey, secondLevelKeyCtx); errs != nil {
		reports = append(reports, errs...)
	}

	return reporting.Safe(reports)
}

func validateFindInMapMapName(builtin parse.IntrinsicFunction, mapName interface{}, ctx PropertyContext) reporting.Reports {
	if mapName == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := mapName.(type) {
	case string:
		// TODO: validate actual map exists
		return nil
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid MapName: %v", mapName)}
}

func validateFindInMapTopLevelKey(builtin parse.IntrinsicFunction, mapName, topLevelKey interface{}, ctx PropertyContext) reporting.Reports {
	if topLevelKey == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := topLevelKey.(type) {
	case string:
		// TODO: validate actual map top level key exists, if possible
		return nil
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid TopLevelKey: %v", topLevelKey)}
}

func validateFindInMapSecondLevelKey(builtin parse.IntrinsicFunction, mapName, topLevelKey, secondLevelKey interface{}, ctx PropertyContext) reporting.Reports {
	if secondLevelKey == nil {
		return reporting.Reports{reporting.NewFailure(ctx, "Cannot be null")}
	}

	switch t := secondLevelKey.(type) {
	case string:
		// TODO: validate actual map second level key exists, if possible
		return nil
	case parse.IntrinsicFunction:
		_, errs := ValidateIntrinsicFunctions(t, ctx, SupportedFunctions{
			parse.FnFindInMap: true,
			parse.FnRef:       true,
		})
		return errs
	}

	return reporting.Reports{reporting.NewFailure(ctx, "Invalid SecondLevelKey: %v", secondLevelKey)}
}
