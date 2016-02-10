# New Type vs. Validation

When should you use a whole new `PropertyType` (e.g. like with `AvailabilityZone`) and when should you use a `ValidateFn` on an existing type?

The rule of thumb is:

> New type if the validation is inherently part of the type itself, and a Validaiton when it's a constraint applied by the resource to an otherwise uninteresting type.

For example, an enum like `AvailabilityZone` has a specific set of known values and despite behaving like a string there's nothing else you can set it to. `AutomaticFailover` is a bool with restrictions on when you can set it to true or false, but these are constraints applied by the owning resource not inherent in the `AutomaticFailover` itself.
