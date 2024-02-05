package models

import (
    "encoding/json"
)

// FullSubscriptionGroupResponse represents a FullSubscriptionGroupResponse struct.
type FullSubscriptionGroupResponse struct {
    Uid                         *string                    `json:"uid,omitempty"`
    Scheme                      *int                       `json:"scheme,omitempty"`
    CustomerId                  *int                       `json:"customer_id,omitempty"`
    PaymentProfileId            *int                       `json:"payment_profile_id,omitempty"`
    SubscriptionIds             []int                      `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId       *int                       `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt            *string                    `json:"next_assessment_at,omitempty"`
    State                       *string                    `json:"state,omitempty"`
    CancelAtEndOfPeriod         *bool                      `json:"cancel_at_end_of_period,omitempty"`
    CurrentBillingAmountInCents *int64                     `json:"current_billing_amount_in_cents,omitempty"`
    Customer                    *SubscriptionGroupCustomer `json:"customer,omitempty"`
    AccountBalances             *SubscriptionGroupBalances `json:"account_balances,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for FullSubscriptionGroupResponse.
// It customizes the JSON marshaling process for FullSubscriptionGroupResponse objects.
func (f *FullSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(f.toMap())
}

// toMap converts the FullSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (f *FullSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if f.Uid != nil {
        structMap["uid"] = f.Uid
    }
    if f.Scheme != nil {
        structMap["scheme"] = f.Scheme
    }
    if f.CustomerId != nil {
        structMap["customer_id"] = f.CustomerId
    }
    if f.PaymentProfileId != nil {
        structMap["payment_profile_id"] = f.PaymentProfileId
    }
    if f.SubscriptionIds != nil {
        structMap["subscription_ids"] = f.SubscriptionIds
    }
    if f.PrimarySubscriptionId != nil {
        structMap["primary_subscription_id"] = f.PrimarySubscriptionId
    }
    if f.NextAssessmentAt != nil {
        structMap["next_assessment_at"] = f.NextAssessmentAt
    }
    if f.State != nil {
        structMap["state"] = f.State
    }
    if f.CancelAtEndOfPeriod != nil {
        structMap["cancel_at_end_of_period"] = f.CancelAtEndOfPeriod
    }
    if f.CurrentBillingAmountInCents != nil {
        structMap["current_billing_amount_in_cents"] = f.CurrentBillingAmountInCents
    }
    if f.Customer != nil {
        structMap["customer"] = f.Customer.toMap()
    }
    if f.AccountBalances != nil {
        structMap["account_balances"] = f.AccountBalances.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FullSubscriptionGroupResponse.
// It customizes the JSON unmarshaling process for FullSubscriptionGroupResponse objects.
func (f *FullSubscriptionGroupResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid                         *string                    `json:"uid,omitempty"`
        Scheme                      *int                       `json:"scheme,omitempty"`
        CustomerId                  *int                       `json:"customer_id,omitempty"`
        PaymentProfileId            *int                       `json:"payment_profile_id,omitempty"`
        SubscriptionIds             []int                      `json:"subscription_ids,omitempty"`
        PrimarySubscriptionId       *int                       `json:"primary_subscription_id,omitempty"`
        NextAssessmentAt            *string                    `json:"next_assessment_at,omitempty"`
        State                       *string                    `json:"state,omitempty"`
        CancelAtEndOfPeriod         *bool                      `json:"cancel_at_end_of_period,omitempty"`
        CurrentBillingAmountInCents *int64                     `json:"current_billing_amount_in_cents,omitempty"`
        Customer                    *SubscriptionGroupCustomer `json:"customer,omitempty"`
        AccountBalances             *SubscriptionGroupBalances `json:"account_balances,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    f.Uid = temp.Uid
    f.Scheme = temp.Scheme
    f.CustomerId = temp.CustomerId
    f.PaymentProfileId = temp.PaymentProfileId
    f.SubscriptionIds = temp.SubscriptionIds
    f.PrimarySubscriptionId = temp.PrimarySubscriptionId
    f.NextAssessmentAt = temp.NextAssessmentAt
    f.State = temp.State
    f.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    f.CurrentBillingAmountInCents = temp.CurrentBillingAmountInCents
    f.Customer = temp.Customer
    f.AccountBalances = temp.AccountBalances
    return nil
}
