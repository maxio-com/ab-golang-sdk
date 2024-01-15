package models

import (
    "encoding/json"
)

// Metafield represents a Metafield struct.
type Metafield struct {
    Id        *int                  `json:"id,omitempty"`
    Name      *string               `json:"name,omitempty"`
    // Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
    Scope     *MetafieldScope       `json:"scope,omitempty"`
    // the amount of subscriptions this metafield has been applied to in Chargify
    DataCount *int                  `json:"data_count,omitempty"`
    InputType *string               `json:"input_type,omitempty"`
    Enum      Optional[interface{}] `json:"enum"`
}

// MarshalJSON implements the json.Marshaler interface for Metafield.
// It customizes the JSON marshaling process for Metafield objects.
func (m *Metafield) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the Metafield object to a map representation for JSON marshaling.
func (m *Metafield) toMap() map[string]any {
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
    if m.DataCount != nil {
        structMap["data_count"] = m.DataCount
    }
    if m.InputType != nil {
        structMap["input_type"] = m.InputType
    }
    if m.Enum.IsValueSet() {
        structMap["enum"] = m.Enum.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metafield.
// It customizes the JSON unmarshaling process for Metafield objects.
func (m *Metafield) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id        *int                  `json:"id,omitempty"`
        Name      *string               `json:"name,omitempty"`
        Scope     *MetafieldScope       `json:"scope,omitempty"`
        DataCount *int                  `json:"data_count,omitempty"`
        InputType *string               `json:"input_type,omitempty"`
        Enum      Optional[interface{}] `json:"enum"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    m.Id = temp.Id
    m.Name = temp.Name
    m.Scope = temp.Scope
    m.DataCount = temp.DataCount
    m.InputType = temp.InputType
    m.Enum = temp.Enum
    return nil
}
