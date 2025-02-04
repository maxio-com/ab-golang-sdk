/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceLineItemPricingDetail represents a InvoiceLineItemPricingDetail struct.
type InvoiceLineItemPricingDetail struct {
    Label                *string                `json:"label,omitempty"`
    Amount               *string                `json:"amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceLineItemPricingDetail,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceLineItemPricingDetail) String() string {
    return fmt.Sprintf(
    	"InvoiceLineItemPricingDetail[Label=%v, Amount=%v, AdditionalProperties=%v]",
    	i.Label, i.Amount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemPricingDetail.
// It customizes the JSON marshaling process for InvoiceLineItemPricingDetail objects.
func (i InvoiceLineItemPricingDetail) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "label", "amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemPricingDetail object to a map representation for JSON marshaling.
func (i InvoiceLineItemPricingDetail) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Label != nil {
        structMap["label"] = i.Label
    }
    if i.Amount != nil {
        structMap["amount"] = i.Amount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItemPricingDetail.
// It customizes the JSON unmarshaling process for InvoiceLineItemPricingDetail objects.
func (i *InvoiceLineItemPricingDetail) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceLineItemPricingDetail
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "label", "amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Label = temp.Label
    i.Amount = temp.Amount
    return nil
}

// tempInvoiceLineItemPricingDetail is a temporary struct used for validating the fields of InvoiceLineItemPricingDetail.
type tempInvoiceLineItemPricingDetail  struct {
    Label  *string `json:"label,omitempty"`
    Amount *string `json:"amount,omitempty"`
}
