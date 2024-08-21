/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// SubscriptionGroupSignupSuccessData represents a SubscriptionGroupSignupSuccessData struct.
type SubscriptionGroupSignupSuccessData struct {
    Uid                   string         `json:"uid"`
    Scheme                int            `json:"scheme"`
    CustomerId            int            `json:"customer_id"`
    PaymentProfileId      int            `json:"payment_profile_id"`
    SubscriptionIds       []int          `json:"subscription_ids"`
    PrimarySubscriptionId int            `json:"primary_subscription_id"`
    NextAssessmentAt      time.Time      `json:"next_assessment_at"`
    State                 string         `json:"state"`
    CancelAtEndOfPeriod   bool           `json:"cancel_at_end_of_period"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupSuccessData.
// It customizes the JSON marshaling process for SubscriptionGroupSignupSuccessData objects.
func (s SubscriptionGroupSignupSuccessData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupSuccessData object to a map representation for JSON marshaling.
func (s SubscriptionGroupSignupSuccessData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["uid"] = s.Uid
    structMap["scheme"] = s.Scheme
    structMap["customer_id"] = s.CustomerId
    structMap["payment_profile_id"] = s.PaymentProfileId
    structMap["subscription_ids"] = s.SubscriptionIds
    structMap["primary_subscription_id"] = s.PrimarySubscriptionId
    structMap["next_assessment_at"] = s.NextAssessmentAt.Format(time.RFC3339)
    structMap["state"] = s.State
    structMap["cancel_at_end_of_period"] = s.CancelAtEndOfPeriod
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupSuccessData.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupSuccessData objects.
func (s *SubscriptionGroupSignupSuccessData) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupSignupSuccessData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "scheme", "customer_id", "payment_profile_id", "subscription_ids", "primary_subscription_id", "next_assessment_at", "state", "cancel_at_end_of_period")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Uid = *temp.Uid
    s.Scheme = *temp.Scheme
    s.CustomerId = *temp.CustomerId
    s.PaymentProfileId = *temp.PaymentProfileId
    s.SubscriptionIds = *temp.SubscriptionIds
    s.PrimarySubscriptionId = *temp.PrimarySubscriptionId
    NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
    if err != nil {
        log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
    }
    s.NextAssessmentAt = NextAssessmentAtVal
    s.State = *temp.State
    s.CancelAtEndOfPeriod = *temp.CancelAtEndOfPeriod
    return nil
}

// tempSubscriptionGroupSignupSuccessData is a temporary struct used for validating the fields of SubscriptionGroupSignupSuccessData.
type tempSubscriptionGroupSignupSuccessData  struct {
    Uid                   *string `json:"uid"`
    Scheme                *int    `json:"scheme"`
    CustomerId            *int    `json:"customer_id"`
    PaymentProfileId      *int    `json:"payment_profile_id"`
    SubscriptionIds       *[]int  `json:"subscription_ids"`
    PrimarySubscriptionId *int    `json:"primary_subscription_id"`
    NextAssessmentAt      *string `json:"next_assessment_at"`
    State                 *string `json:"state"`
    CancelAtEndOfPeriod   *bool   `json:"cancel_at_end_of_period"`
}

func (s *tempSubscriptionGroupSignupSuccessData) validate() error {
    var errs []string
    if s.Uid == nil {
        errs = append(errs, "required field `uid` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.Scheme == nil {
        errs = append(errs, "required field `scheme` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.CustomerId == nil {
        errs = append(errs, "required field `customer_id` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.PaymentProfileId == nil {
        errs = append(errs, "required field `payment_profile_id` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.SubscriptionIds == nil {
        errs = append(errs, "required field `subscription_ids` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.PrimarySubscriptionId == nil {
        errs = append(errs, "required field `primary_subscription_id` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.NextAssessmentAt == nil {
        errs = append(errs, "required field `next_assessment_at` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.State == nil {
        errs = append(errs, "required field `state` is missing for type `Subscription Group Signup Success Data`")
    }
    if s.CancelAtEndOfPeriod == nil {
        errs = append(errs, "required field `cancel_at_end_of_period` is missing for type `Subscription Group Signup Success Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
