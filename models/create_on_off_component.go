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

// CreateOnOffComponent represents a CreateOnOffComponent struct.
type CreateOnOffComponent struct {
    OnOffComponent       OnOffComponent `json:"on_off_component"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOnOffComponent.
// It customizes the JSON marshaling process for CreateOnOffComponent objects.
func (c CreateOnOffComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOnOffComponent object to a map representation for JSON marshaling.
func (c CreateOnOffComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["on_off_component"] = c.OnOffComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOnOffComponent.
// It customizes the JSON unmarshaling process for CreateOnOffComponent objects.
func (c *CreateOnOffComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreateOnOffComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "on_off_component")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.OnOffComponent = *temp.OnOffComponent
    return nil
}

// tempCreateOnOffComponent is a temporary struct used for validating the fields of CreateOnOffComponent.
type tempCreateOnOffComponent  struct {
    OnOffComponent *OnOffComponent `json:"on_off_component"`
}

func (c *tempCreateOnOffComponent) validate() error {
    var errs []string
    if c.OnOffComponent == nil {
        errs = append(errs, "required field `on_off_component` is missing for type `Create On/Off Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
