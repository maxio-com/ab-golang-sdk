package models

import (
    "encoding/json"
)

// UpdateCustomerRequest represents a UpdateCustomerRequest struct.
type UpdateCustomerRequest struct {
    Customer UpdateCustomer `json:"customer"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCustomerRequest.
// It customizes the JSON marshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCustomerRequest object to a map representation for JSON marshaling.
func (u *UpdateCustomerRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["customer"] = u.Customer
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCustomerRequest.
// It customizes the JSON unmarshaling process for UpdateCustomerRequest objects.
func (u *UpdateCustomerRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Customer UpdateCustomer `json:"customer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Customer = temp.Customer
    return nil
}
