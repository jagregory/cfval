# Coercion

Most types in CloudFormation are just strings; however, there are many constraints on what those strings can contain in different circumstances. For example, there are `SecurityGroupName` properties and `SecurityGroupId` properties both which are just strings; if you assign a Name to an Id bad things will happen.

## What cfval should try to do

When a string is supplied as a property value, we should validate it to see if it matches the expected type. e.g. `InstanceId` should look like `i-abcd...`

When a `Ref` is used to a Parameter or Pseudo Parameter, we should validate the Parameter type matches the property it is being assigned to where possible (e.g. if a Parameter is a VPCID then we can check it). If there is a Default for a Parameter, we should validate that conforms to any constraints too.

When a `Ref` is used to a Resource, we should validate the ReturnValue of the Resource to ensure it is the correct type for the property it is being assigned to.

When a `Fn::GetAtt` is used, we should validate the Attribute type matches the Property it is being assigned to.

## Strings to everything

Unfortunately, some properties are Strings when they could be something more sophisticated. E.g. `MaxSize` is a String instead of a Number. In these cases we'll need to coerce strings and do our best to validate their contents.
