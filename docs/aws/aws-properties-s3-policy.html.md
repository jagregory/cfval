AWS::S3::BucketPolicy
=====================

The AWS::S3::BucketPolicy type applies an Amazon S3 bucket policy to an Amazon S3 bucket.

AWS::S3::BucketPolicy Snippet: [Declaring an Amazon S3 Bucket Policy](quickref-iam.html#scenario-bucket-policy "Declaring an Amazon S3 Bucket Policy")

Syntax
------

``` {.programlisting}
      
{
   "Type" : "AWS::S3::BucketPolicy",
   "Properties" : {
      "Bucket" : String,
      "PolicyDocument" : JSON
   }
}     
    
```

Properties
----------

 `Bucket`   
The Amazon S3 bucket that the policy applies to.

*Required*: Yes

*Type*: String

You cannot update this property. If you want to add or remove a bucket from a bucket policy, you must modify your AWS CloudFormation template by creating a new bucket policy resource and removing the old one. Then use the modified template to update your AWS CloudFormation stack.

 `PolicyDocument`   
A policy document containing permissions to add to the specified bucket. For more information, see [Access Policy Language Overview](http://docs.aws.amazon.com/AmazonS3/latest/dev/access-policy-language-overview.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: Yes

*Type*: JSON object

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Examples
--------

**Example Bucket policy that allows GET requests from specific referers**

The following sample is a bucket policy that is attached to the `myExampleBucket` bucket and allows GET requests that originate from `www.example.com` and `example.com`:

``` {.programlisting}
          "SampleBucketPolicy" : {
  "Type" : "AWS::S3::BucketPolicy",
  "Properties" : {
    "Bucket" : {"Ref" : "myExampleBucket"},
    "PolicyDocument": {
      "Statement":[{
        "Action":["s3:GetObject"],
        "Effect":"Allow",
        "Resource": { "Fn::Join" : ["", ["arn:aws:s3:::", { "Ref" : "myExampleBucket" } , "/*" ]]},
        "Principal":"*",
        "Condition":{
          "StringLike":{
            "aws:Referer":[
              "http://www.example.com/*",
              "http://example.com/*"
            ]
          }
        }
      }]
    }
  }
}
        
```


