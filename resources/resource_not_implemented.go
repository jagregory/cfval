package resources

import . "github.com/jagregory/cfval/schema"

func LifecycleHook() Resource {
	return Resource{
		AwsType:    "AWS::AutoScaling::LifecycleHook",
		Properties: map[string]Schema{},
	}
}

func ScalingPolicy() Resource {
	return Resource{
		AwsType:    "AWS::AutoScaling::ScalingPolicy",
		Properties: map[string]Schema{},
	}
}

func ScheduledAction() Resource {
	return Resource{
		AwsType:    "AWS::AutoScaling::ScheduledAction",
		Properties: map[string]Schema{},
	}
}

func Authentication() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Authentication",
		Properties: map[string]Schema{},
	}
}

func CustomResource() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::CustomResource",
		Properties: map[string]Schema{},
	}
}

func Init() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Init",
		Properties: map[string]Schema{},
	}
}

func Interface() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Interface",
		Properties: map[string]Schema{},
	}
}

func Stack() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::Stack",
		Properties: map[string]Schema{},
	}
}

func WaitCondition() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::WaitCondition",
		Properties: map[string]Schema{},
	}
}

func WaitConditionHandle() Resource {
	return Resource{
		AwsType:    "AWS::CloudFormation::WaitConditionHandle",
		Properties: map[string]Schema{},
	}
}

func Trail() Resource {
	return Resource{
		AwsType:    "AWS::CloudTrail::Trail",
		Properties: map[string]Schema{},
	}
}

func CodeDeployApplication() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::Application",
		Properties: map[string]Schema{},
	}
}

func DeploymentConfig() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::DeploymentConfig",
		Properties: map[string]Schema{},
	}
}

func DeploymentGroup() Resource {
	return Resource{
		AwsType:    "AWS::CodeDeploy::DeploymentGroup",
		Properties: map[string]Schema{},
	}
}

func CustomActionType() Resource {
	return Resource{
		AwsType:    "AWS::CodePipeline::CustomActionType",
		Properties: map[string]Schema{},
	}
}

func Pipeline() Resource {
	return Resource{
		AwsType:    "AWS::CodePipeline::Pipeline",
		Properties: map[string]Schema{},
	}
}

func ConfigRule() Resource {
	return Resource{
		AwsType:    "AWS::Config::ConfigRule",
		Properties: map[string]Schema{},
	}
}

func ConfigurationRecorder() Resource {
	return Resource{
		AwsType:    "AWS::Config::ConfigurationRecorder",
		Properties: map[string]Schema{},
	}
}

func DeliveryChannel() Resource {
	return Resource{
		AwsType:    "AWS::Config::DeliveryChannel",
		Properties: map[string]Schema{},
	}
}

func DataPipelinePipeline() Resource {
	return Resource{
		AwsType:    "AWS::DataPipeline::Pipeline",
		Properties: map[string]Schema{},
	}
}

func MicrosoftAD() Resource {
	return Resource{
		AwsType:    "AWS::DirectoryService::MicrosoftAD",
		Properties: map[string]Schema{},
	}
}

func SimpleAD() Resource {
	return Resource{
		AwsType:    "AWS::DirectoryService::SimpleAD",
		Properties: map[string]Schema{},
	}
}

func Table() Resource {
	return Resource{
		AwsType:    "AWS::DynamoDB::Table",
		Properties: map[string]Schema{},
	}
}

func CustomerGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::CustomerGateway",
		Properties: map[string]Schema{},
	}
}

func DHCPOptions() Resource {
	return Resource{
		AwsType:    "AWS::EC2::DHCPOptions",
		Properties: map[string]Schema{},
	}
}

func EIP() Resource {
	return Resource{
		AwsType:    "AWS::EC2::EIP",
		Properties: map[string]Schema{},
	}
}

func EIPAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::EIPAssociation",
		Properties: map[string]Schema{},
	}
}

func NetworkAcl() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkAcl",
		Properties: map[string]Schema{},
	}
}

func NetworkAclEntry() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkAclEntry",
		Properties: map[string]Schema{},
	}
}

func NetworkInterface() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkInterface",
		Properties: map[string]Schema{},
	}
}

func NetworkInterfaceAttachment() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkInterfaceAttachment",
		Properties: map[string]Schema{},
	}
}

func PlacementGroup() Resource {
	return Resource{
		AwsType:    "AWS::EC2::PlacementGroup",
		Properties: map[string]Schema{},
	}
}

func SecurityGroupEgress() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SecurityGroupEgress",
		Properties: map[string]Schema{},
	}
}

func SpotFleet() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SpotFleet",
		Properties: map[string]Schema{},
	}
}

func SubnetNetworkAclAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SubnetNetworkAclAssociation",
		Properties: map[string]Schema{},
	}
}

func Volume() Resource {
	return Resource{
		AwsType:    "AWS::EC2::Volume",
		Properties: map[string]Schema{},
	}
}

func VolumeAttachment() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VolumeAttachment",
		Properties: map[string]Schema{},
	}
}

func VPC() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPC",
		Properties: map[string]Schema{},
	}
}

func VPCDHCPOptionsAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCDHCPOptionsAssociation",
		Properties: map[string]Schema{},
	}
}

func VPCEndpoint() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCEndpoint",
		Properties: map[string]Schema{},
	}
}

func VPCGatewayAttachment() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCGatewayAttachment",
		Properties: map[string]Schema{},
	}
}

func VPCPeeringConnection() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCPeeringConnection",
		Properties: map[string]Schema{},
	}
}

func VPNConnection() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNConnection",
		Properties: map[string]Schema{},
	}
}

func VPNConnectionRoute() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNConnectionRoute",
		Properties: map[string]Schema{},
	}
}

func VPNGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNGateway",
		Properties: map[string]Schema{},
	}
}

func VPNGatewayRoutePropagation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNGatewayRoutePropagation",
		Properties: map[string]Schema{},
	}
}

func Cluster() Resource {
	return Resource{
		AwsType:    "AWS::ECS::Cluster",
		Properties: map[string]Schema{},
	}
}

func Service() Resource {
	return Resource{
		AwsType:    "AWS::ECS::Service",
		Properties: map[string]Schema{},
	}
}

func TaskDefinition() Resource {
	return Resource{
		AwsType:    "AWS::ECS::TaskDefinition",
		Properties: map[string]Schema{},
	}
}

func FileSystem() Resource {
	return Resource{
		AwsType:    "AWS::EFS::FileSystem",
		Properties: map[string]Schema{},
	}
}

func MountTarget() Resource {
	return Resource{
		AwsType:    "AWS::EFS::MountTarget",
		Properties: map[string]Schema{},
	}
}

func ParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::ParameterGroup",
		Properties: map[string]Schema{},
	}
}

func ElastiCacheSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::SecurityGroup",
		Properties: map[string]Schema{},
	}
}

func ElastiCacheSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::SecurityGroupIngress",
		Properties: map[string]Schema{},
	}
}

func AccessKey() Resource {
	return Resource{
		AwsType:    "AWS::IAM::AccessKey",
		Properties: map[string]Schema{},
	}
}

func Group() Resource {
	return Resource{
		AwsType:    "AWS::IAM::Group",
		Properties: map[string]Schema{},
	}
}

func ManagedPolicy() Resource {
	return Resource{
		AwsType:    "AWS::IAM::ManagedPolicy",
		Properties: map[string]Schema{},
	}
}

func User() Resource {
	return Resource{
		AwsType:    "AWS::IAM::User",
		Properties: map[string]Schema{},
	}
}

func UserToGroupAddition() Resource {
	return Resource{
		AwsType:    "AWS::IAM::UserToGroupAddition",
		Properties: map[string]Schema{},
	}
}

func Stream() Resource {
	return Resource{
		AwsType:    "AWS::Kinesis::Stream",
		Properties: map[string]Schema{},
	}
}

func Key() Resource {
	return Resource{
		AwsType:    "AWS::KMS::Key",
		Properties: map[string]Schema{},
	}
}

func EventSourceMapping() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::EventSourceMapping",
		Properties: map[string]Schema{},
	}
}

func Function() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::Function",
		Properties: map[string]Schema{},
	}
}

func Permission() Resource {
	return Resource{
		AwsType:    "AWS::Lambda::Permission",
		Properties: map[string]Schema{},
	}
}

func Destination() Resource {
	return Resource{
		AwsType:    "AWS::Logs::Destination",
		Properties: map[string]Schema{},
	}
}

func LogGroup() Resource {
	return Resource{
		AwsType:    "AWS::Logs::LogGroup",
		Properties: map[string]Schema{},
	}
}

func LogStream() Resource {
	return Resource{
		AwsType:    "AWS::Logs::LogStream",
		Properties: map[string]Schema{},
	}
}

func MetricFilter() Resource {
	return Resource{
		AwsType:    "AWS::Logs::MetricFilter",
		Properties: map[string]Schema{},
	}
}

func SubscriptionFilter() Resource {
	return Resource{
		AwsType:    "AWS::Logs::SubscriptionFilter",
		Properties: map[string]Schema{},
	}
}

func App() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::App",
		Properties: map[string]Schema{},
	}
}

func ElasticLoadBalancerAttachment() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::ElasticLoadBalancerAttachment",
		Properties: map[string]Schema{},
	}
}

func OpsWorksInstance() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Instance",
		Properties: map[string]Schema{},
	}
}

func Layer() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Layer",
		Properties: map[string]Schema{},
	}
}

func OpsWorksStack() Resource {
	return Resource{
		AwsType:    "AWS::OpsWorks::Stack",
		Properties: map[string]Schema{},
	}
}

func DBCluster() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBCluster",
		Properties: map[string]Schema{},
	}
}

func DBClusterParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBClusterParameterGroup",
		Properties: map[string]Schema{},
	}
}

func DBParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBParameterGroup",
		Properties: map[string]Schema{},
	}
}

func DBSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBSecurityGroup",
		Properties: map[string]Schema{},
	}
}

func DBSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBSecurityGroupIngress",
		Properties: map[string]Schema{},
	}
}

func EventSubscription() Resource {
	return Resource{
		AwsType:    "AWS::RDS::EventSubscription",
		Properties: map[string]Schema{},
	}
}

func OptionGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::OptionGroup",
		Properties: map[string]Schema{},
	}
}

func RedshiftCluster() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::Cluster",
		Properties: map[string]Schema{},
	}
}

func ClusterParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterParameterGroup",
		Properties: map[string]Schema{},
	}
}

func ClusterSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSecurityGroup",
		Properties: map[string]Schema{},
	}
}

func ClusterSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSecurityGroupIngress",
		Properties: map[string]Schema{},
	}
}

func ClusterSubnetGroup() Resource {
	return Resource{
		AwsType:    "AWS::Redshift::ClusterSubnetGroup",
		Properties: map[string]Schema{},
	}
}

func HealthCheck() Resource {
	return Resource{
		AwsType:    "AWS::Route53::HealthCheck",
		Properties: map[string]Schema{},
	}
}

func HostedZone() Resource {
	return Resource{
		AwsType:    "AWS::Route53::HostedZone",
		Properties: map[string]Schema{},
	}
}

func RecordSetGroup() Resource {
	return Resource{
		AwsType:    "AWS::Route53::RecordSetGroup",
		Properties: map[string]Schema{},
	}
}

func BucketPolicy() Resource {
	return Resource{
		AwsType:    "AWS::S3::BucketPolicy",
		Properties: map[string]Schema{},
	}
}

func Domain() Resource {
	return Resource{
		AwsType:    "AWS::SDB::Domain",
		Properties: map[string]Schema{},
	}
}

func TopicPolicy() Resource {
	return Resource{
		AwsType:    "AWS::SNS::TopicPolicy",
		Properties: map[string]Schema{},
	}
}

func Queue() Resource {
	return Resource{
		AwsType:    "AWS::SQS::Queue",
		Properties: map[string]Schema{},
	}
}

func QueuePolicy() Resource {
	return Resource{
		AwsType:    "AWS::SQS::QueuePolicy",
		Properties: map[string]Schema{},
	}
}

func Document() Resource {
	return Resource{
		AwsType:    "AWS::SSM::Document",
		Properties: map[string]Schema{},
	}
}

func ByteMatchSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::ByteMatchSet",
		Properties: map[string]Schema{},
	}
}

func IPSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::IPSet",
		Properties: map[string]Schema{},
	}
}

func Rule() Resource {
	return Resource{
		AwsType:    "AWS::WAF::Rule",
		Properties: map[string]Schema{},
	}
}

func SqlInjectionMatchSet() Resource {
	return Resource{
		AwsType:    "AWS::WAF::SqlInjectionMatchSet",
		Properties: map[string]Schema{},
	}
}

func WebACL() Resource {
	return Resource{
		AwsType:    "AWS::WAF::WebACL",
		Properties: map[string]Schema{},
	}
}

func Workspace() Resource {
	return Resource{
		AwsType:    "AWS::WorkSpaces::Workspace",
		Properties: map[string]Schema{},
	}
}
