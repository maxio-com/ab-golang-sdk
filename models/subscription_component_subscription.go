package models

import (
	"encoding/json"
)

// SubscriptionComponentSubscription represents a SubscriptionComponentSubscription struct.
// An optional object, will be returned if provided `include=subscription` query param.
type SubscriptionComponentSubscription struct {
	State     *string `json:"state,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionComponentSubscription.
// It customizes the JSON marshaling process for SubscriptionComponentSubscription objects.
func (s *SubscriptionComponentSubscription) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionComponentSubscription object to a map representation for JSON marshaling.
func (s *SubscriptionComponentSubscription) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.State != nil {
		structMap["state"] = s.State
	}
	if s.UpdatedAt != nil {
		structMap["updated_at"] = s.UpdatedAt
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionComponentSubscription.
// It customizes the JSON unmarshaling process for SubscriptionComponentSubscription objects.
func (s *SubscriptionComponentSubscription) UnmarshalJSON(input []byte) error {
	temp := &struct {
		State     *string `json:"state,omitempty"`
		UpdatedAt *string `json:"updated_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.State = temp.State
	s.UpdatedAt = temp.UpdatedAt
	return nil
}
