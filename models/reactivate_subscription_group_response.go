package models

import (
    "encoding/json"
    "log"
    "time"
)

// ReactivateSubscriptionGroupResponse represents a ReactivateSubscriptionGroupResponse struct.
type ReactivateSubscriptionGroupResponse struct {
    Uid                   *string        `json:"uid,omitempty"`
    Scheme                *int           `json:"scheme,omitempty"`
    CustomerId            *int           `json:"customer_id,omitempty"`
    PaymentProfileId      *int           `json:"payment_profile_id,omitempty"`
    SubscriptionIds       []int          `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId *int           `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt      *time.Time     `json:"next_assessment_at,omitempty"`
    State                 *string        `json:"state,omitempty"`
    CancelAtEndOfPeriod   *bool          `json:"cancel_at_end_of_period,omitempty"`
    AdditionalProperties  map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReactivateSubscriptionGroupResponse.
// It customizes the JSON marshaling process for ReactivateSubscriptionGroupResponse objects.
func (r ReactivateSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivateSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (r ReactivateSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Uid != nil {
        structMap["uid"] = r.Uid
    }
    if r.Scheme != nil {
        structMap["scheme"] = r.Scheme
    }
    if r.CustomerId != nil {
        structMap["customer_id"] = r.CustomerId
    }
    if r.PaymentProfileId != nil {
        structMap["payment_profile_id"] = r.PaymentProfileId
    }
    if r.SubscriptionIds != nil {
        structMap["subscription_ids"] = r.SubscriptionIds
    }
    if r.PrimarySubscriptionId != nil {
        structMap["primary_subscription_id"] = r.PrimarySubscriptionId
    }
    if r.NextAssessmentAt != nil {
        structMap["next_assessment_at"] = r.NextAssessmentAt.Format(time.RFC3339)
    }
    if r.State != nil {
        structMap["state"] = r.State
    }
    if r.CancelAtEndOfPeriod != nil {
        structMap["cancel_at_end_of_period"] = r.CancelAtEndOfPeriod
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivateSubscriptionGroupResponse.
// It customizes the JSON unmarshaling process for ReactivateSubscriptionGroupResponse objects.
func (r *ReactivateSubscriptionGroupResponse) UnmarshalJSON(input []byte) error {
    var temp reactivateSubscriptionGroupResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "scheme", "customer_id", "payment_profile_id", "subscription_ids", "primary_subscription_id", "next_assessment_at", "state", "cancel_at_end_of_period")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Uid = temp.Uid
    r.Scheme = temp.Scheme
    r.CustomerId = temp.CustomerId
    r.PaymentProfileId = temp.PaymentProfileId
    r.SubscriptionIds = temp.SubscriptionIds
    r.PrimarySubscriptionId = temp.PrimarySubscriptionId
    if temp.NextAssessmentAt != nil {
        NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
        }
        r.NextAssessmentAt = &NextAssessmentAtVal
    }
    r.State = temp.State
    r.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    return nil
}

// TODO
type reactivateSubscriptionGroupResponse  struct {
    Uid                   *string `json:"uid,omitempty"`
    Scheme                *int    `json:"scheme,omitempty"`
    CustomerId            *int    `json:"customer_id,omitempty"`
    PaymentProfileId      *int    `json:"payment_profile_id,omitempty"`
    SubscriptionIds       []int   `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId *int    `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt      *string `json:"next_assessment_at,omitempty"`
    State                 *string `json:"state,omitempty"`
    CancelAtEndOfPeriod   *bool   `json:"cancel_at_end_of_period,omitempty"`
}
