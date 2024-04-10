package models

import (
    "encoding/json"
    "log"
    "time"
)

// FullSubscriptionGroupResponse represents a FullSubscriptionGroupResponse struct.
type FullSubscriptionGroupResponse struct {
    Uid                         *string                    `json:"uid,omitempty"`
    Scheme                      *int                       `json:"scheme,omitempty"`
    CustomerId                  *int                       `json:"customer_id,omitempty"`
    PaymentProfileId            *int                       `json:"payment_profile_id,omitempty"`
    SubscriptionIds             []int                      `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId       *int                       `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt            *time.Time                 `json:"next_assessment_at,omitempty"`
    // The state of a subscription.
    // * **Live States**
    //     * `active` - A normal, active subscription. It is not in a trial and is paid and up to date.
    //     * `assessing` - An internal (transient) state that indicates a subscription is in the middle of periodic assessment. Do not base any access decisions in your app on this state, as it may not always be exposed.
    //     * `pending` - An internal (transient) state that indicates a subscription is in the creation process. Do not base any access decisions in your app on this state, as it may not always be exposed.
    //     * `trialing` - A subscription in trialing state has a valid trial subscription. This type of subscription may transition to active once payment is received when the trial has ended. Otherwise, it may go to a Problem or End of Life state.
    //     * `paused` - An internal state that indicates that your account with Advanced Billing is in arrears.
    // * **Problem States**
    //     * `past_due` - Indicates that the most recent payment has failed, and payment is past due for this subscription. If you have enabled our automated dunning, this subscription will be in the dunning process (additional status and callbacks from the dunning process will be available in the future). If you are handling dunning and payment updates yourself, you will want to use this state to initiate a payment update from your customers.
    //     * `soft_failure` - Indicates that normal assessment/processing of the subscription has failed for a reason that cannot be fixed by the Customer. For example, a Soft Fail may result from a timeout at the gateway or incorrect credentials on your part. The subscriptions should be retried automatically. An interface is being built for you to review problems resulting from these events to take manual action when needed.
    //     * `unpaid` - Indicates an unpaid subscription. A subscription is marked unpaid if the retry period expires and you have configured your [Dunning](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405505141005) settings to have a Final Action of `mark the subscription unpaid`.
    // * **End of Life States**
    //     * `canceled` - Indicates a canceled subscription. This may happen at your request (via the API or the web interface) or due to the expiration of the [Dunning](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405505141005) process without payment. See the [Reactivation](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404559291021) documentation for info on how to restart a canceled subscription.
    //     While a subscription is canceled, its period will not advance, it will not accrue any new charges, and Advanced Billing will not attempt to collect the overdue balance.
    //     * `expired` - Indicates a subscription that has expired due to running its normal life cycle. Some products may be configured to have an expiration period. An expired subscription then is one that stayed active until it fulfilled its full period.
    //     * `failed_to_create` - Indicates that signup has failed. (You may see this state in a signup_failure webhook.)
    //     * `on_hold` - Indicates that a subscription’s billing has been temporarily stopped. While it is expected that the subscription will resume and return to active status, this is still treated as an “End of Life” state because the customer is not paying for services during this time.
    //     * `suspended` - Indicates that a prepaid subscription has used up all their prepayment balance. If a prepayment is applied, it will return to an active state.
    //     * `trial_ended` - A subscription in a trial_ended state is a subscription that completed a no-obligation trial and did not have a card on file at the expiration of the trial period. See [Product Pricing – No Obligation Trials](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405246782221) for more details.
    // See [Subscription States](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404222005773) for more info about subscription states and state transitions.
    State                       *SubscriptionState         `json:"state,omitempty"`
    CancelAtEndOfPeriod         *bool                      `json:"cancel_at_end_of_period,omitempty"`
    CurrentBillingAmountInCents *int64                     `json:"current_billing_amount_in_cents,omitempty"`
    Customer                    *SubscriptionGroupCustomer `json:"customer,omitempty"`
    AccountBalances             *SubscriptionGroupBalances `json:"account_balances,omitempty"`
    AdditionalProperties        map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for FullSubscriptionGroupResponse.
// It customizes the JSON marshaling process for FullSubscriptionGroupResponse objects.
func (f FullSubscriptionGroupResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(f.toMap())
}

// toMap converts the FullSubscriptionGroupResponse object to a map representation for JSON marshaling.
func (f FullSubscriptionGroupResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, f.AdditionalProperties)
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
        structMap["next_assessment_at"] = f.NextAssessmentAt.Format(time.RFC3339)
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
    var temp fullSubscriptionGroupResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "scheme", "customer_id", "payment_profile_id", "subscription_ids", "primary_subscription_id", "next_assessment_at", "state", "cancel_at_end_of_period", "current_billing_amount_in_cents", "customer", "account_balances")
    if err != nil {
    	return err
    }
    
    f.AdditionalProperties = additionalProperties
    f.Uid = temp.Uid
    f.Scheme = temp.Scheme
    f.CustomerId = temp.CustomerId
    f.PaymentProfileId = temp.PaymentProfileId
    f.SubscriptionIds = temp.SubscriptionIds
    f.PrimarySubscriptionId = temp.PrimarySubscriptionId
    if temp.NextAssessmentAt != nil {
        NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
        }
        f.NextAssessmentAt = &NextAssessmentAtVal
    }
    f.State = temp.State
    f.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    f.CurrentBillingAmountInCents = temp.CurrentBillingAmountInCents
    f.Customer = temp.Customer
    f.AccountBalances = temp.AccountBalances
    return nil
}

// TODO
type fullSubscriptionGroupResponse  struct {
    Uid                         *string                    `json:"uid,omitempty"`
    Scheme                      *int                       `json:"scheme,omitempty"`
    CustomerId                  *int                       `json:"customer_id,omitempty"`
    PaymentProfileId            *int                       `json:"payment_profile_id,omitempty"`
    SubscriptionIds             []int                      `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId       *int                       `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt            *string                    `json:"next_assessment_at,omitempty"`
    State                       *SubscriptionState         `json:"state,omitempty"`
    CancelAtEndOfPeriod         *bool                      `json:"cancel_at_end_of_period,omitempty"`
    CurrentBillingAmountInCents *int64                     `json:"current_billing_amount_in_cents,omitempty"`
    Customer                    *SubscriptionGroupCustomer `json:"customer,omitempty"`
    AccountBalances             *SubscriptionGroupBalances `json:"account_balances,omitempty"`
}
