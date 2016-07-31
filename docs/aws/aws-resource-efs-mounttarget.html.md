AWS::EFS::MountTarget
=====================

The `AWS::EFS::MountTarget` resource creates a mount target for an Amazon Elastic File System (Amazon EFS) file system ([AWS::EFS::FileSystem](aws-resource-efs-filesystem.html "AWS::EFS::FileSystem")). Use the mount target to mount file systems on Amazon Elastic Compute Cloud (Amazon EC2) instances. For more information, see the [CreateMountTarget](http://docs.aws.amazon.com/efs/latest/ug/API_CreateMountTarget.html) API in the *Amazon Elastic File System User Guide*.

Note

EC2 instances and the mount target that they connect to must be in a VPC with DNS enabled.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EFS::MountTarget",
  "Properties" : {
    "FileSystemId" : String,
    "IpAddress" : String,
    "SecurityGroups" : [ String, ... ],
    "SubnetId" : String
  }
}
    
```

Properties
----------

 `FileSystemId`   
The ID of the file system for which you want to create the mount target.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement). Before updating this property, stop EC2 instances that are using this mount target, and then restart them after the update is complete. This allows the instances to unmount the file system before the mount target is replaced. If you don't stop and restart them, instances or applications that are using those mounts might be disrupted when the mount target is deleted (uncommitted writes might be lost).

 `IpAddress`   
An IPv4 address that is within the address range of the subnet that is specified in the `SubnetId` property. If you don't specify an IP address, Amazon EFS automatically assigns an address that is within the range of the subnet.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement). Before updating this property, stop EC2 instances that are using this mount target, and then restart them after the update is complete. This allows the instances to unmount the file system before the mount target is replaced. If you don't stop and restart them, instances or applications that are using those mounts might be disrupted when the mount target is deleted (uncommitted writes might be lost).

 `SecurityGroups`   
A maximum of five VPC security group IDs that are in the same VPC as the subnet that is specified in the `SubnetId` property. For more information about security groups and mount targets, see [Security](http://docs.aws.amazon.com/efs/latest/ug/security-considerations.html) in the *Amazon Elastic File System User Guide*.

*Required*: Yes

*Type*: List of strings

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `SubnetId`   
The ID of the subnet in which you want to add the mount target.

Note

For each file system, you can create only one mount target per Availability Zone (AZ). All EC2 instances in an AZ share a single mount target for a file system. If you create multiple mount targets for a single file system, do not specify a subnet that is an AZ that already has a mount target associated with the same file system.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement). Before updating this property, stop EC2 instances that are using this mount target and then restart them after the update is complete. That way the instances can unmount the file system before the mount target is replaced. If you don't stop and restart them, instances or applications that are using those mounts might be disrupted when the mount target is deleted (uncommitted writes might be lost).

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource ID, such as `fsmt-55a4413c`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Example
----------------

The following example declares a mount target that is associated with a file system, subnet, and security group, which are all declared in the same template. EC2 instances that are in the same AZ as the mount target can use the mount target to connect to the associated file system. For information about mounting file systems on EC2 instances, see [Mounting File Systems](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs.html) in the *Amazon Elastic File System User Guide*.

``` {.programlisting}
      "MountTarget": {
  "Type": "AWS::EFS::MountTarget",
  "Properties": {
    "FileSystemId": { "Ref": "FileSystem" },
    "SubnetId": { "Ref": "Subnet" },
    "SecurityGroups": [ { "Ref": "MountTargetSecurityGroup" } ]        
  }
}
    
```

Additional Resources
--------------------

For a complete sample template, see [Amazon Elastic File System Sample Template](quickref-efs.html "Amazon Elastic File System Sample Template").

