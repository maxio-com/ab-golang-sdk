package models

import (
    "encoding/json"
    "log"
    "time"
)

// ListPrepaymentsFilter represents a ListPrepaymentsFilter struct.
type ListPrepaymentsFilter struct {
    // The type of filter you would like to apply to your search. `created_at` - Time when prepayment was created. `application_at` - Time when prepayment was applied to invoice. Use in query `filter[date_field]=created_at`.
    DateField            *ListPrepaymentDateField `json:"date_field,omitempty"`
    // The start date (format YYYY-MM-DD) with which to filter the date_field. Returns prepayments with a timestamp at or after midnight (12:00:00 AM) in your site's time zone on the date specified. Use in query: `filter[start_date]=2011-12-15`.
    StartDate            *time.Time               `json:"start_date,omitempty"`
    // The end date (format YYYY-MM-DD) with which to filter the date_field. Returns prepayments with a timestamp up to and including 11:59:59PM in your site's time zone on the date specified. Use in query: `filter[end_date]=2011-12-15`.
    EndDate              *time.Time               `json:"end_date,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListPrepaymentsFilter.
// It customizes the JSON marshaling process for ListPrepaymentsFilter objects.
func (l ListPrepaymentsFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListPrepaymentsFilter object to a map representation for JSON marshaling.
func (l ListPrepaymentsFilter) toMap() map[string]any {
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPrepaymentsFilter.
// It customizes the JSON unmarshaling process for ListPrepaymentsFilter objects.
func (l *ListPrepaymentsFilter) UnmarshalJSON(input []byte) error {
    var temp listPrepaymentsFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "date_field", "start_date", "end_date")
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
    return nil
}

// listPrepaymentsFilter is a temporary struct used for validating the fields of ListPrepaymentsFilter.
type listPrepaymentsFilter  struct {
    DateField *ListPrepaymentDateField `json:"date_field,omitempty"`
    StartDate *string                  `json:"start_date,omitempty"`
    EndDate   *string                  `json:"end_date,omitempty"`
}
