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

// CreateEBBComponent represents a CreateEBBComponent struct.
type CreateEBBComponent struct {
    EventBasedComponent  EBBComponent   `json:"event_based_component"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEBBComponent.
// It customizes the JSON marshaling process for CreateEBBComponent objects.
func (c CreateEBBComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateEBBComponent object to a map representation for JSON marshaling.
func (c CreateEBBComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["event_based_component"] = c.EventBasedComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEBBComponent.
// It customizes the JSON unmarshaling process for CreateEBBComponent objects.
func (c *CreateEBBComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreateEBBComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "event_based_component")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.EventBasedComponent = *temp.EventBasedComponent
    return nil
}

// tempCreateEBBComponent is a temporary struct used for validating the fields of CreateEBBComponent.
type tempCreateEBBComponent  struct {
    EventBasedComponent *EBBComponent `json:"event_based_component"`
}

func (c *tempCreateEBBComponent) validate() error {
    var errs []string
    if c.EventBasedComponent == nil {
        errs = append(errs, "required field `event_based_component` is missing for type `Create EBB Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
