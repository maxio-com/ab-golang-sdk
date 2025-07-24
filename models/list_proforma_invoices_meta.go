// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ListProformaInvoicesMeta represents a ListProformaInvoicesMeta struct.
type ListProformaInvoicesMeta struct {
    TotalCount           *int                   `json:"total_count,omitempty"`
    CurrentPage          *int                   `json:"current_page,omitempty"`
    TotalPages           *int                   `json:"total_pages,omitempty"`
    StatusCode           *int                   `json:"status_code,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListProformaInvoicesMeta,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListProformaInvoicesMeta) String() string {
    return fmt.Sprintf(
    	"ListProformaInvoicesMeta[TotalCount=%v, CurrentPage=%v, TotalPages=%v, StatusCode=%v, AdditionalProperties=%v]",
    	l.TotalCount, l.CurrentPage, l.TotalPages, l.StatusCode, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListProformaInvoicesMeta.
// It customizes the JSON marshaling process for ListProformaInvoicesMeta objects.
func (l ListProformaInvoicesMeta) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "total_count", "current_page", "total_pages", "status_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListProformaInvoicesMeta object to a map representation for JSON marshaling.
func (l ListProformaInvoicesMeta) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp tempListProformaInvoicesMeta
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "total_count", "current_page", "total_pages", "status_code")
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

// tempListProformaInvoicesMeta is a temporary struct used for validating the fields of ListProformaInvoicesMeta.
type tempListProformaInvoicesMeta  struct {
    TotalCount  *int `json:"total_count,omitempty"`
    CurrentPage *int `json:"current_page,omitempty"`
    TotalPages  *int `json:"total_pages,omitempty"`
    StatusCode  *int `json:"status_code,omitempty"`
}
