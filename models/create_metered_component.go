package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateMeteredComponent represents a CreateMeteredComponent struct.
type CreateMeteredComponent struct {
    MeteredComponent     MeteredComponent `json:"metered_component"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateMeteredComponent.
// It customizes the JSON marshaling process for CreateMeteredComponent objects.
func (c CreateMeteredComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateMeteredComponent object to a map representation for JSON marshaling.
func (c CreateMeteredComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["metered_component"] = c.MeteredComponent.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateMeteredComponent.
// It customizes the JSON unmarshaling process for CreateMeteredComponent objects.
func (c *CreateMeteredComponent) UnmarshalJSON(input []byte) error {
    var temp createMeteredComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "metered_component")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.MeteredComponent = *temp.MeteredComponent
    return nil
}

// TODO
type createMeteredComponent  struct {
    MeteredComponent *MeteredComponent `json:"metered_component"`
}

func (c *createMeteredComponent) validate() error {
    var errs []string
    if c.MeteredComponent == nil {
        errs = append(errs, "required field `metered_component` is missing for type `Create Metered Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
