# cfval: The CloudFormation template validator

[![Build Status](https://travis-ci.org/jagregory/cfval.svg?branch=master)](https://travis-ci.org/jagregory/cfval)

> Have you ever waited 15 minutes for CloudFormation to let you know that you forgot to set the Type of a DNS record? Yeah, and that's on a *good day*. Try 45 minutes for your CloudFront Distribution to fail...

`cfval` is a small tool which validates a CloudFormation JSON template and notifies you of any issues it can find. Missing required properties, properties which conflict with others, `Ref`s to parameters which don't exist or incompatible properties of resources, and much more.

## Usage

```
$ cfval validate my-template.json

Resources.MyLaunchConfiguration.UserData.Ref ... Ref 'CloudInitScript' is not a resource or parameter

Fail: 1 failure, 0 warnings
```

## Installation

For the latest stable release on OSX (still pre-release):

    brew install jagregory/tools/cfval

For other operating systems and/or to use the absoltue latest, cfval is installable from source via `go get`.

    go get -v github.com/jagregory/cfval

## Features

`cfval` aims to identify as many possible issues with your CloudFormation templates *before* you try to run them. Issues are categorised as either:

  * **Failure**: things which are definitely wrong, such as a `Ref` pointing to something which doesn't exist, an unexpected resource property, or a unmistakably wrong value assigned to a property (`"hello world"` to a list property or an EC2 Instance ID).

  * **Warning:** things which are likely wrong, but we aren't certain. These are nearly always type coercion issues (a `String` being assigned to a more specific type like an `VpcID`) or unfortunate AWS documentation issues (a resource returning an ID when the docs suggest a Name). Please report any warnings which seem incorrect.

The main high-level features are:

  * Resource type checks (valid Type attribute)
    * Property validations
    * Unexpected properties
    * Required properties
    * Grouped required properties (e.g. must specify X when Y is specified)
    * Alternate required properties (e.g. must specify X when Y isn't specified)
    * Conflicting properties (e.g. can't specify X when Y is specified)
    * Required when certain property values are specified (e.g. must specify X when Y is hello)
  * `Ref` validations
    * Target exists and is a Resource/Parameter/Pseudo-parameter
    * Target actually is usable in a Ref
    * Value from a Ref is compatible with the property it is being assigned to
  * `GetAtt` validations
    * Target resource exists
    * Attribute is available on target resource
    * Attribute type is compatible with the property it is being assigned to
  * Pseudo-parameter validations (type checking)
  * Various type validations
    * IP addresses
    * CIDR ranges
    * Availability zone names
    * etc...

## Known issues

Most of the major/common AWS resources are now supported by `cfval`; however, there are still quite a few outstanding.

Watch this space. Contributors *very welcome*.

See [AWS Resource support](https://github.com/jagregory/cfval/issues/3) for the current status of Resources.

## Contributing

I need help in two ways:

### 1. Implementing more resources

Take a look at the `resources/` directory to see existing examples and go nuts. If there's anything complicated or unusual, write a test.

### 2. Testing

I only have limited CloudFormation templates available to test `cfval` against. The more weird and wonderful templates I have the more accurate I can make `cfval`.

The easiest thing you can do is run `cfval` against your weird and wonderful template and tell me what happens. Raise an issue.

Alternatively, [email me (james@jagregory.com)](mailto:james@jagregory.com) your templates! Sanitise/obfuscate them if necessary.
