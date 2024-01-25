package models

import (
	"encoding/json"
)

// SubscriptionGroupResponse represents a SubscriptionGroupResponse struct.
type SubscriptionGroupResponse struct {
	SubscriptionGroup SubscriptionGroup `json:"subscription_group"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupResponse.
// It customizes the JSON marshaling process for SubscriptionGroupResponse objects.
func (s *SubscriptionGroupResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupResponse object to a map representation for JSON marshaling.
func (s *SubscriptionGroupResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["subscription_group"] = s.SubscriptionGroup
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupResponse.
// It customizes the JSON unmarshaling process for SubscriptionGroupResponse objects.
func (s *SubscriptionGroupResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SubscriptionGroup SubscriptionGroup `json:"subscription_group"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.SubscriptionGroup = temp.SubscriptionGroup
	return nil
}
