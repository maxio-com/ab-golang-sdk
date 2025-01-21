/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CustomerChange represents a CustomerChange struct.
type CustomerChange struct {
    Payer                Optional[CustomerPayerChange]        `json:"payer"`
    ShippingAddress      Optional[AddressChange]              `json:"shipping_address"`
    BillingAddress       Optional[AddressChange]              `json:"billing_address"`
    CustomFields         Optional[CustomerCustomFieldsChange] `json:"custom_fields"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for CustomerChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerChange) String() string {
    return fmt.Sprintf(
    	"CustomerChange[Payer=%v, ShippingAddress=%v, BillingAddress=%v, CustomFields=%v, AdditionalProperties=%v]",
    	c.Payer, c.ShippingAddress, c.BillingAddress, c.CustomFields, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CustomerChange.
// It customizes the JSON marshaling process for CustomerChange objects.
func (c CustomerChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "payer", "shipping_address", "billing_address", "custom_fields"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerChange object to a map representation for JSON marshaling.
func (c CustomerChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCustomerChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payer", "shipping_address", "billing_address", "custom_fields")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Payer = temp.Payer
    c.ShippingAddress = temp.ShippingAddress
    c.BillingAddress = temp.BillingAddress
    c.CustomFields = temp.CustomFields
    return nil
}

// tempCustomerChange is a temporary struct used for validating the fields of CustomerChange.
type tempCustomerChange  struct {
    Payer           Optional[CustomerPayerChange]        `json:"payer"`
    ShippingAddress Optional[AddressChange]              `json:"shipping_address"`
    BillingAddress  Optional[AddressChange]              `json:"billing_address"`
    CustomFields    Optional[CustomerCustomFieldsChange] `json:"custom_fields"`
}
