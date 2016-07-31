AWS Lambda Function Code
========================

`Code` is a property of the [AWS::Lambda::Function](aws-resource-lambda-function.html "AWS::Lambda::Function") resource that enables you to specify the source code of an AWS Lambda (Lambda) function. Source code can be located in a file in an S3 bucket. For `nodejs`, `nodejs4.3`, and `python2.7` runtime environments only, you can provide source code as inline text.

Note

To update a Lambda function whose source code is in an S3 bucket, you must trigger an update by updating the `S3Bucket`, `S3Key`, or `S3ObjectVersion` property. Updating the source code alone doesn't update the function.

Syntax
------

``` {.programlisting}
      {
  "S3Bucket" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String
  "ZipFile" : String
},
    
```

Properties
----------

 `S3Bucket`   
The name of the S3 bucket that contains the source code of your Lambda function. The S3 bucket must be in the same region as the stack.

Note

The `cfn-response` module isn't available for source code stored in S3 buckets. You must write your own functions to send responses.

*Required*: Conditional You must specify both the `S3Bucket` and `S3Key` properties or specify the `ZipFile` property.

*Type*: String

 `S3Key`   
The location and name of the `.zip` file that contains your source code. If you specify this property, you must also specify the `S3Bucket` property.

*Required*: Conditional You must specify both the `S3Bucket` and `S3Key` properties or specify the `ZipFile` property.

*Type*: String

 `S3ObjectVersion`   
If you have S3 versioning enabled, the version ID of the`.zip` file that contains your source code. You can specify this property only if you specify the `S3Bucket` and `S3Key` properties.

*Required*: No

*Type*: String

 `ZipFile`   
For `nodejs`, `nodejs4.3`, and `python2.7` runtime environments, the source code of your Lambda function. You can't use this property with other runtime environments.

You can specify up to 4096 characters. You must precede certain special characters in your source code, such as quotation marks (`"`), newlines (`\n`), and tabs (`\t`), with a backslash (`\`). For a list of special characters, see <http://json.org/>.

If you specify a function that interacts with an AWS CloudFormation custom resource, you don't have to write your own functions to send responses to the custom resource that invoked the function. AWS CloudFormation provides a response module that simplifies sending responses. For more information, see [`cfn-response` Module](aws-properties-lambda-function-code.html#cfn-lambda-function-code-cfnresponsemodule "cfn-response Module").

*Required*: Conditional You must specify both the `S3Bucket` and `S3Key` properties or specify the `ZipFile` property.

*Type*: String

`cfn-response` Module
---------------------

When you use the `ZipFile` property to specify your function's source code and that function interacts with an AWS CloudFormation custom resource, you can load the `cfn-response` module to send responses to those resources. The module contains a `send` method, which sends a [response object](crpg-ref-responses.html "Custom Resource Response Objects") to a custom resource by way of an Amazon S3 pre-signed URL (the `ResponseURL`).

After executing the `send` method, the Lambda function terminates, so anything you write after that method is ignored.

Note

The `cfn-response` module is available only when you use the `ZipFile` property to write your source code. It isn't available for source code stored in S3 buckets. For code in S3 buckets, you must write your own functions to send responses.

### Loading the `cfn-response` Module

For `nodejs` and `nodejs4.3` runtime environments, use the `require()` function to load the `cfn-response` module. For example, the following code example creates a `cfn-response` object with the name `response`:

``` {.programlisting}
        var response = require('cfn-response');
      
```

For `python2.7` environments, use the `import` statement to load the `cfnresponse` module, as shown in the following example:

Note

Use this exact import statement. If you use other variants of the import statement, AWS CloudFormation won't include the response module.

``` {.programlisting}
        import cfnresponse
      
```

### `send` Method Parameters

You can use the following parameters with the `send` method.

 `event`   
The fields in a [custom resource request](crpg-ref-requesttypes.html "Custom Resource Request Types").

 `context`   
An object, specific to Lambda functions, that you can use to specify when the function and any callbacks have completed execution or to access information from within the Lambda execution environment. For more information, see [Programming Model (Node.js)](http://docs.aws.amazon.com/lambda/latest/dg/programming-model.html) in the *AWS Lambda Developer Guide*.

 `responseStatus`   
Whether the function successfully completed. Use the `cfnresponse` module constants to specify the status: `SUCCESS` for successful executions and `FAILED` for failed executions.

 `responseData`   
The `Data` field of a custom resource [response object](crpg-ref-responses.html "Custom Resource Response Objects"). The data is a list of name-value pairs.

 `physicalResourceId`   
Optional. The unique identifier of the custom resource that invoked the function. By default, the module uses the name of the Amazon CloudWatch Logs log stream that is associated with the Lambda function.

### Examples

In the following Node.js example, the inline Lambda function takes an input value and multiplies it by 5. Inline functions are especially useful for smaller functions because they allow you to specify the source code directly in the template instead of creating a package and uploading it to an Amazon S3 bucket. The function uses the `cfn-response` `send` method to send the result back to the custom resource that invoked it.

``` {.programlisting}
        "ZipFile": { "Fn::Join": ["", [
  "var response = require('cfn-response');",
  "exports.handler = function(event, context) {",
  "  var input = parseInt(event.ResourceProperties.Input);",
  "  var responseData = {Value: input * 5};",
  "  response.send(event, context, response.SUCCESS, responseData);",
  "};"
]]}
      
```

As in the preceding example, in the following Python 2.7 example, the inline Lambda function takes an integer value and multiplies it by 5.

``` {.programlisting}
        "ZipFile" : { "Fn::Join" : ["\n", [
  "import json",
  "import cfnresponse",
  "def handler(event, context):",
  "   responseValue = int(event['ResourceProperties']['Input']) * 5",
  "   responseData = {}",
  "   responseData['Data'] = responseValue",
  "   cfnresponse.send(event, context, cfnresponse.SUCCESS, responseData, \"CustomResourcePhysicalID\")"
]]}
      
```

### Module Source Code

The response module source code for the `nodejs` and `nodejs4.3` runtime environments follows. Review it to understand what the module does and for help with implementing your own response functions.

``` {.programlisting}
        /* Copyright 2015 Amazon Web Services, Inc. or its affiliates. All Rights Reserved.
   This file is licensed to you under the AWS Customer Agreement (the "License").
   You may not use this file except in compliance with the License.
   A copy of the License is located at http://aws.amazon.com/agreement/.
   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or implied.
   See the License for the specific language governing permissions and limitations under the License. */
 
exports.SUCCESS = "SUCCESS";
exports.FAILED = "FAILED";
 
exports.send = function(event, context, responseStatus, responseData, physicalResourceId) {
 
    var responseBody = JSON.stringify({
        Status: responseStatus,
        Reason: "See the details in CloudWatch Log Stream: " + context.logStreamName,
        PhysicalResourceId: physicalResourceId || context.logStreamName,
        StackId: event.StackId,
        RequestId: event.RequestId,
        LogicalResourceId: event.LogicalResourceId,
        Data: responseData
    });
 
    console.log("Response body:\n", responseBody);
 
    var https = require("https");
    var url = require("url");
 
    var parsedUrl = url.parse(event.ResponseURL);
    var options = {
        hostname: parsedUrl.hostname,
        port: 443,
        path: parsedUrl.path,
        method: "PUT",
        headers: {
            "content-type": "",
            "content-length": responseBody.length
        }
    };
 
    var request = https.request(options, function(response) {
        console.log("Status code: " + response.statusCode);
        console.log("Status message: " + response.statusMessage);
        context.done();
    });
 
    request.on("error", function(error) {
        console.log("send(..) failed executing https.request(..): " + error);
        context.done();
    });
 
    request.write(responseBody);
    request.end();
}
      
```

The response module source code for the `python2.7` environment follows:

``` {.programlisting}
        #  Copyright 2016 Amazon Web Services, Inc. or its affiliates. All Rights Reserved.
#  This file is licensed to you under the AWS Customer Agreement (the "License").
#  You may not use this file except in compliance with the License.
#  A copy of the License is located at http://aws.amazon.com/agreement/ .
#  This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, express or implied.
#  See the License for the specific language governing permissions and limitations under the License.

from botocore.vendored import requests
import json

SUCCESS = "SUCCESS"
FAILED = "FAILED"

def send(event, context, responseStatus, responseData, physicalResourceId):
    responseUrl = event['ResponseURL']

    print responseUrl

    responseBody = {}
    responseBody['Status'] = responseStatus
    responseBody['Reason'] = 'See the details in CloudWatch Log Stream: ' + context.log_stream_name
    responseBody['PhysicalResourceId'] = physicalResourceId or context.log_stream_name
    responseBody['StackId'] = event['StackId']
    responseBody['RequestId'] = event['RequestId']
    responseBody['LogicalResourceId'] = event['LogicalResourceId']
    responseBody['Data'] = responseData

    json_responseBody = json.dumps(responseBody)
   
    print "Response body:\n" + json_responseBody

    headers = {
        'content-type' : '', 
        'content-length' : str(len(json_responseBody))
    }
    
    try:
        response = requests.put(responseUrl,
                                data=json_responseBody,
                                headers=headers)
        print "Status code: " + response.reason
    except Exception as e:
        print "send(..) failed executing requests.put(..): " + str(e)
      
```

### Related Information

To view a sample template that uses the `ZipFile` property and `cfn-response` module for Node.js, see [Walkthrough: Refer to Resources in Another Stack](walkthrough-custom-resources-lambda-cross-stack-ref.html "Walkthrough: Refer to Resources in Another Stack").

