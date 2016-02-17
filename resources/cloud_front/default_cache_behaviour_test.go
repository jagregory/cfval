package cloud_front

import (
	"testing"

	"github.com/jagregory/cfval/parse"
	"github.com/jagregory/cfval/schema"
)

func TestAllowedMethodsFixedArrays(t *testing.T) {
	res := Distribution()

	definitions := schema.NewResourceDefinitions(map[string]func() schema.Resource{
		"TestResource": func() schema.Resource {
			return res
		},
	})

	testCFDistribution := func(allowedMethods []interface{}) parse.TemplateResource {
		return parse.TemplateResource{
			Type: "TestResource",
			Properties: map[string]interface{}{
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
			},
		}
	}

	ctx := []string{}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"HEAD", "GET"}), definitions, ctx); errs != nil {
		t.Error("Should pass with expected array", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD"}), definitions, ctx); errs != nil {
		t.Error("Should pass with expected array in different order", errs)
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"DELETE", "GET", "HEAD"}), definitions, ctx); errs == nil {
		t.Error("Should fail with random subset")
	}

	if _, errs := res.Validate(testCFDistribution([]interface{}{"GET", "HEAD", "somethingElse"}), definitions, ctx); errs == nil {
		t.Error("Should fail with unexpected item")
	}
}
