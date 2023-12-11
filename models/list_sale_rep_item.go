package models

import (
	"encoding/json"
)

// ListSaleRepItem represents a ListSaleRepItem struct.
type ListSaleRepItem struct {
	Id                 *int                      `json:"id,omitempty"`
	FullName           *string                   `json:"full_name,omitempty"`
	SubscriptionsCount *int                      `json:"subscriptions_count,omitempty"`
	MrrData            map[string]SaleRepItemMrr `json:"mrr_data,omitempty"`
	TestMode           *bool                     `json:"test_mode,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListSaleRepItem.
// It customizes the JSON marshaling process for ListSaleRepItem objects.
func (l *ListSaleRepItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSaleRepItem object to a map representation for JSON marshaling.
func (l *ListSaleRepItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	if l.FullName != nil {
		structMap["full_name"] = l.FullName
	}
	if l.SubscriptionsCount != nil {
		structMap["subscriptions_count"] = l.SubscriptionsCount
	}
	if l.MrrData != nil {
		structMap["mrr_data"] = l.MrrData
	}
	if l.TestMode != nil {
		structMap["test_mode"] = l.TestMode
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSaleRepItem.
// It customizes the JSON unmarshaling process for ListSaleRepItem objects.
func (l *ListSaleRepItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                 *int                      `json:"id,omitempty"`
		FullName           *string                   `json:"full_name,omitempty"`
		SubscriptionsCount *int                      `json:"subscriptions_count,omitempty"`
		MrrData            map[string]SaleRepItemMrr `json:"mrr_data,omitempty"`
		TestMode           *bool                     `json:"test_mode,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Id = temp.Id
	l.FullName = temp.FullName
	l.SubscriptionsCount = temp.SubscriptionsCount
	l.MrrData = temp.MrrData
	l.TestMode = temp.TestMode
	return nil
}
