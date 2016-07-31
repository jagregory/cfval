AWS::ElasticLoadBalancing::LoadBalancer
=======================================

The AWS::ElasticLoadBalancing::LoadBalancer type creates a LoadBalancer.

Note

If this resource has a public IP address and is also in a VPC that is defined in the same template, you must use the `DependsOn` attribute to declare a dependency on the VPC-gateway attachment. For more information, see [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute").

Syntax
------

``` {.programlisting}
      
{
   "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
   "Properties": {
      "AccessLoggingPolicy" : AccessLoggingPolicy,
      "AppCookieStickinessPolicy" : [ AppCookieStickinessPolicy, ... ],
      "AvailabilityZones" : [ String, ... ],
      "ConnectionDrainingPolicy" : ConnectionDrainingPolicy,
      "ConnectionSettings" : ConnectionSettings,
      "CrossZone" : Boolean,
      "HealthCheck" : HealthCheck,
      "Instances" : [ String, ... ],
      "LBCookieStickinessPolicy" : [ LBCookieStickinessPolicy, ... ],
      "LoadBalancerName" : String,
      "Listeners" : [ Listener, ... ],
      "Policies" : [ ElasticLoadBalancing Policy, ... ],
      "Scheme" : String,
      "SecurityGroups" : [ Security Group, ... ],
      "Subnets" : [ String, ... ],
      "Tags" : [ Resource Tag, ... ]
   }
}
    
```

Properties
----------

 `AccessLoggingPolicy`   
Captures detailed information for all requests made to your load balancer, such as the time a request was received, client’s IP address, latencies, request path, and server responses.

*Required*: No

*Type*: [Elastic Load Balancing AccessLoggingPolicy](aws-properties-ec2-elb-accessloggingpolicy.html "Elastic Load Balancing AccessLoggingPolicy")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AppCookieStickinessPolicy`   
Generates one or more stickiness policies with sticky session lifetimes that follow that of an application-generated cookie. These policies can be associated only with HTTP/HTTPS listeners.

*Required*: No

*Type*: A list of [AppCookieStickinessPolicy](aws-properties-ec2-elb-AppCookieStickinessPolicy.html "ElasticLoadBalancing AppCookieStickinessPolicy Type") objects.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `AvailabilityZones`   
The Availability Zones in which to create the load balancer. You can specify the `AvailabilityZones` or `Subnets` property, but not both.

Note

For load balancers that are in a VPC, specify the `Subnets` property.

*Required*: No

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you did not have an Availability Zone specified and you are adding one or if you are removing all Availability Zones. Otherwise, update requires [no interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `ConnectionDrainingPolicy`   
Whether deregistered or unhealthy instances can complete all in-flight requests.

*Required*: No

*Type*: [Elastic Load Balancing ConnectionDrainingPolicy](aws-properties-ec2-elb-connectiondrainingpolicy.html "Elastic Load Balancing ConnectionDrainingPolicy")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `ConnectionSettings`   
Specifies how long front-end and back-end connections of your load balancer can remain idle.

*Required*: No

*Type*: [Elastic Load Balancing ConnectionSettings](aws-properties-ec2-elb-connectionsettings.html "Elastic Load Balancing ConnectionSettings")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `CrossZone`   
Whether cross-zone load balancing is enabled for the load balancer. With cross-zone load balancing, your load balancer nodes route traffic to the back-end instances across all Availability Zones. By default the `CrossZone` property is `false`.

*Required*: No

*Type*: Boolean

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HealthCheck`   
Application health check for the instances.

*Required*: No

*Type*: [ElasticLoadBalancing HealthCheck Type](aws-properties-ec2-elb-health-check.html "ElasticLoadBalancing HealthCheck Type").

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you did not have a health check specified and you are adding one or if you are removing a health check. Otherwise, update requires [no interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

 `Instances`   
A list of EC2 instance IDs for the load balancer.

*Required*: No

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LBCookieStickinessPolicy`   
Generates a stickiness policy with sticky session lifetimes controlled by the lifetime of the browser (user-agent), or by a specified expiration period. This policy can be associated only with HTTP/HTTPS listeners.

*Required*: No

*Type*: A list of [LBCookieStickinessPolicy](aws-properties-ec2-elb-LBCookieStickinessPolicy.html "ElasticLoadBalancing LBCookieStickinessPolicy Type") objects.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LoadBalancerName`   
A name for the load balancer. For valid values, see the `LoadBalancerName` parameter for the [`CreateLoadBalancer`](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_CreateLoadBalancer.html) action in the *Elastic Load Balancing API Reference*.

If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the load balancer. The name must be unique within your set of load balancers. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Listeners`   
One or more listeners for this load balancer. Each listener must be registered for a specific port, and you cannot have more than one listener for a given port.

Important

If you update the property values for a listener specified by the `Listeners` property, AWS CloudFormation will delete the existing listener and create a new one with the updated properties. During the time that AWS CloudFormation is performing this action, clients will not be able to connect to the load balancer.

*Required*: Yes

*Type*: A list of [ElasticLoadBalancing Listener Property Type](aws-properties-ec2-elb-listener.html "ElasticLoadBalancing Listener Property Type") objects.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Policies`   
A list of elastic load balancing policies to apply to this elastic load balancer. Specify only back-end server policies. For more information, see [DescribeLoadBalancerPolicyTypes](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DescribeLoadBalancerPolicyTypes.html) in the *Elastic Load Balancing API Reference*.

*Required*: No

*Type*: A list of [ElasticLoadBalancing policy](aws-properties-ec2-elb-policy.html "ElasticLoadBalancing Policy Type") objects.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Scheme`   
For load balancers attached to an Amazon VPC, this parameter can be used to specify the type of load balancer to use. Specify `internal` to create an internal load balancer with a DNS name that resolves to private IP addresses or `internet-facing` to create a load balancer with a publicly resolvable DNS name, which resolves to public IP addresses.

Note

If you specify `internal`, you must specify subnets to associate with the load balancer, not Availability Zones.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `SecurityGroups`   
*Required*: No

*Type*: A list of security groups assigned to your load balancer within your virtual private cloud (VPC).

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Subnets`   
A list of subnet IDs in your virtual private cloud (VPC) to attach to your load balancer. Do not specify multiple subnets that are in the same Availability Zone. You can specify the `AvailabilityZones` or `Subnets` property, but not both.

For more information about using Elastic Load Balancing in a VPC, see [How Do I Use Elastic Load Balancing in Amazon VPC](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/DeveloperGuide/UserScenariosForVPC.html) in the *Elastic Load Balancing Developer Guide*.

*Required*: No

*Type*: List of strings

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement) if you did not have an subnet specified and you are adding one or if you are removing all subnets. Otherwise, update requires [no interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt). To update the load balancer to another subnet that is in the same Availability Zone, you must do two updates. You must first update the load balancer to use a subnet in different Availability Zone. After the update is complete, update the load balancer to use the new subnet that is in the original Availability Zone.

 `Tags`   
An arbitrary set of tags (key-value pairs) for this load balancer.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example, `mystack-myelb-1WQN7BJGDB5YQ`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `CanonicalHostedZoneName`   
The name of the Amazon Route 53 hosted zone that is associated with the load balancer.

Important

If you specify `internal` for the Elastic Load Balancing scheme, use `DNSName` instead. For an `internal` scheme, the load balancer doesn't have a `CanonicalHostedZoneName` value.

Example: `mystack-myelb-15HMABG9ZCN57-1013119603.us-east-1.elb.amazonaws.com`

 `CanonicalHostedZoneNameID`   
The ID of the Amazon Route 53 hosted zone name that is associated with the load balancer.

Example: `Z3DZXE0Q79N41H`

 `DNSName`   
The DNS name for the load balancer.

Example: `mystack-myelb-15HMABG9ZCN57-1013119603.us-east-1.elb.amazonaws.com`

 `SourceSecurityGroup.GroupName`   
The security group that you can use as part of your inbound rules for your load balancer's back-end Amazon EC2 application instances.

Example: `amazon-elb`

 `SourceSecurityGroup.OwnerAlias`   
The owner of the source security group.

Example: `amazon-elb-sg`

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

### A load balancer with a health check and access logs

``` {.programlisting}
        "ElasticLoadBalancer" : {
  "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
  "Properties" : {
    "AvailabilityZones" : { "Fn::GetAZs" : "" },
    "Instances" : [ { "Ref" : "Ec2Instance1" },{ "Ref" : "Ec2Instance2" } ],
    "Listeners" : [ {
      "LoadBalancerPort" : "80",
      "InstancePort" : { "Ref" : "WebServerPort" },
      "Protocol" : "HTTP"
    } ],
    "HealthCheck" : {
      "Target" : {
        "Fn::Join" : [ "", [ "HTTP:", { "Ref" : "WebServerPort" }, "/" ] ]
      },
      "HealthyThreshold" : "3",
      "UnhealthyThreshold" : "5",
      "Interval" : "30",
      "Timeout" : "5"
    },
    "AccessLoggingPolicy": {
      "S3BucketName": {
        "Ref": "S3LoggingBucket"
      },
      "S3BucketPrefix": "MyELBLogs",
      "Enabled": "true",
      "EmitInterval" : "60"
    },
  "DependsOn": "S3LoggingBucketPolicy"
  }
}
      
```

### A load balancer with access logging enabled

The following sample snippet creates an Amazon S3 bucket with a bucket policy that allows the load balancer to store information in the `Logs/AWSLogs/AWS                   account number`/ folder. The load balancer also includes an explicit dependency on the bucket policy, which is required before the load balancer can write to the bucket.

``` {.programlisting}
        "S3LoggingBucket": {
  "Type": "AWS::S3::Bucket"
},
"S3LoggingBucketPolicy": {
  "Type": "AWS::S3::BucketPolicy",
  "Properties": {
    "Bucket": {
      "Ref": "S3LoggingBucket"
    },
    "PolicyDocument": {
      "Version": "2012-10-17",
      "Statement": [ {
        "Sid": "ELBAccessLogs20130930",
        "Effect": "Allow",
        "Resource": {
          "Fn::Join": [
            "",
            [
              "arn:aws:s3:::",
              { "Ref": "S3LoggingBucket" },
              "/",
              "Logs",
              "/AWSLogs/",
              { "Ref": "AWS::AccountId" },
              "/*"
            ]
          ]
        },
        "Principal": "{ "Ref": "ElasticLoadBalancingAccountID" }",
        "Action": [
          "s3:PutObject"
        ]
      } ]
    }
  }
},
"ElasticLoadBalancer": {
  "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
  "Properties": {
    "AvailabilityZones": { "Fn::GetAZs": "" },
    "Listeners": [{
      "LoadBalancerPort": "80",
      "InstancePort": "80",
      "Protocol": "HTTP"
    }],
    "HealthCheck": {
      "Target": "HTTP:80/",
      "HealthyThreshold": "3",
      "UnhealthyThreshold": "5",
      "Interval": "30",
      "Timeout": "5"
    },
    "AccessLoggingPolicy": {
      "S3BucketName": {
        "Ref": "S3LoggingBucket"
      },
      "S3BucketPrefix": "Logs",
      "Enabled": "true",
      "EmitInterval" : "60"
    }
  },
  "DependsOn": "S3LoggingBucketPolicy"
}
      
```

### A load balancer with a connection draining policy

The following snippet enables a connection draining policy that ends connections to a deregistered or unhealthy instance after 60 seconds.

``` {.programlisting}
        "ElasticLoadBalancer" : {
  "Type" : "AWS::ElasticLoadBalancing::LoadBalancer",
  "Properties" : {
    "AvailabilityZones" : { "Fn::GetAZs" : "" },
    "Instances" : [ { "Ref" : "Ec2Instance1" },{ "Ref" : "Ec2Instance2" } ],
    "Listeners": [{
      "LoadBalancerPort": "80",
      "InstancePort": "80",
      "Protocol": "HTTP"
    }],
    "HealthCheck": {
      "Target": "HTTP:80/",
      "HealthyThreshold": "3",
      "UnhealthyThreshold": "5",
      "Interval": "30",
      "Timeout": "5"
    },
    "ConnectionDrainingPolicy": {
      "Enabled" : "true",
      "Timeout" : "60"
    }
  }
}
      
```

### A load balancer with multiple policies

The following snippet creates a load balancer with listeners on port 80 and 443. The snippet applies a proxy on port 80 and a back-end server authentication policy on port 443.

``` {.programlisting}
        "ElasticLoadBalancer": {
  "Type": "AWS::ElasticLoadBalancing::LoadBalancer",
  "Properties": {
    "SecurityGroups" : { "Ref" : "SecurityGroups" },
    "Scheme" : "internet-facing",
    "AvailabilityZones": { "Fn::GetAZs": "" },
    "Listeners": [
      {
        "LoadBalancerPort": "80",
        "InstancePort": "80",
        "Protocol": "TCP",
        "InstanceProtocol" : "TCP"
      },
      {
        "LoadBalancerPort": "443",
        "InstancePort": "443",
        "Protocol": "HTTPS",
        "SSLCertificateId" : { "Ref" : "CertARN" },
        "PolicyNames" : ["MySSLNegotiationPolicy", "MyAppCookieStickinessPolicy"]
      }
    ],
    "Policies" : [
      {
        "PolicyName" : "MySSLNegotiationPolicy",
        "PolicyType" : "SSLNegotiationPolicyType",
        "Attributes" : [
            { "Name" : "Protocol-TLSv1", "Value" : "true" },
            { "Name" : "Protocol-SSLv2", "Value" : "true" },
            { "Name" : "Protocol-SSLv3", "Value" : "false" },
            { "Name" : "DHE-RSA-AES256-SHA", "Value" : "true" }
        ]
      },
      {
        "PolicyName" : "MyAppCookieStickinessPolicy",
        "PolicyType" : "AppCookieStickinessPolicyType",
        "Attributes" : [
          { "Name" : "CookieName", "Value" : "MyCookie" }
        ]
      },
      {
        "PolicyName" : "MyPublicKeyPolicy",
        "PolicyType" : "PublicKeyPolicyType",
        "Attributes" : [
          { "Name" : "PublicKey", "Value" : { "Fn::Join" : [ "\n", [
            "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDh/51Aohx5VrpmlfGHZCzciMBa",
            "fkHve+MQYYJcxmNUKMdsWnz9WtVfKxxWUU7Cfor4lorYmENGCG8FWqCoLDMFs7pN",
            "yGEtpsrlKhzZWtgY1d7eGrUrBil03bI90E2KW0j4qAwGYAC8xixOkNClicojeEz4",
            "f4rr3sUf+ZBSsuMEuwIDAQAB" ] ] }
          }
        ]
      },
      {
        "PolicyName" : "MyBackendServerAuthenticationPolicy",
        "PolicyType" : "BackendServerAuthenticationPolicyType",
        "Attributes" : [
          { "Name" : "PublicKeyPolicyName", "Value" : "MyPublicKeyPolicy" }
        ],
        "InstancePorts" : [ "443" ]
      },
      {
        "PolicyName" : "EnableProxyProtocol",
        "PolicyType" : "ProxyProtocolPolicyType",
        "Attributes" : [
          { "Name"  : "ProxyProtocol", "Value" : "true" }
        ],
        "InstancePorts" : ["80"]
      }
    ]
  }
}
      
```

### Additional Examples

You can view additional examples from the AWS CloudFormation sample template collection: [Sample Templates](cfn-sample-templates.html "Sample Templates").

