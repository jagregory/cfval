package elasti_cache

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAutomaticFailoverEnabled(t *testing.T) {
	template := &parse.Template{}
	res := ReplicationGroup
	ctx := schema.Context{
		Definitions: schema.NewResourceDefinitions(map[string]schema.Resource{
			"TestResource": res,
		}),
		Path:     []string{},
		Template: template,
	}

	badVersion := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"EngineVersion": "2.7",
			"CacheNodeType": "cache.m3.medium",
		}),
		res,
	}

	badNodeTypeT1 := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.t1.micro",
		}),
		res,
	}

	badNodeTypeT2 := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.t2.micro",
		}),
		res,
	}

	good := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.m3.medium",
		}),
		res,
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badVersion, ctx); errs == nil {
		t.Error("Should fail if has engine less than 2.8")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT1, ctx); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT2, ctx); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, good, ctx); errs != nil {
		t.Error("Should pass if engine is 2.8 or above and node type isn't t1 or t2")
	}
}
