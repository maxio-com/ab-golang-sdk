package models

import (
	"encoding/json"
)

// CustomerBillingAddressChange represents a CustomerBillingAddressChange struct.
type CustomerBillingAddressChange struct {
	Before *interface{} `json:"before,omitempty"`
	After  *interface{} `json:"after,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerBillingAddressChange.
// It customizes the JSON marshaling process for CustomerBillingAddressChange objects.
func (c *CustomerBillingAddressChange) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerBillingAddressChange object to a map representation for JSON marshaling.
func (c *CustomerBillingAddressChange) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Before != nil {
		structMap["before"] = c.Before
	}
	if c.After != nil {
		structMap["after"] = c.After
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerBillingAddressChange.
// It customizes the JSON unmarshaling process for CustomerBillingAddressChange objects.
func (c *CustomerBillingAddressChange) UnmarshalJSON(input []byte) error {
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
