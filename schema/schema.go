package schema

import (
	"fmt"
	"strconv"

	"github.com/jagregory/cfval/reporting"
)

//go:generate stringer -type=ValueType

type ValueType int

const (
	TypeEnum ValueType = iota
	TypeString
	TypeBool
	TypeInteger
	TypeMap
)

type ValidateFunc func(interface{}, TemplateResource, []string) (bool, []reporting.Failure)

type Schema struct {
	Array          bool
	Conflicts      []string
	Required       bool
	RequiredIf     []string
	RequiredUnless []string
	Type           interface{}
	ValidateFunc   ValidateFunc
}

func (s Schema) Validate(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if !s.Required && value == nil {
		return true, nil
	} else if s.Required && value == nil {
		return false, []reporting.Failure{reporting.NewFailure("Required property is missing", context)}
	}

	failures := make([]reporting.Failure, 0, 20)
	pass := true

	if s.Array {
		for i, item := range value.([]interface{}) {
			if ok, errs := validateProperty(s, item, tr, append(context, strconv.Itoa(i))); !ok {
				failures = append(failures, errs...)
				pass = false
			}
		}
	} else {
		if ok, errs := validateProperty(s, value, tr, context); !ok {
			failures = append(failures, errs...)
			pass = false
		}
	}

	return pass, failures
}

func validateResourceProperty(r Resource, value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if properties, ok := value.(map[string]interface{}); ok {
		return TemplateResource{tr.Template, r, properties}.Validate(context)
	}

	return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Invalid type %T for nested resource %s", value, r.AwsType), context)}
}

func validateProperty(s Schema, value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if resource, ok := s.Type.(Resource); ok {
		return validateResourceProperty(resource, value, tr, context)
	}

	if ok := validateValueType(s.Type, value, context); !ok {
		if complex, ok := value.(map[string]interface{}); ok {
			return validateBuiltinFns(complex, tr, context)
		}

		return false, []reporting.Failure{reporting.NewInvalidTypeFailure(s.Type, value, context)}
	}

	if s.ValidateFunc != nil {
		return s.ValidateFunc(value, tr, context)
	}

	return true, nil
}

func validateValueType(valueType interface{}, value interface{}, context []string) bool {
	switch valueType {
	case TypeBool:
		if _, ok := value.(bool); ok {
			return true
		}
	case TypeEnum:
		fallthrough
	case TypeString:
		if _, ok := value.(string); ok {
			return true
		}
	case TypeInteger:
		if _, ok := value.(float64); ok {
			return true
		}
	case TypeMap:
		if _, ok := value.(map[string]interface{}); ok {
			return true
		}
	}

	return false
}

func validateBuiltinFns(value map[string]interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if ref, ok := value["Ref"]; ok {
		return validateRef(ref, tr.Template, append(context, "Ref"))
	}

	if join, ok := value["Fn::Join"]; ok {
		return validateJoin(join, tr, append(context, "Fn::Join"))
	}

	if getatt, ok := value["Fn::GetAtt"]; ok {
		return validateGetAtt(getatt, tr, append(context, "Fn::GetAtt"))
	}

	if find, ok := value["Fn::FindInMap"]; ok {
		return validateFindInMap(find, tr, append(context, "Fn::FindInMap"))
	}

	if base64, ok := value["Fn::Base64"]; ok {
		return validateBase64(base64, tr, append(context, "Fn::Base64"))
	}

	return false, []reporting.Failure{reporting.NewFailure("Value is a map but isn't a builtin", context)}
}

var pseudoParameters = map[string]bool{
	"AWS::AccountId":        true,
	"AWS::NotificationARNs": true,
	"AWS::NoValue":          true,
	"AWS::Region":           true,
	"AWS::StackId":          true,
	"AWS::StackName":        true,
}

func validateRef(value interface{}, t *Template, context []string) (bool, []reporting.Failure) {
	if ref, ok := value.(string); ok {
		if _, ok := t.Resources[ref]; ok {
			// ref is to a resource and we've found it
			// TODO: validate resource ref value is correct type for property
			return true, nil
		} else if _, ok := t.Parameters[ref]; ok {
			// ref is to a parameter and we've found it
			// TODO: validate parameter type is correct for property
			return true, nil
		} else if _, ok := pseudoParameters[ref]; ok {
			// ref is to a cloudformation pseudo parameter and we've found it
			// TODO: validate parameter type is correct for property
			return true, nil
		}

		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Ref '%s' is not a resource or parameter", ref), context)}
	}

	return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Ref has invalid value '%s'", value), context)}
}

// TODO: Supported functions within a function
func validateFindInMap(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	find, ok := value.([]interface{})
	if !ok {
		return false, []reporting.Failure{reporting.NewFailure("Options need to be an array", context)}
	}

	if len(find) != 3 {
		return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Options has wrong number of items, expected: 3, actual: %d", len(find)), context)}
	}

	mapName := find[0]
	_, mapNameIsString := mapName.(string)
	if ok, errs := validateProperty(Schema{Type: TypeString}, mapName, tr, append(context, "0")); !ok {
		return false, errs
	}

	if mapNameIsString {
		// map name is a string, so we can do some further interrogation
		// TODO: lookup whether MapName is a valid Map
	}

	topLevelKey := find[1]
	_, topLevelKeyIsString := topLevelKey.(string)
	if ok, errs := validateProperty(Schema{Type: TypeString}, topLevelKey, tr, append(context, "1")); !ok {
		return false, errs
	}

	if mapNameIsString && topLevelKeyIsString {
		// TODO: lookup whether topLevelKey is in mapName
	}

	secondLevelKey := find[2]
	_, secondLevelKeyIsString := secondLevelKey.(string)
	if ok, errs := validateProperty(Schema{Type: TypeString}, secondLevelKey, tr, append(context, "2")); !ok {
		return false, errs
	}

	if mapNameIsString && topLevelKeyIsString && secondLevelKeyIsString {
		// TODO: lookup whether secondLevelKeyIsString is in topLevelKey
	}

	return true, nil
}

func validateBase64(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	return validateProperty(Schema{Type: TypeString}, value, tr, context)
}

func validateJoin(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if items, ok := value.([]interface{}); ok {
		if len(items) != 2 {
			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Join has incorrect number of arguments (expected: 2, actual: %d)", len(items)), context)}
		}

		_, ok := items[0].(string)
		if !ok {
			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Join '%s' is not a valid delimiter", items[0]), context)}
		}

		parts, ok := items[1].([]interface{})
		if !ok {
			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("Join items are not valid: %s", items[1]), context)}
		}

		failures := make([]reporting.Failure, 0, len(parts))
		for i, part := range parts {
			if ok, errs := validateProperty(Schema{Type: TypeString}, part, tr, append(context, "1", strconv.Itoa(i))); !ok {
				failures = append(failures, errs...)
			}
		}
		return len(failures) == 0, failures
	}

	return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("GetAtt has invalid value '%s'", value), context)}
}

func validateGetAtt(value interface{}, tr TemplateResource, context []string) (bool, []reporting.Failure) {
	if items, ok := value.([]interface{}); ok {
		if len(items) != 2 {
			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("GetAtt has incorrect number of arguments (expected: 2, actual: %d)", len(items)), context)}
		}

		if resourceID, ok := items[0].(string); ok {
			if _, ok := tr.Template.Resources[resourceID]; ok {
				if _, ok := items[1].(string); ok {
					// TODO: Check attr is actually a valid attribute for the resource type
					return true, nil
				}
			} else {
				// resource not found
				return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a resource", resourceID), context)}
			}
		} else {
			// resource not a string
			return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("GetAtt '%s' is not a valid resource name", items[0]), context)}
		}
	}

	return false, []reporting.Failure{reporting.NewFailure(fmt.Sprintf("GetAtt has invalid value '%s'", value), context)}
}
