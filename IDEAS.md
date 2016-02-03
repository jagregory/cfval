# Random ideas for validations

These are generally based on what has failed on me in the past.

1. Creating a VPC AutoScalingGroup with multiple Subnets in VPCIdentifier; if you also specify the AvailabilityZones property all the Subnets in VPCIdentifier must reside within the Availability Zones. Recommend either being explicit and fixing your AvailabilityZones list, or remove it entirely and let AWS work it out.
