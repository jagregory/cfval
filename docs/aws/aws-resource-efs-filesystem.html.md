AWS::EFS::FileSystem
====================

The `AWS::EFS::FileSystem` resource creates a new, empty file system in Amazon Elastic File System (Amazon EFS). You must create a mount target ([AWS::EFS::MountTarget](aws-resource-efs-mounttarget.html "AWS::EFS::MountTarget")) to mount your Amazon EFS file system on an Amazon Elastic Compute Cloud (Amazon EC2) instance. For more information, see the [CreateFileSystem](http://docs.aws.amazon.com/efs/latest/ug/API_CreateFileSystem.html) API in the *Amazon Elastic File System User Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::EFS::FileSystem",
  "Properties" : {
    "FileSystemTags" : [ FileSystemTags, ... ]
  }
}
    
```

Properties
----------

 `FileSystemTags`   
Tags to associate with the file system.

*Required*: No

*Type*: [Amazon Elastic File System FileSystem FileSystemTags](aws-properties-efs-filesystem-filesystemtags.html "Amazon Elastic File System FileSystem FileSystemTags")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource ID, such as `fs-47a2c22e`.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Example
----------------

The following example declares a file system with a tag key `Name` and tag value `TestFileSystem`:

``` {.programlisting}
      "WebServerFileSystem" : {
  "Type" : "AWS::EFS::FileSystem",
  "Properties" : {
    "FileSystemTags" : [
      {
        "Key" : "Name",
        "Value" : "TestFileSystem"
      }
    ]
  }
}
    
```

Additional Resources
--------------------

For a complete sample template, see [Amazon Elastic File System Sample Template](quickref-efs.html "Amazon Elastic File System Sample Template").

