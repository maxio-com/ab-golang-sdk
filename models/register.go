// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// Register represents a Register struct.
type Register struct {
    Id                   *int                   `json:"id,omitempty"`
    MaxioId              *string                `json:"maxio_id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // The ISO 4217 currency code (3 character string) representing the currency of invoice transaction.
    CurrencyCode         *string                `json:"currency_code,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Register,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r Register) String() string {
    return fmt.Sprintf(
    	"Register[Id=%v, MaxioId=%v, Name=%v, CurrencyCode=%v, AdditionalProperties=%v]",
    	r.Id, r.MaxioId, r.Name, r.CurrencyCode, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Register.
// It customizes the JSON marshaling process for Register objects.
func (r Register) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "maxio_id", "name", "currency_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the Register object to a map representation for JSON marshaling.
func (r Register) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.MaxioId != nil {
        structMap["maxio_id"] = r.MaxioId
    }
    if r.Name != nil {
        structMap["name"] = r.Name
    }
    if r.CurrencyCode != nil {
        structMap["currency_code"] = r.CurrencyCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Register.
// It customizes the JSON unmarshaling process for Register objects.
func (r *Register) UnmarshalJSON(input []byte) error {
    var temp tempRegister
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "maxio_id", "name", "currency_code")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = temp.Id
    r.MaxioId = temp.MaxioId
    r.Name = temp.Name
    r.CurrencyCode = temp.CurrencyCode
    return nil
}

// tempRegister is a temporary struct used for validating the fields of Register.
type tempRegister  struct {
    Id           *int    `json:"id,omitempty"`
    MaxioId      *string `json:"maxio_id,omitempty"`
    Name         *string `json:"name,omitempty"`
    CurrencyCode *string `json:"currency_code,omitempty"`
}
