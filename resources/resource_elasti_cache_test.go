package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestAutomaticFailoverEnabled(t *testing.T) {
	template := &schema.Template{}
	context := []string{}

	badVersion := schema.NewTemplateResource(template)
	badVersion.Definition = ReplicationGroup()
	badVersion.Properties = map[string]interface{}{
		"EngineVersion": "2.7",
		"CacheNodeType": "cache.m3.medium",
	}

	badNodeTypeT1 := schema.NewTemplateResource(template)
	badNodeTypeT1.Definition = ReplicationGroup()
	badNodeTypeT1.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.t1.micro",
	}

	badNodeTypeT2 := schema.NewTemplateResource(template)
	badNodeTypeT2.Definition = ReplicationGroup()
	badNodeTypeT2.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.t2.micro",
	}

	good := schema.NewTemplateResource(template)
	good.Definition = ReplicationGroup()
	good.Properties = map[string]interface{}{
		"EngineVersion": "2.8",
		"CacheNodeType": "cache.m3.medium",
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badVersion, context); errs == nil {
		t.Error("Should fail if has engine less than 2.8")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT1, context); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, badNodeTypeT2, context); errs == nil {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if _, errs := automaticFailoverEnabledValidation(schema.Schema{}, true, good, context); errs != nil {
		t.Error("Should pass if engine is 2.8 or above and node type isn't t1 or t2")
	}
}
