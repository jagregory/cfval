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
	"AWS::AutoScaling::LifecycleHook":              resources.LifecycleHook,
	"AWS::AutoScaling::ScalingPolicy":              resources.ScalingPolicy,
	"AWS::AutoScaling::ScheduledAction":            resources.ScheduledAction,
	"AWS::CloudFormation::Authentication":          resources.Authentication,
	"AWS::CloudFormation::CustomResource":          resources.CustomResource,
	"AWS::CloudFormation::Init":                    resources.Init,
	"AWS::CloudFormation::Interface":               resources.Interface,
	"AWS::CloudFormation::Stack":                   resources.Stack,
	"AWS::CloudFormation::WaitCondition":           resources.WaitCondition,
	"AWS::CloudFormation::WaitConditionHandle":     resources.WaitConditionHandle,
	"AWS::CloudFront::Distribution":                resources.Distribution,
	"AWS::CloudTrail::Trail":                       resources.Trail,
	"AWS::CloudWatch::Alarm":                       resources.Alarm,
	"AWS::CodeDeploy::Application":                 resources.CodeDeployApplication,
	"AWS::CodeDeploy::DeploymentConfig":            resources.DeploymentConfig,
	"AWS::CodeDeploy::DeploymentGroup":             resources.DeploymentGroup,
	"AWS::CodePipeline::CustomActionType":          resources.CustomActionType,
	"AWS::CodePipeline::Pipeline":                  resources.Pipeline,
	"AWS::Config::ConfigRule":                      resources.ConfigRule,
	"AWS::Config::ConfigurationRecorder":           resources.ConfigurationRecorder,
	"AWS::Config::DeliveryChannel":                 resources.DeliveryChannel,
	"AWS::DataPipeline::Pipeline":                  resources.DataPipelinePipeline,
	"AWS::DirectoryService::MicrosoftAD":           resources.MicrosoftAD,
	"AWS::DirectoryService::SimpleAD":              resources.SimpleAD,
	"AWS::DynamoDB::Table":                         resources.Table,
	"AWS::EC2::CustomerGateway":                    resources.CustomerGateway,
	"AWS::EC2::DHCPOptions":                        resources.DHCPOptions,
	"AWS::EC2::EIP":                                resources.Eip,
	"AWS::EC2::EIPAssociation":                     resources.EIPAssociation,
	"AWS::EC2::Instance":                           resources.Instance,
	"AWS::EC2::InternetGateway":                    resources.InternetGateway,
	"AWS::EC2::NetworkAcl":                         resources.NetworkAcl,
	"AWS::EC2::NetworkAclEntry":                    resources.NetworkAclEntry,
	"AWS::EC2::NetworkInterface":                   resources.NetworkInterface,
	"AWS::EC2::NetworkInterfaceAttachment":         resources.NetworkInterfaceAttachment,
	"AWS::EC2::PlacementGroup":                     resources.PlacementGroup,
	"AWS::EC2::Route":                              resources.Route,
	"AWS::EC2::RouteTable":                         resources.RouteTable,
	"AWS::EC2::SecurityGroup":                      resources.SecurityGroup,
	"AWS::EC2::SecurityGroupEgress":                resources.SecurityGroupEgress,
	"AWS::EC2::SecurityGroupIngress":               resources.SecurityGroupIngress,
	"AWS::EC2::SpotFleet":                          resources.SpotFleet,
	"AWS::EC2::Subnet":                             resources.Subnet,
	"AWS::EC2::SubnetNetworkAclAssociation":        resources.SubnetNetworkAclAssociation,
	"AWS::EC2::SubnetRouteTableAssociation":        resources.SubnetRouteTableAssociation,
	"AWS::EC2::Volume":                             resources.Volume,
	"AWS::EC2::VolumeAttachment":                   resources.VolumeAttachment,
	"AWS::EC2::VPC":                                resources.VPC,
	"AWS::EC2::VPCDHCPOptionsAssociation":          resources.VPCDHCPOptionsAssociation,
	"AWS::EC2::VPCEndpoint":                        resources.VPCEndpoint,
	"AWS::EC2::VPCGatewayAttachment":               resources.VPCGatewayAttachment,
	"AWS::EC2::VPCPeeringConnection":               resources.VPCPeeringConnection,
	"AWS::EC2::VPNConnection":                      resources.VPNConnection,
	"AWS::EC2::VPNConnectionRoute":                 resources.VPNConnectionRoute,
	"AWS::EC2::VPNGateway":                         resources.VPNGateway,
	"AWS::EC2::VPNGatewayRoutePropagation":         resources.VPNGatewayRoutePropagation,
	"AWS::ECS::Cluster":                            resources.Cluster,
	"AWS::ECS::Service":                            resources.Service,
	"AWS::ECS::TaskDefinition":                     resources.TaskDefinition,
	"AWS::EFS::FileSystem":                         resources.FileSystem,
	"AWS::EFS::MountTarget":                        resources.MountTarget,
	"AWS::ElastiCache::CacheCluster":               resources.CacheCluster,
	"AWS::ElastiCache::ParameterGroup":             resources.ParameterGroup,
	"AWS::ElastiCache::ReplicationGroup":           resources.ReplicationGroup,
	"AWS::ElastiCache::SecurityGroup":              resources.ElastiCacheSecurityGroup,
	"AWS::ElastiCache::SecurityGroupIngress":       resources.ElastiCacheSecurityGroupIngress,
	"AWS::ElastiCache::SubnetGroup":                resources.SubnetGroup,
	"AWS::ElasticBeanstalk::Application":           resources.Application,
	"AWS::ElasticBeanstalk::ApplicationVersion":    resources.ApplicationVersion,
	"AWS::ElasticBeanstalk::ConfigurationTemplate": resources.ConfigurationTemplate,
	"AWS::ElasticBeanstalk::Environment":           resources.Environment,
	"AWS::ElasticLoadBalancing::LoadBalancer":      resources.LoadBalancer,
	"AWS::IAM::AccessKey":                          resources.AccessKey,
	"AWS::IAM::Group":                              resources.Group,
	"AWS::IAM::InstanceProfile":                    resources.InstanceProfile,
	"AWS::IAM::ManagedPolicy":                      resources.ManagedPolicy,
	"AWS::IAM::Policy":                             resources.Policy,
	"AWS::IAM::Role":                               resources.Role,
	"AWS::IAM::User":                               resources.User,
	"AWS::IAM::UserToGroupAddition":                resources.UserToGroupAddition,
	"AWS::Kinesis::Stream":                         resources.Stream,
	"AWS::KMS::Key":                                resources.Key,
	"AWS::Lambda::EventSourceMapping":              resources.EventSourceMapping,
	"AWS::Lambda::Function":                        resources.Function,
	"AWS::Lambda::Permission":                      resources.Permission,
	"AWS::Logs::Destination":                       resources.Destination,
	"AWS::Logs::LogGroup":                          resources.LogGroup,
	"AWS::Logs::LogStream":                         resources.LogStream,
	"AWS::Logs::MetricFilter":                      resources.MetricFilter,
	"AWS::Logs::SubscriptionFilter":                resources.SubscriptionFilter,
	"AWS::OpsWorks::App":                           resources.App,
	"AWS::OpsWorks::ElasticLoadBalancerAttachment": resources.ElasticLoadBalancerAttachment,
	"AWS::OpsWorks::Instance":                      resources.OpsWorksInstance,
	"AWS::OpsWorks::Layer":                         resources.Layer,
	"AWS::OpsWorks::Stack":                         resources.OpsWorksStack,
	"AWS::RDS::DBCluster":                          resources.DBCluster,
	"AWS::RDS::DBClusterParameterGroup":            resources.DBClusterParameterGroup,
	"AWS::RDS::DBInstance":                         resources.DBInstance,
	"AWS::RDS::DBParameterGroup":                   resources.DBParameterGroup,
	"AWS::RDS::DBSecurityGroup":                    resources.DBSecurityGroup,
	"AWS::RDS::DBSecurityGroupIngress":             resources.DBSecurityGroupIngress,
	"AWS::RDS::DBSubnetGroup":                      resources.DBSubnetGroup,
	"AWS::RDS::EventSubscription":                  resources.EventSubscription,
	"AWS::RDS::OptionGroup":                        resources.OptionGroup,
	"AWS::Redshift::Cluster":                       resources.RedshiftCluster,
	"AWS::Redshift::ClusterParameterGroup":         resources.ClusterParameterGroup,
	"AWS::Redshift::ClusterSecurityGroup":          resources.ClusterSecurityGroup,
	"AWS::Redshift::ClusterSecurityGroupIngress":   resources.ClusterSecurityGroupIngress,
	"AWS::Redshift::ClusterSubnetGroup":            resources.ClusterSubnetGroup,
	"AWS::Route53::HealthCheck":                    resources.HealthCheck,
	"AWS::Route53::HostedZone":                     resources.HostedZone,
	"AWS::Route53::RecordSet":                      resources.RecordSet,
	"AWS::Route53::RecordSetGroup":                 resources.RecordSetGroup,
	"AWS::S3::Bucket":                              resources.Bucket,
	"AWS::S3::BucketPolicy":                        resources.BucketPolicy,
	"AWS::SDB::Domain":                             resources.Domain,
	"AWS::SNS::Topic":                              resources.Topic,
	"AWS::SNS::TopicPolicy":                        resources.TopicPolicy,
	"AWS::SQS::Queue":                              resources.Queue,
	"AWS::SQS::QueuePolicy":                        resources.QueuePolicy,
	"AWS::SSM::Document":                           resources.Document,
	"AWS::WAF::ByteMatchSet":                       resources.ByteMatchSet,
	"AWS::WAF::IPSet":                              resources.IPSet,
	"AWS::WAF::Rule":                               resources.Rule,
	"AWS::WAF::SqlInjectionMatchSet":               resources.SqlInjectionMatchSet,
	"AWS::WAF::WebACL":                             resources.WebACL,
	"AWS::WorkSpaces::Workspace":                   resources.Workspace,
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
