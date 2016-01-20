#!/usr/bin/env ruby
require 'json'

class Ref
end

class Tag
end

class Prop
  def initialize(arity, types)
    @arity = arity
    @types = types
  end
end

class Resource
  def initialize(aws_type:, prop_types:)
    @aws_type = aws_type
    @prop_types = prop_types
  end

  def validate(template_resource)
    true
  end
end

RESOURCES = [
  'AWS::AutoScalingGroup::AutoScalingGroup' => Resource.new(
    aws_type: 'AWS::AutoScaling::AutoScalingGroup',
    prop_types: {
      "AvailabilityZones" => Prop.new(:array, [String, Ref]),
      "DesiredCapacity" => Prop.new(:scala, [String, Integer, Ref]),
      "LaunchConfigurationName" => Prop.new(:scala, [String, Ref]),
      "LoadBalancerNames" => Prop.new(:array, [String, Ref]),
      "MaxSize" => Prop.new(:scala, [String, Integer, Ref]),
      "MinSize" => Prop.new(:scala, [String, Integer, Ref]),
      "Tags" => Prop.new(:array, [Tag]),
      "VPCZoneIdentifier" => Prop.new(:array, [String, Ref]),
    }
  )
]

VALID_RESOURCE_TYPES = [
  'AWS::AutoScaling::AutoScalingGroup',
  'AWS::AutoScaling::LaunchConfiguration',
  'AWS::CloudFront::Distribution',
  'AWS::CloudWatch::Alarm',
  'AWS::EC2::EIP',
  'AWS::EC2::Instance',
  'AWS::EC2::InternetGateway',
  'AWS::EC2::Route',
  'AWS::EC2::RouteTable',
  'AWS::EC2::SecurityGroup',
  'AWS::EC2::SecurityGroupIngress',
  'AWS::EC2::Subnet',
  'AWS::EC2::SubnetRouteTableAssociation',
  'AWS::EC2::VPCGatewayAttachment',
  'AWS::ElastiCache::CacheCluster',
  'AWS::ElastiCache::SubnetGroup',
  'AWS::ElasticBeanstalk::Application',
  'AWS::ElasticBeanstalk::ApplicationVersion',
  'AWS::ElasticBeanstalk::ConfigurationTemplate',
  'AWS::ElasticBeanstalk::Environment',
  'AWS::ElasticLoadBalancing::LoadBalancer',
  'AWS::IAM::InstanceProfile',
  'AWS::IAM::Policy',
  'AWS::IAM::Role',
  'AWS::Route53::RecordSet',
  'AWS::S3::Bucket',
  'AWS::SNS::Topic',
]

# see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html
GETTABLE_ATTRIBUTES = {
  'AWS::CloudFormation::WaitCondition' => [
    'Data'
  ],
  'AWS::CloudFormation::Stack' => [
    'Outputs.NestedStackOutputName', # Yeah this won't work...
  ],
  'AWS::CloudFront::Distribution' => [
    'DomainName'
  ],
  'AWS::Config::ConfigRule' => [
    'Arn',
    'ConfigRuleId',
    'Compliance.Type'
  ],
  'AWS::DirectoryService::MicrosoftAD' => [
    'Alias',
    'DnsIpAddresses'
  ],
  'AWS::DirectoryService::SimpleAD' => [
    'Alias',
    'DnsIpAddresses'
  ],
  'AWS::DynamoDB::Table' => [
    'StreamArn'
  ],
  'AWS::EC2::EIP' => [
    'AllocationId'
  ],
  'AWS::EC2::Instance' => [
    'AvailabilityZone',
    'PrivateDnsName',
    'PublicDnsName',
    'PrivateIp',
    'PublicIp'
  ],
  'AWS::EC2::NetworkInterface' => [
    'PrimaryPrivateIpAddress',
    'SecondaryPrivateIpAddresses'
  ],
  'AWS::EC2::SecurityGroup' => [
    'GroupId'
  ],
  'AWS::EC2::Subnet' => [
    'AvailabilityZone'
  ],
  'AWS::EC2::SubnetNetworkAclAssociation' => [
    'AssociationId'
  ],
  'AWS::EC2::VPC' => [
    'CidrBlock',
    'DefaultNetworkAcl',
    'DefaultSecurityGroup'
  ],
  'AWS::ElastiCache::CacheCluster' => [
    'ConfigurationEndpoint.Address',
    'ConfigurationEndpoint.Port'
  ],
  'AWS::ElastiCache::ReplicationGroup' => [
    'PrimaryEndPoint.Address',
    'PrimaryEndPoint.Port',
    'ReadEndPoint.Addresses',
    'ReadEndPoint.Ports',
    'ReadEndPoint.Addresses.List',
    'ReadEndPoint.Ports.List'
  ],
  'AWS::ElasticBeanstalk::Environment' => [
    'EndpointURL'
  ],
  'AWS::ElasticLoadBalancing::LoadBalancer' => [
    'CanonicalHostedZoneName',
    'CanonicalHostedZoneNameID',
    'DNSName',
    'SourceSecurityGroup.GroupName',
    'SourceSecurityGroup.OwnerAlias'
  ],
  'AWS::IAM::AccessKey' => [
    'SecretAccessKey'
  ],
  'AWS::IAM::Group' => [
    'Arn'
  ],
  'AWS::IAM::InstanceProfile' => [
    'Arn'
  ],
  'AWS::IAM::Role' => [
    'Arn'
  ],
  'AWS::IAM::User' => [
    'Arn'
  ],
  'AWS::Kinesis::Stream' => [
    'Arn'
  ],
  'AWS::Lambda::Function' => [
    'Arn'
  ],
  'AWS::Logs::LogGroup' => [
    'Arn'
  ],
  'AWS::Redshift::Cluster' => [
    'Endpoint.Address',
    'Endpoint.Port'
  ],
  'AWS::RDS::DBCluster' => [
    'Endpoint.Address',
    'Endpoint.Port'
  ],
  'AWS::RDS::DBInstance' => [
    'Endpoint.Address',
    'Endpoint.Port',
  ],
  'AWS::S3::Bucket' => [
    'DomainName',
    'WebsiteURL',
  ],
  'AWS::SNS::Topic' => [
    'TopicName'
  ],
  'AWS::SQS::Queue' => [
    'Arn',
    'QueueName'
  ]
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
    @json = json
    @results = []
  end

  def verify
    verify_resources
    # TODO: verify resource metadata
    # TODO: verify_parameters
    # TODO: verify_mappings
    verify_outputs
    @results
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

  def resources
    @json['Resources']
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

  def verify_resources
    walk(resources, []) do |key,value,context|
      matcher = RESOURCE_MATCHERS[key]
      next unless matcher
      result = [matcher.call(self, value, context)].flatten
      @results = @results + result
    end
  end

  def verify_outputs
    return unless @json['Outputs']

    outputs = @json['Outputs']

    if outputs.is_a? Hash
      outputs.each do |key,value|
        @results += [
          verify_output_description(key, value),
          verify_output_value(key, value),
          # TODO: verify_output_condition(key, value),
        ].flatten
      end
    else
      @results << Result.fail('Output', 'Outputs', 'Not a Hash', [])
    end
  end

  def verify_output_description(logical_id, output)
    context = ['Outputs', logical_id]
    description = output['Description']

    results = []

    if description
      results << Result.fail('Output', logical_id, 'Incorrect type for description', context) unless description.is_a? String
      results << Result.fail('Output', logical_id, 'Description too long', context) if description.length > 4000
    end

    results
  end

  def verify_output_value(logical_id, output)
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
  "Resources.#{(context || []).join('.')}"
end

fail_only = true

begin
  json = JSON.parse STDIN.read
rescue => ex
  puts "Error parsing JSON: #{ex}"
  exit 1
end

results = CfTemplate.new(json).verify

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
