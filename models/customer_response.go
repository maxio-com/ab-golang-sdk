/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CustomerResponse represents a CustomerResponse struct.
type CustomerResponse struct {
    Customer             Customer               `json:"customer"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CustomerResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerResponse) String() string {
    return fmt.Sprintf(
    	"CustomerResponse[Customer=%v, AdditionalProperties=%v]",
    	c.Customer, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CustomerResponse.
// It customizes the JSON marshaling process for CustomerResponse objects.
func (c CustomerResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "customer"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerResponse object to a map representation for JSON marshaling.
func (c CustomerResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["customer"] = c.Customer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerResponse.
// It customizes the JSON unmarshaling process for CustomerResponse objects.
func (c *CustomerResponse) UnmarshalJSON(input []byte) error {
    var temp tempCustomerResponse
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

// tempCustomerResponse is a temporary struct used for validating the fields of CustomerResponse.
type tempCustomerResponse  struct {
    Customer *Customer `json:"customer"`
}

func (c *tempCustomerResponse) validate() error {
    var errs []string
    if c.Customer == nil {
        errs = append(errs, "required field `customer` is missing for type `Customer Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
