// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateMeteredComponent represents a CreateMeteredComponent struct.
type CreateMeteredComponent struct {
    MeteredComponent     MeteredComponent       `json:"metered_component"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateMeteredComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateMeteredComponent) String() string {
    return fmt.Sprintf(
    	"CreateMeteredComponent[MeteredComponent=%v, AdditionalProperties=%v]",
    	c.MeteredComponent, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateMeteredComponent.
// It customizes the JSON marshaling process for CreateMeteredComponent objects.
func (c CreateMeteredComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "metered_component"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMeteredComponent object to a map representation for JSON marshaling.
func (c CreateMeteredComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["metered_component"] = c.MeteredComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMeteredComponent.
// It customizes the JSON unmarshaling process for CreateMeteredComponent objects.
func (c *CreateMeteredComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreateMeteredComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metered_component")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.MeteredComponent = *temp.MeteredComponent
    return nil
}

// tempCreateMeteredComponent is a temporary struct used for validating the fields of CreateMeteredComponent.
type tempCreateMeteredComponent  struct {
    MeteredComponent *MeteredComponent `json:"metered_component"`
}

func (c *tempCreateMeteredComponent) validate() error {
    var errs []string
    if c.MeteredComponent == nil {
        errs = append(errs, "required field `metered_component` is missing for type `Create Metered Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
