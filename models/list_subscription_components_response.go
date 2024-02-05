package models

import (
    "encoding/json"
)

// ListSubscriptionComponentsResponse represents a ListSubscriptionComponentsResponse struct.
type ListSubscriptionComponentsResponse struct {
    SubscriptionsComponents []SubscriptionComponent `json:"subscriptions_components"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON marshaling process for ListSubscriptionComponentsResponse objects.
func (l *ListSubscriptionComponentsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsResponse object to a map representation for JSON marshaling.
func (l *ListSubscriptionComponentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["subscriptions_components"] = l.SubscriptionsComponents
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsResponse objects.
func (l *ListSubscriptionComponentsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        SubscriptionsComponents []SubscriptionComponent `json:"subscriptions_components"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.SubscriptionsComponents = temp.SubscriptionsComponents
    return nil
}
