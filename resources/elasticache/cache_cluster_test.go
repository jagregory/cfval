package elasticache

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestAZModeValidate(t *testing.T) {
	prop := schema.Schema{}
	ctx := []string{}

	singleAZ := schema.TemplateResource{
		Properties: map[string]interface{}{
			"PreferredAvailabilityZones": []interface{}{"one"},
		},
	}

	multiAZ := schema.TemplateResource{
		Properties: map[string]interface{}{
			"PreferredAvailabilityZones": []interface{}{"one", "two"},
		},
	}

	if _, errs := azModeValidate(prop, "cross-az", singleAZ, ctx); errs == nil {
		t.Error("Should fail if cross-az with single availability zone", errs)
	}

	if _, errs := azModeValidate(prop, "cross-az", multiAZ, ctx); errs != nil {
		t.Error("Should pass if cross-az with multiple availability zones", errs)
	}
}

func TestNumCacheNodesValidate(t *testing.T) {
	prop := schema.Schema{}
	ctx := []string{}

	redis := schema.TemplateResource{
		Properties: map[string]interface{}{
			"Engine": "redis",
		},
	}

	memcached := schema.TemplateResource{
		Properties: map[string]interface{}{
			"Engine": "memcached",
		},
	}

	if _, errs := numCacheNodesValidate(prop, float64(1), redis, ctx); errs != nil {
		t.Error("Should pass with 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(2), redis, ctx); errs == nil {
		t.Error("Should fail with more than 1 redis node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(1), memcached, ctx); errs != nil {
		t.Error("Should pass with 1 memcached node", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(20), memcached, ctx); errs != nil {
		t.Error("Should pass with 20 memcached nodes", errs)
	}

	if _, errs := numCacheNodesValidate(prop, float64(21), memcached, ctx); errs == nil {
		t.Error("Should fail with 21 memcached nodes", errs)
	}
}
