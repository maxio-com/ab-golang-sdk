package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateCustomerRequest represents a UpdateCustomerRequest struct.
type UpdateCustomerRequest struct {
    Customer             UpdateCustomer `json:"customer"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCustomerRequest.
// It customizes the JSON marshaling process for UpdateCustomerRequest objects.
func (u UpdateCustomerRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCustomerRequest object to a map representation for JSON marshaling.
func (u UpdateCustomerRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["customer"] = u.Customer.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCustomerRequest.
// It customizes the JSON unmarshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) UnmarshalJSON(input []byte) error {
    var temp updateCustomerRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "customer")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Customer = *temp.Customer
    return nil
}

// TODO
type updateCustomerRequest  struct {
    Customer *UpdateCustomer `json:"customer"`
}

func (u *updateCustomerRequest) validate() error {
    var errs []string
    if u.Customer == nil {
        errs = append(errs, "required field `customer` is missing for type `Update Customer Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
