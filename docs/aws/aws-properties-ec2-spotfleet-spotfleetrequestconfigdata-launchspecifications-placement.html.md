Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications Placement
================================================================================================

`Placement` is a property of the [Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications](aws-properties-ec2-spotfleet-spotfleetrequestconfigdata-launchspecifications.html "Amazon Elastic Compute Cloud SpotFleet SpotFleetRequestConfigData LaunchSpecifications") property that defines the placement group for the Spot instances.

Syntax
------

``` {.programlisting}
      {
  "AvailabilityZone" : String,
  "GroupName" : String
}
    
```

Properties
----------

 `AvailabilityZone`   
The Availability Zone (AZ) of the placement group.

*Required*: No

*Type*: String

 `GroupName`   
The name of the placement group (for cluster instances).

*Required*: No

*Type*: String


