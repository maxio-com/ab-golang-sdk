package models

import (
    "encoding/json"
)

// CustomerError represents a CustomerError struct.
type CustomerError struct {
    Customer             *string        `json:"customer,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerError.
// It customizes the JSON marshaling process for CustomerError objects.
func (c CustomerError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerError object to a map representation for JSON marshaling.
func (c CustomerError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Customer != nil {
        structMap["customer"] = c.Customer
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerError.
// It customizes the JSON unmarshaling process for CustomerError objects.
func (c *CustomerError) UnmarshalJSON(input []byte) error {
    var temp customerError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "customer")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Customer = temp.Customer
    return nil
}

// TODO
type customerError  struct {
    Customer *string `json:"customer,omitempty"`
}
