// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
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
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for Metafield,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Metafield) String() string {
    return fmt.Sprintf(
    	"Metafield[Id=%v, Name=%v, Scope=%v, DataCount=%v, InputType=%v, Enum=%v, AdditionalProperties=%v]",
    	m.Id, m.Name, m.Scope, m.DataCount, m.InputType, m.Enum, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Metafield.
// It customizes the JSON marshaling process for Metafield objects.
func (m Metafield) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "id", "name", "scope", "data_count", "input_type", "enum"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Metafield object to a map representation for JSON marshaling.
func (m Metafield) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    var temp tempMetafield
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "scope", "data_count", "input_type", "enum")
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

// tempMetafield is a temporary struct used for validating the fields of Metafield.
type tempMetafield  struct {
    Id        *int                    `json:"id,omitempty"`
    Name      *string                 `json:"name,omitempty"`
    Scope     *MetafieldScope         `json:"scope,omitempty"`
    DataCount *int                    `json:"data_count,omitempty"`
    InputType *MetafieldInput         `json:"input_type,omitempty"`
    Enum      Optional[MetafieldEnum] `json:"enum"`
}
