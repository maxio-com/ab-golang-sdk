/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// CreditNoteApplication represents a CreditNoteApplication struct.
type CreditNoteApplication struct {
    Uid                  *string        `json:"uid,omitempty"`
    TransactionTime      *time.Time     `json:"transaction_time,omitempty"`
    InvoiceUid           *string        `json:"invoice_uid,omitempty"`
    Memo                 *string        `json:"memo,omitempty"`
    AppliedAmount        *string        `json:"applied_amount,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreditNoteApplication.
// It customizes the JSON marshaling process for CreditNoteApplication objects.
func (c CreditNoteApplication) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreditNoteApplication object to a map representation for JSON marshaling.
func (c CreditNoteApplication) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Uid != nil {
        structMap["uid"] = c.Uid
    }
    if c.TransactionTime != nil {
        structMap["transaction_time"] = c.TransactionTime.Format(time.RFC3339)
    }
    if c.InvoiceUid != nil {
        structMap["invoice_uid"] = c.InvoiceUid
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.AppliedAmount != nil {
        structMap["applied_amount"] = c.AppliedAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditNoteApplication.
// It customizes the JSON unmarshaling process for CreditNoteApplication objects.
func (c *CreditNoteApplication) UnmarshalJSON(input []byte) error {
    var temp tempCreditNoteApplication
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "transaction_time", "invoice_uid", "memo", "applied_amount")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Uid = temp.Uid
    if temp.TransactionTime != nil {
        TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
        if err != nil {
            log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
        }
        c.TransactionTime = &TransactionTimeVal
    }
    c.InvoiceUid = temp.InvoiceUid
    c.Memo = temp.Memo
    c.AppliedAmount = temp.AppliedAmount
    return nil
}

// tempCreditNoteApplication is a temporary struct used for validating the fields of CreditNoteApplication.
type tempCreditNoteApplication  struct {
    Uid             *string `json:"uid,omitempty"`
    TransactionTime *string `json:"transaction_time,omitempty"`
    InvoiceUid      *string `json:"invoice_uid,omitempty"`
    Memo            *string `json:"memo,omitempty"`
    AppliedAmount   *string `json:"applied_amount,omitempty"`
}
