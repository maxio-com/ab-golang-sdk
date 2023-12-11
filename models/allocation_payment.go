package models

import (
	"encoding/json"
)

// AllocationPayment represents a AllocationPayment struct.
// Information for captured payment, if applicable
type AllocationPayment struct {
	Id            *int    `json:"id,omitempty"`
	AmountInCents *int64  `json:"amount_in_cents,omitempty"`
	Success       *bool   `json:"success,omitempty"`
	Memo          *string `json:"memo,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPayment.
// It customizes the JSON marshaling process for AllocationPayment objects.
func (a *AllocationPayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AllocationPayment object to a map representation for JSON marshaling.
func (a *AllocationPayment) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.Id != nil {
		structMap["id"] = a.Id
	}
	if a.AmountInCents != nil {
		structMap["amount_in_cents"] = a.AmountInCents
	}
	if a.Success != nil {
		structMap["success"] = a.Success
	}
	if a.Memo != nil {
		structMap["memo"] = a.Memo
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPayment.
// It customizes the JSON unmarshaling process for AllocationPayment objects.
func (a *AllocationPayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            *int    `json:"id,omitempty"`
		AmountInCents *int64  `json:"amount_in_cents,omitempty"`
		Success       *bool   `json:"success,omitempty"`
		Memo          *string `json:"memo,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	a.Id = temp.Id
	a.AmountInCents = temp.AmountInCents
	a.Success = temp.Success
	a.Memo = temp.Memo
	return nil
}
