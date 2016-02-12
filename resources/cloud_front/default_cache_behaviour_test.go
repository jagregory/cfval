package cloud_front

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestAllowedMethodsFixedArrays(t *testing.T) {
	testCFDistribution := func(allowedMethods []interface{}) schema.TemplateResource {
		return schema.TemplateResource{
			Definition: Distribution(),
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

	if _, errs := testCFDistribution([]interface{}{"HEAD", "GET"}).Validate(ctx); errs != nil {
		t.Error("Should pass with expected array", errs)
	}

	if _, errs := testCFDistribution([]interface{}{"GET", "HEAD"}).Validate(ctx); errs != nil {
		t.Error("Should pass with expected array in different order", errs)
	}

	if _, errs := testCFDistribution([]interface{}{"DELETE", "GET", "HEAD"}).Validate(ctx); errs == nil {
		t.Error("Should fail with random subset")
	}

	if _, errs := testCFDistribution([]interface{}{"GET", "HEAD", "somethingElse"}).Validate(ctx); errs == nil {
		t.Error("Should fail with unexpected item")
	}
}
