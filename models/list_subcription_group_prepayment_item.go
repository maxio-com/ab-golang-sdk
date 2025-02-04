/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ListSubcriptionGroupPrepaymentItem represents a ListSubcriptionGroupPrepaymentItem struct.
type ListSubcriptionGroupPrepaymentItem struct {
    Id                     *int                   `json:"id,omitempty"`
    SubscriptionGroupUid   *string                `json:"subscription_group_uid,omitempty"`
    AmountInCents          *int64                 `json:"amount_in_cents,omitempty"`
    RemainingAmountInCents *int64                 `json:"remaining_amount_in_cents,omitempty"`
    Details                *string                `json:"details,omitempty"`
    External               *bool                  `json:"external,omitempty"`
    Memo                   *string                `json:"memo,omitempty"`
    PaymentType            *PrepaymentMethod      `json:"payment_type,omitempty"`
    CreatedAt              *time.Time             `json:"created_at,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubcriptionGroupPrepaymentItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubcriptionGroupPrepaymentItem) String() string {
    return fmt.Sprintf(
    	"ListSubcriptionGroupPrepaymentItem[Id=%v, SubscriptionGroupUid=%v, AmountInCents=%v, RemainingAmountInCents=%v, Details=%v, External=%v, Memo=%v, PaymentType=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	l.Id, l.SubscriptionGroupUid, l.AmountInCents, l.RemainingAmountInCents, l.Details, l.External, l.Memo, l.PaymentType, l.CreatedAt, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubcriptionGroupPrepaymentItem.
// It customizes the JSON marshaling process for ListSubcriptionGroupPrepaymentItem objects.
func (l ListSubcriptionGroupPrepaymentItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "id", "subscription_group_uid", "amount_in_cents", "remaining_amount_in_cents", "details", "external", "memo", "payment_type", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubcriptionGroupPrepaymentItem object to a map representation for JSON marshaling.
func (l ListSubcriptionGroupPrepaymentItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp tempListSubcriptionGroupPrepaymentItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "subscription_group_uid", "amount_in_cents", "remaining_amount_in_cents", "details", "external", "memo", "payment_type", "created_at")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
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

// tempListSubcriptionGroupPrepaymentItem is a temporary struct used for validating the fields of ListSubcriptionGroupPrepaymentItem.
type tempListSubcriptionGroupPrepaymentItem  struct {
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
