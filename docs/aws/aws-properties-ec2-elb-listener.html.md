ElasticLoadBalancing Listener Property Type
===========================================

The Listener property is an embedded property of the [AWS::ElasticLoadBalancing::LoadBalancer](aws-properties-ec2-elb.html "AWS::ElasticLoadBalancing::LoadBalancer") type.

Syntax
------

``` {.programlisting}
      
{
   "InstancePort" : String,
   "InstanceProtocol" : String,
   "LoadBalancerPort" : String,
   "PolicyNames" :  [ String, ... ],
   "Protocol" : String,
   "SSLCertificateId" : String
}
    
```

Properties
----------

 `InstancePort`   
Specifies the TCP port on which the instance server is listening. This property cannot be modified for the life of the load balancer.

*Required*: Yes

*Type*: String

 `InstanceProtocol`   
Specifies the protocol to use for routing traffic to back-end instances—HTTP, HTTPS, TCP, or SSL. This property cannot be modified for the life of the load balancer.

*Required*: No

*Type*: String

Note

-   If the front-end protocol is HTTP or HTTPS, *`InstanceProtocol`* has to be at the same protocol layer, i.e., HTTP or HTTPS. Likewise, if the front-end protocol is TCP or SSL, *`InstanceProtocol`* has to be TCP or SSL.

-   If there is another listener with the same InstancePort whose *`InstanceProtocol`* is secure, i.e., HTTPS or SSL, the listener's *`InstanceProtocol`* has to be secure, i.e., HTTPS or SSL. If there is another listener with the same InstancePort whose *`InstanceProtocol`* is HTTP or TCP, the listener's *`InstanceProtocol`* must be either HTTP or TCP.

 `LoadBalancerPort`   
Specifies the external load balancer port number. This property cannot be modified for the life of the load balancer.

*Required*: Yes

*Type*: String

 `PolicyNames`   
A list of [ElasticLoadBalancing policy](aws-properties-ec2-elb-policy.html "ElasticLoadBalancing Policy Type") names to associate with the listener. Specify only policies that are compatible with listeners. For more information, see [DescribeLoadBalancerPolicyTypes](http://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DescribeLoadBalancerPolicyTypes.html) in the *Elastic Load Balancing API Reference*.

*Required*: No

*Type*: List of strings

 `Protocol`   
Specifies the load balancer transport protocol to use for routing — HTTP, HTTPS, TCP or SSL. This property cannot be modified for the life of the load balancer.

*Required*: Yes

*Type*: String

 `SSLCertificateId`   
The ARN of the SSL certificate to use. For more information about SSL certificates, see [Managing Server Certificates](http://docs.aws.amazon.com/IAM/latest/UserGuide/ManagingServerCerts.html) in the AWS Identity and Access Management documentation.

*Required*: No

*Type*: String


