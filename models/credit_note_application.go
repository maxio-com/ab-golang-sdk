package models

import (
    "encoding/json"
)

// CreditNoteApplication represents a CreditNoteApplication struct.
type CreditNoteApplication struct {
    Uid             *string `json:"uid,omitempty"`
    TransactionTime *string `json:"transaction_time,omitempty"`
    InvoiceUid      *string `json:"invoice_uid,omitempty"`
    Memo            *string `json:"memo,omitempty"`
    AppliedAmount   *string `json:"applied_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreditNoteApplication.
// It customizes the JSON marshaling process for CreditNoteApplication objects.
func (c *CreditNoteApplication) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreditNoteApplication object to a map representation for JSON marshaling.
func (c *CreditNoteApplication) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Uid != nil {
        structMap["uid"] = c.Uid
    }
    if c.TransactionTime != nil {
        structMap["transaction_time"] = c.TransactionTime
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
    temp := &struct {
        Uid             *string `json:"uid,omitempty"`
        TransactionTime *string `json:"transaction_time,omitempty"`
        InvoiceUid      *string `json:"invoice_uid,omitempty"`
        Memo            *string `json:"memo,omitempty"`
        AppliedAmount   *string `json:"applied_amount,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Uid = temp.Uid
    c.TransactionTime = temp.TransactionTime
    c.InvoiceUid = temp.InvoiceUid
    c.Memo = temp.Memo
    c.AppliedAmount = temp.AppliedAmount
    return nil
}
