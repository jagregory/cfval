package main

import "encoding/json"
import (
	"github.com/jagregory/cfval/resources/auto_scaling"
	"github.com/jagregory/cfval/resources/cloud_front"
	"github.com/jagregory/cfval/resources/cloud_trail"
	"github.com/jagregory/cfval/resources/cloud_watch"
	"github.com/jagregory/cfval/resources/dynamo_db"
	"github.com/jagregory/cfval/resources/ec2"
	"github.com/jagregory/cfval/resources/ecs"
	"github.com/jagregory/cfval/resources/efs"
	"github.com/jagregory/cfval/resources/elasti_cache"
	"github.com/jagregory/cfval/resources/elastic_beanstalk"
	"github.com/jagregory/cfval/resources/elastic_load_balancing"
	"github.com/jagregory/cfval/resources/iam"
	"github.com/jagregory/cfval/resources/kinesis"
	"github.com/jagregory/cfval/resources/kms"
	"github.com/jagregory/cfval/resources/lambda"
	"github.com/jagregory/cfval/resources/logs"
	"github.com/jagregory/cfval/resources/not_supported"
	"github.com/jagregory/cfval/resources/rds"
	"github.com/jagregory/cfval/resources/route_53"
	"github.com/jagregory/cfval/resources/s3"
	"github.com/jagregory/cfval/resources/sns"
	"github.com/jagregory/cfval/resources/sqs"
	"github.com/jagregory/cfval/schema"
)

type resourceCtor func() schema.Resource

var typeCtors = map[string]resourceCtor{
	"AWS::AutoScaling::AutoScalingGroup":           auto_scaling.AutoScalingGroup,
	"AWS::AutoScaling::LaunchConfiguration":        auto_scaling.LaunchConfiguration,
	"AWS::AutoScaling::LifecycleHook":              auto_scaling.LifecycleHook,
	"AWS::AutoScaling::ScalingPolicy":              auto_scaling.ScalingPolicy,
	"AWS::AutoScaling::ScheduledAction":            auto_scaling.ScheduledAction,
	"AWS::CloudFormation::Authentication":          not_supported.Authentication,
	"AWS::CloudFormation::CustomResource":          not_supported.CustomResource,
	"AWS::CloudFormation::Init":                    not_supported.Init,
	"AWS::CloudFormation::Interface":               not_supported.Interface,
	"AWS::CloudFormation::Stack":                   not_supported.Stack,
	"AWS::CloudFormation::WaitCondition":           not_supported.WaitCondition,
	"AWS::CloudFormation::WaitConditionHandle":     not_supported.WaitConditionHandle,
	"AWS::CloudFront::Distribution":                cloud_front.Distribution,
	"AWS::CloudTrail::Trail":                       cloud_trail.Trail,
	"AWS::CloudWatch::Alarm":                       cloud_watch.Alarm,
	"AWS::CodeDeploy::Application":                 not_supported.CodeDeployApplication,
	"AWS::CodeDeploy::DeploymentConfig":            not_supported.DeploymentConfig,
	"AWS::CodeDeploy::DeploymentGroup":             not_supported.DeploymentGroup,
	"AWS::CodePipeline::CustomActionType":          not_supported.CustomActionType,
	"AWS::CodePipeline::Pipeline":                  not_supported.Pipeline,
	"AWS::Config::ConfigRule":                      not_supported.ConfigRule,
	"AWS::Config::ConfigurationRecorder":           not_supported.ConfigurationRecorder,
	"AWS::Config::DeliveryChannel":                 not_supported.DeliveryChannel,
	"AWS::DataPipeline::Pipeline":                  not_supported.DataPipelinePipeline,
	"AWS::DirectoryService::MicrosoftAD":           not_supported.MicrosoftAD,
	"AWS::DirectoryService::SimpleAD":              not_supported.SimpleAD,
	"AWS::DynamoDB::Table":                         dynamo_db.Table,
	"AWS::EC2::CustomerGateway":                    ec2.CustomerGateway,
	"AWS::EC2::DHCPOptions":                        ec2.DHCPOptions,
	"AWS::EC2::EIP":                                ec2.Eip,
	"AWS::EC2::EIPAssociation":                     ec2.EIPAssociation,
	"AWS::EC2::Instance":                           ec2.Instance,
	"AWS::EC2::InternetGateway":                    ec2.InternetGateway,
	"AWS::EC2::NetworkAcl":                         ec2.NetworkACL,
	"AWS::EC2::NetworkAclEntry":                    ec2.NetworkACLEntry,
	"AWS::EC2::NetworkInterface":                   ec2.NetworkInterface,
	"AWS::EC2::NetworkInterfaceAttachment":         ec2.NetworkInterfaceAttachment,
	"AWS::EC2::PlacementGroup":                     ec2.PlacementGroup,
	"AWS::EC2::Route":                              ec2.Route,
	"AWS::EC2::RouteTable":                         ec2.RouteTable,
	"AWS::EC2::SecurityGroup":                      ec2.SecurityGroup,
	"AWS::EC2::SecurityGroupEgress":                ec2.SecurityGroupEgress,
	"AWS::EC2::SecurityGroupIngress":               ec2.SecurityGroupIngress,
	"AWS::EC2::SpotFleet":                          ec2.SpotFleet,
	"AWS::EC2::Subnet":                             ec2.Subnet,
	"AWS::EC2::SubnetNetworkAclAssociation":        ec2.SubnetNetworkACLAssociation,
	"AWS::EC2::SubnetRouteTableAssociation":        ec2.SubnetRouteTableAssociation,
	"AWS::EC2::Volume":                             ec2.Volume,
	"AWS::EC2::VolumeAttachment":                   ec2.VolumeAttachment,
	"AWS::EC2::VPC":                                ec2.VPC,
	"AWS::EC2::VPCDHCPOptionsAssociation":          ec2.VPCDHCPOptionsAssociation,
	"AWS::EC2::VPCEndpoint":                        ec2.VPCEndpoint,
	"AWS::EC2::VPCGatewayAttachment":               ec2.VPCGatewayAttachment,
	"AWS::EC2::VPCPeeringConnection":               ec2.VPCPeeringConnection,
	"AWS::EC2::VPNConnection":                      ec2.VPNConnection,
	"AWS::EC2::VPNConnectionRoute":                 ec2.VPNConnectionRoute,
	"AWS::EC2::VPNGateway":                         ec2.VPNGateway,
	"AWS::EC2::VPNGatewayRoutePropagation":         ec2.VPNGatewayRoutePropagation,
	"AWS::ECS::Cluster":                            ecs.Cluster,
	"AWS::ECS::Service":                            ecs.Service,
	"AWS::ECS::TaskDefinition":                     ecs.TaskDefinition,
	"AWS::EFS::FileSystem":                         efs.FileSystem,
	"AWS::EFS::MountTarget":                        efs.MountTarget,
	"AWS::ElastiCache::CacheCluster":               elasti_cache.CacheCluster,
	"AWS::ElastiCache::ParameterGroup":             elasti_cache.ParameterGroup,
	"AWS::ElastiCache::ReplicationGroup":           elasti_cache.ReplicationGroup,
	"AWS::ElastiCache::SecurityGroup":              elasti_cache.SecurityGroup,
	"AWS::ElastiCache::SecurityGroupIngress":       elasti_cache.SecurityGroupIngress,
	"AWS::ElastiCache::SubnetGroup":                elasti_cache.SubnetGroup,
	"AWS::ElasticBeanstalk::Application":           elastic_beanstalk.Application,
	"AWS::ElasticBeanstalk::ApplicationVersion":    elastic_beanstalk.ApplicationVersion,
	"AWS::ElasticBeanstalk::ConfigurationTemplate": elastic_beanstalk.ConfigurationTemplate,
	"AWS::ElasticBeanstalk::Environment":           elastic_beanstalk.Environment,
	"AWS::ElasticLoadBalancing::LoadBalancer":      elastic_load_balancing.LoadBalancer,
	"AWS::IAM::AccessKey":                          iam.AccessKey,
	"AWS::IAM::Group":                              iam.Group,
	"AWS::IAM::InstanceProfile":                    iam.InstanceProfile,
	"AWS::IAM::ManagedPolicy":                      iam.ManagedPolicy,
	"AWS::IAM::Policy":                             iam.Policy,
	"AWS::IAM::Role":                               iam.Role,
	"AWS::IAM::User":                               iam.User,
	"AWS::IAM::UserToGroupAddition":                iam.UserToGroupAddition,
	"AWS::Kinesis::Stream":                         kinesis.Stream,
	"AWS::KMS::Key":                                kms.Key,
	"AWS::Lambda::EventSourceMapping":              lambda.EventSourceMapping,
	"AWS::Lambda::Function":                        lambda.Function,
	"AWS::Lambda::Permission":                      lambda.Permission,
	"AWS::Logs::Destination":                       logs.Destination,
	"AWS::Logs::LogGroup":                          logs.LogGroup,
	"AWS::Logs::LogStream":                         logs.LogStream,
	"AWS::Logs::MetricFilter":                      logs.MetricFilter,
	"AWS::Logs::SubscriptionFilter":                logs.SubscriptionFilter,
	"AWS::OpsWorks::App":                           not_supported.App,
	"AWS::OpsWorks::ElasticLoadBalancerAttachment": not_supported.ElasticLoadBalancerAttachment,
	"AWS::OpsWorks::Instance":                      not_supported.OpsWorksInstance,
	"AWS::OpsWorks::Layer":                         not_supported.Layer,
	"AWS::OpsWorks::Stack":                         not_supported.OpsWorksStack,
	"AWS::RDS::DBCluster":                          rds.DBCluster,
	"AWS::RDS::DBClusterParameterGroup":            rds.DBClusterParameterGroup,
	"AWS::RDS::DBInstance":                         rds.DBInstance,
	"AWS::RDS::DBParameterGroup":                   rds.DBParameterGroup,
	"AWS::RDS::DBSecurityGroup":                    rds.DBSecurityGroup,
	"AWS::RDS::DBSecurityGroupIngress":             rds.DBSecurityGroupIngress,
	"AWS::RDS::DBSubnetGroup":                      rds.DBSubnetGroup,
	"AWS::RDS::EventSubscription":                  rds.EventSubscription,
	"AWS::RDS::OptionGroup":                        rds.OptionGroup,
	"AWS::Redshift::Cluster":                       not_supported.RedshiftCluster,
	"AWS::Redshift::ClusterParameterGroup":         not_supported.ClusterParameterGroup,
	"AWS::Redshift::ClusterSecurityGroup":          not_supported.ClusterSecurityGroup,
	"AWS::Redshift::ClusterSecurityGroupIngress":   not_supported.ClusterSecurityGroupIngress,
	"AWS::Redshift::ClusterSubnetGroup":            not_supported.ClusterSubnetGroup,
	"AWS::Route53::HealthCheck":                    route_53.HealthCheck,
	"AWS::Route53::HostedZone":                     route_53.HostedZone,
	"AWS::Route53::RecordSet":                      route_53.RecordSet,
	"AWS::Route53::RecordSetGroup":                 route_53.RecordSetGroup,
	"AWS::S3::Bucket":                              s3.Bucket,
	"AWS::S3::BucketPolicy":                        s3.BucketPolicy,
	"AWS::SDB::Domain":                             not_supported.Domain,
	"AWS::SNS::Topic":                              sns.Topic,
	"AWS::SNS::TopicPolicy":                        sns.TopicPolicy,
	"AWS::SQS::Queue":                              sqs.Queue,
	"AWS::SQS::QueuePolicy":                        sqs.QueuePolicy,
	"AWS::SSM::Document":                           not_supported.Document,
	"AWS::WAF::ByteMatchSet":                       not_supported.ByteMatchSet,
	"AWS::WAF::IPSet":                              not_supported.IPSet,
	"AWS::WAF::Rule":                               not_supported.Rule,
	"AWS::WAF::SqlInjectionMatchSet":               not_supported.SqlInjectionMatchSet,
	"AWS::WAF::WebACL":                             not_supported.WebACL,
	"AWS::WorkSpaces::Workspace":                   not_supported.Workspace,
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
