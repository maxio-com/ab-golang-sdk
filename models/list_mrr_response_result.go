/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ListMRRResponseResult represents a ListMRRResponseResult struct.
type ListMRRResponseResult struct {
    Page                 *int                   `json:"page,omitempty"`
    PerPage              *int                   `json:"per_page,omitempty"`
    TotalPages           *int                   `json:"total_pages,omitempty"`
    TotalEntries         *int                   `json:"total_entries,omitempty"`
    Currency             *string                `json:"currency,omitempty"`
    CurrencySymbol       *string                `json:"currency_symbol,omitempty"`
    Movements            []Movement             `json:"movements,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListMRRResponseResult,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListMRRResponseResult) String() string {
    return fmt.Sprintf(
    	"ListMRRResponseResult[Page=%v, PerPage=%v, TotalPages=%v, TotalEntries=%v, Currency=%v, CurrencySymbol=%v, Movements=%v, AdditionalProperties=%v]",
    	l.Page, l.PerPage, l.TotalPages, l.TotalEntries, l.Currency, l.CurrencySymbol, l.Movements, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListMRRResponseResult.
// It customizes the JSON marshaling process for ListMRRResponseResult objects.
func (l ListMRRResponseResult) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "page", "per_page", "total_pages", "total_entries", "currency", "currency_symbol", "movements"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListMRRResponseResult object to a map representation for JSON marshaling.
func (l ListMRRResponseResult) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Page != nil {
        structMap["page"] = l.Page
    }
    if l.PerPage != nil {
        structMap["per_page"] = l.PerPage
    }
    if l.TotalPages != nil {
        structMap["total_pages"] = l.TotalPages
    }
    if l.TotalEntries != nil {
        structMap["total_entries"] = l.TotalEntries
    }
    if l.Currency != nil {
        structMap["currency"] = l.Currency
    }
    if l.CurrencySymbol != nil {
        structMap["currency_symbol"] = l.CurrencySymbol
    }
    if l.Movements != nil {
        structMap["movements"] = l.Movements
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMRRResponseResult.
// It customizes the JSON unmarshaling process for ListMRRResponseResult objects.
func (l *ListMRRResponseResult) UnmarshalJSON(input []byte) error {
    var temp tempListMRRResponseResult
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "page", "per_page", "total_pages", "total_entries", "currency", "currency_symbol", "movements")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Page = temp.Page
    l.PerPage = temp.PerPage
    l.TotalPages = temp.TotalPages
    l.TotalEntries = temp.TotalEntries
    l.Currency = temp.Currency
    l.CurrencySymbol = temp.CurrencySymbol
    l.Movements = temp.Movements
    return nil
}

// tempListMRRResponseResult is a temporary struct used for validating the fields of ListMRRResponseResult.
type tempListMRRResponseResult  struct {
    Page           *int       `json:"page,omitempty"`
    PerPage        *int       `json:"per_page,omitempty"`
    TotalPages     *int       `json:"total_pages,omitempty"`
    TotalEntries   *int       `json:"total_entries,omitempty"`
    Currency       *string    `json:"currency,omitempty"`
    CurrencySymbol *string    `json:"currency_symbol,omitempty"`
    Movements      []Movement `json:"movements,omitempty"`
}
