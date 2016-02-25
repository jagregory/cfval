package schema

import (
	"strconv"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/reporting"
)

func validateJoin(join parse.Join, ctx PropertyContext) (reporting.ValidateResult, reporting.Reports) {
	joinValue, found := join.UnderlyingMap["Fn::Join"]
	if !found || joinValue == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Missing \"Fn::Join\" key")}
	}

	// TODO: this will fail with { "Fn::Join": { "Fn::GetAZs": "" }} and such
	items, ok := joinValue.([]interface{})
	if !ok || items == nil {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Invalid \"Fn::Join\" key: %s", items)}
	}

	if len(join.UnderlyingMap) > 1 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Unexpected extra keys: %s", keysExcept(join.UnderlyingMap, "Fn::Join"))}
	}

	if len(items) != 2 {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join has incorrect number of arguments (expected: 2, actual: %d)", len(items))}
	}

	if _, ok := items[0].(string); !ok {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join '%s' is not a valid delimiter", items[0])}
	}

	parts, ok := items[1].([]interface{})
	if !ok {
		return reporting.ValidateAbort, reporting.Reports{reporting.NewFailure(ctx, "Join items are not valid: %s", items[1])}
	}

	joinItemContext := NewPropertyContext(ctx, Schema{Type: ValueString})
	failures := make(reporting.Reports, 0, len(parts))
	for i, part := range parts {
		if _, errs := ValueString.Validate(part, PropertyContextAdd(joinItemContext, "1", strconv.Itoa(i))); errs != nil {
			failures = append(failures, errs...)
		}
	}

	if len(failures) == 0 {
		return reporting.ValidateAbort, nil
	}

	return reporting.ValidateAbort, failures
}
