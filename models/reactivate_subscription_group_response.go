package models

import (
    "encoding/json"
)

// ReactivateSubscriptionGroupResponse represents a ReactivateSubscriptionGroupResponse struct.
type ReactivateSubscriptionGroupResponse struct {
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

// MarshalJSON implements the json.Marshaler interface for ReactivateSubscriptionGroupResponse.
// It customizes the JSON marshaling process for ReactivateSubscriptionGroupResponse objects.
func (r *ReactivateSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivateSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (r *ReactivateSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
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
        structMap["next_assessment_at"] = r.NextAssessmentAt
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
    temp := &struct {
        Uid                   *string `json:"uid,omitempty"`
        Scheme                *int    `json:"scheme,omitempty"`
        CustomerId            *int    `json:"customer_id,omitempty"`
        PaymentProfileId      *int    `json:"payment_profile_id,omitempty"`
        SubscriptionIds       []int   `json:"subscription_ids,omitempty"`
        PrimarySubscriptionId *int    `json:"primary_subscription_id,omitempty"`
        NextAssessmentAt      *string `json:"next_assessment_at,omitempty"`
        State                 *string `json:"state,omitempty"`
        CancelAtEndOfPeriod   *bool   `json:"cancel_at_end_of_period,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.Uid = temp.Uid
    r.Scheme = temp.Scheme
    r.CustomerId = temp.CustomerId
    r.PaymentProfileId = temp.PaymentProfileId
    r.SubscriptionIds = temp.SubscriptionIds
    r.PrimarySubscriptionId = temp.PrimarySubscriptionId
    r.NextAssessmentAt = temp.NextAssessmentAt
    r.State = temp.State
    r.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    return nil
}
