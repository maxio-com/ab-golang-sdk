// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// MetafieldScope represents a MetafieldScope struct.
// Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
type MetafieldScope struct {
    // Include (1) or exclude (0) metafields from the csv export.
    Csv                  *IncludeOption         `json:"csv,omitempty"`
    // Include (1) or exclude (0) metafields from invoices.
    Invoices             *IncludeOption         `json:"invoices,omitempty"`
    // Include (1) or exclude (0) metafields from statements.
    Statements           *IncludeOption         `json:"statements,omitempty"`
    // Include (1) or exclude (0) metafields from the portal.
    Portal               *IncludeOption         `json:"portal,omitempty"`
    // Include (1) or exclude (0) metafields from being viewable by your ecosystem.
    PublicShow           *IncludeOption         `json:"public_show,omitempty"`
    // Include (1) or exclude (0) metafields from being edited by your ecosystem.
    PublicEdit           *IncludeOption         `json:"public_edit,omitempty"`
    Hosted               []string               `json:"hosted,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MetafieldScope,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MetafieldScope) String() string {
    return fmt.Sprintf(
    	"MetafieldScope[Csv=%v, Invoices=%v, Statements=%v, Portal=%v, PublicShow=%v, PublicEdit=%v, Hosted=%v, AdditionalProperties=%v]",
    	m.Csv, m.Invoices, m.Statements, m.Portal, m.PublicShow, m.PublicEdit, m.Hosted, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MetafieldScope.
// It customizes the JSON marshaling process for MetafieldScope objects.
func (m MetafieldScope) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "csv", "invoices", "statements", "portal", "public_show", "public_edit", "hosted"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MetafieldScope object to a map representation for JSON marshaling.
func (m MetafieldScope) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Csv != nil {
        structMap["csv"] = m.Csv
    }
    if m.Invoices != nil {
        structMap["invoices"] = m.Invoices
    }
    if m.Statements != nil {
        structMap["statements"] = m.Statements
    }
    if m.Portal != nil {
        structMap["portal"] = m.Portal
    }
    if m.PublicShow != nil {
        structMap["public_show"] = m.PublicShow
    }
    if m.PublicEdit != nil {
        structMap["public_edit"] = m.PublicEdit
    }
    if m.Hosted != nil {
        structMap["hosted"] = m.Hosted
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MetafieldScope.
// It customizes the JSON unmarshaling process for MetafieldScope objects.
func (m *MetafieldScope) UnmarshalJSON(input []byte) error {
    var temp tempMetafieldScope
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "csv", "invoices", "statements", "portal", "public_show", "public_edit", "hosted")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Csv = temp.Csv
    m.Invoices = temp.Invoices
    m.Statements = temp.Statements
    m.Portal = temp.Portal
    m.PublicShow = temp.PublicShow
    m.PublicEdit = temp.PublicEdit
    m.Hosted = temp.Hosted
    return nil
}

// tempMetafieldScope is a temporary struct used for validating the fields of MetafieldScope.
type tempMetafieldScope  struct {
    Csv        *IncludeOption `json:"csv,omitempty"`
    Invoices   *IncludeOption `json:"invoices,omitempty"`
    Statements *IncludeOption `json:"statements,omitempty"`
    Portal     *IncludeOption `json:"portal,omitempty"`
    PublicShow *IncludeOption `json:"public_show,omitempty"`
    PublicEdit *IncludeOption `json:"public_edit,omitempty"`
    Hosted     []string       `json:"hosted,omitempty"`
}
