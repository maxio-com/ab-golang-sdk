package models

import (
    "encoding/json"
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
    // :- When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions.
    PaymentType            *PrepaymentMethod `json:"payment_type,omitempty"`
    CreatedAt              *string           `json:"created_at,omitempty"`
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
        structMap["created_at"] = l.CreatedAt
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubcriptionGroupPrepaymentItem.
// It customizes the JSON unmarshaling process for ListSubcriptionGroupPrepaymentItem objects.
func (l *ListSubcriptionGroupPrepaymentItem) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                     *int              `json:"id,omitempty"`
        SubscriptionGroupUid   *string           `json:"subscription_group_uid,omitempty"`
        AmountInCents          *int64            `json:"amount_in_cents,omitempty"`
        RemainingAmountInCents *int64            `json:"remaining_amount_in_cents,omitempty"`
        Details                *string           `json:"details,omitempty"`
        External               *bool             `json:"external,omitempty"`
        Memo                   *string           `json:"memo,omitempty"`
        PaymentType            *PrepaymentMethod `json:"payment_type,omitempty"`
        CreatedAt              *string           `json:"created_at,omitempty"`
    }{}
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
    l.CreatedAt = temp.CreatedAt
    return nil
}
