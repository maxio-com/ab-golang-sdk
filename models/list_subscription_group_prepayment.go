package models

import (
	"encoding/json"
	"errors"
	"strings"
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
	structMap["prepayment"] = l.Prepayment.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepayment.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepayment objects.
func (l *ListSubscriptionGroupPrepayment) UnmarshalJSON(input []byte) error {
	var temp listSubscriptionGroupPrepayment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	l.Prepayment = *temp.Prepayment
	return nil
}

// TODO
type listSubscriptionGroupPrepayment struct {
	Prepayment *ListSubcriptionGroupPrepaymentItem `json:"prepayment"`
}

func (l *listSubscriptionGroupPrepayment) validate() error {
	var errs []string
	if l.Prepayment == nil {
		errs = append(errs, "required field `prepayment` is missing for type `List Subscription Group Prepayment`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
