/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionGroupPrepayment represents a SubscriptionGroupPrepayment struct.
type SubscriptionGroupPrepayment struct {
    Amount               int                               `json:"amount"`
    Details              string                            `json:"details"`
    Memo                 string                            `json:"memo"`
    Method               SubscriptionGroupPrepaymentMethod `json:"method"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepayment.
// It customizes the JSON marshaling process for SubscriptionGroupPrepayment objects.
func (s SubscriptionGroupPrepayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepayment object to a map representation for JSON marshaling.
func (s SubscriptionGroupPrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["amount"] = s.Amount
    structMap["details"] = s.Details
    structMap["memo"] = s.Memo
    structMap["method"] = s.Method
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepayment.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepayment objects.
func (s *SubscriptionGroupPrepayment) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "details", "memo", "method")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Amount = *temp.Amount
    s.Details = *temp.Details
    s.Memo = *temp.Memo
    s.Method = *temp.Method
    return nil
}

// tempSubscriptionGroupPrepayment is a temporary struct used for validating the fields of SubscriptionGroupPrepayment.
type tempSubscriptionGroupPrepayment  struct {
    Amount  *int                               `json:"amount"`
    Details *string                            `json:"details"`
    Memo    *string                            `json:"memo"`
    Method  *SubscriptionGroupPrepaymentMethod `json:"method"`
}

func (s *tempSubscriptionGroupPrepayment) validate() error {
    var errs []string
    if s.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Subscription Group Prepayment`")
    }
    if s.Details == nil {
        errs = append(errs, "required field `details` is missing for type `Subscription Group Prepayment`")
    }
    if s.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Subscription Group Prepayment`")
    }
    if s.Method == nil {
        errs = append(errs, "required field `method` is missing for type `Subscription Group Prepayment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
