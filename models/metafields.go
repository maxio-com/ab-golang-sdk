package models

import (
	"encoding/json"
)

// Metafields represents a Metafields struct.
type Metafields struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
	Scope *MetafieldScope `json:"scope,omitempty"`
	// Indicates how data should be added to the metafield. For example, a text type is just a string, so a given metafield of this type can have any value attached. On the other hand, dropdown and radio have a set of allowed values that can be input, and appear differently on a Public Signup Page.
	InputType *MetafieldInputEnum `json:"input_type,omitempty"`
	// Only applicable when input_type is radio or dropdown
	Enum []string `json:"enum,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Metafields.
// It customizes the JSON marshaling process for Metafields objects.
func (m *Metafields) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Metafields object to a map representation for JSON marshaling.
func (m *Metafields) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.Scope != nil {
		structMap["scope"] = m.Scope
	}
	if m.InputType != nil {
		structMap["input_type"] = m.InputType
	}
	if m.Enum != nil {
		structMap["enum"] = m.Enum
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metafields.
// It customizes the JSON unmarshaling process for Metafields objects.
func (m *Metafields) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id        *int                `json:"id,omitempty"`
		Name      *string             `json:"name,omitempty"`
		Scope     *MetafieldScope     `json:"scope,omitempty"`
		InputType *MetafieldInputEnum `json:"input_type,omitempty"`
		Enum      []string            `json:"enum,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Id = temp.Id
	m.Name = temp.Name
	m.Scope = temp.Scope
	m.InputType = temp.InputType
	m.Enum = temp.Enum
	return nil
}
