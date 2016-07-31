Amazon S3 Lifecycle Rule
========================

Describes lifecycle rules for the [Amazon S3 Lifecycle Configuration](aws-properties-s3-bucket-lifecycleconfig.html "Amazon S3 Lifecycle Configuration") property.

Syntax
------

``` {.programlisting}
      {
  "ExpirationDate" : String,
  "ExpirationInDays" : Integer,
  "Id" : String,
  "NoncurrentVersionExpirationInDays" : Integer,
  "NoncurrentVersionTransition (deprecated)" : NoncurrentVersionTransition type,
  "NoncurrentVersionTransitions" : [ NoncurrentVersionTransition type, ... ],
  "Prefix" : String,
  "Status" : String,
  "Transition (deprecated)" : Transition type,
  "Transitions" : [ Transition type, ... ]
}
    
```

Properties
----------

 `ExpirationDate`   
Indicates when objects are deleted from Amazon S3 and Amazon Glacier. The date value must be in ISO 8601 format. The time is always midnight UTC. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: String

 `ExpirationInDays`   
Indicates the number of days after creation when objects are deleted from Amazon S3 and Amazon Glacier. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: Integer

 `Id`   
A unique identifier for this rule. The value cannot be more than 255 characters.

*Required*: No

*Type*: String

 `NoncurrentVersionExpirationInDays`   
For buckets with versioning enabled (or suspended), specifies the time, in days, between when a new version of the object is uploaded to the bucket and when old versions of the object expire. When object versions expire, Amazon S3 permanently deletes them. If you specify a transition and expiration time, the expiration time must be later than the transition time.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: Integer

 `NoncurrentVersionTransition` (deprecated)   
For buckets with versioning enabled (or suspended), specifies when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransitions` property.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: [Amazon S3 Lifecycle Rule NoncurrentVersionTransition](aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html "Amazon S3 Lifecycle Rule NoncurrentVersionTransition")

 `NoncurrentVersionTransitions`   
For buckets with versioning enabled (or suspended), one or more transition rules that specify when non-current objects transition to a specified storage class. If you specify a transition and expiration time, the expiration time must be later than the transition time. If you specify this property, don't specify the `NoncurrentVersionTransition` property.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: List of [Amazon S3 Lifecycle Rule NoncurrentVersionTransition](aws-properties-s3-bucket-lifecycleconfig-rule-noncurrentversiontransition.html "Amazon S3 Lifecycle Rule NoncurrentVersionTransition")

 `Prefix`   
Object key prefix that identifies one or more objects to which this rule applies.

*Required*: No

*Type*: String

 `Status`   
Specify either `Enabled` or `Disabled`. If you specify `Enabled`, Amazon S3 executes this rule as scheduled. If you specify `Disabled`, Amazon S3 ignores this rule.

*Required*: Yes

*Type*: String

 `Transition` (deprecated)   
Specifies when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transitions` property.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: [Amazon S3 Lifecycle Rule Transition](aws-properties-s3-bucket-lifecycleconfig-rule-transition.html "Amazon S3 Lifecycle Rule Transition")

 `Transitions`   
One or more transition rules that specify when an object transitions to a specified storage class. If you specify an expiration and transition time, you must use the same time unit for both properties (either in days or by date). The expiration time must also be later than the transition time. If you specify this property, don't specify the `Transition` property.

*Required*: Conditional. You must specify at least one of the following properties: `ExpirationDate`, `ExpirationInDays`, `NoncurrentVersionExpirationInDays`, `NoncurrentVersionTransition`, `NoncurrentVersionTransitions`, `Transition`, or `Transitions`.

*Type*: List of [Amazon S3 Lifecycle Rule Transition](aws-properties-s3-bucket-lifecycleconfig-rule-transition.html "Amazon S3 Lifecycle Rule Transition")


