package models

import (
	"encoding/json"
)

// SubscriptionGroupSignupResponse represents a SubscriptionGroupSignupResponse struct.
type SubscriptionGroupSignupResponse struct {
	Uid                   *string                 `json:"uid,omitempty"`
	Scheme                *int                    `json:"scheme,omitempty"`
	CustomerId            *int                    `json:"customer_id,omitempty"`
	PaymentProfileId      *int                    `json:"payment_profile_id,omitempty"`
	SubscriptionIds       []int                   `json:"subscription_ids,omitempty"`
	PrimarySubscriptionId *int                    `json:"primary_subscription_id,omitempty"`
	NextAssessmentAt      *string                 `json:"next_assessment_at,omitempty"`
	State                 *string                 `json:"state,omitempty"`
	CancelAtEndOfPeriod   *bool                   `json:"cancel_at_end_of_period,omitempty"`
	Subscriptions         []SubscriptionGroupItem `json:"subscriptions,omitempty"`
	// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
	PaymentCollectionMethod *PaymentCollectionMethod `json:"payment_collection_method,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupResponse.
// It customizes the JSON marshaling process for SubscriptionGroupSignupResponse objects.
func (s *SubscriptionGroupSignupResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupResponse object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Uid != nil {
		structMap["uid"] = s.Uid
	}
	if s.Scheme != nil {
		structMap["scheme"] = s.Scheme
	}
	if s.CustomerId != nil {
		structMap["customer_id"] = s.CustomerId
	}
	if s.PaymentProfileId != nil {
		structMap["payment_profile_id"] = s.PaymentProfileId
	}
	if s.SubscriptionIds != nil {
		structMap["subscription_ids"] = s.SubscriptionIds
	}
	if s.PrimarySubscriptionId != nil {
		structMap["primary_subscription_id"] = s.PrimarySubscriptionId
	}
	if s.NextAssessmentAt != nil {
		structMap["next_assessment_at"] = s.NextAssessmentAt
	}
	if s.State != nil {
		structMap["state"] = s.State
	}
	if s.CancelAtEndOfPeriod != nil {
		structMap["cancel_at_end_of_period"] = s.CancelAtEndOfPeriod
	}
	if s.Subscriptions != nil {
		structMap["subscriptions"] = s.Subscriptions
	}
	if s.PaymentCollectionMethod != nil {
		structMap["payment_collection_method"] = s.PaymentCollectionMethod
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupResponse.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupResponse objects.
func (s *SubscriptionGroupSignupResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Uid                     *string                  `json:"uid,omitempty"`
		Scheme                  *int                     `json:"scheme,omitempty"`
		CustomerId              *int                     `json:"customer_id,omitempty"`
		PaymentProfileId        *int                     `json:"payment_profile_id,omitempty"`
		SubscriptionIds         []int                    `json:"subscription_ids,omitempty"`
		PrimarySubscriptionId   *int                     `json:"primary_subscription_id,omitempty"`
		NextAssessmentAt        *string                  `json:"next_assessment_at,omitempty"`
		State                   *string                  `json:"state,omitempty"`
		CancelAtEndOfPeriod     *bool                    `json:"cancel_at_end_of_period,omitempty"`
		Subscriptions           []SubscriptionGroupItem  `json:"subscriptions,omitempty"`
		PaymentCollectionMethod *PaymentCollectionMethod `json:"payment_collection_method,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Uid = temp.Uid
	s.Scheme = temp.Scheme
	s.CustomerId = temp.CustomerId
	s.PaymentProfileId = temp.PaymentProfileId
	s.SubscriptionIds = temp.SubscriptionIds
	s.PrimarySubscriptionId = temp.PrimarySubscriptionId
	s.NextAssessmentAt = temp.NextAssessmentAt
	s.State = temp.State
	s.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
	s.Subscriptions = temp.Subscriptions
	s.PaymentCollectionMethod = temp.PaymentCollectionMethod
	return nil
}
