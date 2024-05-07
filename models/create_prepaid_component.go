package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePrepaidComponent represents a CreatePrepaidComponent struct.
type CreatePrepaidComponent struct {
    PrepaidUsageComponent PrepaidUsageComponent `json:"prepaid_usage_component"`
    AdditionalProperties  map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaidComponent.
// It customizes the JSON marshaling process for CreatePrepaidComponent objects.
func (c CreatePrepaidComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaidComponent object to a map representation for JSON marshaling.
func (c CreatePrepaidComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepaid_usage_component"] = c.PrepaidUsageComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaidComponent.
// It customizes the JSON unmarshaling process for CreatePrepaidComponent objects.
func (c *CreatePrepaidComponent) UnmarshalJSON(input []byte) error {
    var temp createPrepaidComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepaid_usage_component")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.PrepaidUsageComponent = *temp.PrepaidUsageComponent
    return nil
}

// createPrepaidComponent is a temporary struct used for validating the fields of CreatePrepaidComponent.
type createPrepaidComponent  struct {
    PrepaidUsageComponent *PrepaidUsageComponent `json:"prepaid_usage_component"`
}

func (c *createPrepaidComponent) validate() error {
    var errs []string
    if c.PrepaidUsageComponent == nil {
        errs = append(errs, "required field `prepaid_usage_component` is missing for type `Create Prepaid Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
