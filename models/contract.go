// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Contract represents a Contract struct.
// Contract linked to the scheduled renewal configuration.
type Contract struct {
    Id                   *int                   `json:"id,omitempty"`
    MaxioId              *string                `json:"maxio_id,omitempty"`
    Number               Optional[string]       `json:"number"`
    Register             *Register              `json:"register,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Contract,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c Contract) String() string {
    return fmt.Sprintf(
    	"Contract[Id=%v, MaxioId=%v, Number=%v, Register=%v, AdditionalProperties=%v]",
    	c.Id, c.MaxioId, c.Number, c.Register, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Contract.
// It customizes the JSON marshaling process for Contract objects.
func (c Contract) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "maxio_id", "number", "register"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the Contract object to a map representation for JSON marshaling.
func (c Contract) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.MaxioId != nil {
        structMap["maxio_id"] = c.MaxioId
    }
    if c.Number.IsValueSet() {
        if c.Number.Value() != nil {
            structMap["number"] = c.Number.Value()
        } else {
            structMap["number"] = nil
        }
    }
    if c.Register != nil {
        structMap["register"] = c.Register.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Contract.
// It customizes the JSON unmarshaling process for Contract objects.
func (c *Contract) UnmarshalJSON(input []byte) error {
    var temp tempContract
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "maxio_id", "number", "register")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.MaxioId = temp.MaxioId
    c.Number = temp.Number
    c.Register = temp.Register
    return nil
}

// tempContract is a temporary struct used for validating the fields of Contract.
type tempContract  struct {
    Id       *int             `json:"id,omitempty"`
    MaxioId  *string          `json:"maxio_id,omitempty"`
    Number   Optional[string] `json:"number"`
    Register *Register        `json:"register,omitempty"`
}
