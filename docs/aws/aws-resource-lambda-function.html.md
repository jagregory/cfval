AWS::Lambda::Function
=====================

The `AWS::Lambda::Function` resource creates an AWS Lambda (Lambda) function that can run code in response to events. For more information, see [CreateFunction](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html) in the *AWS Lambda Developer Guide*.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::Lambda::Function",
  "Properties" : {
    "Code" : Code,
    "Description" : String,
    "FunctionName" : String,
    "Handler" : String,
    "MemorySize" : Integer,
    "Role" : String,
    "Runtime" : String,
    "Timeout" : Integer,
    "VpcConfig" : VPCConfig
  }
}
    
```

Properties
----------

 `Code`   
The source code of your Lambda function. You can point to a file in an Amazon Simple Storage Service (Amazon S3) bucket or specify your source code as inline text.

*Required*: Yes

*Type*: [AWS Lambda Function Code](aws-properties-lambda-function-code.html "AWS Lambda Function Code")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Description`   
A description of the function.

*Required*: No

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `FunctionName`   
A name for the function. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the function's name. For more information, see [Name Type](aws-properties-name.html "Name Type").

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Handler`   
The name of the function (within your source code) that Lambda calls to start running your code. For more information, see the [Handler](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html) property in the *AWS Lambda Developer Guide*.

Note

If you specify your source code as inline text by specifying the `ZipFile` property within the `Code` property, specify `index.function_name` as the handler.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `MemorySize`   
The amount of memory, in MB, that is allocated to your Lambda function. Lambda uses this value to proportionally allocate the amount of CPU power. For more information, see [Resource Model](http://docs.aws.amazon.com/lambda/latest/dg/resource-model.html) in the *AWS Lambda Developer Guide*.

Your function use case determines your CPU and memory requirements. For example, a database operation might need less memory than an image processing function. You must specify a value that is greater than or equal to `128`, and it must be a multiple of 64. You cannot specify a size larger than `1536`. The default value is 128 MB.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Role`   
The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) execution role that Lambda assumes when it runs your code to access AWS services.

*Required*: Yes

*Type*: String

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Runtime`   
The runtime environment for the Lambda function that you are uploading. For valid values, see the [Runtime](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) property in the *AWS Lambda Developer Guide*.

*Required*: Yes

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `Timeout`   
The function execution time (in seconds) after which Lambda terminates the function. Because the execution time affects cost, set this value based on the function's expected execution time. By default, `Timeout` is set to `3` seconds.

*Required*: No

*Type*: Integer

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VpcConfig`   
If the Lambda function requires access to resources in a VPC, specify a VPC configuration that Lambda uses to set up an elastic network interface (ENI). The ENI enables your function to connect to other resources in your VPC, but it doesn't provide public Internet access. If your function requires Internet access (for example, to access AWS services that don't have VPC endpoints), configure a Network Address Translation (NAT) instance inside your VPC or use an Amazon Virtual Private Cloud (Amazon VPC) NAT gateway. For more information, see [NAT Gateways](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/vpc-nat-gateway.html) in the Amazon VPC User Guide.

*Required*: No

*Type*: [AWS Lambda Function VPCConfig](aws-properties-lambda-function-vpcconfig.html "AWS Lambda Function VPCConfig")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

In the following sample, the `Ref` function returns the name of the `AMILookUp` function, such as `MyStack-AMILookUp-NT5EUXTNTXXD`.

``` {.programlisting}
        { "Ref": "AMILookUp" }
      
```

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `Arn`   
The ARN of the Lambda function, such as `arn:aws:lambda:us-west-2:123456789012:MyStack-AMILookUp-NT5EUXTNTXXD`.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Example
-------

The following example uses a packaged file in an S3 bucket to create a Lambda function.

``` {.programlisting}
      "AMIIDLookup": {
  "Type": "AWS::Lambda::Function",
  "Properties": {
    "Handler": "index.handler",
    "Role": { "Fn::GetAtt" : ["LambdaExecutionRole", "Arn"] },
    "Code": {
      "S3Bucket": "lambda-functions",
      "S3Key": "amilookup.zip"
    },
    "Runtime": "nodejs",
    "Timeout": "25"
  }
}
    
```

Related Resources
-----------------

For more information about how you can use a Lambda function with AWS CloudFormation custom resources, see [AWS Lambda-backed Custom Resources](template-custom-resources-lambda.html "AWS Lambda-backed Custom Resources").

For a sample template, see [AWS Lambda Template](quickref-lambda.html "AWS Lambda Template").

