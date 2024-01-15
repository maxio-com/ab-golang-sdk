package models

import (
    "encoding/json"
)

// CustomerShippingAddressChange represents a CustomerShippingAddressChange struct.
type CustomerShippingAddressChange struct {
    Before *interface{} `json:"before,omitempty"`
    After  *interface{} `json:"after,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerShippingAddressChange.
// It customizes the JSON marshaling process for CustomerShippingAddressChange objects.
func (c *CustomerShippingAddressChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerShippingAddressChange object to a map representation for JSON marshaling.
func (c *CustomerShippingAddressChange) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Before != nil {
        structMap["before"] = c.Before
    }
    if c.After != nil {
        structMap["after"] = c.After
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerShippingAddressChange.
// It customizes the JSON unmarshaling process for CustomerShippingAddressChange objects.
func (c *CustomerShippingAddressChange) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Before *interface{} `json:"before,omitempty"`
        After  *interface{} `json:"after,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Before = temp.Before
    c.After = temp.After
    return nil
}
