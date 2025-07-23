// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UpdateCustomerRequest represents a UpdateCustomerRequest struct.
type UpdateCustomerRequest struct {
	Customer             UpdateCustomer         `json:"customer"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateCustomerRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateCustomerRequest) String() string {
	return fmt.Sprintf(
		"UpdateCustomerRequest[Customer=%v, AdditionalProperties=%v]",
		u.Customer, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateCustomerRequest.
// It customizes the JSON marshaling process for UpdateCustomerRequest objects.
func (u UpdateCustomerRequest) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"customer"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCustomerRequest object to a map representation for JSON marshaling.
func (u UpdateCustomerRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	structMap["customer"] = u.Customer.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCustomerRequest.
// It customizes the JSON unmarshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) UnmarshalJSON(input []byte) error {
	var temp tempUpdateCustomerRequest
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
	u.AdditionalProperties = additionalProperties

	u.Customer = *temp.Customer
	return nil
}

// tempUpdateCustomerRequest is a temporary struct used for validating the fields of UpdateCustomerRequest.
type tempUpdateCustomerRequest struct {
	Customer *UpdateCustomer `json:"customer"`
}

func (u *tempUpdateCustomerRequest) validate() error {
	var errs []string
	if u.Customer == nil {
		errs = append(errs, "required field `customer` is missing for type `Update Customer Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
