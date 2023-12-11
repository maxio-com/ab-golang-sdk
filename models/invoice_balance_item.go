package models

import (
	"encoding/json"
)

// InvoiceBalanceItem represents a InvoiceBalanceItem struct.
type InvoiceBalanceItem struct {
	Uid               *string `json:"uid,omitempty"`
	Number            *string `json:"number,omitempty"`
	OutstandingAmount *string `json:"outstanding_amount,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceBalanceItem.
// It customizes the JSON marshaling process for InvoiceBalanceItem objects.
func (i *InvoiceBalanceItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceBalanceItem object to a map representation for JSON marshaling.
func (i *InvoiceBalanceItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Uid != nil {
		structMap["uid"] = i.Uid
	}
	if i.Number != nil {
		structMap["number"] = i.Number
	}
	if i.OutstandingAmount != nil {
		structMap["outstanding_amount"] = i.OutstandingAmount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceBalanceItem.
// It customizes the JSON unmarshaling process for InvoiceBalanceItem objects.
func (i *InvoiceBalanceItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Uid               *string `json:"uid,omitempty"`
		Number            *string `json:"number,omitempty"`
		OutstandingAmount *string `json:"outstanding_amount,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.Uid = temp.Uid
	i.Number = temp.Number
	i.OutstandingAmount = temp.OutstandingAmount
	return nil
}
