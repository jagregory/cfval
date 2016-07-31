AWS::CloudFormation::WaitCondition
==================================

Important

For Amazon EC2 and Auto Scaling resources, we recommend that you use a CreationPolicy attribute instead of wait conditions. Add a CreationPolicy attribute to those resources and use the cfn-signal helper script to signal when an instance has been successfully created.

You can use a wait condition for situations like the following:

-   To coordinate stack resource creation with configuration actions that are external to the stack creation

-   To track the status of a configuration process

For these situations, we recommend that you associate a [CreationPolicy](aws-attribute-creationpolicy.html "CreationPolicy Attribute") attribute with the wait condition so that you don't have to use a wait condition handle. For more information and an example, see [Creating Wait Conditions in a Template](using-cfn-waitcondition.html "Creating Wait Conditions in a Template"). If you use a CreationPolicy with a wait condition, do not specify any of the wait condition's properties.

Note

If you use the [VPC endpoint](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/vpc-endpoints.html) feature, resources in the VPC that respond to wait conditions must have access to AWS CloudFormation-specific Amazon Simple Storage Service (Amazon S3) buckets. Resources must send wait condition responses to a pre-signed Amazon S3 URL. If they can't send responses to Amazon S3, AWS CloudFormation won't receive a response and the stack operation fails. For more information, see [AWS CloudFormation and VPC Endpoints](cfn-vpce-bucketnames.html "AWS CloudFormation and VPC Endpoints").

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::CloudFormation::WaitCondition",
   "Properties" : {
      "Count" : String,
      "Handle" : String,
      "Timeout" : String
   }
}
    
```

Properties
----------

 `Count`   
The number of success signals that AWS CloudFormation must receive before it continues the stack creation process. When the wait condition receives the requisite number of success signals, AWS CloudFormation resumes the creation of the stack. If the wait condition does not receive the specified number of success signals before the Timeout period expires, AWS CloudFormation assumes that the wait condition has failed and rolls the stack back.

*Required*: No

*Type*: String

*Update requires*: Updates are not supported.

 `Handle`   
A reference to the wait condition handle used to signal this wait condition. Use the `Ref` intrinsic function to specify an [AWS::CloudFormation::WaitConditionHandle](aws-properties-waitconditionhandle.html "AWS::CloudFormation::WaitConditionHandle") resource.

Anytime you add a WaitCondition resource during a stack update, you must associate the wait condition with a new WaitConditionHandle resource. Do not reuse an old wait condition handle that has already been defined in the template. If you reuse a wait condition handle, the wait condition might evaluate old signals from a previous create or update stack command.

*Required*: Yes

*Type*: String

*Update requires*: Updates are not supported.

 `Timeout`   
The length of time (in seconds) to wait for the number of signals that the Count property specifies. `Timeout` is a minimum-bound property, meaning the timeout occurs no sooner than the time you specify, but can occur shortly thereafter. The maximum time that can be specified for this property is 12 hours (43200 seconds).

*Required*: Yes

*Type*: String

*Update requires*: Updates are not supported.

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Data`   
*Returns*: A JSON object that contains the *`UniqueId`* and *`Data`* values from the wait condition signal(s) for the specified wait condition. For more information about wait condition signals, see [Wait Condition Signal JSON Format](using-cfn-waitcondition.html#using-cfn-waitcondition-signaljson "Wait Condition Signal JSON Format").

Example return value for a wait condition with 2 signals:

``` {.programlisting}
              
{ "Signal1" : "Step 1 complete." , "Signal2" : "Step 2 complete." } 
            
```

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

**Example WaitCondition that waits for the desired number of instances in a web server group**

``` {.programlisting}
          
"WebServerGroup" : {
   "Type" : "AWS::AutoScaling::AutoScalingGroup",
   "Properties" : {
      "AvailabilityZones" : { "Fn::GetAZs" : "" },
      "LaunchConfigurationName" : { "Ref" : "LaunchConfig" },
      "MinSize" : "1",
      "MaxSize" : "5",
      "DesiredCapacity" : { "Ref" : "WebServerCapacity" },
      "LoadBalancerNames" : [ { "Ref" : "ElasticLoadBalancer" } ]
   }
},

"WaitHandle" : {
   "Type" : "AWS::CloudFormation::WaitConditionHandle"
},

"WaitCondition" : {
   "Type" : "AWS::CloudFormation::WaitCondition",
   "DependsOn" : "WebServerGroup",
   "Properties" : {
      "Handle"  : { "Ref" : "WaitHandle" },
      "Timeout" : "300",
      "Count"   : { "Ref" : "WebServerCapacity" }
   }
}        
        
```

See Also
--------

-   [Creating Wait Conditions in a Template](using-cfn-waitcondition.html "Creating Wait Conditions in a Template")

-   [DependsOn Attribute](aws-attribute-dependson.html "DependsOn Attribute")


