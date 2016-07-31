AWS::EC2::SpotFleet
===================

The AWS::EC2::SpotFleet resource creates a request for a collection of Spot instances. The Spot fleet attempts to launch the number of Spot instances to meet the target capacity that you specified. For more information, see [Spot Instances](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html) in the *Amazon EC2 User Guide for Linux Instances*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EC2::SpotFleet",
  "Properties" : {
    "SpotFleetRequestConfigData" : SpotFleetRequestConfigData
  }
}
    
```

Properties
----------

 `SpotFleetRequestConfigData`   
The configuration for a Spot fleet request.

*Required*: Yes

*Type*: [Amazon EC2 SpotFleet SpotFleetRequestConfigData](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata.html "Amazon EC2 SpotFleet SpotFleetRequestConfigData")

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

The following example creates a Spot fleet with two launch specifications. The weighted capacities are the same, so Amazon EC2 launches the same number of instances for each specification. For more information, see [How Spot Fleet Works](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances*.

``` {.programlisting}
      "SpotFleet": {
  "Type": "AWS::EC2::SpotFleet",
  "Properties": {
    "SpotFleetRequestConfigData": {
      "IamFleetRole": { "Ref": "IAMFleetRole" },
      "SpotPrice": "1000",
      "TargetCapacity": { "Ref": "TargetCapacity" },
      "LaunchSpecifications": [
      {
        "EbsOptimized": "false",
        "InstanceType": { "Ref": "InstanceType" },
        "ImageId": { "Fn::FindInMap": [ "AWSRegionArch2AMI", { "Ref": "AWS::Region" },
                     { "Fn::FindInMap": [ "AWSInstanceType2Arch", { "Ref": "InstanceType" }, "Arch" ] }
                   ]},
        "SubnetId": { "Ref": "Subnet1" },
        "WeightedCapacity": "8"
      },
      {
        "EbsOptimized": "true",
        "InstanceType": { "Ref": "InstanceType" },
        "ImageId": { "Fn::FindInMap": [ "AWSRegionArch2AMI", { "Ref": "AWS::Region" },
                     { "Fn::FindInMap": [ "AWSInstanceType2Arch", { "Ref": "InstanceType" }, "Arch" ] }
                   ]},
        "Monitoring": { "Enabled": "true" },
        "SecurityGroups": [ { "GroupId": { "Fn::GetAtt": [ "SG0", "GroupId" ] } } ],
        "SubnetId": { "Ref": "Subnet0" },
        "IamInstanceProfile": { "Arn": { "Fn::GetAtt": [ "RootInstanceProfile", "Arn" ] } },
        "WeightedCapacity": "8"
      }
      ]
    }
  }
}
    
```
