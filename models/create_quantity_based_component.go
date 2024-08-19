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

// CreateQuantityBasedComponent represents a CreateQuantityBasedComponent struct.
type CreateQuantityBasedComponent struct {
    QuantityBasedComponent QuantityBasedComponent `json:"quantity_based_component"`
    AdditionalProperties   map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateQuantityBasedComponent.
// It customizes the JSON marshaling process for CreateQuantityBasedComponent objects.
func (c CreateQuantityBasedComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateQuantityBasedComponent object to a map representation for JSON marshaling.
func (c CreateQuantityBasedComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["quantity_based_component"] = c.QuantityBasedComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateQuantityBasedComponent.
// It customizes the JSON unmarshaling process for CreateQuantityBasedComponent objects.
func (c *CreateQuantityBasedComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreateQuantityBasedComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "quantity_based_component")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.QuantityBasedComponent = *temp.QuantityBasedComponent
    return nil
}

// tempCreateQuantityBasedComponent is a temporary struct used for validating the fields of CreateQuantityBasedComponent.
type tempCreateQuantityBasedComponent  struct {
    QuantityBasedComponent *QuantityBasedComponent `json:"quantity_based_component"`
}

func (c *tempCreateQuantityBasedComponent) validate() error {
    var errs []string
    if c.QuantityBasedComponent == nil {
        errs = append(errs, "required field `quantity_based_component` is missing for type `Create Quantity Based Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
