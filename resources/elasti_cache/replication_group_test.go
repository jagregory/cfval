package elasti_cache

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAutomaticFailoverEnabled(t *testing.T) {
	template := &parse.Template{}
	context := []string{}

	res := ReplicationGroup

	definitions := schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	})

	badVersion := parse.NewTemplateResource(template)
	badVersion.Type = "TestResource"
	badVersion.Properties = map[string]interface{}{
		"EngineVersion": "2.7",
		"CacheNodeType": "cache.m3.medium",
	}

	badNodeTypeT1 := parse.NewTemplateResource(template)
	badNodeTypeT1.Type = "TestResource"
	badNodeTypeT1.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.t1.micro",
	}

	badNodeTypeT2 := parse.NewTemplateResource(template)
	badNodeTypeT2.Type = "TestResource"
	badNodeTypeT2.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.t2.micro",
	}

	good := parse.NewTemplateResource(template)
	good.Type = "TestResource"
	good.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.m3.medium",
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badVersion, definitions, context); errs == nil {
		t.Error("Should fail if has engine less than 2.8")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT1, definitions, context); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT2, definitions, context); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, good, definitions, context); errs != nil {
		t.Error("Should pass if engine is 2.8 or above and node type isn't t1 or t2")
	}
}
