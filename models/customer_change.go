package models

import (
    "encoding/json"
)

// CustomerChange represents a CustomerChange struct.
type CustomerChange struct {
    Payer           Optional[CustomerPayerChange]        `json:"payer"`
    ShippingAddress Optional[AddressChange]              `json:"shipping_address"`
    BillingAddress  Optional[AddressChange]              `json:"billing_address"`
    CustomFields    Optional[CustomerCustomFieldsChange] `json:"custom_fields"`
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
        structMap["payer"] = c.Payer.Value()
    }
    if c.ShippingAddress.IsValueSet() {
        structMap["shipping_address"] = c.ShippingAddress.Value()
    }
    if c.BillingAddress.IsValueSet() {
        structMap["billing_address"] = c.BillingAddress.Value()
    }
    if c.CustomFields.IsValueSet() {
        structMap["custom_fields"] = c.CustomFields.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerChange.
// It customizes the JSON unmarshaling process for CustomerChange objects.
func (c *CustomerChange) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Payer           Optional[CustomerPayerChange]        `json:"payer"`
        ShippingAddress Optional[AddressChange]              `json:"shipping_address"`
        BillingAddress  Optional[AddressChange]              `json:"billing_address"`
        CustomFields    Optional[CustomerCustomFieldsChange] `json:"custom_fields"`
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
