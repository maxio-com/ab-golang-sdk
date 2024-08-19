/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SendInvoiceRequest represents a SendInvoiceRequest struct.
type SendInvoiceRequest struct {
    RecipientEmails      []string       `json:"recipient_emails,omitempty"`
    CcRecipientEmails    []string       `json:"cc_recipient_emails,omitempty"`
    BccRecipientEmails   []string       `json:"bcc_recipient_emails,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SendInvoiceRequest.
// It customizes the JSON marshaling process for SendInvoiceRequest objects.
func (s SendInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SendInvoiceRequest object to a map representation for JSON marshaling.
func (s SendInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.RecipientEmails != nil {
        structMap["recipient_emails"] = s.RecipientEmails
    }
    if s.CcRecipientEmails != nil {
        structMap["cc_recipient_emails"] = s.CcRecipientEmails
    }
    if s.BccRecipientEmails != nil {
        structMap["bcc_recipient_emails"] = s.BccRecipientEmails
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SendInvoiceRequest.
// It customizes the JSON unmarshaling process for SendInvoiceRequest objects.
func (s *SendInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempSendInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "recipient_emails", "cc_recipient_emails", "bcc_recipient_emails")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.RecipientEmails = temp.RecipientEmails
    s.CcRecipientEmails = temp.CcRecipientEmails
    s.BccRecipientEmails = temp.BccRecipientEmails
    return nil
}

// tempSendInvoiceRequest is a temporary struct used for validating the fields of SendInvoiceRequest.
type tempSendInvoiceRequest  struct {
    RecipientEmails    []string `json:"recipient_emails,omitempty"`
    CcRecipientEmails  []string `json:"cc_recipient_emails,omitempty"`
    BccRecipientEmails []string `json:"bcc_recipient_emails,omitempty"`
}
