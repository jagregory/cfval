Amazon Elastic File System FileSystem FileSystemTags
====================================================

FileSystemTags is a property of the [AWS::EFS::FileSystem](aws-resource-efs-filesystem.html "AWS::EFS::FileSystem") resource that associates key-value pairs with a file system. You can use any of the following Unicode characters for keys and values: letters, digits, whitespace, \_, ., /, =, +, and -.

Syntax
------

``` {.programlisting}
      {
  "Key" : String,
  "Value" : String
}
    
```

Properties
----------

 `Key`   
The key name of the tag. You can specify a value that is from 1 to 128 Unicode characters in length, but you cannot use the prefix `aws:`.

*Required*: No

*Type*: String

 `Value`   
The value of the tag key. You can specify a value that is from 0 to 128 Unicode characters in length.

*Required*: No

*Type*: String


