package models

import (
	"encoding/json"
)

// InvoiceCredit represents a InvoiceCredit struct.
type InvoiceCredit struct {
	Uid              *string `json:"uid,omitempty"`
	CreditNoteNumber *string `json:"credit_note_number,omitempty"`
	CreditNoteUid    *string `json:"credit_note_uid,omitempty"`
	TransactionTime  *string `json:"transaction_time,omitempty"`
	Memo             *string `json:"memo,omitempty"`
	OriginalAmount   *string `json:"original_amount,omitempty"`
	AppliedAmount    *string `json:"applied_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceCredit.
// It customizes the JSON marshaling process for InvoiceCredit objects.
func (i *InvoiceCredit) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceCredit object to a map representation for JSON marshaling.
func (i *InvoiceCredit) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Uid != nil {
		structMap["uid"] = i.Uid
	}
	if i.CreditNoteNumber != nil {
		structMap["credit_note_number"] = i.CreditNoteNumber
	}
	if i.CreditNoteUid != nil {
		structMap["credit_note_uid"] = i.CreditNoteUid
	}
	if i.TransactionTime != nil {
		structMap["transaction_time"] = i.TransactionTime
	}
	if i.Memo != nil {
		structMap["memo"] = i.Memo
	}
	if i.OriginalAmount != nil {
		structMap["original_amount"] = i.OriginalAmount
	}
	if i.AppliedAmount != nil {
		structMap["applied_amount"] = i.AppliedAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceCredit.
// It customizes the JSON unmarshaling process for InvoiceCredit objects.
func (i *InvoiceCredit) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Uid              *string `json:"uid,omitempty"`
		CreditNoteNumber *string `json:"credit_note_number,omitempty"`
		CreditNoteUid    *string `json:"credit_note_uid,omitempty"`
		TransactionTime  *string `json:"transaction_time,omitempty"`
		Memo             *string `json:"memo,omitempty"`
		OriginalAmount   *string `json:"original_amount,omitempty"`
		AppliedAmount    *string `json:"applied_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.Uid = temp.Uid
	i.CreditNoteNumber = temp.CreditNoteNumber
	i.CreditNoteUid = temp.CreditNoteUid
	i.TransactionTime = temp.TransactionTime
	i.Memo = temp.Memo
	i.OriginalAmount = temp.OriginalAmount
	i.AppliedAmount = temp.AppliedAmount
	return nil
}
