package models

import (
	"encoding/json"
)

// ListInvoiceEventsResponse represents a ListInvoiceEventsResponse struct.
type ListInvoiceEventsResponse struct {
	Events     []InvoiceEvent `json:"events,omitempty"`
	Page       *int           `json:"page,omitempty"`
	PerPage    *int           `json:"per_page,omitempty"`
	TotalPages *int           `json:"total_pages,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListInvoiceEventsResponse.
// It customizes the JSON marshaling process for ListInvoiceEventsResponse objects.
func (l *ListInvoiceEventsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListInvoiceEventsResponse object to a map representation for JSON marshaling.
func (l *ListInvoiceEventsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Events != nil {
		structMap["events"] = l.Events
	}
	if l.Page != nil {
		structMap["page"] = l.Page
	}
	if l.PerPage != nil {
		structMap["per_page"] = l.PerPage
	}
	if l.TotalPages != nil {
		structMap["total_pages"] = l.TotalPages
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListInvoiceEventsResponse.
// It customizes the JSON unmarshaling process for ListInvoiceEventsResponse objects.
func (l *ListInvoiceEventsResponse) UnmarshalJSON(input []byte) error {
	var temp listInvoiceEventsResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	l.Events = temp.Events
	l.Page = temp.Page
	l.PerPage = temp.PerPage
	l.TotalPages = temp.TotalPages
	return nil
}

// TODO
type listInvoiceEventsResponse struct {
	Events     []InvoiceEvent `json:"events,omitempty"`
	Page       *int           `json:"page,omitempty"`
	PerPage    *int           `json:"per_page,omitempty"`
	TotalPages *int           `json:"total_pages,omitempty"`
}
