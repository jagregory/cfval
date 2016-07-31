AWS::EC2::SubnetNetworkAclAssociation
=====================================

Associates a subnet with a network ACL.

For more information, go to [ReplaceNetworkAclAssociation](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-ReplaceNetworkAclAssociation.html) in the *Amazon EC2 API Reference*.

Note

The EC2 API Reference refers to the *`SubnetId`* parameter as the *`AssociationId`*.

Syntax
------

``` {.programlisting}
      
"Type" : "AWS::EC2::SubnetNetworkAclAssociation",
"Properties" : {
   "SubnetId" : { String },
   "NetworkAclId" : { String }
}

    
```

Properties
----------

 `SubnetId`   
The ID representing the current association between the original network ACL and the subnet.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `NetworkAclId`   
The ID of the new ACL to associate with the subnet.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `AssociationId`   
Returns the value of this object's [SubnetId](aws-resource-ec2-subnet-network-acl-assoc.html#cfn-ec2-subnetnetworkaclassociation-associationid) property.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Template Examples
-----------------

**Example**

``` {.programlisting}
          
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "mySubnetNetworkAclAssociation" : {
         "Type" : "AWS::EC2::SubnetNetworkAclAssociation",
         "Properties" : {
            "SubnetId" : { "Ref" : "mySubnet" },
            "NetworkAclId" : { "Ref" : "myNetworkAcl" }
         }
      }
   }
}
        
```


