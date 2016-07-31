AWS::EC2::InternetGateway
=========================

Creates a new Internet gateway in your AWS account. After creating the Internet gateway, you then attach it to a VPC.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::EC2::InternetGateway",
   "Properties" : {
      "Tags" : [ Resource Tag, ... ]
   }
}     
    
```

Properties
----------

 `Tags`   
An arbitrary set of tags (key–value pairs) for this resource.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt).

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

``` {.programlisting}
      
{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myInternetGateway" : {
         "Type" : "AWS::EC2::InternetGateway",
         "Properties" : {
            "Tags" : [ {"Key" : "foo", "Value" : "bar"}]
         }
      }
   }
}     
    
```

Related Information
-------------------

-   [CreateInternetGateway](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-CreateInternetGateway.html) in the *Amazon EC2 API Reference*.

-   Use the [AWS::EC2::VPCGatewayAttachment](aws-resource-ec2-vpc-gateway-attachment.html "AWS::EC2::VPCGatewayAttachment") resource to associate an Internet gateway with a VPC.


