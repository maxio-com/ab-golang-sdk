// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// SubscriptionFilter represents a SubscriptionFilter struct.
// Nested filter used for List Subscription Components For Site Filter
type SubscriptionFilter struct {
	// Allows fetching components allocations that belong to the subscription with matching states based on provided values. To use this filter you also have to include the following param in the request `include=subscription`. Use in query `filter[subscription][states]=active,canceled&include=subscription`.
	States []SubscriptionStateFilter `json:"states,omitempty"`
	// The type of filter you'd like to apply to your search. To use this filter you also have to include the following param in the request `include=subscription`.
	DateField *SubscriptionListDateField `json:"date_field,omitempty"`
	// The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. To use this filter you also have to include the following param in the request `include=subscription`.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components that belong to the subscription with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. To use this filter you also have to include the following param in the request `include=subscription`.
	EndDate *time.Time `json:"end_date,omitempty"`
	// The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. To use this filter you also have to include the following param in the request `include=subscription`.
	StartDatetime *time.Time `json:"start_datetime,omitempty"`
	// The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components that belong to the subscription with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. To use this filter you also have to include the following param in the request `include=subscription`.
	EndDatetime          *time.Time             `json:"end_datetime,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionFilter) String() string {
	return fmt.Sprintf(
		"SubscriptionFilter[States=%v, DateField=%v, StartDate=%v, EndDate=%v, StartDatetime=%v, EndDatetime=%v, AdditionalProperties=%v]",
		s.States, s.DateField, s.StartDate, s.EndDate, s.StartDatetime, s.EndDatetime, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionFilter.
// It customizes the JSON marshaling process for SubscriptionFilter objects.
func (s SubscriptionFilter) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"states", "date_field", "start_date", "end_date", "start_datetime", "end_datetime"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionFilter object to a map representation for JSON marshaling.
func (s SubscriptionFilter) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.States != nil {
		structMap["states"] = s.States
	}
	if s.DateField != nil {
		structMap["date_field"] = s.DateField
	}
	if s.StartDate != nil {
		structMap["start_date"] = s.StartDate.Format(DEFAULT_DATE)
	}
	if s.EndDate != nil {
		structMap["end_date"] = s.EndDate.Format(DEFAULT_DATE)
	}
	if s.StartDatetime != nil {
		structMap["start_datetime"] = s.StartDatetime.Format(time.RFC3339)
	}
	if s.EndDatetime != nil {
		structMap["end_datetime"] = s.EndDatetime.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionFilter.
// It customizes the JSON unmarshaling process for SubscriptionFilter objects.
func (s *SubscriptionFilter) UnmarshalJSON(input []byte) error {
	var temp tempSubscriptionFilter
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "states", "date_field", "start_date", "end_date", "start_datetime", "end_datetime")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.States = temp.States
	s.DateField = temp.DateField
	if temp.StartDate != nil {
		StartDateVal, err := time.Parse(DEFAULT_DATE, *temp.StartDate)
		if err != nil {
			log.Fatalf("Cannot Parse start_date as % s format.", DEFAULT_DATE)
		}
		s.StartDate = &StartDateVal
	}
	if temp.EndDate != nil {
		EndDateVal, err := time.Parse(DEFAULT_DATE, *temp.EndDate)
		if err != nil {
			log.Fatalf("Cannot Parse end_date as % s format.", DEFAULT_DATE)
		}
		s.EndDate = &EndDateVal
	}
	if temp.StartDatetime != nil {
		StartDatetimeVal, err := time.Parse(time.RFC3339, *temp.StartDatetime)
		if err != nil {
			log.Fatalf("Cannot Parse start_datetime as % s format.", time.RFC3339)
		}
		s.StartDatetime = &StartDatetimeVal
	}
	if temp.EndDatetime != nil {
		EndDatetimeVal, err := time.Parse(time.RFC3339, *temp.EndDatetime)
		if err != nil {
			log.Fatalf("Cannot Parse end_datetime as % s format.", time.RFC3339)
		}
		s.EndDatetime = &EndDatetimeVal
	}
	return nil
}

// tempSubscriptionFilter is a temporary struct used for validating the fields of SubscriptionFilter.
type tempSubscriptionFilter struct {
	States        []SubscriptionStateFilter  `json:"states,omitempty"`
	DateField     *SubscriptionListDateField `json:"date_field,omitempty"`
	StartDate     *string                    `json:"start_date,omitempty"`
	EndDate       *string                    `json:"end_date,omitempty"`
	StartDatetime *string                    `json:"start_datetime,omitempty"`
	EndDatetime   *string                    `json:"end_datetime,omitempty"`
}
