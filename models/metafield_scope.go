package models

import (
	"encoding/json"
)

// MetafieldScope represents a MetafieldScope struct.
// Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
type MetafieldScope struct {
	// Include (1) or exclude (0) metafields from the csv export.
	Csv *IncludeOptionEnum `json:"csv,omitempty"`
	// Include (1) or exclude (0) metafields from invoices.
	Invoices *IncludeOptionEnum `json:"invoices,omitempty"`
	// Include (1) or exclude (0) metafields from statements.
	Statements *IncludeOptionEnum `json:"statements,omitempty"`
	// Include (1) or exclude (0) metafields from the portal.
	Portal *IncludeOptionEnum `json:"portal,omitempty"`
	// Include (1) or exclude (0) metafields from being viewable by your ecosystem.
	PublicShow *IncludeOptionEnum `json:"public_show,omitempty"`
	// Include (1) or exclude (0) metafields from being edited by your ecosystem.
	PublicEdit *IncludeOptionEnum `json:"public_edit,omitempty"`
	Hosted     []string           `json:"hosted,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for MetafieldScope.
// It customizes the JSON marshaling process for MetafieldScope objects.
func (m *MetafieldScope) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MetafieldScope object to a map representation for JSON marshaling.
func (m *MetafieldScope) toMap() map[string]any {
	structMap := make(map[string]any)
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
	temp := &struct {
		Csv        *IncludeOptionEnum `json:"csv,omitempty"`
		Invoices   *IncludeOptionEnum `json:"invoices,omitempty"`
		Statements *IncludeOptionEnum `json:"statements,omitempty"`
		Portal     *IncludeOptionEnum `json:"portal,omitempty"`
		PublicShow *IncludeOptionEnum `json:"public_show,omitempty"`
		PublicEdit *IncludeOptionEnum `json:"public_edit,omitempty"`
		Hosted     []string           `json:"hosted,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Csv = temp.Csv
	m.Invoices = temp.Invoices
	m.Statements = temp.Statements
	m.Portal = temp.Portal
	m.PublicShow = temp.PublicShow
	m.PublicEdit = temp.PublicEdit
	m.Hosted = temp.Hosted
	return nil
}
