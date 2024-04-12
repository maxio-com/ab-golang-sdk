package models

import (
    "encoding/json"
)

// Metafield represents a Metafield struct.
type Metafield struct {
    Id                   *int                    `json:"id,omitempty"`
    Name                 *string                 `json:"name,omitempty"`
    // Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
    Scope                *MetafieldScope         `json:"scope,omitempty"`
    // the amount of subscriptions this metafield has been applied to in Chargify
    DataCount            *int                    `json:"data_count,omitempty"`
    // Indicates how data should be added to the metafield. For example, a text type is just a string, so a given metafield of this type can have any value attached. On the other hand, dropdown and radio have a set of allowed values that can be input, and appear differently on a Public Signup Page. Defaults to 'text'
    InputType            *MetafieldInput         `json:"input_type,omitempty"`
    Enum                 Optional[MetafieldEnum] `json:"enum"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Metafield.
// It customizes the JSON marshaling process for Metafield objects.
func (m Metafield) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the Metafield object to a map representation for JSON marshaling.
func (m Metafield) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Scope != nil {
        structMap["scope"] = m.Scope.toMap()
    }
    if m.DataCount != nil {
        structMap["data_count"] = m.DataCount
    }
    if m.InputType != nil {
        structMap["input_type"] = m.InputType
    }
    if m.Enum.IsValueSet() {
        if m.Enum.Value() != nil {
            structMap["enum"] = m.Enum.Value().toMap()
        } else {
            structMap["enum"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metafield.
// It customizes the JSON unmarshaling process for Metafield objects.
func (m *Metafield) UnmarshalJSON(input []byte) error {
    var temp metafield
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "name", "scope", "data_count", "input_type", "enum")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Id = temp.Id
    m.Name = temp.Name
    m.Scope = temp.Scope
    m.DataCount = temp.DataCount
    m.InputType = temp.InputType
    m.Enum = temp.Enum
    return nil
}

// TODO
type metafield  struct {
    Id        *int                    `json:"id,omitempty"`
    Name      *string                 `json:"name,omitempty"`
    Scope     *MetafieldScope         `json:"scope,omitempty"`
    DataCount *int                    `json:"data_count,omitempty"`
    InputType *MetafieldInput         `json:"input_type,omitempty"`
    Enum      Optional[MetafieldEnum] `json:"enum"`
}
