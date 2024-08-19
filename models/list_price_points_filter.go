/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "log"
    "time"
)

// ListPricePointsFilter represents a ListPricePointsFilter struct.
type ListPricePointsFilter struct {
    // The type of filter you would like to apply to your search. Use in query: `filter[date_field]=created_at`.
    DateField            *BasicDateField       `json:"date_field,omitempty"`
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified.
    StartDate            *time.Time            `json:"start_date,omitempty"`
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns price points with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified.
    EndDate              *time.Time            `json:"end_date,omitempty"`
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date.
    StartDatetime        *time.Time            `json:"start_datetime,omitempty"`
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns price points with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date.
    EndDatetime          *time.Time            `json:"end_datetime,omitempty"`
    // Allows fetching price points with matching type. Use in query: `filter[type]=custom,catalog`.
    Type                 []PricePointType      `json:"type,omitempty"`
    // Allows fetching price points with matching id based on provided values. Use in query: `filter[ids]=1,2,3`.
    Ids                  []int                 `json:"ids,omitempty"`
    // Allows fetching price points only if archived_at is present or not. Use in query: `filter[archived_at]=not_null`.
    ArchivedAt           *IncludeNullOrNotNull `json:"archived_at,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListPricePointsFilter.
// It customizes the JSON marshaling process for ListPricePointsFilter objects.
func (l ListPricePointsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListPricePointsFilter object to a map representation for JSON marshaling.
func (l ListPricePointsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.DateField != nil {
        structMap["date_field"] = l.DateField
    }
    if l.StartDate != nil {
        structMap["start_date"] = l.StartDate.Format(DEFAULT_DATE)
    }
    if l.EndDate != nil {
        structMap["end_date"] = l.EndDate.Format(DEFAULT_DATE)
    }
    if l.StartDatetime != nil {
        structMap["start_datetime"] = l.StartDatetime.Format(time.RFC3339)
    }
    if l.EndDatetime != nil {
        structMap["end_datetime"] = l.EndDatetime.Format(time.RFC3339)
    }
    if l.Type != nil {
        structMap["type"] = l.Type
    }
    if l.Ids != nil {
        structMap["ids"] = l.Ids
    }
    if l.ArchivedAt != nil {
        structMap["archived_at"] = l.ArchivedAt
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPricePointsFilter.
// It customizes the JSON unmarshaling process for ListPricePointsFilter objects.
func (l *ListPricePointsFilter) UnmarshalJSON(input []byte) error {
    var temp tempListPricePointsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "date_field", "start_date", "end_date", "start_datetime", "end_datetime", "type", "ids", "archived_at")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.DateField = temp.DateField
    if temp.StartDate != nil {
        StartDateVal, err := time.Parse(DEFAULT_DATE, *temp.StartDate)
        if err != nil {
            log.Fatalf("Cannot Parse start_date as % s format.", DEFAULT_DATE)
        }
        l.StartDate = &StartDateVal
    }
    if temp.EndDate != nil {
        EndDateVal, err := time.Parse(DEFAULT_DATE, *temp.EndDate)
        if err != nil {
            log.Fatalf("Cannot Parse end_date as % s format.", DEFAULT_DATE)
        }
        l.EndDate = &EndDateVal
    }
    if temp.StartDatetime != nil {
        StartDatetimeVal, err := time.Parse(time.RFC3339, *temp.StartDatetime)
        if err != nil {
            log.Fatalf("Cannot Parse start_datetime as % s format.", time.RFC3339)
        }
        l.StartDatetime = &StartDatetimeVal
    }
    if temp.EndDatetime != nil {
        EndDatetimeVal, err := time.Parse(time.RFC3339, *temp.EndDatetime)
        if err != nil {
            log.Fatalf("Cannot Parse end_datetime as % s format.", time.RFC3339)
        }
        l.EndDatetime = &EndDatetimeVal
    }
    l.Type = temp.Type
    l.Ids = temp.Ids
    l.ArchivedAt = temp.ArchivedAt
    return nil
}

// tempListPricePointsFilter is a temporary struct used for validating the fields of ListPricePointsFilter.
type tempListPricePointsFilter  struct {
    DateField     *BasicDateField       `json:"date_field,omitempty"`
    StartDate     *string               `json:"start_date,omitempty"`
    EndDate       *string               `json:"end_date,omitempty"`
    StartDatetime *string               `json:"start_datetime,omitempty"`
    EndDatetime   *string               `json:"end_datetime,omitempty"`
    Type          []PricePointType      `json:"type,omitempty"`
    Ids           []int                 `json:"ids,omitempty"`
    ArchivedAt    *IncludeNullOrNotNull `json:"archived_at,omitempty"`
}
