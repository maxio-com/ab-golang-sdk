package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// DeductServiceCredit represents a DeductServiceCredit struct.
type DeductServiceCredit struct {
	Amount DeductServiceCreditAmount `json:"amount"`
	Memo   string                    `json:"memo"`
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
	structMap["amount"] = d.Amount.toMap()
	structMap["memo"] = d.Memo
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCredit.
// It customizes the JSON unmarshaling process for DeductServiceCredit objects.
func (d *DeductServiceCredit) UnmarshalJSON(input []byte) error {
	var temp deductServiceCredit
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	d.Amount = *temp.Amount
	d.Memo = *temp.Memo
	return nil
}

// TODO
type deductServiceCredit struct {
	Amount *DeductServiceCreditAmount `json:"amount"`
	Memo   *string                    `json:"memo"`
}

func (d *deductServiceCredit) validate() error {
	var errs []string
	if d.Amount == nil {
		errs = append(errs, "required field `amount` is missing for type `Deduct Service Credit`")
	}
	if d.Memo == nil {
		errs = append(errs, "required field `memo` is missing for type `Deduct Service Credit`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
