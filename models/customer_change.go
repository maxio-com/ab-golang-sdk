package models

import (
	"encoding/json"
)

// CustomerChange represents a CustomerChange struct.
type CustomerChange struct {
	Payer           *CustomerPayerChange           `json:"payer,omitempty"`
	ShippingAddress *CustomerShippingAddressChange `json:"shipping_address,omitempty"`
	BillingAddress  *CustomerBillingAddressChange  `json:"billing_address,omitempty"`
	CustomFields    *CustomerCustomFieldsChange    `json:"custom_fields,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerChange.
// It customizes the JSON marshaling process for CustomerChange objects.
func (c *CustomerChange) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerChange object to a map representation for JSON marshaling.
func (c *CustomerChange) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Payer != nil {
		structMap["payer"] = c.Payer
	}
	if c.ShippingAddress != nil {
		structMap["shipping_address"] = c.ShippingAddress
	}
	if c.BillingAddress != nil {
		structMap["billing_address"] = c.BillingAddress
	}
	if c.CustomFields != nil {
		structMap["custom_fields"] = c.CustomFields
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChange.
// It customizes the JSON unmarshaling process for CustomerChange objects.
func (c *CustomerChange) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Payer           *CustomerPayerChange           `json:"payer,omitempty"`
		ShippingAddress *CustomerShippingAddressChange `json:"shipping_address,omitempty"`
		BillingAddress  *CustomerBillingAddressChange  `json:"billing_address,omitempty"`
		CustomFields    *CustomerCustomFieldsChange    `json:"custom_fields,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Payer = temp.Payer
	c.ShippingAddress = temp.ShippingAddress
	c.BillingAddress = temp.BillingAddress
	c.CustomFields = temp.CustomFields
	return nil
}
