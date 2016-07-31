AWS::EC2::NetworkInterfaceAttachment
====================================

Attaches an elastic network interface (ENI) to an Amazon EC2 instance. You can use this resource type to attach additional network interfaces to an instances without interruption.

Syntax
------

``` {.programlisting}
      {
   "Type" : "AWS::EC2::NetworkInterfaceAttachment",
   "Properties" : {
      "DeleteOnTermination": Boolean,
      "DeviceIndex": String,
      "InstanceId": String,
      "NetworkInterfaceId": String
   }
}
    
```

Properties
----------

 `DeleteOnTermination`   
Whether to delete the network interface when the instance terminates. By default, this value is set to `True`.

*Required*: No

*Type*: Boolean.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `DeviceIndex`   
The network interface's position in the attachment order. For example, the first attached network interface has a `DeviceIndex` of `0`.

*Required*: Yes.

*Type*: String.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `InstanceId`   
The ID of the instance to which you will attach the ENI.

*Required*: Yes.

*Type*: String.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

 `NetworkInterfaceId`   
The ID of the ENI that you want to attach.

*Required*: Yes.

*Type*: String.

*Update requires*: [No interruption](using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

Return Values
-------------

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [Ref](intrinsic-function-reference-ref.html "Ref").

Example
-------

**Example Attaching `MyNetworkInterface` to `MyInstance`**

``` {.programlisting}
          "NetworkInterfaceAttachment" : {
    "Type" : "AWS::EC2::NetworkInterfaceAttachment",
        "Properties" : {
            "InstanceId" : {"Ref" : "MyInstance"},
            "NetworkInterfaceId" : {"Ref" : "MyNetworkInterface"},
            "DeviceIndex" : "1" 
        }
}
        
```


