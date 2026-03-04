// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// DeliverProformaInvoiceRequest represents a DeliverProformaInvoiceRequest struct.
type DeliverProformaInvoiceRequest struct {
    RecipientEmails      []string               `json:"recipient_emails,omitempty"`
    CcRecipientEmails    []string               `json:"cc_recipient_emails,omitempty"`
    BccRecipientEmails   []string               `json:"bcc_recipient_emails,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DeliverProformaInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DeliverProformaInvoiceRequest) String() string {
    return fmt.Sprintf(
    	"DeliverProformaInvoiceRequest[RecipientEmails=%v, CcRecipientEmails=%v, BccRecipientEmails=%v, AdditionalProperties=%v]",
    	d.RecipientEmails, d.CcRecipientEmails, d.BccRecipientEmails, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DeliverProformaInvoiceRequest.
// It customizes the JSON marshaling process for DeliverProformaInvoiceRequest objects.
func (d DeliverProformaInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "recipient_emails", "cc_recipient_emails", "bcc_recipient_emails"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DeliverProformaInvoiceRequest object to a map representation for JSON marshaling.
func (d DeliverProformaInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.RecipientEmails != nil {
        structMap["recipient_emails"] = d.RecipientEmails
    }
    if d.CcRecipientEmails != nil {
        structMap["cc_recipient_emails"] = d.CcRecipientEmails
    }
    if d.BccRecipientEmails != nil {
        structMap["bcc_recipient_emails"] = d.BccRecipientEmails
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeliverProformaInvoiceRequest.
// It customizes the JSON unmarshaling process for DeliverProformaInvoiceRequest objects.
func (d *DeliverProformaInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempDeliverProformaInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "recipient_emails", "cc_recipient_emails", "bcc_recipient_emails")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.RecipientEmails = temp.RecipientEmails
    d.CcRecipientEmails = temp.CcRecipientEmails
    d.BccRecipientEmails = temp.BccRecipientEmails
    return nil
}

// tempDeliverProformaInvoiceRequest is a temporary struct used for validating the fields of DeliverProformaInvoiceRequest.
type tempDeliverProformaInvoiceRequest  struct {
    RecipientEmails    []string `json:"recipient_emails,omitempty"`
    CcRecipientEmails  []string `json:"cc_recipient_emails,omitempty"`
    BccRecipientEmails []string `json:"bcc_recipient_emails,omitempty"`
}
