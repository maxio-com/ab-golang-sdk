/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// InvoiceDiscountBreakout represents a InvoiceDiscountBreakout struct.
type InvoiceDiscountBreakout struct {
    Uid                  *string                `json:"uid,omitempty"`
    EligibleAmount       *string                `json:"eligible_amount,omitempty"`
    DiscountAmount       *string                `json:"discount_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceDiscountBreakout.
// It customizes the JSON marshaling process for InvoiceDiscountBreakout objects.
func (i InvoiceDiscountBreakout) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "eligible_amount", "discount_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceDiscountBreakout object to a map representation for JSON marshaling.
func (i InvoiceDiscountBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.EligibleAmount != nil {
        structMap["eligible_amount"] = i.EligibleAmount
    }
    if i.DiscountAmount != nil {
        structMap["discount_amount"] = i.DiscountAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDiscountBreakout.
// It customizes the JSON unmarshaling process for InvoiceDiscountBreakout objects.
func (i *InvoiceDiscountBreakout) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceDiscountBreakout
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "eligible_amount", "discount_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.EligibleAmount = temp.EligibleAmount
    i.DiscountAmount = temp.DiscountAmount
    return nil
}

// tempInvoiceDiscountBreakout is a temporary struct used for validating the fields of InvoiceDiscountBreakout.
type tempInvoiceDiscountBreakout  struct {
    Uid            *string `json:"uid,omitempty"`
    EligibleAmount *string `json:"eligible_amount,omitempty"`
    DiscountAmount *string `json:"discount_amount,omitempty"`
}
