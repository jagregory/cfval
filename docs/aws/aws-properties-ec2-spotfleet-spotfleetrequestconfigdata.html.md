Amazon EC2 SpotFleet SpotFleetRequestConfigData
===============================================

`SpotFleetRequestConfigData` is a property of the [AWS::EC2::SpotFleet](aws-resource-ec2-spotfleet.html "AWS::EC2::SpotFleet") resource that defines the configuration of a Spot fleet request.

Syntax
------

``` {.programlisting}
      {
  "AllocationStrategy" : String,
  "ExcessCapacityTerminationPolicy" : String,
  "IamFleetRole" : String,
  "LaunchSpecifications" : [ LaunchSpecifications, ... ],
  "SpotPrice" : String,
  "TargetCapacity" : Integer,
  "TerminateInstancesWithExpiration" : Boolean,
  "ValidFrom" : String,
  "ValidUntil" : String
}
    
```

Properties
----------

 `AllocationStrategy`   
Indicates how to allocate the target capacity across the Spot pools that you specified in the Spot fleet request. For valid values, see [SpotFleetRequestConfigData](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfigData.html) in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: String

 `ExcessCapacityTerminationPolicy`   
Indicates whether running Spot instances are terminated if you decrease the target capacity of the Spot fleet request below the current size of the Spot fleet. For valid values, see [SpotFleetRequestConfigData](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotFleetRequestConfigData.html) in the *Amazon EC2 API Reference*.

*Required*: No

*Type*: String

 `IamFleetRole`   
The Amazon Resource Name (ARN) of an AWS Identity and Access Management (IAM) role that grants the Spot fleet the ability to bid on, launch, and terminate instances on your behalf. For more information, see [Spot Fleet Prerequisites](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-requests.html#spot-fleet-prerequisites) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: Yes

*Type*: String

 `LaunchSpecifications`   
The launch specifications for the Spot fleet request.

*Required*: Yes

*Type*: List of [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications")

 `SpotPrice`   
The bid price per unit hour. For more information, see [How Spot Fleet Works](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: Yes

*Type*: String

 `TargetCapacity`   
The number of units to request for the spot fleet. You can choose to set the target capacity as the number of instances or as a performance characteristic that is important to your application workload, such as vCPUs, memory, or I/O. For more information, see [How Spot Fleet Works](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet.html) in the *Amazon EC2 User Guide for Linux Instances*.

*Required*: Yes

*Type*: Integer

 `TerminateInstancesWithExpiration`   
Indicates whether running Spot instances are terminated when the Spot fleet request expires.

*Required*: No

*Type*: Boolean

 `ValidFrom`   
The start date and time of the request, in UTC format (*`YYYY`*-*`MM`*-*`DD`*T*`HH`*:*`MM`*:*`SS`*Z). By default, Amazon Elastic Compute Cloud (Amazon EC2 ) starts fulfilling the request immediately.

*Required*: No

*Type*: String

 `ValidUntil`   
The end date and time of the request, in UTC format (*`YYYY`*-*`MM`*-*`DD`*T*`HH`*:*`MM`*:*`SS`*Z). After the end date and time, Amazon EC2 doesn't request new Spot instances or enable them to fulfill the request.

*Required*: No

*Type*: String


