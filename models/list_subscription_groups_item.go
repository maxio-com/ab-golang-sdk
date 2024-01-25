package models

import (
	"encoding/json"
)

// ListSubscriptionGroupsItem represents a ListSubscriptionGroupsItem struct.
type ListSubscriptionGroupsItem struct {
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
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsItem.
// It customizes the JSON marshaling process for ListSubscriptionGroupsItem objects.
func (l *ListSubscriptionGroupsItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsItem object to a map representation for JSON marshaling.
func (l *ListSubscriptionGroupsItem) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["next_assessment_at"] = l.NextAssessmentAt
	}
	if l.State != nil {
		structMap["state"] = l.State
	}
	if l.CancelAtEndOfPeriod != nil {
		structMap["cancel_at_end_of_period"] = l.CancelAtEndOfPeriod
	}
	if l.AccountBalances != nil {
		structMap["account_balances"] = l.AccountBalances
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupsItem.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupsItem objects.
func (l *ListSubscriptionGroupsItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
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
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Uid = temp.Uid
	l.Scheme = temp.Scheme
	l.CustomerId = temp.CustomerId
	l.PaymentProfileId = temp.PaymentProfileId
	l.SubscriptionIds = temp.SubscriptionIds
	l.PrimarySubscriptionId = temp.PrimarySubscriptionId
	l.NextAssessmentAt = temp.NextAssessmentAt
	l.State = temp.State
	l.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
	l.AccountBalances = temp.AccountBalances
	return nil
}
