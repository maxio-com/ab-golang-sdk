package models

import (
    "encoding/json"
)

// CreateProductFamily represents a CreateProductFamily struct.
type CreateProductFamily struct {
    Name                 *string          `json:"name,omitempty"`
    Description          Optional[string] `json:"description"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductFamily.
// It customizes the JSON marshaling process for CreateProductFamily objects.
func (c CreateProductFamily) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateProductFamily object to a map representation for JSON marshaling.
func (c CreateProductFamily) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Description.IsValueSet() {
        if c.Description.Value() != nil {
            structMap["description"] = c.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductFamily.
// It customizes the JSON unmarshaling process for CreateProductFamily objects.
func (c *CreateProductFamily) UnmarshalJSON(input []byte) error {
    var temp createProductFamily
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "description")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = temp.Name
    c.Description = temp.Description
    return nil
}

// createProductFamily is a temporary struct used for validating the fields of CreateProductFamily.
type createProductFamily  struct {
    Name        *string          `json:"name,omitempty"`
    Description Optional[string] `json:"description"`
}
