/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SubscriptionGroupPrepaymentResponse represents a SubscriptionGroupPrepaymentResponse struct.
type SubscriptionGroupPrepaymentResponse struct {
    Id                   *int                   `json:"id,omitempty"`
    // The amount in cents of the entry.
    AmountInCents        *int64                 `json:"amount_in_cents,omitempty"`
    // The ending balance in cents of the account.
    EndingBalanceInCents *int64                 `json:"ending_balance_in_cents,omitempty"`
    // The type of entry
    EntryType            *ServiceCreditType     `json:"entry_type,omitempty"`
    // A memo attached to the entry.
    Memo                 Optional[string]       `json:"memo"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepaymentResponse.
// It customizes the JSON marshaling process for SubscriptionGroupPrepaymentResponse objects.
func (s SubscriptionGroupPrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepaymentResponse object to a map representation for JSON marshaling.
func (s SubscriptionGroupPrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.AmountInCents != nil {
        structMap["amount_in_cents"] = s.AmountInCents
    }
    if s.EndingBalanceInCents != nil {
        structMap["ending_balance_in_cents"] = s.EndingBalanceInCents
    }
    if s.EntryType != nil {
        structMap["entry_type"] = s.EntryType
    }
    if s.Memo.IsValueSet() {
        if s.Memo.Value() != nil {
            structMap["memo"] = s.Memo.Value()
        } else {
            structMap["memo"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepaymentResponse.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepaymentResponse objects.
func (s *SubscriptionGroupPrepaymentResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupPrepaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "amount_in_cents", "ending_balance_in_cents", "entry_type", "memo")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.AmountInCents = temp.AmountInCents
    s.EndingBalanceInCents = temp.EndingBalanceInCents
    s.EntryType = temp.EntryType
    s.Memo = temp.Memo
    return nil
}

// tempSubscriptionGroupPrepaymentResponse is a temporary struct used for validating the fields of SubscriptionGroupPrepaymentResponse.
type tempSubscriptionGroupPrepaymentResponse  struct {
    Id                   *int               `json:"id,omitempty"`
    AmountInCents        *int64             `json:"amount_in_cents,omitempty"`
    EndingBalanceInCents *int64             `json:"ending_balance_in_cents,omitempty"`
    EntryType            *ServiceCreditType `json:"entry_type,omitempty"`
    Memo                 Optional[string]   `json:"memo"`
}
