package models

import (
	"encoding/json"
)

// RefundPrepayment represents a RefundPrepayment struct.
type RefundPrepayment struct {
	// `amount` is not required if you pass `amount_in_cents`.
	AmountInCents int64 `json:"amount_in_cents"`
	// `amount_in_cents` is not required if you pass `amount`.
	Amount interface{} `json:"amount"`
	Memo   string      `json:"memo"`
	// Specify the type of refund you wish to initiate. When the prepayment is external, the `external` flag is optional. But if the prepayment was made through a payment profile, the `external` flag is required.
	External *bool `json:"external,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepayment.
// It customizes the JSON marshaling process for RefundPrepayment objects.
func (r *RefundPrepayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepayment object to a map representation for JSON marshaling.
func (r *RefundPrepayment) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount_in_cents"] = r.AmountInCents
	structMap["amount"] = r.Amount
	structMap["memo"] = r.Memo
	if r.External != nil {
		structMap["external"] = r.External
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepayment.
// It customizes the JSON unmarshaling process for RefundPrepayment objects.
func (r *RefundPrepayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AmountInCents int64       `json:"amount_in_cents"`
		Amount        interface{} `json:"amount"`
		Memo          string      `json:"memo"`
		External      *bool       `json:"external,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.AmountInCents = temp.AmountInCents
	r.Amount = temp.Amount
	r.Memo = temp.Memo
	r.External = temp.External
	return nil
}
