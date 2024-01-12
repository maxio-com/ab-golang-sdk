package models

import (
    "encoding/json"
)

// CustomerResponse represents a CustomerResponse struct.
type CustomerResponse struct {
    Customer Customer `json:"customer"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerResponse.
// It customizes the JSON marshaling process for CustomerResponse objects.
func (c *CustomerResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerResponse object to a map representation for JSON marshaling.
func (c *CustomerResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["customer"] = c.Customer
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerResponse.
// It customizes the JSON unmarshaling process for CustomerResponse objects.
func (c *CustomerResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Customer Customer `json:"customer"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Customer = temp.Customer
    return nil
}
