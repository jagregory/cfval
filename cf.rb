#!/usr/bin/env ruby
require 'json'

class Ref
end

class Join
end

class FindInMap
end

class GetAtt
end

class Cidr
end

class VpcId
end

class KeyName
end

class UnrecognisedProp
  def initialize(resource_type:, name:)
    @resource_type = resource_type
    @name = name
  end

  def validate(context)
    Result.fail(:property, @name, "Unrecognised property for resource #@resource_type", context + [@name])
  end
end

class Prop
  def initialize(definition:, name:, value:)
    @definition = definition
    @name = name
    @value = value
  end

  def validate(context)
    if !@definition.arity_matches_type?(@value)
      Result.fail(:property, @name, 'Invalid type', context + [@name])
    else
      Result.pass(:property, @name, nil, context + [@name])
    end
  end
end

class UnrecognisedPropDefinition
  def initialize(resource_type:, name:)
    @resource_type = resource_type
    @name = name
  end

  def to_prop(name, value)
    UnrecognisedProp.new(resource_type: @resource_type, name: @name)
  end
end

class PropDefinition
  def initialize(name:, arity: :scala, types: [])
    @name = name
    @arity = arity
    @types = types
  end

  def arity_matches_type?(prop)
    (@arity == :array && prop.is_a?(Array)) || true
  end

  def to_prop(name, value)
    Prop.new(definition: self, name: name, value: value)
  end

  def self.unrecognised(resource_type:, name:)
    UnrecognisedPropDefinition.new(
      resource_type: resource_type,
      name: name
    )
  end
end

class ResourceDefinition
  def initialize(aws_type:, prop_types:)
    @aws_type = aws_type
    @prop_types = prop_types
  end

  def to_resource(logical_id:, json:)
    Resource.new(
      definition: self,
      logical_id: logical_id,
      properties: parse_resource_properties(json['Properties'])
    )
  end

  def self.unrecognised(type)
    ResourceDefinition.new(aws_type: type, prop_types: {})
  end

  private
  def parse_resource_properties(json)
    return {} unless json

    json.map do |name,value|
      definition = @prop_types[name] || PropDefinition.unrecognised(resource_type: @aws_type, name: name)
      [name, definition.to_prop(name, value)]
    end
  end
end

class Resource
  def initialize(definition:, logical_id:, properties:)
    @definition = definition
    @logical_id = logical_id
    @properties = properties
  end

  def validate(context)
    @properties
      .map {|name,value| value.validate(context + [@logical_id, 'Properties']) }
      .flatten
  end
end

class AggregateTypeMatcher
  def initialize(matchers)
    @matchers = matchers
  end

  def match?
    raise 'AggregateTypeMatcher not implemented'
  end
end

class TypeMatcher
  def initialize(type)
    @type = type
  end

  def match?
    raise 'Type matcher not implemented'
  end
end

class EnumTypeMatcher
  def initialize(options)
    @options = options
  end

  def match?
    raise 'EnumTypeMatcher not implemented'
  end
end


def built_in_fns_and(*types)
  matchers = (types + [Ref, Join, FindInMap, GetAtt])
    .flatten
    .map {|type| TypeMatcher.new type }
  AggregateTypeMatcher.new matchers
end

def enum_of(options)
  EnumTypeMatcher.new options
end

class ResourceDefinitionBuilder
  def initialize(type)
    @aws_type = type
    @properties = {}
  end

  def property(name, type, required: :optional, update: :no_interruption)
    types = [type]
      .flatten
      .map {|type| type.is_a?(Class) ? TypeMatcher.new(type) : type }
    @properties[name] = PropDefinition.new(name: name, types: types)
  end

  def array_property(name, type, required: :optional, update: :no_interruption)
    types = [type]
      .flatten
      .map {|type| type.is_a?(Class) ? TypeMatcher.new(type) : type }

    @properties[name] = PropDefinition.new(name: name, types: types, arity: :array)
  end

  def ref(nickname, type)
    # TODO: ref return value for resource
  end

  def to_resource_definition
    ResourceDefinition.new(aws_type: @aws_type, prop_types: @properties)
  end
end

def resource(type, &block)
  builder = ResourceDefinitionBuilder.new type
  block.call builder
  builder.to_resource_definition
end

alias_target = resource('AliasTarget') do |res|
  res.property 'DNSName', built_in_fns_and(String), required: :present
  res.property 'EvaluateTargetHealth', built_in_fns_and(TrueClass, FalseClass)
  res.property 'HostedZoneId', built_in_fns_and(String), required: :present
end

geo_location = resource('GeoLocation') do |res|
  res.property 'CountryCode', built_in_fns_and(enum_of(%W{AF AN AS EU OC NA SA})),
    required: -> (resource) { resource.has_prop?('ContinentCode') ? :absent : :present }

  res.property 'ContinentCode', built_in_fns_and(enum_of(%W{
    AO BF BI BJ BW CD CF CG CI CM CV DJ DZ EG ER ET GA GH GM GN GQ GW KE KM LR LS LY MA MG ML MR MU MW MZ NA NE NG RE RW SC SD SH SL SN SO SS ST SZ TD TG TN TZ UG YT ZA ZM ZW
    AQ GS TF
    AE AF AM AZ BD BH BN BT CC CN GE HK ID IL IN IO IQ IR JO JP KG KH KP KR KW KZ LA LB LK MM MN MO MV MY NP OM PH PK PS QA SA SG SY TH TJ TM TR TW UZ VN YE
    AD AL AT AX BA BE BG BY CH CY CZ DE DK EE ES FI FO FR GB GG GI GR HR HU IE IM IS IT JE LI LT LU LV MC MD ME MK MT NL NO PL PT RO RS RU SE SI SJ SK SM UA VA XK
    AG AI AW BB BL BM BQ BS BZ CA CR CU CW DM DO GD GL GP GT HN HT JM KN KY LC MF MQ MS MX NI PA PM PR SV SX TC TT US VC VG VI
    AS AU CK FJ FM GU KI MH MP NC NF NR NU NZ PF PG PN PW SB TK TL TO TV UM VU WF WS
    AR BO BR CL CO EC FK GF GY PE PY SR UY VE
  })),
    required: -> (resource) { resource.has_prop?('CountryCode') ? :absent : :present }

  res.property 'SubdivisionCode', built_in_fns_and(enum_of(%W{AK AL AR AZ CA CO CT DC DE FL GA HI IA ID IL IN KS KY LA MA MD ME MI MN MO MS MT NC ND NE NH NJ NM NV NY OH OK OR PA RI SC SD TN TX UT VA VT WA WI WV WY})),
    required: -> (resource) { resource.has_prop?('CountryCode', 'US') ? :optional : :absent }
end

metrics_collection = resource('MetricsCollection') do |res|
  res.property 'Granularity', built_in_fns_and(String), required: :present
  res.array_property 'Metrics', built_in_fns_and(String)
end

notification_configurations = resource('NotificationConfigurations') do |res|
  res.array_property 'NotificationTypes', built_in_fns_and(enum_of(%W{autoscaling:EC2_INSTANCE_LAUNCH autoscaling:EC2_INSTANCE_LAUNCH_ERROR autoscaling:EC2_INSTANCE_TERMINATE autoscaling:EC2_INSTANCE_TERMINATE_ERROR autoscaling:TEST_NOTIFICATION})),
    required: :present
  res.property 'TopicARN', built_in_fns_and(String), required: :present
end

auto_scaling_tag = resource('AutoScalingTag') do |res|
  res.property 'Key', built_in_fns_and(String), required: :present
  res.property 'Value', built_in_fns_and(String), required: :present
  res.property 'PropagateAtLaunch', built_in_fns_and(TrueClass, FalseClass), required: :present
end

resource_tag = resource('ResourceTag') do |res|
  res.property 'Key', built_in_fns_and(String), required: :present
  res.property 'Value', built_in_fns_and(String), required: :present
end

ebs_block_device = resource('EbsBlockDevice') do |res|
  res.property 'DeleteOnTermination', built_in_fns_and(TrueClass, FalseClass)
  res.property 'Encrypted', built_in_fns_and(TrueClass, FalseClass)
  res.property 'Iops', built_in_fns_and(Integer)
  res.property 'SnapshotId', built_in_fns_and(String)
  res.property 'VolumeSize', built_in_fns_and(Integer), update: :interruptions
  res.property 'VolumeType', built_in_fns_and('String')
end

block_device_mapping = resource('BlockDeviceMapping') do |res|
  res.property 'DeviceName', built_in_fns_and(String), required: :present
  res.property 'Ebs', ebs_block_device,
    required: -> (resource) { resource.has_prop?('VirtualName') ? :absent : :present }
  res.property 'NoDevice', built_in_fns_and(TrueClass, FalseClass)
  res.property 'VirtualName', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('Ebs') ? :absent : :present }
end

security_group_ingress_rule = resource('SecurityGroupIngressRule') do |res|
  res.property 'CidrIp', built_in_fns_and(Cidr),
    required: -> (resource) { resource.has_any_prop?(%W{SourceSecurityGroupName SourceSecurityGroupId}) ? :absent : :present }
  res.property 'FromPort', built_in_fns_and(Integer), required: :present
  res.property 'IpProtocol', built_in_fns_and(enum_of(%W{tcp udp icmp -1})), required: :present
  res.property 'SourceSecurityGroupId', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('CidrIp') ? :absent : :present }
  res.property 'SourceSecurityGroupName', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('CidrIp') ? :absent : :present }
  res.property 'SourceSecurityGroupOwnerId', built_in_fns_and(String)
  res.property 'ToPort', built_in_fns_and(Integer), required: :present
end

security_group_egress_rule = resource('SecurityGroupEgressRule') do |res|
  res.property 'CidrIp', built_in_fns_and(Cidr),
    required: -> (resource) { resource.has_any_prop?(%W{SourceSecurityGroupName SourceSecurityGroupId}) ? :absent : :present }
  res.property 'DestinationSecurityGroupId', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('CidrIp') ? :absent : :present }
  res.property 'FromPort', built_in_fns_and(Integer), required: :present
  res.property 'IpProtocol', built_in_fns_and(enum_of(%W{tcp udp icmp -1})), required: :present
  res.property 'ToPort', built_in_fns_and(Integer), required: :present
end

private_ip_address_specification = resource('PrivateIpAddressSpecification') do |res|
  res.property 'PrivateIpAddress', built_in_fns_and(String), required: :present
  res.property 'Primary', built_in_fns_and(TrueClass, FalseClass), required: :present
end

network_interface = resource('NetworkInterface') do |res|
  res.property 'AssociatePublicIpAddress', built_in_fns_and(TrueClass, FalseClass)
  res.property 'DeleteOnTermination', built_in_fns_and(TrueClass, FalseClass)
  res.property 'Description', built_in_fns_and(String)
  res.property 'DeviceIndex', built_in_fns_and(String), required: :present
  res.array_property 'GroupSet', built_in_fns_and(String)
  res.property 'NetworkInterfaceId', built_in_fns_and(String)
  res.property 'PrivateIpAddress', built_in_fns_and(String)
  res.array_property 'PrivateIpAddresses', private_ip_address_specification
  res.property 'SecondaryPrivateIpAddressCount', built_in_fns_and(Integer)
  res.property 'SubnetId', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('NetworkInterfaceId') ? :optional : :present }
end

ssm_association_parameter = resource('SsmAssociationParameter') do |res|
  res.property 'Key', built_in_fns_and(String), required: :present
  res.array_property 'Value', built_in_fns_and(String), required: :present
end

ssm_association = resource('SsmAssociation') do |res|
  res.array_property 'AssociationParameters', ssm_association_parameter
  res.property 'DocumentName', built_in_fns_and(String), required: :present
end

mount_point = resource('MountPoint') do |res|
  res.property 'Device', built_in_fns_and(String), required: :present
  res.property 'VolumeId', built_in_fns_and(String), required: :present
end

s3_cors_configuration_rule = resource('S3CorsConfigurationRule') do |res|
  res.array_property 'AllowedHeaders', built_in_fns_and(String)
  res.array_property 'AllowedMethods', built_in_fns_and(String), required: :present
  res.array_property 'AllowedOrigins', built_in_fns_and(String), required: :present
  res.array_property 'ExposedHeaders', built_in_fns_and(String)
  res.property 'Id', built_in_fns_and(String)
  res.property 'MaxAge', built_in_fns_and(Integer)
end

s3_cors_configuration = resource('S3CorsConfiguration') do |res|
  res.array_property 'CorsRules', s3_cors_configuration_rule, required: :present
end

s3_lifecycle_rule_noncurrent_version_transition = resource('S3LifecycleRuleNoncurrentVersionTransition') do |res|
  res.property 'StorageClass', built_in_fns_and(enum_of(%W{STANDARD_IA GLACIER})), required: :present
  res.property 'TransitionInDays', built_in_fns_and(Integer), required: :present
end

s3_lifecycle_rule_transition = resource('S3LifecycleRuleTransition') do |res|
  res.property 'StorageClass', built_in_fns_and(enum_of(%W{STANDARD_IA GLACIER})), required: :present
  res.property 'TransitionDate', built_in_fns_and(String)
  res.property 'TransitionInDays', built_in_fns_and(Integer)
end

s3_lifecycle_rule = resource('S3LifecycleRule') do |res|
  res.property 'ExpirationDate', built_in_fns_and(String),
    required: -> (resource) { resource.has_any_prop?(%W{ExpirationInDays NoncurrentVersionExpirationInDays NoncurrentVersionTransition Transition}) ? :optional : :present }
  res.property 'ExpirationInDays', built_in_fns_and(Integer),
    required: -> (resource) { resource.has_any_prop?(%W{ExpirationDate NoncurrentVersionExpirationInDays NoncurrentVersionTransition Transition}) ? :optional : :present }
  res.property 'Id', built_in_fns_and(String)
  res.property 'NoncurrentVersionExpirationInDays', built_in_fns_and(Integer),
    required: -> (resource) { resource.has_any_prop?(%W{ExpirationDate ExpirationInDays NoncurrentVersionTransition Transition}) ? :optional : :present }
  res.property 'NoncurrentVersionTransition', s3_lifecycle_rule_noncurrent_version_transition,
    required: -> (resource) { resource.has_any_prop?(%W{ExpirationDate ExpirationInDays NoncurrentVersionExpirationInDays Transition}) ? :optional : :present }
  res.property 'Prefix', built_in_fns_and(String)
  res.property 'Status', built_in_fns_and(String), required: :present
  res.property 'Transition', s3_lifecycle_rule_transition,
    required: -> (resource) { resource.has_any_prop?(%W{ExpirationDate ExpirationInDays NoncurrentVersionExpirationInDays NoncurrentVersionTransition}) ? :optional : :present }
end

s3_lifecycle_configuration = resource('S3LifecycleConfiguration') do |res|
  res.array_property 'Rules', s3_lifecycle_rule, required: :present
end

s3_logging_configuration = resource('S3LoggingConfiguration') do |res|
  res.property 'DestinationBucketName', built_in_fns_and(String)
  res.property 'LogFilePrefix', built_in_fns_and(String)
end

s3_notification_configuration_config_filter_s3key_rule = resource('S3NotificationConfigurationConfigFilterS3KeyRule') do |res|
  res.property 'Name', built_in_fns_and(String), required: :present
  res.property 'Value', built_in_fns_and(String), required: :present
end

s3_notification_configuration_config_filter_s3key = resource('S3NotificationConfigurationConfigFilterS3Key') do |res|
  res.array_property 'Rules', s3_notification_configuration_config_filter_s3key_rule, required: :present
end

s3_notification_configuration_config_filter = resource('S3NotificationConfigurationConfigFilter') do |res|
  res.property 'S3Key', s3_notification_configuration_config_filter_s3key, required: :present
end

s3_notification_configuration_lambda = resource('S3NotificationConfigurationLambdaConfiguration') do |res|
  res.property 'Event', built_in_fns_and(String), required: :present
  res.property 'Filter', s3_notification_configuration_config_filter
  res.property 'Function', built_in_fns_and(String), required: :present
end

s3_notification_configuration_queue = resource('S3NotificationConfigurationQueueConfiguration') do |res|
  res.property 'Event', built_in_fns_and(String), required: :present
  res.property 'Filter', s3_notification_configuration_config_filter
  res.property 'Queue', built_in_fns_and(String), required: :present
end

s3_notification_configuration_topic = resource('S3NotificationConfigurationTopicConfiguration') do |res|
  res.property 'Event', built_in_fns_and(String), required: :present
  res.property 'Filter', s3_notification_configuration_config_filter
  res.property 'Topic', built_in_fns_and(String), required: :present
end

s3_notification_configuration = resource('S3NotificationConfiguration') do |res|
  res.array_property 'LambdaConfigurations', s3_notification_configuration_lambda
  res.array_property 'QueueConfigurations', s3_notification_configuration_queue
  res.array_property 'TopicConfigurations', s3_notification_configuration_topic
end

s3_replication_configuration_rule_destination = resource('S3ReplicationConfigurationRuleDestination') do |res|
  res.property 'Bucket', built_in_fns_and(String), required: :present
  res.property 'StorageClass', built_in_fns_and(enum_of(%W{STANDARD_IA GLACIER}))
end

s3_replication_configuration_rule = resource('S3ReplicationConfigurationRule') do |res|
  res.property 'Destination', s3_replication_configuration_rule_destination, required: :present
  res.property 'Id', built_in_fns_and(String)
  res.property 'Prefix', built_in_fns_and(String), required: :present
  res.property 'Status', built_in_fns_and(enum_of(%W{Enabled Disabled})), required: :present
end

s3_replication_configuration = resource('S3ReplicationConfiguration') do |res|
  res.property 'Role', built_in_fns_and(String), required: :present
  res.array_property 'Rules', s3_replication_configuration_rule, required: :present
end

s3_versioning_configuration = resource('S3VersioningConfiguration') do |res|
  res.property 'Status', built_in_fns_and(String), required: :present
end

s3_website_configuration_redirect = resource('S3WebsiteConfigurationRedirect') do |res|
  res.property 'HostName', built_in_fns_and(String), required: :present
  res.property 'Protocol', built_in_fns_and(enum_of(%W{http https}))
end

s3_website_configuration_routing_rule_redirect = resource('S3WebsiteConfigurationRoutingRuleRedirect') do |res|
  res.property 'HostName', built_in_fns_and(String)
  res.property 'HttpRedirectCode', built_in_fns_and(String)
  res.property 'Protocol', built_in_fns_and(enum_of(%W{http https}))
  res.property 'ReplaceKeyPrefixWith', built_in_fns_and(String),
    required: -> (resource) { resource.has_property?('ReplaceKeyWith') ? :absent : :optional }
  res.property 'ReplaceKeyWith', built_in_fns_and(String),
    required: -> (resource) { resource.has_property?('ReplaceKeyPrefixWith') ? :absent : :optional }
end

s3_website_configuration_routing_rule_condition = resource('S3WebsiteConfigurationRoutingRuleCondition') do |res|
  res.property 'HttpErrorCodeReturnedEquals', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('KeyPrefixEquals') ? :optional : :present}
  res.property 'KeyPrefixEquals', built_in_fns_and(String),
    required: -> (resource) { resource.has_prop?('HttpErrorCodeReturnedEquals') ? :optional : :present}
end

s3_website_configuration_routing_rule = resource('S3WebsiteConfigurationRoutingRule') do |res|
  res.property 'RedirectRule', s3_website_configuration_routing_rule_redirect, required: :present
  res.property 'RoutingRuleCondition', s3_website_configuration_routing_rule_condition
end

s3_website_configuration = resource('S3WebsiteConfiguration') do |res|
  not_redirect = -> (resource) { resource.has_prop?('RedirectAllRequestsTo') ? :absent : :optional }

  res.property 'ErrorDocument', built_in_fns_and(String), required: not_redirect
  res.property 'IndexDocument', built_in_fns_and(String), required: not_redirect
  res.property 'RedirectAllRequestsTo', s3_website_configuration_redirect
  res.array_property 'RoutingRules', s3_website_configuration_routing_rule, required: not_redirect
end

RESOURCES = {
  'AWS::AutoScaling::AutoScalingGroup' => resource('AWS::AutoScaling::AutoScalingGroup') do |res|
    res.array_property 'AvailabilityZones', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('VPCZoneIdentifier') ? :optional : :present }
    res.property 'Cooldown', built_in_fns_and(String)
    res.property 'DesiredCapacity', built_in_fns_and(String)
    res.property 'HealthCheckGracePeriod', built_in_fns_and(Integer)
    res.property 'HealthCheckType', built_in_fns_and(enum_of(%W{EC2 ELB}))
    res.property 'InstanceId', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('LaunchConfigurationName') ? :absent : :present },
      update: :replace
    res.property 'LaunchConfigurationName', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('InstanceId') ? :absent : :present }
    res.array_property 'LoadBalancerNames', built_in_fns_and(String),
      update: :replace
    res.property 'MaxSize', built_in_fns_and(String),
      required: :present
    res.array_property 'MetricsCollection', metrics_collection
    res.property 'MinSize', built_in_fns_and(String),
      required: :present
    res.array_property 'NotificationConfigurations', notification_configurations
    res.property 'PlacementGroup', built_in_fns_and(String)
    res.array_property 'Tags', auto_scaling_tag
    res.array_property 'TerminationPolicies', built_in_fns_and(String)
    res.array_property 'VPCZoneIdentifier', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('AvailabilityZones') ? :optional : :present },
      update: :interrupt

    res.ref 'Name', String
  end,

  'AWS::AutoScaling::LaunchConfiguration' => resource('AWS::AutoScaling::LaunchConfiguration') do |res|
    res.property 'AssociatePublicIpAddress', built_in_fns_and(TrueClass, FalseClass), update: :replace
    res.array_property 'BlockDeviceMappings', block_device_mapping, update: :replace
    res.property 'ClassicLinkVPCId', built_in_fns_and(String), update: :replace
    res.array_property 'ClassicLinkVPCSecurityGroups', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('ClassicLinkVPCId') ? :present : :absent },
      update: :replace
    res.property 'EbsOptimized', built_in_fns_and(TrueClass, FalseClass), update: :replace
    res.property 'IamInstanceProfile', built_in_fns_and(String), update: :replace
    res.property 'ImageId', built_in_fns_and(String), required: :present, update: :replace
    res.property 'InstanceId', built_in_fns_and(String), update: :replace
    res.property 'InstanceMonitoring', built_in_fns_and(TrueClass, FalseClass), update: :replace
    res.property 'InstanceType', built_in_fns_and(String), required: :present, update: :replace
    res.property 'KernelId', built_in_fns_and(String), update: :replace
    res.property 'KeyName', built_in_fns_and(KeyName), update: :replace
    res.property 'PlacementTenancy', built_in_fns_and(String), update: :replace
    res.property 'RamDiskId', built_in_fns_and(String), update: :replace
    res.array_property 'SecurityGroups', built_in_fns_and(String), update: :replace
    res.property 'SpotPrice', built_in_fns_and(String), update: :replace
    res.property 'UserData', built_in_fns_and(String), update: :replace
  end,

  'AWS::EC2::EIP' => resource('AWS::EC2::EIP') do |res|
    res.property 'InstanceId', built_in_fns_and(String)
    res.property 'Domain', built_in_fns_and(String), update: :replace
  end,

  'AWS::EC2::Instance' => resource('AWS::EC2::Instance') do |res|
    is_instance_store = lambda do |resource|
      return false unless resource.has_prop? 'BlockDeviceMappings'
      resource.properties['BlockDeviceMappings'].any? { |mapping| mapping.has_prop? 'VirtualName' }
    end
    is_ebs = lambda do |resource|
      return false unless resource.has_prop? 'BlockDeviceMappings'
      resource.properties['BlockDeviceMappings'].any? { |mapping| mapping.has_prop? 'Ebs' }
    end
    replace_if_ebs_interrupt_if_instance_store = -> (resource) { is_instance_store.call(resource) ? :interrupt : is_ebs.call(resource) ? :replace : :no_interruption }

    res.property 'AvailabilityZone', built_in_fns_and(String), update: :replace
    res.array_property 'BlockDeviceMappings', block_device_mapping, update: :replace
    res.property 'DisableApiTermination', built_in_fns_and(TrueClass, FalseClass)
    res.property 'EbsOptimized', built_in_fns_and(TrueClass, FalseClass), update: replace_if_ebs_interrupt_if_instance_store
    res.property 'IamInstanceProfile', built_in_fns_and(String), update: :replace
    res.property 'ImageId', built_in_fns_and(String), required: :present, update: :replace
    res.property 'InstanceInitiatedShutdownBehavior', built_in_fns_and(String)
    res.property 'InstanceType', built_in_fns_and(String), update: replace_if_ebs_interrupt_if_instance_store
    res.property 'KernelId', built_in_fns_and(String), update: replace_if_ebs_interrupt_if_instance_store
    res.property 'KeyName', built_in_fns_and(KeyName), update: :replace
    res.property 'Monitoring', built_in_fns_and(TrueClass, FalseClass)
    res.array_property 'NetworkInterfaces', network_interface, update: :replace
    res.property 'PlacementGroupName', built_in_fns_and(String), update: :replace
    res.property 'PrivateIpAddress', built_in_fns_and(String), update: :replace
    res.property 'RamdiskId', built_in_fns_and(String), update: replace_if_ebs_interrupt_if_instance_store
    res.array_property 'SecurityGroupIds', built_in_fns_and(String)
    res.array_property 'SecurityGroups', built_in_fns_and(String), update: :replace
    res.property 'SourceDestCheck', built_in_fns_and(TrueClass, FalseClass)
    res.array_property 'SsmAssociations', ssm_association
    res.property 'SubnetId', built_in_fns_and(String), update: :replace
    res.array_property 'Tags', resource_tag
    res.property 'Tenancy', built_in_fns_and(String), update: :replace
    res.property 'UserData', built_in_fns_and(String),
      update: -> (resource) { is_ebs.call(resource) ? :interrupt : :no_interruption }
    res.array_property 'Volumes', mount_point
    res.property 'AdditionalInfo', built_in_fns_and(String), update: replace_if_ebs_interrupt_if_instance_store
  end,

  'AWS::EC2::Route' => resource('AWS::EC2::Route') do |res|
    res.property 'DestinationCidrBlock', built_in_fns_and(Cidr), required: :present, update: :replace
    res.property 'GatewayId', built_in_fns_and(String),
      required: -> (resource) { resource.has_any_prop?(%W{InstanceId NetworkInterfaceId VpcPeeringConnectionId}) ? :absent : :present }
    res.property 'InstanceId', built_in_fns_and(String),
      required: -> (resource) { resource.has_any_prop?(%W{GatewayId NetworkInterfaceId VpcPeeringConnectionId}) ? :absent : :present }
    res.property 'NetworkInterfaceId', built_in_fns_and(String),
      required: -> (resource) { resource.has_any_prop?(%W{GatewayId InstanceId VpcPeeringConnectionId}) ? :absent : :present }
    res.property 'RouteTableId', built_in_fns_and(String), required: :present, update: :replacement
    res.property 'VpcPeeringConnectionId', built_in_fns_and(String),
      required: -> (resource) { resource.has_any_prop?(%W{GatewayId InstanceId NetworkInterfaceId}) ? :absent : :present }
  end,

  'AWS::EC2::RouteTable' => resource('AWS::EC2::RouteTable') do |res|
    res.property 'VpcId', VpcId, required: :present, update: :replace
    res.array_property 'Tags', resource_tag
  end,

  'AWS::EC2::SecurityGroup' => resource('AWS::EC2::SecurityGroup') do |res|
    res.property 'GroupDescription', built_in_fns_and(String), required: :present, update: :replace
    res.array_property 'SecurityGroupEgress', security_group_egress_rule
    res.array_property 'SecurityGroupIngress', security_group_ingress_rule
    res.array_property 'Tags', resource_tag
    res.property 'VpcId', built_in_fns_and(String), update: :replace
  end,

  'AWS::Route53::RecordSet' => resource('AWS::Route53::RecordSet') do |res|
    res.property 'AliasTarget', alias_target
    res.property 'Failover', built_in_fns_and(String)
    res.property 'GeoLocation', geo_location
    res.property 'HealthCheckId', built_in_fns_and(String)
    res.property 'HostedZoneId', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('HostedZoneName') ? :absent : :present },
      update: :replace
    res.property 'HostedZoneName', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('HostedZoneId') ? :absent : :present },
      update: :replace
    res.property 'Name', built_in_fns_and(String),
      required: :present
    res.property 'Region', built_in_fns_and(String)
    res.array_property 'ResourceRecords', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('AliasTarget') ? :absent : :present }
    res.property 'SetIdentifier', built_in_fns_and(String),
      required: -> (resource) { resource.has_any_prop?('Weight', 'TTL', 'Failover', 'GeoLocation') ? :present : :absent }
    res.property 'TTL', built_in_fns_and(String, Integer),
      required: -> (resource) { resource.has_prop?('AliasTarget') ? :absent : :present }
    res.property 'Type', built_in_fns_and(enum_of(%W{A AAAA CNAME MX NS PTR SOA SPF SRV TXT})),
      required: :present
    res.property 'Weight', built_in_fns_and(Integer)

    res.ref 'DomainName', String
  end,

  'AWS::EC2::Subnet' => resource('AWS::EC2::Subnet') do |res|
    res.property 'AvailabilityZone', built_in_fns_and(String),
      update: :replace
    res.property 'CidrBlock', built_in_fns_and(Cidr),
      required: :present,
      update: :replace
    res.property 'MapPublicIpOnLaunch', built_in_fns_and(TrueClass, FalseClass)
    res.array_property 'Tags', resource_tag
    res.property 'VpcId', built_in_fns_and(VpcId),
      required: :present,
      update: :replace
  end,

  'AWS::EC2::SubnetRouteTableAssociation' => resource('AWS::EC2::SubnetRouteTableAssociation') do |res|
    res.property 'RouteTableId', built_in_fns_and(String), required: :present
    res.property 'SubnetId', built_in_fns_and(String), required: :present, update: :replace
  end,

  'AWS::EC2::VPCGatewayAttachment' => resource('AWS::EC2::VPCGatewayAttachment') do |res|
    res.property 'InternetGatewayId', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('VpnGatewayId') ? :absent : :present }
    res.property 'VpcId', built_in_fns_and(String),
      required: :present
    res.property 'VpnGatewayId', built_in_fns_and(String),
      required: -> (resource) { resource.has_prop?('InternetGatewayId') ? :absent : :present }
  end,

  'AWS::S3::Bucket' => resource('AWS::S3::Bucket') do |res|
    res.property 'AccessControl', built_in_fns_and(enum_of(%W{AuthenticatedRead AwsExecRead BucketOwnerRead BucketOwnerFullControl LogDeliveryWrite Private PublicRead PublicReadWrite}))
    res.property 'BucketName', built_in_fns_and(String), update: :replace
    res.property 'CorsConfiguration', s3_cors_configuration
    res.property 'LifecycleConfiguration', s3_lifecycle_configuration
    res.property 'LoggingConfiguration', s3_logging_configuration
    res.property 'NotificationConfiguration', s3_notification_configuration
    res.property 'ReplicationConfiguration', s3_replication_configuration
    res.array_property 'Tags', resource_tag
    res.property 'VersioningConfiguration', s3_versioning_configuration
    res.property 'WebsiteConfiguration', s3_website_configuration
  end,
}

def is_valid_attribute_for_resource(resource, attribute)
  possible_attributes = GETTABLE_ATTRIBUTES[resource['Type']] || []

  possible_attributes.include? attribute
end

def assert_cidr(template,cidr,context)
  return unless cidr.is_a? String # skip refs/getatts etc...

  if cidr =~ /^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\/([0-9]|[1-2][0-9]|3[0-2]))$/
    Result.pass :cidrblock, cidr, nil, context
  else
    Result.fail :cidrblock, cidr, "\"#{cidr}\" is not a valid CIDR block", context
  end
end

class Result
  attr_reader :success, :type, :name, :data, :context

  def initialize(success, type, name, data, context)
    @success = success
    @type = type
    @name = name
    @data = data
    @context = context
  end

  def pass?
    @success == :pass
  end

  def fail?
    @success == :fail
  end

  def to_s
    "#@type[#@name]"
  end

  def self.pass(type, name, data, context)
    Result.new(:pass, type, name, data, context)
  end

  def self.fail(type, name, data, context)
    Result.new(:fail, type, name, data, context)
  end
end

PROP_REF_TYPES = {
  'SubnetId' => 'AWS::EC2::Subnet',
  'RouteTableId' => 'AWS::EC2::RouteTable',
  'InstanceId' => 'AWS::EC2::Instance',
  'LaunchConfigurationName' => 'AWS::AutoScaling::LaunchConfiguration',
}

def validate_resource_reference(resource, ref, context)
  matching_prop = PROP_REF_TYPES[context.last]

  if matching_prop && matching_prop != resource['Type']
    Result.fail 'Ref', ref, "\"#{ref}\" is not an #{matching_prop}", context
  else
    Result.pass 'Ref', ref, :resource, context
  end
end

# http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/parameters-section-structure.html
PARAM_TYPES = [
  'String',
  'Number',
  'List<Number>',
  'CommaDelimitedList',
  'AWS::EC2::AvailabilityZone::Name',
  'AWS::EC2::Image::Id',
  'AWS::EC2::Instance::Id',
  'AWS::EC2::KeyPair::KeyName',
  'AWS::EC2::SecurityGroup::GroupName',
  'AWS::EC2::SecurityGroup::Id',
  'AWS::EC2::Subnet::Id',
  'AWS::EC2::Volume::Id',
  'AWS::EC2::VPC::Id',
  'AWS::Route53::HostedZone::Id',
  'List<AWS::EC2::AvailabilityZone::Name>',
  'List<AWS::EC2::Image::Id>',
  'List<AWS::EC2::Instance::Id>',
  'List<AWS::EC2::SecurityGroup::GroupName>',
  'List<AWS::EC2::SecurityGroup::Id>',
  'List<AWS::EC2::Subnet::Id>',
  'List<AWS::EC2::Volume::Id>',
  'List<AWS::EC2::VPC::Id>',
  'List<AWS::Route53::HostedZone::Id>',
]
PARAM_REF_TYPES = {
  'KeyName' => 'AWS::EC2::KeyPair::KeyName',
  'VpcId' => 'AWS::EC2::VPC::Id',
  'AvailabilityZone' => 'AWS::EC2::AvailabilityZone::Name'
}

def validate_parameter_reference(parameter, ref, context)
  matching_prop = PARAM_REF_TYPES[context.last]

  if matching_prop && matching_prop != parameter['Type']
    Result.fail 'Ref', ref, "\"#{ref}\" parameter is not an #{matching_prop}", context
  else
    Result.pass 'Ref', ref, :parameter, context
  end
end

PSEUDO_PARAMETERS = [
  'AWS::AccountId',
  'AWS::NotificationARNs',
  'AWS::NoValue',
  'AWS::Region',
  'AWS::StackId',
  'AWS::StackName',
]

def assert_ref(template, ref, context)
  if (parameter = template.parameters[ref])
    validate_parameter_reference(parameter, ref, context)
  elsif (resource = template.resources[ref])
    validate_resource_reference(resource, ref, context)
  elsif PSEUDO_PARAMETERS.include? ref
    Result.pass 'Ref', ref, :pseudo, context
  else
    Result.fail 'Ref', ref, "\"#{ref}\" is not a known resource, parameter, or pseudo parameter", context
  end
end

def assert_join(template,join,context)
  if join.length < 2
    Result.fail 'Fn::Join', join.to_json, "Join has too few parameters [sep, [val...]]", context
  elsif join.length > 2
    Result.fail 'Fn::Join', join.to_json, "Join has too many parameters [sep, [val...]]", context
  elsif not join[1].is_a? Array
    Result.fail 'Fn::Join', join.to_json, "Second parameter should be an array [sep, [val...]]", context
  else
    Result.pass 'Fn::Join', join.to_json, nil, context
  end
end

def assert_get_att(template,att,context)
  resource,attribute = att
  name = "#{resource || '?'}.#{attribute || '?'}"

  if att.length == 0
    Result.fail 'Fn::GetAtt', name, "GetAtt is missing resource and attribute", context
  elsif att.length == 1
    Result.fail 'Fn::GetAtt', name, "GetAtt is missing attribute", context
  elsif att.length > 2
    Result.fail 'Fn::GetAtt', name, "GetAtt has too many parameters [resource, attribute]", context
  elsif not template.resource_names.include? att[0]
    Result.fail 'Fn::GetAtt', name, "\"#{att[0]}\" is not a known resource", context
  elsif not is_valid_attribute_for_resource(template.resources[att[0]], att[1])
    Result.fail 'Fn::GetAtt', name, "\"#{att[1]}\" is not an attribute of \"#{template.resources[att[0]]['Type']}\"", context
  else
    Result.pass 'Fn::GetAtt', name, nil, context
  end
end

def assert_find_in_map(template,find,context)
  # TODO: FindInMap shouldn't allow GetAtt inside
  name = find.join('.')
  if find.length == 0
    Result.fail 'Fn::FindInMap', name, "Missing MapName, TopLevelKey, and SecondLevelKey", context
  elsif find.length == 1
    Result.fail 'Fn::FindInMap', name, "Missing TopLevelKey, and SecondLevelKey", context
  elsif find.length == 2
    Result.fail 'Fn::FindInMap', name, "Missing SecondLevelKey", context
  elsif find.length > 3
    Result.fail 'Fn::FindInMap', name, "Too many parameters [MapName, TopLevelKey, SecondLevelKey]", context
  else
    map_name, top_level_key, second_level_key = find
    if map_name.is_a?(String) && !template.mappings.keys.include?(map_name)
      Result.fail 'Fn::FindInMap', name, "\"#{map_name}\" is not a known Mapping", context
    elsif top_level_key.is_a?(String) && !template.mappings[map_name].keys.include?(top_level_key)
      Result.fail 'Fn::FindInMap', name, "\"#{top_level_key}\" is not a known key in #{map_name} Mapping", context
    elsif top_level_key.is_a?(String) && second_level_key.is_a?(String) && !template.mappings[map_name][top_level_key].keys.include?(second_level_key)
      Result.fail 'Fn::FindInMap', name, "\"#{second_level_key}\" is not a known key in #{map_name}.#{top_level_key} Mapping", context
    else
      Result.pass 'Fn::FindInMap', name, nil, context
    end
  end
end

RESOURCE_MATCHERS = {
  'Ref' => method(:assert_ref),

  'DependsOn' => lambda do |template,ref,context|
    if template.resource_names.include? ref
      Result.pass 'DependsOn', ref, nil, context
    else
      Result.fail 'DependsOn', ref, "\"#{ref}\" is not a known resource", context
    end
  end,

  'Fn::Join' => method(:assert_join),
  'Fn::GetAtt' => method(:assert_get_att),
  'Fn::FindInMap' => method(:assert_find_in_map),

  'CidrBlock' => method(:assert_cidr),
  'CidrIp' => method(:assert_cidr),
  'DestinationCidrBlock' => method(:assert_cidr),
  'Tags' => lambda do |template,tags,context|
    tags.map do |tag|
      key = tag['Key']
      value = tag['Value']
      if key.nil?
        Result.fail 'Tag', '[unknown]', "Missing Key", context
      elsif key.strip == ''
        Result.fail 'Tag', '[unknown]', "Key is blank or whitespace", context
      elsif key.start_with? 'aws:'
        Result.fail 'Tag', key, "\"#{key}\" cannot begin with aws:, these are reserved", context
      elsif value.nil?
        Result.fail 'Tag', key, "\"#{key}\" is missing value, use blank string if no value is needed", context
      else
        Result.pass 'Tag', key, nil, context
      end
    end
  end,
  'Type' => lambda do |template,type,context|
    if VALID_RESOURCE_TYPES.include? type
      Result.pass 'Type', context.last, nil, context
    else
      Result.fail 'Type', context.last, "\"#{type}\" is not a known resource type", context
    end
  end
}

class CfTemplate
  def initialize(json)
    @resources = parse_resources(json['Resources'])
    @results = []
  end

  def parse_resources(json)
    return {} unless json

    Hash[json.map do |logical_id,resource_json|
      type = resource_json['Type']
      definition = RESOURCES[type] || ResourceDefinition.unrecognised(type)

      [
        logical_id,
        definition.to_resource(logical_id: logical_id, json: resource_json)
      ]
    end]
  end

  def validate
    validate_resources(['Resources'])
    # TODO: validate resource metadata
    # TODO: validate_parameters
    # TODO: validate_mappings
    # validate_outputs
  end

  def parameter_names
    parameters.keys
  end

  def parameters
    @json['Parameters']
  end

  def resource_names
    resources.keys
  end

  def mappings
    @json['Mappings'] || {}
  end

  private
  def walk(hash, context, &block)
    hash.each do |key,value|
      if value.is_a?(Hash)
        walk(value, context + [key], &block)
      else
        # TODO: handle recursing into arrays better
        if value.is_a? Array
          value.each do |v|
            if v.is_a? Hash
              walk(v, context + [key], &block)
            else
              block.call(key, value, context)
            end
          end
        else
          block.call(key, value, context)
        end
      end
    end
  end

  def validate_resources(context)
    @resources
      .map { |_,resource| resource.validate(context) }
      .flatten
  end

  def validate_outputs
    return unless @json['Outputs']

    outputs = @json['Outputs']

    if outputs.is_a? Hash
      outputs.each do |key,value|
        @results += [
          validate_output_description(key, value),
          validate_output_value(key, value),
          # TODO: validate_output_condition(key, value),
        ].flatten
      end
    else
      @results << Result.fail('Output', 'Outputs', 'Not a Hash', [])
    end
  end

  def validate_output_description(logical_id, output)
    context = ['Outputs', logical_id]
    description = output['Description']

    results = []

    if description
      results << Result.fail('Output', logical_id, 'Incorrect type for description', context) unless description.is_a? String
      results << Result.fail('Output', logical_id, 'Description too long', context) if description.length > 4000
    end

    results
  end

  def validate_output_value(logical_id, output)
    context = ['Outputs', logical_id]
    value = output['Value']

    results = []

    if value
      results << assert_ref(self, value['Ref'], context) if is_a_ref value
      results << assert_join(self, value['Fn::Join'], context) if is_a_join value
      results << assert_get_att(self, value['Fn::GetAtt'], context) if is_a_get_att value
      results << assert_find_in_map(self, value['Fn::FindInMap'], context) if is_a_find_in_map value
    else
      results << Result.fail('Output', logical_id, 'Missing Value', context)
    end

    results
  end

  def is_a_ref(obj)
    obj.keys.length == 1 && obj.keys.first == 'Ref'
  end

  def is_a_join(obj)
    obj.keys.length == 1 && obj.keys.first == 'Fn::Join'
  end

  def is_a_get_att(obj)
    obj.keys.length == 1 && obj.keys.first == 'Fn::GetAtt'
  end

  def is_a_find_in_map(obj)
    obj.keys.length == 1 && obj.keys.first == 'Fn::FindInMap'
  end
end

def format_context(context)
  (context || []).join('.')
end

fail_only = true

begin
  json = JSON.parse STDIN.read
rescue => ex
  puts "Error parsing JSON: #{ex}"
  exit 1
end

results = CfTemplate.new(json).validate

max_name_length = results
  .select { |result| fail_only ? result.fail? : true }
  .map { |result| result.to_s.length }.max

results.each do |result|
  next if result.pass? and fail_only

  print ("#{result} " + '.' * (max_name_length + 4))[0, max_name_length + 4]

  if result.pass?
    puts " OK"
  else
    puts " FAIL (#{result.data} see: #{format_context result.context})"
  end
end

puts ""
puts "#{results.length} assertions, #{results.select {|result| result.pass? }.length} passes, #{results.select {|result| result.fail? }.length} failures"
