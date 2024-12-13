/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CustomerError represents a CustomerError struct.
type CustomerError struct {
    Customer             *string                `json:"customer,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerError.
// It customizes the JSON marshaling process for CustomerError objects.
func (c CustomerError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "customer"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerError object to a map representation for JSON marshaling.
func (c CustomerError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Customer != nil {
        structMap["customer"] = c.Customer
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerError.
// It customizes the JSON unmarshaling process for CustomerError objects.
func (c *CustomerError) UnmarshalJSON(input []byte) error {
    var temp tempCustomerError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "customer")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Customer = temp.Customer
    return nil
}

// tempCustomerError is a temporary struct used for validating the fields of CustomerError.
type tempCustomerError  struct {
    Customer *string `json:"customer,omitempty"`
}
