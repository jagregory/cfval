package elasti_cache

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAZModeValidate(t *testing.T) {
	template := &parse.Template{}
	prop := schema.Schema{}
	ctx := schema.NewInitialContext(template, schema.NewResourceDefinitions(nil))

	singleAZCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource(template, "", map[string]interface{}{
				"PreferredAvailabilityZones": []interface{}{"one"},
			}),
			schema.Resource{},
		}),
		prop)

	multiAZCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource(template, "", map[string]interface{}{
				"PreferredAvailabilityZones": []interface{}{"one", "two"},
			}),
			schema.Resource{},
		}),
		prop)

	if _, errs := azModeValidate("cross-az", singleAZCtx); errs == nil {
		t.Error("Should fail if cross-az with single availability zone", errs)
	}

	if _, errs := azModeValidate("cross-az", multiAZCtx); errs != nil {
		t.Error("Should pass if cross-az with multiple availability zones", errs)
	}
}

func TestNumCacheNodesValidate(t *testing.T) {
	template := &parse.Template{}
	prop := schema.Schema{}
	ctx := schema.NewInitialContext(template, schema.NewResourceDefinitions(nil))

	redisCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource(template, "", map[string]interface{}{
				"Engine": "redis",
			}),
			schema.Resource{},
		}),
		prop)

	memcachedCtx := schema.NewPropertyContext(
		schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource(template, "", map[string]interface{}{
				"Engine": "memcached",
			}),
			schema.Resource{},
		}),
		prop)

	if _, errs := numCacheNodesValidate(float64(1), redisCtx); errs != nil {
		t.Error("Should pass with 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(float64(2), redisCtx); errs == nil {
		t.Error("Should fail with more than 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(float64(1), memcachedCtx); errs != nil {
		t.Error("Should pass with 1 memcached node", errs)
	}

	if _, errs := numCacheNodesValidate(float64(20), memcachedCtx); errs != nil {
		t.Error("Should pass with 20 memcached nodes", errs)
	}

	if _, errs := numCacheNodesValidate(float64(21), memcachedCtx); errs == nil {
		t.Error("Should fail with 21 memcached nodes", errs)
	}
}
