# cfval: The CloudFormation template validator

> Have you ever waited 15 minutes for CloudFormation to let you know that you forgot to set the Type of a DNS record? Yeah, and that's on a *good day*. Try 45 minutes for your CloudFront Distribution to fail...
>
> After getting very tired of this process, and with a large infrastructure refactor looming, I decided some time could be better spent catching issues sooner in the process. Hence, cfval.

`cfval` is a small tool which validates a CloudFormation JSON template and notifies you of any issues it can find. Missing required properties, properties which conflict with others, `Ref`s to parameters which don't exist or properties of resources, and much more.

```
$ cfval validate my-template.json

Resources.MyLaunchConfiguration.UserData.Ref ... Ref 'CloudInitScript' is not a resource or parameter

1 failure
```

## Installation

For now cfval is only installable via `go get`. This will change once development stabalises and I can push releases out.

`go get -v github.com/jagregory/cfval`

## Known issues

Heaps of resource types aren't supported at the moment. `cfval` currently only supports the resources I've specifically created for my current infrastructure.

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
