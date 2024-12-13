/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreatePrepaidComponent represents a CreatePrepaidComponent struct.
type CreatePrepaidComponent struct {
    PrepaidUsageComponent PrepaidUsageComponent  `json:"prepaid_usage_component"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaidComponent.
// It customizes the JSON marshaling process for CreatePrepaidComponent objects.
func (c CreatePrepaidComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "prepaid_usage_component"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaidComponent object to a map representation for JSON marshaling.
func (c CreatePrepaidComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["prepaid_usage_component"] = c.PrepaidUsageComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaidComponent.
// It customizes the JSON unmarshaling process for CreatePrepaidComponent objects.
func (c *CreatePrepaidComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreatePrepaidComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepaid_usage_component")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PrepaidUsageComponent = *temp.PrepaidUsageComponent
    return nil
}

// tempCreatePrepaidComponent is a temporary struct used for validating the fields of CreatePrepaidComponent.
type tempCreatePrepaidComponent  struct {
    PrepaidUsageComponent *PrepaidUsageComponent `json:"prepaid_usage_component"`
}

func (c *tempCreatePrepaidComponent) validate() error {
    var errs []string
    if c.PrepaidUsageComponent == nil {
        errs = append(errs, "required field `prepaid_usage_component` is missing for type `Create Prepaid Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
