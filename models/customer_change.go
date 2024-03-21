package models

import (
	"encoding/json"
)

// CustomerChange represents a CustomerChange struct.
type CustomerChange struct {
	Payer           Optional[CustomerChangePayer]           `json:"payer"`
	ShippingAddress Optional[CustomerChangeShippingAddress] `json:"shipping_address"`
	BillingAddress  Optional[CustomerChangeBillingAddress]  `json:"billing_address"`
	CustomFields    Optional[CustomerChangeCustomFields]    `json:"custom_fields"`
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
	if c.Payer.IsValueSet() {
		if c.Payer.Value() != nil {
			structMap["payer"] = c.Payer.Value().toMap()
		} else {
			structMap["payer"] = nil
		}
	}
	if c.ShippingAddress.IsValueSet() {
		if c.ShippingAddress.Value() != nil {
			structMap["shipping_address"] = c.ShippingAddress.Value().toMap()
		} else {
			structMap["shipping_address"] = nil
		}
	}
	if c.BillingAddress.IsValueSet() {
		if c.BillingAddress.Value() != nil {
			structMap["billing_address"] = c.BillingAddress.Value().toMap()
		} else {
			structMap["billing_address"] = nil
		}
	}
	if c.CustomFields.IsValueSet() {
		if c.CustomFields.Value() != nil {
			structMap["custom_fields"] = c.CustomFields.Value().toMap()
		} else {
			structMap["custom_fields"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChange.
// It customizes the JSON unmarshaling process for CustomerChange objects.
func (c *CustomerChange) UnmarshalJSON(input []byte) error {
	var temp customerChange
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

// TODO
type customerChange struct {
	Payer           Optional[CustomerChangePayer]           `json:"payer"`
	ShippingAddress Optional[CustomerChangeShippingAddress] `json:"shipping_address"`
	BillingAddress  Optional[CustomerChangeBillingAddress]  `json:"billing_address"`
	CustomFields    Optional[CustomerChangeCustomFields]    `json:"custom_fields"`
}
