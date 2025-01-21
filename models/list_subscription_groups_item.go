/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ListSubscriptionGroupsItem represents a ListSubscriptionGroupsItem struct.
type ListSubscriptionGroupsItem struct {
    Uid                   *string                    `json:"uid,omitempty"`
    Scheme                *int                       `json:"scheme,omitempty"`
    CustomerId            *int                       `json:"customer_id,omitempty"`
    PaymentProfileId      *int                       `json:"payment_profile_id,omitempty"`
    SubscriptionIds       []int                      `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId *int                       `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt      *time.Time                 `json:"next_assessment_at,omitempty"`
    State                 *string                    `json:"state,omitempty"`
    CancelAtEndOfPeriod   *bool                      `json:"cancel_at_end_of_period,omitempty"`
    AccountBalances       *SubscriptionGroupBalances `json:"account_balances,omitempty"`
    GroupType             *GroupType                 `json:"group_type,omitempty"`
    AdditionalProperties  map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionGroupsItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionGroupsItem) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionGroupsItem[Uid=%v, Scheme=%v, CustomerId=%v, PaymentProfileId=%v, SubscriptionIds=%v, PrimarySubscriptionId=%v, NextAssessmentAt=%v, State=%v, CancelAtEndOfPeriod=%v, AccountBalances=%v, GroupType=%v, AdditionalProperties=%v]",
    	l.Uid, l.Scheme, l.CustomerId, l.PaymentProfileId, l.SubscriptionIds, l.PrimarySubscriptionId, l.NextAssessmentAt, l.State, l.CancelAtEndOfPeriod, l.AccountBalances, l.GroupType, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsItem.
// It customizes the JSON marshaling process for ListSubscriptionGroupsItem objects.
func (l ListSubscriptionGroupsItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "uid", "scheme", "customer_id", "payment_profile_id", "subscription_ids", "primary_subscription_id", "next_assessment_at", "state", "cancel_at_end_of_period", "account_balances", "group_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsItem object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupsItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Uid != nil {
        structMap["uid"] = l.Uid
    }
    if l.Scheme != nil {
        structMap["scheme"] = l.Scheme
    }
    if l.CustomerId != nil {
        structMap["customer_id"] = l.CustomerId
    }
    if l.PaymentProfileId != nil {
        structMap["payment_profile_id"] = l.PaymentProfileId
    }
    if l.SubscriptionIds != nil {
        structMap["subscription_ids"] = l.SubscriptionIds
    }
    if l.PrimarySubscriptionId != nil {
        structMap["primary_subscription_id"] = l.PrimarySubscriptionId
    }
    if l.NextAssessmentAt != nil {
        structMap["next_assessment_at"] = l.NextAssessmentAt.Format(time.RFC3339)
    }
    if l.State != nil {
        structMap["state"] = l.State
    }
    if l.CancelAtEndOfPeriod != nil {
        structMap["cancel_at_end_of_period"] = l.CancelAtEndOfPeriod
    }
    if l.AccountBalances != nil {
        structMap["account_balances"] = l.AccountBalances.toMap()
    }
    if l.GroupType != nil {
        structMap["group_type"] = l.GroupType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupsItem.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupsItem objects.
func (l *ListSubscriptionGroupsItem) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionGroupsItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "scheme", "customer_id", "payment_profile_id", "subscription_ids", "primary_subscription_id", "next_assessment_at", "state", "cancel_at_end_of_period", "account_balances", "group_type")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Uid = temp.Uid
    l.Scheme = temp.Scheme
    l.CustomerId = temp.CustomerId
    l.PaymentProfileId = temp.PaymentProfileId
    l.SubscriptionIds = temp.SubscriptionIds
    l.PrimarySubscriptionId = temp.PrimarySubscriptionId
    if temp.NextAssessmentAt != nil {
        NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
        }
        l.NextAssessmentAt = &NextAssessmentAtVal
    }
    l.State = temp.State
    l.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
    l.AccountBalances = temp.AccountBalances
    l.GroupType = temp.GroupType
    return nil
}

// tempListSubscriptionGroupsItem is a temporary struct used for validating the fields of ListSubscriptionGroupsItem.
type tempListSubscriptionGroupsItem  struct {
    Uid                   *string                    `json:"uid,omitempty"`
    Scheme                *int                       `json:"scheme,omitempty"`
    CustomerId            *int                       `json:"customer_id,omitempty"`
    PaymentProfileId      *int                       `json:"payment_profile_id,omitempty"`
    SubscriptionIds       []int                      `json:"subscription_ids,omitempty"`
    PrimarySubscriptionId *int                       `json:"primary_subscription_id,omitempty"`
    NextAssessmentAt      *string                    `json:"next_assessment_at,omitempty"`
    State                 *string                    `json:"state,omitempty"`
    CancelAtEndOfPeriod   *bool                      `json:"cancel_at_end_of_period,omitempty"`
    AccountBalances       *SubscriptionGroupBalances `json:"account_balances,omitempty"`
    GroupType             *GroupType                 `json:"group_type,omitempty"`
}
