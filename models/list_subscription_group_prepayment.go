package models

import (
	"encoding/json"
)

// ListSubscriptionGroupPrepayment represents a ListSubscriptionGroupPrepayment struct.
type ListSubscriptionGroupPrepayment struct {
	Prepayment ListSubcriptionGroupPrepaymentItem `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupPrepayment.
// It customizes the JSON marshaling process for ListSubscriptionGroupPrepayment objects.
func (l *ListSubscriptionGroupPrepayment) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupPrepayment object to a map representation for JSON marshaling.
func (l *ListSubscriptionGroupPrepayment) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["prepayment"] = l.Prepayment
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepayment.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepayment objects.
func (l *ListSubscriptionGroupPrepayment) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Prepayment ListSubcriptionGroupPrepaymentItem `json:"prepayment"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Prepayment = temp.Prepayment
	return nil
}
