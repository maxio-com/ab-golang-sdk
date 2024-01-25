package models

import (
	"encoding/json"
)

// DeductServiceCredit represents a DeductServiceCredit struct.
type DeductServiceCredit struct {
	Amount interface{} `json:"amount"`
	Memo   string      `json:"memo"`
}

// MarshalJSON implements the json.Marshaler interface for DeductServiceCredit.
// It customizes the JSON marshaling process for DeductServiceCredit objects.
func (d *DeductServiceCredit) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DeductServiceCredit object to a map representation for JSON marshaling.
func (d *DeductServiceCredit) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = d.Amount
	structMap["memo"] = d.Memo
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCredit.
// It customizes the JSON unmarshaling process for DeductServiceCredit objects.
func (d *DeductServiceCredit) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount interface{} `json:"amount"`
		Memo   string      `json:"memo"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	d.Amount = temp.Amount
	d.Memo = temp.Memo
	return nil
}
