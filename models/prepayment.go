package models

import (
	"encoding/json"
)

// Prepayment represents a Prepayment struct.
type Prepayment struct {
	Id                     int     `json:"id"`
	SubscriptionId         int     `json:"subscription_id"`
	AmountInCents          int64   `json:"amount_in_cents"`
	RemainingAmountInCents int64   `json:"remaining_amount_in_cents"`
	RefundedAmountInCents  *int64  `json:"refunded_amount_in_cents,omitempty"`
	Details                *string `json:"details,omitempty"`
	External               bool    `json:"external"`
	Memo                   string  `json:"memo"`
	// The payment type of the prepayment.
	PaymentType *PrepaymentMethodEnum `json:"payment_type,omitempty"`
	CreatedAt   string                `json:"created_at"`
}

// MarshalJSON implements the json.Marshaler interface for Prepayment.
// It customizes the JSON marshaling process for Prepayment objects.
func (p *Prepayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the Prepayment object to a map representation for JSON marshaling.
func (p *Prepayment) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = p.Id
	structMap["subscription_id"] = p.SubscriptionId
	structMap["amount_in_cents"] = p.AmountInCents
	structMap["remaining_amount_in_cents"] = p.RemainingAmountInCents
	if p.RefundedAmountInCents != nil {
		structMap["refunded_amount_in_cents"] = p.RefundedAmountInCents
	}
	if p.Details != nil {
		structMap["details"] = p.Details
	}
	structMap["external"] = p.External
	structMap["memo"] = p.Memo
	if p.PaymentType != nil {
		structMap["payment_type"] = p.PaymentType
	}
	structMap["created_at"] = p.CreatedAt
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prepayment.
// It customizes the JSON unmarshaling process for Prepayment objects.
func (p *Prepayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                     int                   `json:"id"`
		SubscriptionId         int                   `json:"subscription_id"`
		AmountInCents          int64                 `json:"amount_in_cents"`
		RemainingAmountInCents int64                 `json:"remaining_amount_in_cents"`
		RefundedAmountInCents  *int64                `json:"refunded_amount_in_cents,omitempty"`
		Details                *string               `json:"details,omitempty"`
		External               bool                  `json:"external"`
		Memo                   string                `json:"memo"`
		PaymentType            *PrepaymentMethodEnum `json:"payment_type,omitempty"`
		CreatedAt              string                `json:"created_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Id = temp.Id
	p.SubscriptionId = temp.SubscriptionId
	p.AmountInCents = temp.AmountInCents
	p.RemainingAmountInCents = temp.RemainingAmountInCents
	p.RefundedAmountInCents = temp.RefundedAmountInCents
	p.Details = temp.Details
	p.External = temp.External
	p.Memo = temp.Memo
	p.PaymentType = temp.PaymentType
	p.CreatedAt = temp.CreatedAt
	return nil
}
