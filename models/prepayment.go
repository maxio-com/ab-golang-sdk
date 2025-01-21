/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// Prepayment represents a Prepayment struct.
type Prepayment struct {
    Id                     int                    `json:"id"`
    SubscriptionId         int                    `json:"subscription_id"`
    AmountInCents          int64                  `json:"amount_in_cents"`
    RemainingAmountInCents int64                  `json:"remaining_amount_in_cents"`
    RefundedAmountInCents  *int64                 `json:"refunded_amount_in_cents,omitempty"`
    Details                *string                `json:"details,omitempty"`
    External               bool                   `json:"external"`
    Memo                   string                 `json:"memo"`
    // The payment type of the prepayment.
    PaymentType            *PrepaymentMethod      `json:"payment_type,omitempty"`
    CreatedAt              time.Time              `json:"created_at"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Prepayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p Prepayment) String() string {
    return fmt.Sprintf(
    	"Prepayment[Id=%v, SubscriptionId=%v, AmountInCents=%v, RemainingAmountInCents=%v, RefundedAmountInCents=%v, Details=%v, External=%v, Memo=%v, PaymentType=%v, CreatedAt=%v, AdditionalProperties=%v]",
    	p.Id, p.SubscriptionId, p.AmountInCents, p.RemainingAmountInCents, p.RefundedAmountInCents, p.Details, p.External, p.Memo, p.PaymentType, p.CreatedAt, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Prepayment.
// It customizes the JSON marshaling process for Prepayment objects.
func (p Prepayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "id", "subscription_id", "amount_in_cents", "remaining_amount_in_cents", "refunded_amount_in_cents", "details", "external", "memo", "payment_type", "created_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the Prepayment object to a map representation for JSON marshaling.
func (p Prepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
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
    structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prepayment.
// It customizes the JSON unmarshaling process for Prepayment objects.
func (p *Prepayment) UnmarshalJSON(input []byte) error {
    var temp tempPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "subscription_id", "amount_in_cents", "remaining_amount_in_cents", "refunded_amount_in_cents", "details", "external", "memo", "payment_type", "created_at")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Id = *temp.Id
    p.SubscriptionId = *temp.SubscriptionId
    p.AmountInCents = *temp.AmountInCents
    p.RemainingAmountInCents = *temp.RemainingAmountInCents
    p.RefundedAmountInCents = temp.RefundedAmountInCents
    p.Details = temp.Details
    p.External = *temp.External
    p.Memo = *temp.Memo
    p.PaymentType = temp.PaymentType
    CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
    }
    p.CreatedAt = CreatedAtVal
    return nil
}

// tempPrepayment is a temporary struct used for validating the fields of Prepayment.
type tempPrepayment  struct {
    Id                     *int              `json:"id"`
    SubscriptionId         *int              `json:"subscription_id"`
    AmountInCents          *int64            `json:"amount_in_cents"`
    RemainingAmountInCents *int64            `json:"remaining_amount_in_cents"`
    RefundedAmountInCents  *int64            `json:"refunded_amount_in_cents,omitempty"`
    Details                *string           `json:"details,omitempty"`
    External               *bool             `json:"external"`
    Memo                   *string           `json:"memo"`
    PaymentType            *PrepaymentMethod `json:"payment_type,omitempty"`
    CreatedAt              *string           `json:"created_at"`
}

func (p *tempPrepayment) validate() error {
    var errs []string
    if p.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Prepayment`")
    }
    if p.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Prepayment`")
    }
    if p.AmountInCents == nil {
        errs = append(errs, "required field `amount_in_cents` is missing for type `Prepayment`")
    }
    if p.RemainingAmountInCents == nil {
        errs = append(errs, "required field `remaining_amount_in_cents` is missing for type `Prepayment`")
    }
    if p.External == nil {
        errs = append(errs, "required field `external` is missing for type `Prepayment`")
    }
    if p.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Prepayment`")
    }
    if p.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `Prepayment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
