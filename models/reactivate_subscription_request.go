package models

import (
    "encoding/json"
)

// ReactivateSubscriptionRequest represents a ReactivateSubscriptionRequest struct.
type ReactivateSubscriptionRequest struct {
    // These values are only applicable to subscriptions using calendar billing
    CalendarBilling          *ReactivationBilling                 `json:"calendar_billing,omitempty"`
    // If `true` is sent, the reactivated Subscription will include a trial if one is available. If `false` is sent, the trial period will be ignored.
    IncludeTrial             *bool                                `json:"include_trial,omitempty"`
    // If `true` is passed, the existing subscription balance will NOT be cleared/reset before adding the additional reactivation charges.
    PreserveBalance          *bool                                `json:"preserve_balance,omitempty"`
    // The coupon code to be applied during reactivation.
    CouponCode               *string                              `json:"coupon_code,omitempty"`
    // If true is sent, Chargify will use service credits and prepayments upon reactivation. If false is sent, the service credits and prepayments will be ignored.
    UseCreditsAndPrepayments *bool                                `json:"use_credits_and_prepayments,omitempty"`
    // If `true`, Chargify will attempt to resume the subscription's billing period. if not resumable, the subscription will be reactivated with a new billing period. If `false`: Chargify will only attempt to reactivate the subscription.
    Resume                   *ReactivateSubscriptionRequestResume `json:"resume,omitempty"`
    AdditionalProperties     map[string]any                       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReactivateSubscriptionRequest.
// It customizes the JSON marshaling process for ReactivateSubscriptionRequest objects.
func (r ReactivateSubscriptionRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReactivateSubscriptionRequest object to a map representation for JSON marshaling.
func (r ReactivateSubscriptionRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.CalendarBilling != nil {
        structMap["calendar_billing"] = r.CalendarBilling.toMap()
    }
    if r.IncludeTrial != nil {
        structMap["include_trial"] = r.IncludeTrial
    }
    if r.PreserveBalance != nil {
        structMap["preserve_balance"] = r.PreserveBalance
    }
    if r.CouponCode != nil {
        structMap["coupon_code"] = r.CouponCode
    }
    if r.UseCreditsAndPrepayments != nil {
        structMap["use_credits_and_prepayments"] = r.UseCreditsAndPrepayments
    }
    if r.Resume != nil {
        structMap["resume"] = r.Resume.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivateSubscriptionRequest.
// It customizes the JSON unmarshaling process for ReactivateSubscriptionRequest objects.
func (r *ReactivateSubscriptionRequest) UnmarshalJSON(input []byte) error {
    var temp reactivateSubscriptionRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "calendar_billing", "include_trial", "preserve_balance", "coupon_code", "use_credits_and_prepayments", "resume")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.CalendarBilling = temp.CalendarBilling
    r.IncludeTrial = temp.IncludeTrial
    r.PreserveBalance = temp.PreserveBalance
    r.CouponCode = temp.CouponCode
    r.UseCreditsAndPrepayments = temp.UseCreditsAndPrepayments
    r.Resume = temp.Resume
    return nil
}

// TODO
type reactivateSubscriptionRequest  struct {
    CalendarBilling          *ReactivationBilling                 `json:"calendar_billing,omitempty"`
    IncludeTrial             *bool                                `json:"include_trial,omitempty"`
    PreserveBalance          *bool                                `json:"preserve_balance,omitempty"`
    CouponCode               *string                              `json:"coupon_code,omitempty"`
    UseCreditsAndPrepayments *bool                                `json:"use_credits_and_prepayments,omitempty"`
    Resume                   *ReactivateSubscriptionRequestResume `json:"resume,omitempty"`
}
