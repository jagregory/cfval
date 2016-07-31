AWS::S3::Bucket
===============

The AWS::S3::Bucket type creates an Amazon S3 bucket.

You can set a deletion policy for your bucket to control how AWS CloudFormation handles the bucket when the stack is deleted. For Amazon S3 buckets, you can choose to *retain* the bucket or to *delete* the bucket. For more information, see [DeletionPolicy Attribute](aws-attribute-deletionpolicy.html "DeletionPolicy Attribute").

Important

Only Amazon S3 buckets that are empty can be deleted. Deletion will fail for buckets that have contents.

Syntax
------

``` {.programlisting}
      {
  "Type" : "AWS::S3::Bucket",
  "Properties" : {
    "AccessControl" : String,
    "BucketName" : String,
    "CorsConfiguration" : CORS Configuration,
    "LifecycleConfiguration" : Lifecycle Configuration,
    "LoggingConfiguration" : Logging Configuration,
    "NotificationConfiguration" : Notification Configuration,
    "ReplicationConfiguration" : Replication Configuration,
    "Tags" : [ Resource Tag, ... ],
    "VersioningConfiguration" : Versioning Configuration,
    "WebsiteConfiguration" : Website Configuration Type
  }
}
    
```

Properties
----------

 `AccessControl`   
A canned access control list (ACL) that grants predefined permissions to the bucket. For more information about canned ACLs, see [Canned ACLs in the Amazon S3 documentation](http://docs.aws.amazon.com/AmazonS3/latest/dev/CannedACL.html).

*Required*: No

*Type*: String

*Valid values*: `AuthenticatedRead` | `AwsExecRead` | `BucketOwnerRead` | `BucketOwnerFullControl` | `LogDeliveryWrite` | `Private` | `PublicRead` | `PublicReadWrite`

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `BucketName`   
A name for the bucket. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name. For more information, see [Name Type](aws-properties-name.html "Name Type"). The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-).

Important

If you specify a name, you cannot do updates that require this resource to be replaced. You can still do updates that require no or some interruption. If you must replace the resource, specify a new name.

*Required*: No

*Type*: String

*Update requires*: [Replacement](using-cfn-updating-stacks-update-behaviors.html#update-replacement)

 `CorsConfiguration`   
Rules that define cross-origin resource sharing of objects in this bucket. For more information, see [Enabling Cross-Origin Resource Sharing](http://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: No

*Type*: [Amazon S3 Cors Configuration](aws-properties-s3-bucket-cors.html "Amazon S3 Cors Configuration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LifecycleConfiguration`   
Rules that define how Amazon S3 manages objects during their lifetime. For more information, see [Object Lifecycle Management](http://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html) in the *Amazon Simple Storage Service Developer Guide*.

*Required*: No

*Type*: [Amazon S3 Lifecycle Configuration](aws-properties-s3-bucket-lifecycleconfig.html "Amazon S3 Lifecycle Configuration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `LoggingConfiguration`   
Settings that defines where logs are stored.

*Required*: No

*Type*: [Amazon S3 Logging Configuration](aws-properties-s3-bucket-loggingconfig.html "Amazon S3 Logging Configuration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NotificationConfiguration`   
Configuration that defines which Amazon SNS topic to send messages to and what events to report.

*Required*: No

*Type*: [Amazon S3 NotificationConfiguration](aws-properties-s3-bucket-notificationconfig.html "Amazon S3 NotificationConfiguration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `ReplicationConfiguration`   
Configuration for replicating objects in an S3 bucket. To enable replication, you must also enable versioning by using the `VersioningConfiguration` property.

Amazon S3 can store replicated objects in only one destination (S3 bucket). You cannot send replicated objects to multiple S3 buckets.

*Required*: No

*Type*: [Amazon S3 ReplicationConfiguration](aws-properties-s3-bucket-replicationconfiguration.html "Amazon S3 ReplicationConfiguration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `Tags`   
An arbitrary set of tags (key-value pairs) for this Amazon S3 bucket.

*Required*: No

*Type*: [AWS CloudFormation Resource Tags](aws-properties-resource-tags.html "AWS CloudFormation Resource Tags Type")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `VersioningConfiguration`   
Enables multiple variants of all objects in this bucket. You might enable versioning to prevent objects from being deleted or overwritten by mistake or to archive objects so that you can retrieve previous versions of them.

*Required*: No

*Type*: [Amazon S3 Versioning Configuration](aws-properties-s3-bucket-versioningconfig.html "Amazon S3 Versioning Configuration")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `WebsiteConfiguration`   
Information used to configure the bucket as a static website. For more information, see [Hosting Websites on Amazon S3](http://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html).

*Required*: No

*Type*: [Website Configuration Type](aws-properties-s3-websiteconfiguration.html "Amazon S3 Website Configuration Property")

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

Example: mystack-mybucket-kdwwxmddtr2g

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. This section lists the available attributes and sample return values.

 `DomainName`   
Returns the DNS name of the specified bucket.

Example: mystack-mybucket-kdwwxmddtr2g.s3.amazonaws.com

 `WebsiteURL`   
Amazon S3 website endpoint for the specified bucket.

Example: http://mystack-mybucket-kdwwxmddtr2g.s3-website-us-east-1.amazonaws.com/

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](intrinsic-function-reference-getatt.html "Fn::GetAtt").

Examples
--------

**Example Static website configuration with a routing rule**

In this example, AWS::S3::Bucket's `Fn::GetAtt` values are used to provide outputs. The routing rule redirects requests to an Amazon EC2 instance in the event of an HTTP 404 error and inserts a object key prefix `report-404/` in the redirect. For example, if you request a page `ExamplePage.html` and it results in a HTTP 404 error, the request is routed to a page `report-404/ExamplePage.html` on the specified instance. For all other HTTP error codes, `error.html` is returned.

``` {.programlisting}
          "Resources" : {
   "S3Bucket" : {
      "Type" : "AWS::S3::Bucket",
      "Properties" : {
         "AccessControl" : "PublicRead",
         "BucketName" : "PublicBucket",
         "WebsiteConfiguration" : {
            "IndexDocument" : "index.html",
            "ErrorDocument" : "error.html",
            "RoutingRules": [
                {
                    "RoutingRuleCondition": {
                        "HttpErrorCodeReturnedEquals": "404",
                        "KeyPrefixEquals": "out1/"
                    },
                    "RedirectRule": {
                        "HostName": "ec2-11-22-333-44.compute-1.amazonaws.com",
                        "ReplaceKeyPrefixWith": "report-404/"
                    }
                }
            ]
         }
      },
      "DeletionPolicy" : "Retain"
   }
},

"Outputs" : {
   "WebsiteURL" : {
      "Value" : { "Fn::GetAtt" : [ "S3Bucket", "WebsiteURL" ] },
      "Description" : "URL for website hosted on S3"
   },
   "S3BucketSecureURL" : {
      "Value" : { "Fn::Join" : [
         "", [ "https://", { "Fn::GetAtt" : [ "S3Bucket", "DomainName" ] } ]
      ] },
      "Description" : "Name of S3 bucket to hold website content"
   }
}
        
```

**Example Enable cross-origin resource sharing**

The following sample template shows an Amazon S3 bucket with two cross-origin resource sharing rules.

``` {.programlisting}
          {
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicReadWrite",
                "CorsConfiguration": {
                    "CorsRules": [
                        {
                            "AllowedHeaders": [
                                "*"
                            ],
                            "AllowedMethods": [
                                "GET"
                            ],
                            "AllowedOrigins": [
                                "*"
                            ],
                            "ExposedHeaders": [
                                "Date"
                            ],
                            "Id": "myCORSRuleId1",
                            "MaxAge": "3600"
                        },
                        {
                            "AllowedHeaders": [
                                "x-amz-*"
                            ],
                            "AllowedMethods": [
                                "DELETE"
                            ],
                            "AllowedOrigins": [
                                "http://www.example1.com",
                                "http://www.example2.com"
                            ],
                            "ExposedHeaders": [
                                "Connection",
                                "Server",
                                "Date"
                            ],
                            "Id": "myCORSRuleId2",
                            "MaxAge": "1800"
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with CORS enabled."
        }
    }
}
        
```

**Example Manage the lifecycle for Amazon S3 objects**

The following sample template shows an Amazon S3 bucket with a lifecycle configuration rule. The rule applies to all objects with the `glacier` key prefix. The objects are transitioned to Amazon Glacier after one day and deleted after one year.

``` {.programlisting}
          {
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicReadWrite",
                "LifecycleConfiguration": {
                    "Rules": [
                        {
                            "Id": "GlacierRule",
                            "Prefix": "glacier",
                            "Status": "Enabled",
                            "ExpirationInDays": "365",
                            "Transition": {
                                "TransitionInDays": "1",
                                "StorageClass": "Glacier"
                            }
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a lifecycle configuration."
        }
    }
}
        
```

**Example Log access requests for a specific bucket**

The following sample template creates two Amazon S3 buckets. The `LoggingBucket` bucket store the logs from the `S3Bucket` bucket. The logging bucket requires log delivery write permissions in order receive logs from the `S3Bucket` bucket.

``` {.programlisting}
          {
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicRead",
                "LoggingConfiguration": {
                    "DestinationBucketName": {"Ref" : "LoggingBucket"},
                    "LogFilePrefix": "testing-logs"
                }
            }
        },
        "LoggingBucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "LogDeliveryWrite"
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a logging configuration."
        }
    }
}
        
```

**Example Receive bucket notifications to an Amazon SNS topic**

The following sample template shows an Amazon S3 bucket with a notification configuration that sends an event to the specified topic when Amazon S3 has lost all replicas of an object.

``` {.programlisting}
          {
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "S3Bucket": {
            "Type": "AWS::S3::Bucket",
            "Properties": {
                "AccessControl": "PublicReadWrite",
                "NotificationConfiguration": {
                    "TopicConfigurations": [
                        {
                            "Topic": "arn:aws:sns:us-east-1:123456789012:TestTopic",
                            "Event": "s3:ReducedRedundancyLostObject"
                        }
                    ]
                }
            }
        }
    },
    "Outputs": {
        "BucketName": {
            "Value": {
                "Ref": "S3Bucket"
            },
            "Description": "Name of the sample Amazon S3 bucket with a notification configuration."
        }
    }
}
        
```

**Example Replicate objects and store them in another S3 bucket**

The following sample includes two replication rules. Amazon S3 replicates objects with the `MyPrefix` or `MyOtherPrefix` prefixes and stores them in the `my-replication-bucket` bucket, which must be in a different region than the `S3Bucket` bucket.

``` {.programlisting}
          "S3Bucket": {
  "Type": "AWS::S3::Bucket",
  "Properties": {
    "VersioningConfiguration":{
      "Status":"Enabled"
    },
    "ReplicationConfiguration": {
      "Role": "arn:aws:iam::123456789012:role/replication_role",
      "Rules": [
        {
          "Id": "MyRule1",
          "Status": "Enabled",
          "Prefix": "MyPrefix",
          "Destination": {
            "Bucket": "arn:aws:s3:::my-replication-bucket",
            "StorageClass": "STANDARD"
          }
        },
        {
          "Status": "Enabled",
          "Prefix": "MyOtherPrefix",
          "Destination": {
            "Bucket": "arn:aws:s3:::my-replication-bucket"
          }
        }
      ]
    }
  }
}
        
```

For more examples, see [Amazon S3 Template Snippets](quickref-s3.html "Amazon S3 Template Snippets").

See Also
--------

-   [DeletionPolicy Attribute](aws-attribute-deletionpolicy.html "DeletionPolicy Attribute")

-   [Access Control List (ACL) Overview](http://docs.aws.amazon.com/AmazonS3/latest/dev/CannedACL.html) in the *Amazon Simple Storage Service Developer Guide*

-   [Hosting a Static Website on Amazon S3](http://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) in the *Amazon Simple Storage Service Developer Guide*


