package models

import (
	"encoding/json"
)

// ListSubscriptionGroupsResponse represents a ListSubscriptionGroupsResponse struct.
type ListSubscriptionGroupsResponse struct {
	SubscriptionGroups []ListSubscriptionGroupsItem `json:"subscription_groups,omitempty"`
	Meta               *ListSubscriptionGroupsMeta  `json:"meta,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupsResponse.
// It customizes the JSON marshaling process for ListSubscriptionGroupsResponse objects.
func (l *ListSubscriptionGroupsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupsResponse object to a map representation for JSON marshaling.
func (l *ListSubscriptionGroupsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.SubscriptionGroups != nil {
		structMap["subscription_groups"] = l.SubscriptionGroups
	}
	if l.Meta != nil {
		structMap["meta"] = l.Meta
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupsResponse objects.
func (l *ListSubscriptionGroupsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SubscriptionGroups []ListSubscriptionGroupsItem `json:"subscription_groups,omitempty"`
		Meta               *ListSubscriptionGroupsMeta  `json:"meta,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.SubscriptionGroups = temp.SubscriptionGroups
	l.Meta = temp.Meta
	return nil
}
