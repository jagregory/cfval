package resources

import . "github.com/jagregory/cfval/schema"

func Authentication() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Authentication",
		Properties: Properties{},
	}
}

func CustomResource() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::CustomResource",
		Properties: Properties{},
	}
}

func Init() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Init",
		Properties: Properties{},
	}
}

func Interface() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Interface",
		Properties: Properties{},
	}
}

func Stack() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Stack",
		Properties: Properties{},
	}
}

func WaitCondition() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::WaitCondition",
		Properties: Properties{},
	}
}

func WaitConditionHandle() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::WaitConditionHandle",
		Properties: Properties{},
	}
}

func CodeDeployApplication() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::Application",
		Properties: Properties{},
	}
}

func DeploymentConfig() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::DeploymentConfig",
		Properties: Properties{},
	}
}

func DeploymentGroup() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::DeploymentGroup",
		Properties: Properties{},
	}
}

func CustomActionType() Resource {
	return Resource{
		AwsType:    "AWS::CodePipeline::CustomActionType",
		Properties: Properties{},
	}
}

func Pipeline() Resource {
	return Resource{
		AwsType:    "AWS::CodePipeline::Pipeline",
		Properties: Properties{},
	}
}

func ConfigRule() Resource {
	return Resource{
		AwsType:    "AWS::Config::ConfigRule",
		Properties: Properties{},
	}
}

func ConfigurationRecorder() Resource {
	return Resource{
		AwsType:    "AWS::Config::ConfigurationRecorder",
		Properties: Properties{},
	}
}

func DeliveryChannel() Resource {
	return Resource{
		AwsType:    "AWS::Config::DeliveryChannel",
		Properties: Properties{},
	}
}

func DataPipelinePipeline() Resource {
	return Resource{
		AwsType:    "AWS::DataPipeline::Pipeline",
		Properties: Properties{},
	}
}

func MicrosoftAD() Resource {
	return Resource{
		AwsType:    "AWS::DirectoryService::MicrosoftAD",
		Properties: Properties{},
	}
}

func SimpleAD() Resource {
	return Resource{
		AwsType:    "AWS::DirectoryService::SimpleAD",
		Properties: Properties{},
	}
}

func Cluster() Resource {
	return Resource{
		AwsType:    "AWS::ECS::Cluster",
		Properties: Properties{},
	}
}

func Service() Resource {
	return Resource{
		AwsType:    "AWS::ECS::Service",
		Properties: Properties{},
	}
}

func TaskDefinition() Resource {
	return Resource{
		AwsType:    "AWS::ECS::TaskDefinition",
		Properties: Properties{},
	}
}

func FileSystem() Resource {
	return Resource{
		AwsType:    "AWS::EFS::FileSystem",
		Properties: Properties{},
	}
}

func MountTarget() Resource {
	return Resource{
		AwsType:    "AWS::EFS::MountTarget",
		Properties: Properties{},
	}
}

func Stream() Resource {
	return Resource{
		AwsType:    "AWS::Kinesis::Stream",
		Properties: Properties{},
	}
}

func Key() Resource {
	return Resource{
		AwsType:    "AWS::KMS::Key",
		Properties: Properties{},
	}
}

func EventSourceMapping() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::EventSourceMapping",
		Properties: Properties{},
	}
}

func Function() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::Function",
		Properties: Properties{},
	}
}

func Permission() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::Permission",
		Properties: Properties{},
	}
}

func Destination() Resource {
	return Resource{
		AwsType:    "AWS::Logs::Destination",
		Properties: Properties{},
	}
}

func LogGroup() Resource {
	return Resource{
		AwsType:    "AWS::Logs::LogGroup",
		Properties: Properties{},
	}
}

func LogStream() Resource {
	return Resource{
		AwsType:    "AWS::Logs::LogStream",
		Properties: Properties{},
	}
}

func MetricFilter() Resource {
	return Resource{
		AwsType:    "AWS::Logs::MetricFilter",
		Properties: Properties{},
	}
}

func SubscriptionFilter() Resource {
	return Resource{
		AwsType:    "AWS::Logs::SubscriptionFilter",
		Properties: Properties{},
	}
}

func App() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::App",
		Properties: Properties{},
	}
}

func ElasticLoadBalancerAttachment() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::ElasticLoadBalancerAttachment",
		Properties: Properties{},
	}
}

func OpsWorksInstance() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Instance",
		Properties: Properties{},
	}
}

func Layer() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Layer",
		Properties: Properties{},
	}
}

func OpsWorksStack() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Stack",
		Properties: Properties{},
	}
}

func RedshiftCluster() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::Cluster",
		Properties: Properties{},
	}
}

func ClusterParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterParameterGroup",
		Properties: Properties{},
	}
}

func ClusterSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSecurityGroup",
		Properties: Properties{},
	}
}

func ClusterSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSecurityGroupIngress",
		Properties: Properties{},
	}
}

func ClusterSubnetGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSubnetGroup",
		Properties: Properties{},
	}
}

func Domain() Resource {
	return Resource{
		AwsType:    "AWS::SDB::Domain",
		Properties: Properties{},
	}
}

func Document() Resource {
	return Resource{
		AwsType:    "AWS::SSM::Document",
		Properties: Properties{},
	}
}

func ByteMatchSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::ByteMatchSet",
		Properties: Properties{},
	}
}

func IPSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::IPSet",
		Properties: Properties{},
	}
}

func Rule() Resource {
	return Resource{
		AwsType:    "AWS::WAF::Rule",
		Properties: Properties{},
	}
}

func SqlInjectionMatchSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::SqlInjectionMatchSet",
		Properties: Properties{},
	}
}

func WebACL() Resource {
	return Resource{
		AwsType:    "AWS::WAF::WebACL",
		Properties: Properties{},
	}
}

func Workspace() Resource {
	return Resource{
		AwsType:    "AWS::WorkSpaces::Workspace",
		Properties: Properties{},
	}
}
