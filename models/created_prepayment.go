/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// CreatedPrepayment represents a CreatedPrepayment struct.
type CreatedPrepayment struct {
    Id                     *int64                 `json:"id,omitempty"`
    SubscriptionId         *int                   `json:"subscription_id,omitempty"`
    AmountInCents          *int64                 `json:"amount_in_cents,omitempty"`
    Memo                   *string                `json:"memo,omitempty"`
    CreatedAt              *time.Time             `json:"created_at,omitempty"`
    StartingBalanceInCents *int64                 `json:"starting_balance_in_cents,omitempty"`
    EndingBalanceInCents   *int64                 `json:"ending_balance_in_cents,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatedPrepayment.
// It customizes the JSON marshaling process for CreatedPrepayment objects.
func (c CreatedPrepayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "subscription_id", "amount_in_cents", "memo", "created_at", "starting_balance_in_cents", "ending_balance_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatedPrepayment object to a map representation for JSON marshaling.
func (c CreatedPrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.SubscriptionId != nil {
        structMap["subscription_id"] = c.SubscriptionId
    }
    if c.AmountInCents != nil {
        structMap["amount_in_cents"] = c.AmountInCents
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.StartingBalanceInCents != nil {
        structMap["starting_balance_in_cents"] = c.StartingBalanceInCents
    }
    if c.EndingBalanceInCents != nil {
        structMap["ending_balance_in_cents"] = c.EndingBalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatedPrepayment.
// It customizes the JSON unmarshaling process for CreatedPrepayment objects.
func (c *CreatedPrepayment) UnmarshalJSON(input []byte) error {
    var temp tempCreatedPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "subscription_id", "amount_in_cents", "memo", "created_at", "starting_balance_in_cents", "ending_balance_in_cents")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.SubscriptionId = temp.SubscriptionId
    c.AmountInCents = temp.AmountInCents
    c.Memo = temp.Memo
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    c.StartingBalanceInCents = temp.StartingBalanceInCents
    c.EndingBalanceInCents = temp.EndingBalanceInCents
    return nil
}

// tempCreatedPrepayment is a temporary struct used for validating the fields of CreatedPrepayment.
type tempCreatedPrepayment  struct {
    Id                     *int64  `json:"id,omitempty"`
    SubscriptionId         *int    `json:"subscription_id,omitempty"`
    AmountInCents          *int64  `json:"amount_in_cents,omitempty"`
    Memo                   *string `json:"memo,omitempty"`
    CreatedAt              *string `json:"created_at,omitempty"`
    StartingBalanceInCents *int64  `json:"starting_balance_in_cents,omitempty"`
    EndingBalanceInCents   *int64  `json:"ending_balance_in_cents,omitempty"`
}
