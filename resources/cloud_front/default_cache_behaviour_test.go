package cloud_front

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAllowedMethodsFixedArrays(t *testing.T) {
	res := Distribution

	definitions := schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	})

	template := &parse.Template{}

	testCFDistribution := func(allowedMethods []interface{}) schema.ResourceWithDefinition {
		return schema.ResourceWithDefinition{
			parse.NewTemplateResource(template, "TestResource", map[string]interface{}{
				"DistributionConfig": map[string]interface{}{
					"Enabled": true,
					"DefaultCacheBehavior": map[string]interface{}{
						"AllowedMethods":       allowedMethods,
						"TargetOriginId":       "test",
						"ViewerProtocolPolicy": "test",
						"ForwardedValues": map[string]interface{}{
							"QueryString": true,
						},
					},
					"Origins": []interface{}{
						map[string]interface{}{
							"Id":         "test",
							"DomainName": "test",
							"CustomOriginConfig": map[string]interface{}{
								"OriginProtocolPolicy": "test",
							},
						},
					},
				},
			}),
			res,
		}
	}

	ctx := []string{}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"HEAD", "GET"}), template, definitions, ctx); errs != nil {
		t.Error("Should pass with expected array", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD"}), template, definitions, ctx); errs != nil {
		t.Error("Should pass with expected array in different order", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"DELETE", "GET", "HEAD"}), template, definitions, ctx); errs == nil {
		t.Error("Should fail with random subset")
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD", "somethingElse"}), template, definitions, ctx); errs == nil {
		t.Error("Should fail with unexpected item")
	}
}
