package models

import (
	"encoding/json"
	"log"
	"time"
)

// ListSubcriptionGroupPrepaymentItem represents a ListSubcriptionGroupPrepaymentItem struct.
type ListSubcriptionGroupPrepaymentItem struct {
	Id                     *int              `json:"id,omitempty"`
	SubscriptionGroupUid   *string           `json:"subscription_group_uid,omitempty"`
	AmountInCents          *int64            `json:"amount_in_cents,omitempty"`
	RemainingAmountInCents *int64            `json:"remaining_amount_in_cents,omitempty"`
	Details                *string           `json:"details,omitempty"`
	External               *bool             `json:"external,omitempty"`
	Memo                   *string           `json:"memo,omitempty"`
	PaymentType            *PrepaymentMethod `json:"payment_type,omitempty"`
	CreatedAt              *time.Time        `json:"created_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubcriptionGroupPrepaymentItem.
// It customizes the JSON marshaling process for ListSubcriptionGroupPrepaymentItem objects.
func (l *ListSubcriptionGroupPrepaymentItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubcriptionGroupPrepaymentItem object to a map representation for JSON marshaling.
func (l *ListSubcriptionGroupPrepaymentItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	if l.SubscriptionGroupUid != nil {
		structMap["subscription_group_uid"] = l.SubscriptionGroupUid
	}
	if l.AmountInCents != nil {
		structMap["amount_in_cents"] = l.AmountInCents
	}
	if l.RemainingAmountInCents != nil {
		structMap["remaining_amount_in_cents"] = l.RemainingAmountInCents
	}
	if l.Details != nil {
		structMap["details"] = l.Details
	}
	if l.External != nil {
		structMap["external"] = l.External
	}
	if l.Memo != nil {
		structMap["memo"] = l.Memo
	}
	if l.PaymentType != nil {
		structMap["payment_type"] = l.PaymentType
	}
	if l.CreatedAt != nil {
		structMap["created_at"] = l.CreatedAt.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubcriptionGroupPrepaymentItem.
// It customizes the JSON unmarshaling process for ListSubcriptionGroupPrepaymentItem objects.
func (l *ListSubcriptionGroupPrepaymentItem) UnmarshalJSON(input []byte) error {
	var temp listSubcriptionGroupPrepaymentItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	l.Id = temp.Id
	l.SubscriptionGroupUid = temp.SubscriptionGroupUid
	l.AmountInCents = temp.AmountInCents
	l.RemainingAmountInCents = temp.RemainingAmountInCents
	l.Details = temp.Details
	l.External = temp.External
	l.Memo = temp.Memo
	l.PaymentType = temp.PaymentType
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		l.CreatedAt = &CreatedAtVal
	}
	return nil
}

// TODO
type listSubcriptionGroupPrepaymentItem struct {
	Id                     *int              `json:"id,omitempty"`
	SubscriptionGroupUid   *string           `json:"subscription_group_uid,omitempty"`
	AmountInCents          *int64            `json:"amount_in_cents,omitempty"`
	RemainingAmountInCents *int64            `json:"remaining_amount_in_cents,omitempty"`
	Details                *string           `json:"details,omitempty"`
	External               *bool             `json:"external,omitempty"`
	Memo                   *string           `json:"memo,omitempty"`
	PaymentType            *PrepaymentMethod `json:"payment_type,omitempty"`
	CreatedAt              *string           `json:"created_at,omitempty"`
}
