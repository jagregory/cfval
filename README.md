# cfval: The CloudFormation template validator

> Have you ever waited 15 minutes for CloudFormation to let you know that you forgot to set the Type of a DNS record? Yeah, and that's on a *good day*. Try 45 minutes for your CloudFront Distribution to fail...
>
> After getting very tired of this process, and with a large infrastructure refactor looming, I decided some time could be better spent catching issues sooner in the process. Hence, cfval.

`cfval` is a small tool which validates a CloudFormation JSON template and notifies you of any issues it can find. Missing required properties, properties which conflict with others, `Ref`s to parameters which don't exist or properties of resources, and much more.

```
$ cfval < my-template.json

Resources.MyLaunchConfiguration.UserData.Ref ... Ref 'CloudInitScript' is not a resource or parameter

1 failure
```

## Installation

For now cfval is only installable via `go get`. This will change once development stabalises and I can push releases out.

`go get -v github.com/jagregory/cfval`

## Known issues

Heaps of resource types aren't supported at the moment. `cfval` currently only supports the resources I've specifically created for my current infrastructure. There are no resources which are currently completely defined.

Watch this space. Contributors *very welcome*.

The following resources are either implemented or partially implemented:

  - ✅ AWS::AutoScaling::AutoScalingGroup
  - 🕒 AWS::AutoScaling::LaunchConfiguration
  - [ ] AWS::AutoScaling::LifecycleHook
  - [ ] AWS::AutoScaling::ScalingPolicy
  - [ ] AWS::AutoScaling::ScheduledAction
  - [ ] AWS::CloudFormation::Authentication
  - [ ] AWS::CloudFormation::CustomResource
  - [ ] AWS::CloudFormation::Init
  - [ ] AWS::CloudFormation::Interface
  - [ ] AWS::CloudFormation::Stack
  - [ ] AWS::CloudFormation::WaitCondition
  - [ ] AWS::CloudFormation::WaitConditionHandle
  - 🕒 AWS::CloudFront::Distribution
  - [ ] AWS::CloudTrail::Trail
  - 🕒 AWS::CloudWatch::Alarm
  - [ ] AWS::CodeDeploy::Application
  - [ ] AWS::CodeDeploy::DeploymentConfig
  - [ ] AWS::CodeDeploy::DeploymentGroup
  - [ ] AWS::CodePipeline::CustomActionType
  - [ ] AWS::CodePipeline::Pipeline
  - [ ] AWS::Config::ConfigRule
  - [ ] AWS::Config::ConfigurationRecorder
  - [ ] AWS::Config::DeliveryChannel
  - [ ] AWS::DataPipeline::Pipeline
  - [ ] AWS::DirectoryService::MicrosoftAD
  - [ ] AWS::DirectoryService::SimpleAD
  - [ ] AWS::DynamoDB::Table
  - [ ] AWS::EC2::CustomerGateway
  - [ ] AWS::EC2::DHCPOptions
  - 🕒 AWS::EC2::EIP
  - [ ] AWS::EC2::EIPAssociation
  - 🕒 AWS::EC2::Instance
  - 🕒 AWS::EC2::InternetGateway
  - [ ] AWS::EC2::NetworkAcl
  - [ ] AWS::EC2::NetworkAclEntry
  - [ ] AWS::EC2::NetworkInterface
  - [ ] AWS::EC2::NetworkInterfaceAttachment
  - [ ] AWS::EC2::PlacementGroup
  - [ ] AWS::EC2::Route
  - [ ] AWS::EC2::RouteTable
  - [ ] AWS::EC2::SecurityGroup
  - [ ] AWS::EC2::SecurityGroupEgress
  - [ ] AWS::EC2::SecurityGroupIngress
  - [ ] AWS::EC2::SpotFleet
  - [ ] AWS::EC2::Subnet
  - [ ] AWS::EC2::SubnetNetworkAclAssociation
  - [ ] AWS::EC2::SubnetRouteTableAssociation
  - [ ] AWS::EC2::Volume
  - [ ] AWS::EC2::VolumeAttachment
  - [ ] AWS::EC2::VPC
  - [ ] AWS::EC2::VPCDHCPOptionsAssociation
  - [ ] AWS::EC2::VPCEndpoint
  - [ ] AWS::EC2::VPCGatewayAttachment
  - [ ] AWS::EC2::VPCPeeringConnection
  - [ ] AWS::EC2::VPNConnection
  - [ ] AWS::EC2::VPNConnectionRoute
  - [ ] AWS::EC2::VPNGateway
  - [ ] AWS::EC2::VPNGatewayRoutePropagation
  - [ ] AWS::ECS::Cluster
  - [ ] AWS::ECS::Service
  - [ ] AWS::ECS::TaskDefinition
  - [ ] AWS::EFS::FileSystem
  - [ ] AWS::EFS::MountTarget
  - [ ] AWS::ElastiCache::CacheCluster
  - [ ] AWS::ElastiCache::ParameterGroup
  - [ ] AWS::ElastiCache::ReplicationGroup
  - [ ] AWS::ElastiCache::SecurityGroup
  - [ ] AWS::ElastiCache::SecurityGroupIngress
  - [ ] AWS::ElastiCache::SubnetGroup
  - [ ] AWS::ElasticBeanstalk::Application
  - [ ] AWS::ElasticBeanstalk::ApplicationVersion
  - [ ] AWS::ElasticBeanstalk::ConfigurationTemplate
  - [ ] AWS::ElasticBeanstalk::Environment
  - [ ] AWS::ElasticLoadBalancing::LoadBalancer
  - [ ] AWS::IAM::AccessKey
  - [ ] AWS::IAM::Group
  - [ ] AWS::IAM::InstanceProfile
  - [ ] AWS::IAM::ManagedPolicy
  - [ ] AWS::IAM::Policy
  - [ ] AWS::IAM::Role
  - [ ] AWS::IAM::User
  - [ ] AWS::IAM::UserToGroupAddition
  - [ ] AWS::Kinesis::Stream
  - [ ] AWS::KMS::Key
  - [ ] AWS::Lambda::EventSourceMapping
  - [ ] AWS::Lambda::Function
  - [ ] AWS::Lambda::Permission
  - [ ] AWS::Logs::Destination
  - [ ] AWS::Logs::LogGroup
  - [ ] AWS::Logs::LogStream
  - [ ] AWS::Logs::MetricFilter
  - [ ] AWS::Logs::SubscriptionFilter
  - [ ] AWS::OpsWorks::App
  - [ ] AWS::OpsWorks::ElasticLoadBalancerAttachment
  - [ ] AWS::OpsWorks::Instance
  - [ ] AWS::OpsWorks::Layer
  - [ ] AWS::OpsWorks::Stack
  - [ ] AWS::RDS::DBCluster
  - [ ] AWS::RDS::DBClusterParameterGroup
  - [ ] AWS::RDS::DBInstance
  - [ ] AWS::RDS::DBParameterGroup
  - [ ] AWS::RDS::DBSecurityGroup
  - [ ] AWS::RDS::DBSecurityGroupIngress
  - [ ] AWS::RDS::DBSubnetGroup
  - [ ] AWS::RDS::EventSubscription
  - [ ] AWS::RDS::OptionGroup
  - [ ] AWS::Redshift::Cluster
  - [ ] AWS::Redshift::ClusterParameterGroup
  - [ ] AWS::Redshift::ClusterSecurityGroup
  - [ ] AWS::Redshift::ClusterSecurityGroupIngress
  - [ ] AWS::Redshift::ClusterSubnetGroup
  - [ ] AWS::Route53::HealthCheck
  - [ ] AWS::Route53::HostedZone
  - [ ] AWS::Route53::RecordSet
  - [ ] AWS::Route53::RecordSetGroup
  - [ ] AWS::S3::Bucket
  - [ ] AWS::S3::BucketPolicy
  - [ ] AWS::SDB::Domain
  - [ ] AWS::SNS::Topic
  - [ ] AWS::SNS::TopicPolicy
  - [ ] AWS::SQS::Queue
  - [ ] AWS::SQS::QueuePolicy
  - [ ] AWS::SSM::Document
  - [ ] AWS::WAF::ByteMatchSet
  - [ ] AWS::WAF::IPSet
  - [ ] AWS::WAF::Rule
  - [ ] AWS::WAF::SqlInjectionMatchSet
  - [ ] AWS::WAF::WebACL
  - [ ] AWS::WorkSpaces::Workspace

## Contributing

I need help in two ways:

### 1. Implementing more resources

Take a look at the `resources/` directory to see existing examples and go nuts. If there's anything complicated or unusual, write a test.

### 2. Testing

I only have limited CloudFormation templates available to test `cfval` against. The more weird and wonderful templates I have the more accurate I can make `cfval`.

The easiest thing you can do is run `cfval` against your weird and wonderful template and tell me what happens. Raise an issue.

Alternatively, [email me (james@jagregory.com)](mailto:james@jagregory.com) your templates! Sanitise/obfuscate them if necessary.
