package elasti_cache

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAZModeValidate(t *testing.T) {
	template := &parse.Template{}
	prop := schema.Schema{}
	ctx := []string{}

	singleAZ := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "", map[string]interface{}{
			"PreferredAvailabilityZones": []interface{}{"one"},
		}),
		schema.Resource{},
	}

	multiAZ := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "", map[string]interface{}{
			"PreferredAvailabilityZones": []interface{}{"one", "two"},
		}),
		schema.Resource{},
	}

	if _, errs := azModeValidate(prop, "cross-az", singleAZ, template, nil, ctx); errs == nil {
		t.Error("Should fail if cross-az with single availability zone", errs)
	}

	if _, errs := azModeValidate(prop, "cross-az", multiAZ, template, nil, ctx); errs != nil {
		t.Error("Should pass if cross-az with multiple availability zones", errs)
	}
}

func TestNumCacheNodesValidate(t *testing.T) {
	template := &parse.Template{}
	prop := schema.Schema{}
	ctx := []string{}

	redis := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "", map[string]interface{}{
			"Engine": "redis",
		}),
		schema.Resource{},
	}

	memcached := schema.ResourceWithDefinition{
		parse.NewTemplateResource(template, "", map[string]interface{}{
			"Engine": "memcached",
		}),
		schema.Resource{},
	}

	if _, errs := numCacheNodesValidate(prop, float64(1), redis, template, nil, ctx); errs != nil {
		t.Error("Should pass with 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(2), redis, template, nil, ctx); errs == nil {
		t.Error("Should fail with more than 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(1), memcached, template, nil, ctx); errs != nil {
		t.Error("Should pass with 1 memcached node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(20), memcached, template, nil, ctx); errs != nil {
		t.Error("Should pass with 20 memcached nodes", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(21), memcached, template, nil, ctx); errs == nil {
		t.Error("Should fail with 21 memcached nodes", errs)
	}
}
