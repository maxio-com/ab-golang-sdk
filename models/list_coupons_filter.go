// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ListCouponsFilter represents a ListCouponsFilter struct.
type ListCouponsFilter struct {
    // The type of filter you would like to apply to your search. Use in query `filter[date_field]=created_at`.
    DateField            *BasicDateField        `json:"date_field,omitempty"`
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns coupons with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. Use in query `filter[start_date]=2011-12-17`.
    StartDate            *time.Time             `json:"start_date,omitempty"`
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns coupons with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. Use in query `filter[end_date]=2011-12-15`.
    EndDate              *time.Time             `json:"end_date,omitempty"`
    // The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns coupons with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of start_date. Use in query `filter[start_datetime]=2011-12-19T10:15:30+01:00`.
    StartDatetime        *time.Time             `json:"start_datetime,omitempty"`
    // The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns coupons with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site's time zone will be used. If provided, this parameter will be used instead of end_date. Use in query `filter[end_datetime]=2011-12-1T10:15:30+01:00`.
    EndDatetime          *time.Time             `json:"end_datetime,omitempty"`
    // Allows fetching coupons with matching id based on provided values. Use in query `filter[ids]=1,2,3`.
    Ids                  []int                  `json:"ids,omitempty"`
    // Allows fetching coupons with matching codes based on provided values. Use in query `filter[codes]=free,free_trial`.
    Codes                []string               `json:"codes,omitempty"`
    // Allows fetching coupons with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`.
    UseSiteExchangeRate  *bool                  `json:"use_site_exchange_rate,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListCouponsFilter,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListCouponsFilter) String() string {
    return fmt.Sprintf(
    	"ListCouponsFilter[DateField=%v, StartDate=%v, EndDate=%v, StartDatetime=%v, EndDatetime=%v, Ids=%v, Codes=%v, UseSiteExchangeRate=%v, AdditionalProperties=%v]",
    	l.DateField, l.StartDate, l.EndDate, l.StartDatetime, l.EndDatetime, l.Ids, l.Codes, l.UseSiteExchangeRate, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListCouponsFilter.
// It customizes the JSON marshaling process for ListCouponsFilter objects.
func (l ListCouponsFilter) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "date_field", "start_date", "end_date", "start_datetime", "end_datetime", "ids", "codes", "use_site_exchange_rate"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListCouponsFilter object to a map representation for JSON marshaling.
func (l ListCouponsFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    if l.Ids != nil {
        structMap["ids"] = l.Ids
    }
    if l.Codes != nil {
        structMap["codes"] = l.Codes
    }
    if l.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = l.UseSiteExchangeRate
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListCouponsFilter.
// It customizes the JSON unmarshaling process for ListCouponsFilter objects.
func (l *ListCouponsFilter) UnmarshalJSON(input []byte) error {
    var temp tempListCouponsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "date_field", "start_date", "end_date", "start_datetime", "end_datetime", "ids", "codes", "use_site_exchange_rate")
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
    l.Ids = temp.Ids
    l.Codes = temp.Codes
    l.UseSiteExchangeRate = temp.UseSiteExchangeRate
    return nil
}

// tempListCouponsFilter is a temporary struct used for validating the fields of ListCouponsFilter.
type tempListCouponsFilter  struct {
    DateField           *BasicDateField `json:"date_field,omitempty"`
    StartDate           *string         `json:"start_date,omitempty"`
    EndDate             *string         `json:"end_date,omitempty"`
    StartDatetime       *string         `json:"start_datetime,omitempty"`
    EndDatetime         *string         `json:"end_datetime,omitempty"`
    Ids                 []int           `json:"ids,omitempty"`
    Codes               []string        `json:"codes,omitempty"`
    UseSiteExchangeRate *bool           `json:"use_site_exchange_rate,omitempty"`
}
