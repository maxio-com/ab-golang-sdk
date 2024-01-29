package models

import (
	"encoding/json"
)

// ListSubscriptionGroupPrepaymentResponse represents a ListSubscriptionGroupPrepaymentResponse struct.
type ListSubscriptionGroupPrepaymentResponse struct {
	Prepayments []ListSubscriptionGroupPrepayment `json:"prepayments"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON marshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l *ListSubscriptionGroupPrepaymentResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupPrepaymentResponse object to a map representation for JSON marshaling.
func (l *ListSubscriptionGroupPrepaymentResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["prepayments"] = l.Prepayments
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l *ListSubscriptionGroupPrepaymentResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Prepayments []ListSubscriptionGroupPrepayment `json:"prepayments"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Prepayments = temp.Prepayments
	return nil
}
