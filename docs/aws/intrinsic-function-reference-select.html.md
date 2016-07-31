Fn::Select
==========

The intrinsic function `Fn::Select` returns a single object from a list of objects by index.

Important

Fn::Select does not check for null values or if the index is out of bounds of the array. Both conditions will result in a stack error, so you should be certain that the index you choose is valid, and that the list contains non-null values.

Declaration
-----------

{ "Fn::Select" : [ *`index`*, *`listOfObjects`* ] }

Parameters
----------

 index   
The index of the object to retrieve. This must be a value from zero to N-1, where N represents the number of elements in the array.

 listOfObjects   
The list of objects to select from. This list must not be null, nor can it have null entries.

Return Value
------------

The selected object.

Examples
--------

``` {.programlisting}
      
{ "Fn::Select" : [ "1", [ "apples", "grapes", "oranges", "mangoes" ] ] }
    
```

This example returns: "grapes".

### Comma-delimited List Parameter Type

You can use `Fn::Select` to select an object from a `CommaDelimitedList` parameter. You might use a `CommaDelimitedList` parameter to combine the values of related parameters, which reduces the total number of parameters in your template. For example, the following parameter specifies a comma-delimited list of three CIDR blocks:

``` {.programlisting}
        "Parameters" : {
  "DbSubnetIpBlocks": {
    "Description": "Comma-delimited list of three CIDR blocks",
    "Type": "CommaDelimitedList",
      "Default": "10.0.48.0/24, 10.0.112.0/24, 10.0.176.0/24"
  }
}
      
```

To specify one of the three CIDR blocks, use `Fn::Select` in the Resources section of the same template, as shown in the following sample snippet:

``` {.programlisting}
        "Subnet0": {
  "Type": "AWS::EC2::Subnet",
    "Properties": {
      "VpcId": { "Ref": "VPC" },
      "CidrBlock": { "Fn::Select" : [ "0", {"Ref": "DbSubnetIpBlocks"} ] }
    }
},
      
```

Supported Functions
-------------------

For the `Fn::Select` index value, you can use the `Ref` and `Fn::FindInMap` functions.

For the `Fn::Select` list of objects, you can use the following functions:

-   `Fn::FindInMap`

-   `Fn::GetAtt`

-   `Fn::GetAZs`

-   `Fn::If`

-   `Ref`


