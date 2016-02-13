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

func Table() Resource {
	return Resource{
		AwsType:    "AWS::DynamoDB::Table",
		Properties: Properties{},
	}
}

func CustomerGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::CustomerGateway",
		Properties: Properties{},
	}
}

func DHCPOptions() Resource {
	return Resource{
		AwsType:    "AWS::EC2::DHCPOptions",
		Properties: Properties{},
	}
}

func EIP() Resource {
	return Resource{
		AwsType:    "AWS::EC2::EIP",
		Properties: Properties{},
	}
}

func EIPAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::EIPAssociation",
		Properties: Properties{},
	}
}

func NetworkAcl() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkAcl",
		Properties: Properties{},
	}
}

func NetworkAclEntry() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkAclEntry",
		Properties: Properties{},
	}
}

func NetworkInterface() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkInterface",
		Properties: Properties{},
	}
}

func NetworkInterfaceAttachment() Resource {
	return Resource{
		AwsType:    "AWS::EC2::NetworkInterfaceAttachment",
		Properties: Properties{},
	}
}

func PlacementGroup() Resource {
	return Resource{
		AwsType:    "AWS::EC2::PlacementGroup",
		Properties: Properties{},
	}
}

func SecurityGroupEgress() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SecurityGroupEgress",
		Properties: Properties{},
	}
}

func SpotFleet() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SpotFleet",
		Properties: Properties{},
	}
}

func SubnetNetworkAclAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::SubnetNetworkAclAssociation",
		Properties: Properties{},
	}
}

func Volume() Resource {
	return Resource{
		AwsType:    "AWS::EC2::Volume",
		Properties: Properties{},
	}
}

func VolumeAttachment() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VolumeAttachment",
		Properties: Properties{},
	}
}

func VPC() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPC",
		Properties: Properties{},
	}
}

func VPCDHCPOptionsAssociation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCDHCPOptionsAssociation",
		Properties: Properties{},
	}
}

func VPCEndpoint() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCEndpoint",
		Properties: Properties{},
	}
}

func VPCPeeringConnection() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPCPeeringConnection",
		Properties: Properties{},
	}
}

func VPNConnection() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNConnection",
		Properties: Properties{},
	}
}

func VPNConnectionRoute() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNConnectionRoute",
		Properties: Properties{},
	}
}

func VPNGateway() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNGateway",
		Properties: Properties{},
	}
}

func VPNGatewayRoutePropagation() Resource {
	return Resource{
		AwsType:    "AWS::EC2::VPNGatewayRoutePropagation",
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

func ParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::ParameterGroup",
		Properties: Properties{},
	}
}

func ElastiCacheSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::SecurityGroup",
		Properties: Properties{},
	}
}

func ElastiCacheSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::ElastiCache::SecurityGroupIngress",
		Properties: Properties{},
	}
}

func AccessKey() Resource {
	return Resource{
		AwsType:    "AWS::IAM::AccessKey",
		Properties: Properties{},
	}
}

func Group() Resource {
	return Resource{
		AwsType:    "AWS::IAM::Group",
		Properties: Properties{},
	}
}

func ManagedPolicy() Resource {
	return Resource{
		AwsType:    "AWS::IAM::ManagedPolicy",
		Properties: Properties{},
	}
}

func User() Resource {
	return Resource{
		AwsType:    "AWS::IAM::User",
		Properties: Properties{},
	}
}

func UserToGroupAddition() Resource {
	return Resource{
		AwsType:    "AWS::IAM::UserToGroupAddition",
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

func DBCluster() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBCluster",
		Properties: Properties{},
	}
}

func DBClusterParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBClusterParameterGroup",
		Properties: Properties{},
	}
}

func DBParameterGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBParameterGroup",
		Properties: Properties{},
	}
}

func DBSecurityGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBSecurityGroup",
		Properties: Properties{},
	}
}

func DBSecurityGroupIngress() Resource {
	return Resource{
		AwsType:    "AWS::RDS::DBSecurityGroupIngress",
		Properties: Properties{},
	}
}

func EventSubscription() Resource {
	return Resource{
		AwsType:    "AWS::RDS::EventSubscription",
		Properties: Properties{},
	}
}

func OptionGroup() Resource {
	return Resource{
		AwsType:    "AWS::RDS::OptionGroup",
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

func HealthCheck() Resource {
	return Resource{
		AwsType:    "AWS::Route53::HealthCheck",
		Properties: Properties{},
	}
}

func HostedZone() Resource {
	return Resource{
		AwsType:    "AWS::Route53::HostedZone",
		Properties: Properties{},
	}
}

func RecordSetGroup() Resource {
	return Resource{
		AwsType:    "AWS::Route53::RecordSetGroup",
		Properties: Properties{},
	}
}

func BucketPolicy() Resource {
	return Resource{
		AwsType:    "AWS::S3::BucketPolicy",
		Properties: Properties{},
	}
}

func Domain() Resource {
	return Resource{
		AwsType:    "AWS::SDB::Domain",
		Properties: Properties{},
	}
}

func TopicPolicy() Resource {
	return Resource{
		AwsType:    "AWS::SNS::TopicPolicy",
		Properties: Properties{},
	}
}

func Queue() Resource {
	return Resource{
		AwsType:    "AWS::SQS::Queue",
		Properties: Properties{},
	}
}

func QueuePolicy() Resource {
	return Resource{
		AwsType:    "AWS::SQS::QueuePolicy",
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
