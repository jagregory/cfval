package main

import "encoding/json"
import (
	"github.com/jagregory/cfval/resources"
	"github.com/jagregory/cfval/schema"
)

type resourceCtor func() schema.Resource

var typeCtors = map[string]resourceCtor{
	"AWS::AutoScaling::AutoScalingGroup":           resources.AutoScalingGroup,
	"AWS::AutoScaling::LaunchConfiguration":        resources.LaunchConfiguration,
	"AWS::CloudFront::Distribution":                resources.Distribution,
	"AWS::CloudWatch::Alarm":                       resources.Alarm,
	"AWS::RDS::DBInstance":                         resources.DBInstance,
	"AWS::RDS::DBSubnetGroup":                      resources.DBSubnetGroup,
	"AWS::EC2::EIP":                                resources.Eip,
	"AWS::EC2::Instance":                           resources.Instance,
	"AWS::EC2::InternetGateway":                    resources.InternetGateway,
	"AWS::EC2::Route":                              resources.Route,
	"AWS::EC2::RouteTable":                         resources.RouteTable,
	"AWS::EC2::SecurityGroup":                      resources.SecurityGroup,
	"AWS::EC2::SecurityGroupIngress":               resources.SecurityGroupIngress,
	"AWS::EC2::Subnet":                             resources.Subnet,
	"AWS::EC2::SubnetRouteTableAssociation":        resources.SubnetRouteTableAssociation,
	"AWS::EC2::VPCGatewayAttachment":               resources.VpcGatewayAttachment,
	"AWS::ElastiCache::CacheCluster":               resources.CacheCluster,
	"AWS::ElastiCache::ReplicationGroup":           resources.ReplicationGroup,
	"AWS::ElastiCache::SubnetGroup":                resources.SubnetGroup,
	"AWS::ElasticBeanstalk::Application":           resources.Application,
	"AWS::ElasticBeanstalk::ApplicationVersion":    resources.ApplicationVersion,
	"AWS::ElasticBeanstalk::ConfigurationTemplate": resources.ConfigurationTemplate,
	"AWS::ElasticBeanstalk::Environment":           resources.Environment,
	"AWS::ElasticLoadBalancing::LoadBalancer":      resources.LoadBalancer,
	"AWS::IAM::InstanceProfile":                    resources.InstanceProfile,
	"AWS::IAM::Policy":                             resources.Policy,
	"AWS::IAM::Role":                               resources.Role,
	"AWS::Route53::RecordSet":                      resources.RecordSet,
	"AWS::S3::Bucket":                              resources.Bucket,
	"AWS::SNS::Topic":                              resources.Topic,
}

func parseTemplateJSON(data []byte, forgiving bool) (*schema.Template, error) {
	var temp struct {
		Parameters map[string]schema.Parameter
		Outputs    map[string]schema.Output
		Resources  map[string]struct {
			Type       string
			Properties map[string]interface{}
			Metadata   map[string]interface{}
		}
	}

	err := json.Unmarshal(data, &temp)

	if err != nil {
		return nil, err
	}

	template := &schema.Template{
		Resources: make(map[string]schema.TemplateResource),
	}
	template.Parameters = temp.Parameters
	template.Outputs = temp.Outputs

	for key, rawResource := range temp.Resources {
		if ctor, ok := typeCtors[rawResource.Type]; ok {
			template.Resources[key] = schema.TemplateResource{
				Template:   template,
				Definition: ctor(),
				Properties: rawResource.Properties,
				Metadata:   rawResource.Metadata,
			}
		} else if !forgiving {
			template.Resources[key] = schema.NewUnrecognisedResource(template, rawResource.Type)
		}
	}

	return template, nil
}
