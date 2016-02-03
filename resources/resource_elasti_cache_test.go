package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestAutomaticFailoverEnabled(t *testing.T) {
	template := &schema.Template{}
	context := []string{}

	badVersion := schema.TemplateResource{
		Template:   template,
		Definition: ReplicationGroup(),
		Properties: map[string]interface{}{
			"EngineVersion": "2.7",
			"CacheNodeType": "cache.m3.medium",
		},
	}

	badNodeTypeT1 := schema.TemplateResource{
		Template:   template,
		Definition: ReplicationGroup(),
		Properties: map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.t1.micro",
		},
	}

	badNodeTypeT2 := schema.TemplateResource{
		Template:   template,
		Definition: ReplicationGroup(),
		Properties: map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.t2.micro",
		},
	}

	good := schema.TemplateResource{
		Template:   template,
		Definition: ReplicationGroup(),
		Properties: map[string]interface{}{
			"EngineVersion": "2.8",
			"CacheNodeType": "cache.m3.medium",
		},
	}

	if ok, _ := automaticFailoverEnabled(true, badVersion, context); ok {
		t.Error("Should fail if has engine less than 2.8")
	}

	if ok, _ := automaticFailoverEnabled(true, badNodeTypeT1, context); ok {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if ok, _ := automaticFailoverEnabled(true, badNodeTypeT2, context); ok {
		t.Error("Should fail if has node type of t1 or t2")
	}

	if ok, _ := automaticFailoverEnabled(true, good, context); !ok {
		t.Error("Should pass if engine is 2.8 or above and node type isn't t1 or t2")
	}
}
