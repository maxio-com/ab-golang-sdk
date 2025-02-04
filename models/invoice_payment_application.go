/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoicePaymentApplication represents a InvoicePaymentApplication struct.
type InvoicePaymentApplication struct {
    // Unique identifier for the paid invoice. It has the prefix "inv_" followed by alphanumeric characters.
    InvoiceUid           *string                `json:"invoice_uid,omitempty"`
    // Unique identifier for the payment. It has the prefix "pmt_" followed by alphanumeric characters.
    ApplicationUid       *string                `json:"application_uid,omitempty"`
    // Dollar amount of the paid invoice.
    AppliedAmount        *string                `json:"applied_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoicePaymentApplication,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoicePaymentApplication) String() string {
    return fmt.Sprintf(
    	"InvoicePaymentApplication[InvoiceUid=%v, ApplicationUid=%v, AppliedAmount=%v, AdditionalProperties=%v]",
    	i.InvoiceUid, i.ApplicationUid, i.AppliedAmount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoicePaymentApplication.
// It customizes the JSON marshaling process for InvoicePaymentApplication objects.
func (i InvoicePaymentApplication) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "invoice_uid", "application_uid", "applied_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoicePaymentApplication object to a map representation for JSON marshaling.
func (i InvoicePaymentApplication) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.InvoiceUid != nil {
        structMap["invoice_uid"] = i.InvoiceUid
    }
    if i.ApplicationUid != nil {
        structMap["application_uid"] = i.ApplicationUid
    }
    if i.AppliedAmount != nil {
        structMap["applied_amount"] = i.AppliedAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePaymentApplication.
// It customizes the JSON unmarshaling process for InvoicePaymentApplication objects.
func (i *InvoicePaymentApplication) UnmarshalJSON(input []byte) error {
    var temp tempInvoicePaymentApplication
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoice_uid", "application_uid", "applied_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.InvoiceUid = temp.InvoiceUid
    i.ApplicationUid = temp.ApplicationUid
    i.AppliedAmount = temp.AppliedAmount
    return nil
}

// tempInvoicePaymentApplication is a temporary struct used for validating the fields of InvoicePaymentApplication.
type tempInvoicePaymentApplication  struct {
    InvoiceUid     *string `json:"invoice_uid,omitempty"`
    ApplicationUid *string `json:"application_uid,omitempty"`
    AppliedAmount  *string `json:"applied_amount,omitempty"`
}
