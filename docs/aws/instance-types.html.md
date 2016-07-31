Instance Types
==============

When you launch an instance, the *instance type* that you specify determines the hardware of the host computer used for your instance. Each instance type offers different compute, memory, and storage capabilities and are grouped in instance families based on these capabilities. Select an instance type based on the requirements of the application or software that you plan to run on your instance.

Amazon EC2 provides each instance with a consistent and predictable amount of CPU capacity, regardless of its underlying hardware.

Amazon EC2 dedicates some resources of the host computer, such as CPU, memory, and instance storage, to a particular instance. Amazon EC2 shares other resources of the host computer, such as the network and the disk subsystem, among instances. If each instance on a host computer tries to use as much of one of these shared resources as possible, each receives an equal share of that resource. However, when a resource is under-utilized, an instance can consume a higher share of that resource while it's available.

Each instance type provides higher or lower minimum performance from a shared resource. For example, instance types with high I/O performance have a larger allocation of shared resources. Allocating a larger share of shared resources also reduces the variance of I/O performance. For most applications, moderate I/O performance is more than enough. However, for applications that require greater or more consistent I/O performance, consider an instance type with higher I/O performance.

**Contents**

-   [Available Instance Types](instance-types.html#AvailableInstanceTypes "Available Instance Types")

-   [Hardware Specifications](instance-types.html#instance-hardware-specs "Hardware Specifications")

-   [Virtualization Types](instance-types.html#instance-virtualization-type "Virtualization Types")

-   [Networking and Storage Features](instance-types.html#instance-networking-storage "Networking and Storage Features")

-   [Instance Limits](instance-types.html#instance-type-limits "Instance Limits")

Available Instance Types
------------------------

Amazon EC2 provides the instance types listed in the following tables.

### Current Generation Instances

For the best performance, we recommend that you use the current generation instance types when you launch new instances. For more information about the current generation instance types, see [Amazon EC2 Instances](http://aws.amazon.com/ec2/instance-types/).

|Instance Family
Current Generation Instance Types|
|:--------------------------------|
|General purpose

`t2.nano` | `t2.micro` | `t2.small` | `t2.medium` | `t2.large` | `m4.large` | `m4.xlarge` | `m4.2xlarge` | `m4.4xlarge` | `m4.10xlarge` | `m3.medium` | `m3.large` | `m3.xlarge` | `m3.2xlarge`|Compute optimized

`c4.large` | `c4.xlarge` | `c4.2xlarge` | `c4.4xlarge` | `c4.8xlarge` | `c3.large` | `c3.xlarge` | `c3.2xlarge` | `c3.4xlarge` | `c3.8xlarge`|Memory optimized

`r3.large` | `r3.xlarge` | `r3.2xlarge` | `r3.4xlarge` | `r3.8xlarge` | `x1.32xlarge`|Storage optimized

`i2.xlarge` | `i2.2xlarge` | `i2.4xlarge` | `i2.8xlarge` | `d2.xlarge` | `d2.2xlarge` | `d2.4xlarge` | `d2.8xlarge`|GPU instances

`g2.2xlarge` | `g2.8xlarge`|

### Previous Generation Instances

Amazon Web Services offers previous generation instances for users who have optimized their applications around these instances and have yet to upgrade. We encourage you to use the latest generation of instances to get the best performance, but we will continue to support these previous generation instances. If you are currently using a previous generation instance, you can see which current generation instance would be a suitable upgrade. For more information, see [Previous Generation Instances](http://aws.amazon.com/ec2/previous-generation/).

|Instance Family
Previous Generation Instance Types|
|:---------------------------------|
|General purpose

`m1.small` | `m1.medium` | `m1.large` | `m1.xlarge`|Compute optimized

`c1.medium` | `c1.xlarge` | `cc2.8xlarge`|Memory optimized

`m2.xlarge` | `m2.2xlarge` | `m2.4xlarge` | `cr1.8xlarge`|Storage optimized

`hi1.4xlarge` | `hs1.8xlarge`|GPU instances

`cg1.4xlarge`|Micro instances

`t1.micro`|

Hardware Specifications
-----------------------

For more information about the hardware specifications for each Amazon EC2 instance type, see [Amazon EC2 Instances](http://aws.amazon.com/ec2/instance-types/).

To determine which instance type best meets your needs, we recommend that you launch an instance and use your own benchmark application. Because you pay by the instance hour, it's convenient and inexpensive to test multiple instance types before making a decision.

Even after you make a decision, if your needs change, you can resize your instance later on. For more information, see [Resizing Your Instance](ec2-instance-resize.html "Resizing Your Instance").

Virtualization Types
--------------------

Each instance type supports one or both of the following types of virtualization: paravirtual (PV) or hardware virtual machine (HVM). The virtualization type of your instance is determined by the AMI that you use to launch it.

For best performance, we recommend that you use an HVM AMI. In addition, HVM AMIs are required to take advantage of enhanced networking. HVM virtualization uses hardware-assist technology provided by the AWS platform. With HVM virtualization, the guest VM runs as if it were on a native hardware platform, except that it still uses PV network and storage drivers for improved performance. For more information, see [Linux AMI Virtualization Types](virtualization_types.html "Linux AMI Virtualization Types").

Networking and Storage Features
-------------------------------

When you select an instance type, this determines which of the following networking and storage features are available:

-   Some instance types are not available in EC2-Classic, so you must launch them in a VPC. By launching an instance in a VPC, you can leverage features that are not available in EC2-Classic, such as enhanced networking, assigning multiple private IP addresses to the instance, and changing the security groups assigned to your instance. For more information, see [Instance Types Available Only in a VPC](using-vpc.html#vpc-only-instance-types "Instance Types Available Only in a VPC").

-   Some instance types support EBS volumes and instance store volumes, while other instance types support only EBS volumes. Some instances that support instance store volumes use solid state drives (SSD) to deliver very high random I/O performance. For more information, see [Storage](Storage.html "Storage").

-   To obtain additional, dedicated capacity for Amazon EBS I/O, you can launch some instance types as EBS–optimized instances. Some instance types are EBS–optimized by default. For more information, see [Amazon EBS–Optimized Instances](EBSOptimized.html "Amazon EBS–Optimized Instances").

-   To optimize your instances for high performance computing (HPC) applications, you can launch some instance types in a placement group. For more information, see [Placement Groups](placement-groups.html "Placement Groups").

-   To get significantly higher packet per second (PPS) performance, lower network jitter, and lower latencies, you can enable enhanced networking for some current generation instance types. For more information, see [Enhanced Networking on Linux](enhanced-networking.html "Enhanced Networking on Linux").

-   The maximum supported MTU varies across instance types. All Amazon EC2 instance types support standard Ethernet V2 1500 MTU frames. All current generation instances support 9001 MTU, or jumbo frames, and some previous generation instances support them as well. For more information, see [Network Maximum Transmission Unit (MTU) for Your EC2 Instance](network_mtu.html "Network Maximum Transmission Unit (MTU) for Your EC2 Instance").

The following table summarizes the networking and storage features supported by the current generation instance types.

| 
VPC only
EBS only
SSD volumes
Placement group
HVM only
Enhanced networking|
|:------------------|
|C3

Yes

Yes

Intel 82599 VF|C4

Yes

Yes

Yes

Yes

Intel 82599 VF|D2

Yes

Yes

Intel 82599 VF|G2

Yes

Yes

Yes|I2

Yes

Yes

Yes

Intel 82599 VF|M3

Yes|M4

Yes

Yes

Yes

Yes

Intel 82599 VF|R3

Yes

Yes

Yes

Intel 82599 VF|T2

Yes

Yes

Yes|X1
Yes

Yes
Yes
Yes
ENA|

Instance Limits
---------------

There is a limit on the total number of instances that you can launch in a region, and there are additional limits on some instance types.

For more information about the default limits, see [How many instances can I run in Amazon EC2?](http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2)

For more information about viewing your current limits or requesting an increase in your current limits, see [Amazon EC2 Service Limits](ec2-resource-limits.html "Amazon EC2 Service Limits").

