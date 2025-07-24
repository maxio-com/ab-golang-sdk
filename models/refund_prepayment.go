// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RefundPrepayment represents a RefundPrepayment struct.
type RefundPrepayment struct {
    // `amount` is not required if you pass `amount_in_cents`.
    AmountInCents        *int64                 `json:"amount_in_cents"`
    // `amount_in_cents` is not required if you pass `amount`.
    Amount               RefundPrepaymentAmount `json:"amount"`
    Memo                 string                 `json:"memo"`
    // Specify the type of refund you wish to initiate. When the prepayment is external, the `external` flag is optional. But if the prepayment was made through a payment profile, the `external` flag is required.
    External             *bool                  `json:"external,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RefundPrepayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundPrepayment) String() string {
    return fmt.Sprintf(
    	"RefundPrepayment[AmountInCents=%v, Amount=%v, Memo=%v, External=%v, AdditionalProperties=%v]",
    	r.AmountInCents, r.Amount, r.Memo, r.External, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepayment.
// It customizes the JSON marshaling process for RefundPrepayment objects.
func (r RefundPrepayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "amount_in_cents", "amount", "memo", "external"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepayment object to a map representation for JSON marshaling.
func (r RefundPrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AmountInCents != nil {
        structMap["amount_in_cents"] = r.AmountInCents
    } else {
        structMap["amount_in_cents"] = nil
    }
    structMap["amount"] = r.Amount.toMap()
    structMap["memo"] = r.Memo
    if r.External != nil {
        structMap["external"] = r.External
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepayment.
// It customizes the JSON unmarshaling process for RefundPrepayment objects.
func (r *RefundPrepayment) UnmarshalJSON(input []byte) error {
    var temp tempRefundPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount_in_cents", "amount", "memo", "external")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.AmountInCents = temp.AmountInCents
    r.Amount = *temp.Amount
    r.Memo = *temp.Memo
    r.External = temp.External
    return nil
}

// tempRefundPrepayment is a temporary struct used for validating the fields of RefundPrepayment.
type tempRefundPrepayment  struct {
    AmountInCents *int64                  `json:"amount_in_cents"`
    Amount        *RefundPrepaymentAmount `json:"amount"`
    Memo          *string                 `json:"memo"`
    External      *bool                   `json:"external,omitempty"`
}

func (r *tempRefundPrepayment) validate() error {
    var errs []string
    if r.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Refund Prepayment`")
    }
    if r.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Refund Prepayment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
