package infrastracture

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/nhannvt/resume/internal/domain/api_error"
	"github.com/nhannvt/resume/internal/domain/validation"

	go_validator "gopkg.in/go-playground/validator.v8"
)

// validator validates given parameters with a context.
// Validation will be done with validation rule defined in context struct.
type validator struct {
	validationStructs map[string]interface{}
}

// NewValidator create validator instance with validation structs.
// validationStructs is map type contains context structs defined in domain layer.
func NewValidator(validationStructs map[string]interface{}) validation.Validator {
	return &validator{validationStructs}
}

// Validate check given parameters if they are valid in accordance with context struct.
func (v *validator) Validate(params map[string]interface{}, context string) error {
	contextStruct, err := v.getContextStruct(context)
	if err != nil {
		return err
	}

	cs, err := v.mapToContextStruct(params, contextStruct)
	if err != nil {
		return &api_error.ValidationError{
			fmt.Sprintf("Invalid Parameter: %s", err.Error()),
		}
	}

	err = v.checkContainUnkownParam(cs, params)
	if err != nil {
		return err
	}

	return v.doValidate(cs)
}

// mapToContextStruct converts map type parameters to specified context struct.
func (v *validator) mapToContextStruct(m map[string]interface{}, contextStruct interface{}) (interface{}, error) {

	tmp, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	t := reflect.TypeOf(contextStruct)
	s := reflect.New(t).Interface()

	err = json.Unmarshal(tmp, &s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// getContextStruct returns context struct which matches given context name.
func (v *validator) getContextStruct(context string) (interface{}, error) {
	if vs, ok := v.validationStructs[context]; ok {
		return vs, nil
	}

	return nil, &api_error.ValidationError{"Invalid	context"}
}

// checkContainUnkownParams check whethere there is the parameter which is not defined in context struct.
func (v *validator) checkContainUnkownParam(contextStruct interface{}, params map[string]interface{}) error {
	var tmp map[string]interface{}
	inrec, _ := json.Marshal(contextStruct)
	json.Unmarshal(inrec, &tmp)
	for k, _ := range params {
		if _, ok := tmp[k]; !ok {
			return &api_error.ValidationError{
				fmt.Sprintf("Unknown parameter[%s] is given", k),
			}
		}
	}

	return nil
}

// doValidate executes validation with given context Struct which set parameter to be validated.
func (v *validator) doValidate(contextStruct interface{}) error {
	config := &go_validator.Config{TagName: "validate"}
	validator := go_validator.New(config)

	if err := validator.Struct(contextStruct); err != nil {
		// TODO: fix error message so that client can make sure the reason of validation error
		return &api_error.ValidationError{
			fmt.Sprintf("Invalid Parameter: %s", err.Error()),
		}
	}

	return nil
}
