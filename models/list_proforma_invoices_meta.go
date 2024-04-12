package models

import (
    "encoding/json"
)

// ListProformaInvoicesMeta represents a ListProformaInvoicesMeta struct.
type ListProformaInvoicesMeta struct {
    TotalCount           *int           `json:"total_count,omitempty"`
    CurrentPage          *int           `json:"current_page,omitempty"`
    TotalPages           *int           `json:"total_pages,omitempty"`
    StatusCode           *int           `json:"status_code,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListProformaInvoicesMeta.
// It customizes the JSON marshaling process for ListProformaInvoicesMeta objects.
func (l ListProformaInvoicesMeta) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListProformaInvoicesMeta object to a map representation for JSON marshaling.
func (l ListProformaInvoicesMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    if l.TotalCount != nil {
        structMap["total_count"] = l.TotalCount
    }
    if l.CurrentPage != nil {
        structMap["current_page"] = l.CurrentPage
    }
    if l.TotalPages != nil {
        structMap["total_pages"] = l.TotalPages
    }
    if l.StatusCode != nil {
        structMap["status_code"] = l.StatusCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProformaInvoicesMeta.
// It customizes the JSON unmarshaling process for ListProformaInvoicesMeta objects.
func (l *ListProformaInvoicesMeta) UnmarshalJSON(input []byte) error {
    var temp listProformaInvoicesMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "total_count", "current_page", "total_pages", "status_code")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.TotalCount = temp.TotalCount
    l.CurrentPage = temp.CurrentPage
    l.TotalPages = temp.TotalPages
    l.StatusCode = temp.StatusCode
    return nil
}

// TODO
type listProformaInvoicesMeta  struct {
    TotalCount  *int `json:"total_count,omitempty"`
    CurrentPage *int `json:"current_page,omitempty"`
    TotalPages  *int `json:"total_pages,omitempty"`
    StatusCode  *int `json:"status_code,omitempty"`
}
