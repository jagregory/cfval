package main

import "encoding/json"
import (
	"github.com/jagregory/cfval/resources"
	"github.com/jagregory/cfval/resources/auto_scaling"
	"github.com/jagregory/cfval/resources/cloud_front"
	"github.com/jagregory/cfval/resources/cloud_trail"
	"github.com/jagregory/cfval/resources/cloud_watch"
	"github.com/jagregory/cfval/resources/ec2"
	"github.com/jagregory/cfval/resources/elasti_cache"
	"github.com/jagregory/cfval/resources/elastic_beanstalk"
	"github.com/jagregory/cfval/resources/elastic_load_balancing"
	"github.com/jagregory/cfval/resources/iam"
	"github.com/jagregory/cfval/resources/rds"
	"github.com/jagregory/cfval/resources/route_53"
	"github.com/jagregory/cfval/resources/s3"
	"github.com/jagregory/cfval/resources/sns"
	"github.com/jagregory/cfval/schema"
)

type resourceCtor func() schema.Resource

var typeCtors = map[string]resourceCtor{
	"AWS::AutoScaling::AutoScalingGroup":           auto_scaling.AutoScalingGroup,
	"AWS::AutoScaling::LaunchConfiguration":        auto_scaling.LaunchConfiguration,
	"AWS::AutoScaling::LifecycleHook":              auto_scaling.LifecycleHook,
	"AWS::AutoScaling::ScalingPolicy":              auto_scaling.ScalingPolicy,
	"AWS::AutoScaling::ScheduledAction":            auto_scaling.ScheduledAction,
	"AWS::CloudFormation::Authentication":          resources.Authentication,
	"AWS::CloudFormation::CustomResource":          resources.CustomResource,
	"AWS::CloudFormation::Init":                    resources.Init,
	"AWS::CloudFormation::Interface":               resources.Interface,
	"AWS::CloudFormation::Stack":                   resources.Stack,
	"AWS::CloudFormation::WaitCondition":           resources.WaitCondition,
	"AWS::CloudFormation::WaitConditionHandle":     resources.WaitConditionHandle,
	"AWS::CloudFront::Distribution":                cloud_front.Distribution,
	"AWS::CloudTrail::Trail":                       cloud_trail.Trail,
	"AWS::CloudWatch::Alarm":                       cloud_watch.Alarm,
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
	"AWS::EC2::EIP":                                ec2.Eip,
	"AWS::EC2::EIPAssociation":                     resources.EIPAssociation,
	"AWS::EC2::Instance":                           ec2.Instance,
	"AWS::EC2::InternetGateway":                    ec2.InternetGateway,
	"AWS::EC2::NetworkAcl":                         resources.NetworkAcl,
	"AWS::EC2::NetworkAclEntry":                    resources.NetworkAclEntry,
	"AWS::EC2::NetworkInterface":                   resources.NetworkInterface,
	"AWS::EC2::NetworkInterfaceAttachment":         resources.NetworkInterfaceAttachment,
	"AWS::EC2::PlacementGroup":                     resources.PlacementGroup,
	"AWS::EC2::Route":                              ec2.Route,
	"AWS::EC2::RouteTable":                         ec2.RouteTable,
	"AWS::EC2::SecurityGroup":                      ec2.SecurityGroup,
	"AWS::EC2::SecurityGroupEgress":                resources.SecurityGroupEgress,
	"AWS::EC2::SecurityGroupIngress":               ec2.SecurityGroupIngress,
	"AWS::EC2::SpotFleet":                          resources.SpotFleet,
	"AWS::EC2::Subnet":                             ec2.Subnet,
	"AWS::EC2::SubnetNetworkAclAssociation":        resources.SubnetNetworkAclAssociation,
	"AWS::EC2::SubnetRouteTableAssociation":        ec2.SubnetRouteTableAssociation,
	"AWS::EC2::Volume":                             resources.Volume,
	"AWS::EC2::VolumeAttachment":                   resources.VolumeAttachment,
	"AWS::EC2::VPC":                                resources.VPC,
	"AWS::EC2::VPCDHCPOptionsAssociation":          resources.VPCDHCPOptionsAssociation,
	"AWS::EC2::VPCEndpoint":                        resources.VPCEndpoint,
	"AWS::EC2::VPCGatewayAttachment":               ec2.VPCGatewayAttachment,
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
	"AWS::ElastiCache::CacheCluster":               elasti_cache.CacheCluster,
	"AWS::ElastiCache::ParameterGroup":             resources.ParameterGroup,
	"AWS::ElastiCache::ReplicationGroup":           elasti_cache.ReplicationGroup,
	"AWS::ElastiCache::SecurityGroup":              resources.ElastiCacheSecurityGroup,
	"AWS::ElastiCache::SecurityGroupIngress":       resources.ElastiCacheSecurityGroupIngress,
	"AWS::ElastiCache::SubnetGroup":                elasti_cache.SubnetGroup,
	"AWS::ElasticBeanstalk::Application":           elastic_beanstalk.Application,
	"AWS::ElasticBeanstalk::ApplicationVersion":    elastic_beanstalk.ApplicationVersion,
	"AWS::ElasticBeanstalk::ConfigurationTemplate": elastic_beanstalk.ConfigurationTemplate,
	"AWS::ElasticBeanstalk::Environment":           elastic_beanstalk.Environment,
	"AWS::ElasticLoadBalancing::LoadBalancer":      elastic_load_balancing.LoadBalancer,
	"AWS::IAM::AccessKey":                          resources.AccessKey,
	"AWS::IAM::Group":                              resources.Group,
	"AWS::IAM::InstanceProfile":                    iam.InstanceProfile,
	"AWS::IAM::ManagedPolicy":                      resources.ManagedPolicy,
	"AWS::IAM::Policy":                             iam.Policy,
	"AWS::IAM::Role":                               iam.Role,
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
	"AWS::RDS::DBInstance":                         rds.DBInstance,
	"AWS::RDS::DBParameterGroup":                   resources.DBParameterGroup,
	"AWS::RDS::DBSecurityGroup":                    resources.DBSecurityGroup,
	"AWS::RDS::DBSecurityGroupIngress":             resources.DBSecurityGroupIngress,
	"AWS::RDS::DBSubnetGroup":                      rds.DBSubnetGroup,
	"AWS::RDS::EventSubscription":                  resources.EventSubscription,
	"AWS::RDS::OptionGroup":                        resources.OptionGroup,
	"AWS::Redshift::Cluster":                       resources.RedshiftCluster,
	"AWS::Redshift::ClusterParameterGroup":         resources.ClusterParameterGroup,
	"AWS::Redshift::ClusterSecurityGroup":          resources.ClusterSecurityGroup,
	"AWS::Redshift::ClusterSecurityGroupIngress":   resources.ClusterSecurityGroupIngress,
	"AWS::Redshift::ClusterSubnetGroup":            resources.ClusterSubnetGroup,
	"AWS::Route53::HealthCheck":                    resources.HealthCheck,
	"AWS::Route53::HostedZone":                     resources.HostedZone,
	"AWS::Route53::RecordSet":                      route_53.RecordSet,
	"AWS::Route53::RecordSetGroup":                 resources.RecordSetGroup,
	"AWS::S3::Bucket":                              s3.Bucket,
	"AWS::S3::BucketPolicy":                        s3.BucketPolicy,
	"AWS::SDB::Domain":                             resources.Domain,
	"AWS::SNS::Topic":                              sns.Topic,
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
			tr := schema.NewTemplateResource(template)
			tr.Definition = ctor()
			tr.Properties = rawResource.Properties
			tr.Metadata = rawResource.Metadata
			template.Resources[key] = tr
		} else if !forgiving {
			template.Resources[key] = schema.NewUnrecognisedResource(template, rawResource.Type)
		}
	}

	return template, nil
}
