Fn::GetAZs
==========

The intrinsic function `Fn::GetAZs` returns an array that lists Availability Zones for a specified region. Because customers have access to different Availability Zones, the intrinsic function `Fn::GetAZs` enables template authors to write templates that adapt to the calling user's access. That way you don't have to hard-code a full list of Availability Zones for a specified region.

Note

For the EC2-Classic platform, the `Fn::GetAZs` function returns all Availability Zones for a region. For the [EC2-VPC](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) platform, the `Fn::GetAZs` function returns only Availability Zones that have a default subnet unless none of the Availability Zones has a default subnet; in that case, all Availability Zones are returned.

IAM permissions

The permissions that you need in order to use the `Fn::GetAZs` function depend on the platform in which you're launching Amazon EC2 instances. For both platforms, you need permissions to the Amazon EC2 `DescribeAvailabilityZones` and `DescribeAccountAttributes` actions. For EC2-VPC, you also need permissions to the Amazon EC2 `DescribeSubnets` action.

Declaration
-----------

"Fn::GetAZs" : "*`region`*"

Parameters
----------

 region   
The name of the region for which you want to get the Availability Zones.

You can use the *`AWS::Region`* pseudo parameter to specify the region in which the stack is created. Specifying an empty string is equivalent to specifying *`AWS::Region`*.

Return Value
------------

The list of Availability Zones for the region.

Examples
--------

``` {.programlisting}
      { "Fn::GetAZs" : "" }
    
```

``` {.programlisting}
      { "Fn::GetAZs" : { "Ref" : "AWS::Region" } }
    
```

``` {.programlisting}
      { "Fn::GetAZs" : "us-east-1" }
    
```

For the previous examples, AWS CloudFormation evaluates Fn::GetAZs to the following array—assuming that the user has created the stack in the us-east-1 region:

``` {.programlisting}
      [ "us-east-1a", "us-east-1b", "us-east-1c" ]
    
```

### Specify a Subnet's Availability Zone

The following example uses `Fn::GetAZs` to specify a subnet's Availability Zone:

``` {.programlisting}
        "mySubnet" : {
  "Type" : "AWS::EC2::Subnet",
  "Properties" : {
    "VpcId" : { "Ref" : "VPC" },
    "CidrBlock" : "10.0.0.0/24",
    "AvailabilityZone" : {
      "Fn::Select" : [ "0", { "Fn::GetAZs" : "" } ]
    }
  }
}
      
```

Supported Functions
-------------------

You can use the `Ref` function in the `Fn::GetAZs` function.

