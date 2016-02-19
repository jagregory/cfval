package cloud_front

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAllowedMethodsFixedArrays(t *testing.T) {
	res := Distribution
	template := &parse.Template{}

	ctx := schema.NewInitialContext(template, schema.NewResourceDefinitions(map[string]schema.Resource{
		"TestResource": res,
	}))

	testCFDistribution := func(allowedMethods []interface{}) schema.ResourceContext {
		return schema.NewResourceContext(ctx, schema.ResourceWithDefinition{
			parse.NewTemplateResource("TestResource", map[string]interface{}{
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
		})
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"HEAD", "GET"})); errs != nil {
		t.Error("Should pass with expected array", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD"})); errs != nil {
		t.Error("Should pass with expected array in different order", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"DELETE", "GET", "HEAD"})); errs == nil {
		t.Error("Should fail with random subset")
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD", "somethingElse"})); errs == nil {
		t.Error("Should fail with unexpected item")
	}
}
