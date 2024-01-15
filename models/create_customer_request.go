package models

import (
    "encoding/json"
)

// CreateCustomerRequest represents a CreateCustomerRequest struct.
type CreateCustomerRequest struct {
    Customer CreateCustomer `json:"customer"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCustomerRequest.
// It customizes the JSON marshaling process for CreateCustomerRequest objects.
func (c *CreateCustomerRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateCustomerRequest object to a map representation for JSON marshaling.
func (c *CreateCustomerRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["customer"] = c.Customer
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCustomerRequest.
// It customizes the JSON unmarshaling process for CreateCustomerRequest objects.
func (c *CreateCustomerRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Customer CreateCustomer `json:"customer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Customer = temp.Customer
    return nil
}
