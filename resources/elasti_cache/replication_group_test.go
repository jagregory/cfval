package elasti_cache

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAutomaticFailoverEnabled(t *testing.T) {
	template := &parse.Template{}
	res := ReplicationGroup
	ctx := schema.NewInitialContext(template, schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	}), schema.ValidationOptions{})

	badVersionCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource("TestResource", map[string]interface{}{
				"EngineVersion": "2.7",
				"CacheNodeType": "cache.m3.medium",
			}),
			res,
		}),
		schema.Schema{})

	badNodeTypeT1Ctx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource("TestResource", map[string]interface{}{
				"EngineVersion": "2.8",
				"CacheNodeType": "cache.t1.micro",
			}),
			res,
		}),
		schema.Schema{})

	badNodeTypeT2Ctx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource("TestResource", map[string]interface{}{
				"EngineVersion": "2.8",
				"CacheNodeType": "cache.t2.micro",
			}),
			res,
		}),
		schema.Schema{})

	goodCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource("TestResource", map[string]interface{}{
				"EngineVersion": "2.8",
				"CacheNodeType": "cache.m3.medium",
			}),
			res,
		}),
		schema.Schema{})

	if _, errs := automaticFailoverEnabledValidation(true, badVersionCtx); errs == nil {
		t.Error("Should fail if has engine less than 2.8")
	}

	if _, errs := automaticFailoverEnabledValidation(true, badNodeTypeT1Ctx); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(true, badNodeTypeT2Ctx); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(true, goodCtx); errs != nil {
		t.Error("Should pass if engine is 2.8 or above and node type isn't t1 or t2")
	}
}
