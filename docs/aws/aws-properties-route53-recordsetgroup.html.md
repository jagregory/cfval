AWS::Route53::RecordSetGroup
============================

The AWS::Route53::RecordSetGroup resource creates record sets for a hosted zone. For more information about constraints and values for each property, see [POST CreateHostedZone](http://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateHostedZone.html) for hosted zones and [POST ChangeResourceRecordSet](http://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) for resource record sets.

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::Route53::RecordSetGroup",
   "Properties" : {
      "Comment" : String,
      "HostedZoneId" : String,
      "HostedZoneName" : String,
      "RecordSets" : [ RecordSet1, ... ]
   }
}
    
```

Properties
----------

 `Comment`   
Any comments you want to include about the hosted zone.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `HostedZoneId`   
The ID of the hosted zone.

*Required*: Conditional: You must specify either the *`HostedZoneName`* or *`HostedZoneId`*, but you cannot specify both.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `HostedZoneName`   
The name of the domain for the hosted zone where you want to add the record set.

When you create a stack using an AWS::Route53::RecordSet that specifies *`HostedZoneName`*, AWS CloudFormation attempts to find a hosted zone whose name matches the HostedZoneName. If AWS CloudFormation cannot find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name, AWS CloudFormation will not create the stack.

If you have multiple hosted zones with the same domain name, you must explicitly specify the hosted zone using *`HostedZoneId`*.

*Required*: Conditional. You must specify either the *`HostedZoneName`* or *`HostedZoneId`*, but you cannot specify both.

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `RecordSets`   
List of resource record sets to add.

*Required*: Yes

*Type*:: List of [AWS::Route53::RecordSet](aws-properties-route53-recordset.html "AWS::Route53::RecordSet") objects, as shown in the following example:

``` {.programlisting}
            "RecordSets" : [
  {
    "Name" : "mysite.example.com.",
    "Type" : "CNAME",
    "TTL" : "900",
    "SetIdentifier" : "Frontend One",
    "Weight" : "4",
    "ResourceRecords" : ["example-ec2.amazonaws.com"]
  },
  {
    "Name" : "mysite.example.com.",
    "Type" : "CNAME",
    "TTL" : "900",
    "SetIdentifier" : "Frontend Two",
    "Weight" : "6",
    "ResourceRecords" : ["example-ec2-larger.amazonaws.com"]
  }
]
          
```

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Value
------------

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name. For example:

``` {.programlisting}
      { "Ref": "MyRecordSetGroup" }
    
```

For the resource with the logical ID "MyRecordSetGroup", `Ref` will return the AWS resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Template Examples
-----------------

For AWS::Route53::RecordSetGroup snippets, see [Amazon Route 53 Template Snippets](quickref-route53.html "Amazon Route 53 Template Snippets").

