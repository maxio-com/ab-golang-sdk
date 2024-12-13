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

// CreateCustomerRequest represents a CreateCustomerRequest struct.
type CreateCustomerRequest struct {
    Customer             CreateCustomer         `json:"customer"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCustomerRequest.
// It customizes the JSON marshaling process for CreateCustomerRequest objects.
func (c CreateCustomerRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "customer"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCustomerRequest object to a map representation for JSON marshaling.
func (c CreateCustomerRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["customer"] = c.Customer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCustomerRequest.
// It customizes the JSON unmarshaling process for CreateCustomerRequest objects.
func (c *CreateCustomerRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateCustomerRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "customer")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Customer = *temp.Customer
    return nil
}

// tempCreateCustomerRequest is a temporary struct used for validating the fields of CreateCustomerRequest.
type tempCreateCustomerRequest  struct {
    Customer *CreateCustomer `json:"customer"`
}

func (c *tempCreateCustomerRequest) validate() error {
    var errs []string
    if c.Customer == nil {
        errs = append(errs, "required field `customer` is missing for type `Create Customer Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
