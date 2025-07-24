// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceTaxBreakout represents a InvoiceTaxBreakout struct.
type InvoiceTaxBreakout struct {
    Uid                  *string                `json:"uid,omitempty"`
    TaxableAmount        *string                `json:"taxable_amount,omitempty"`
    TaxAmount            *string                `json:"tax_amount,omitempty"`
    TaxExemptAmount      *string                `json:"tax_exempt_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceTaxBreakout,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceTaxBreakout) String() string {
    return fmt.Sprintf(
    	"InvoiceTaxBreakout[Uid=%v, TaxableAmount=%v, TaxAmount=%v, TaxExemptAmount=%v, AdditionalProperties=%v]",
    	i.Uid, i.TaxableAmount, i.TaxAmount, i.TaxExemptAmount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceTaxBreakout.
// It customizes the JSON marshaling process for InvoiceTaxBreakout objects.
func (i InvoiceTaxBreakout) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "taxable_amount", "tax_amount", "tax_exempt_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceTaxBreakout object to a map representation for JSON marshaling.
func (i InvoiceTaxBreakout) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.TaxableAmount != nil {
        structMap["taxable_amount"] = i.TaxableAmount
    }
    if i.TaxAmount != nil {
        structMap["tax_amount"] = i.TaxAmount
    }
    if i.TaxExemptAmount != nil {
        structMap["tax_exempt_amount"] = i.TaxExemptAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceTaxBreakout.
// It customizes the JSON unmarshaling process for InvoiceTaxBreakout objects.
func (i *InvoiceTaxBreakout) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceTaxBreakout
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "taxable_amount", "tax_amount", "tax_exempt_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.TaxableAmount = temp.TaxableAmount
    i.TaxAmount = temp.TaxAmount
    i.TaxExemptAmount = temp.TaxExemptAmount
    return nil
}

// tempInvoiceTaxBreakout is a temporary struct used for validating the fields of InvoiceTaxBreakout.
type tempInvoiceTaxBreakout  struct {
    Uid             *string `json:"uid,omitempty"`
    TaxableAmount   *string `json:"taxable_amount,omitempty"`
    TaxAmount       *string `json:"tax_amount,omitempty"`
    TaxExemptAmount *string `json:"tax_exempt_amount,omitempty"`
}
