/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// UpdateMetafield represents a UpdateMetafield struct.
type UpdateMetafield struct {
    CurrentName          *string                `json:"current_name,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // Warning: When updating a metafield's scope attribute, all scope attributes must be passed. Partially complete scope attributes will override the existing settings.
    Scope                *MetafieldScope        `json:"scope,omitempty"`
    // Indicates how data should be added to the metafield. For example, a text type is just a string, so a given metafield of this type can have any value attached. On the other hand, dropdown and radio have a set of allowed values that can be input, and appear differently on a Public Signup Page. Defaults to 'text'
    InputType            *MetafieldInput        `json:"input_type,omitempty"`
    // Only applicable when input_type is radio or dropdown
    Enum                 []string               `json:"enum,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateMetafield,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateMetafield) String() string {
    return fmt.Sprintf(
    	"UpdateMetafield[CurrentName=%v, Name=%v, Scope=%v, InputType=%v, Enum=%v, AdditionalProperties=%v]",
    	u.CurrentName, u.Name, u.Scope, u.InputType, u.Enum, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetafield.
// It customizes the JSON marshaling process for UpdateMetafield objects.
func (u UpdateMetafield) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "current_name", "name", "scope", "input_type", "enum"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetafield object to a map representation for JSON marshaling.
func (u UpdateMetafield) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CurrentName != nil {
        structMap["current_name"] = u.CurrentName
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Scope != nil {
        structMap["scope"] = u.Scope.toMap()
    }
    if u.InputType != nil {
        structMap["input_type"] = u.InputType
    }
    if u.Enum != nil {
        structMap["enum"] = u.Enum
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetafield.
// It customizes the JSON unmarshaling process for UpdateMetafield objects.
func (u *UpdateMetafield) UnmarshalJSON(input []byte) error {
    var temp tempUpdateMetafield
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "current_name", "name", "scope", "input_type", "enum")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.CurrentName = temp.CurrentName
    u.Name = temp.Name
    u.Scope = temp.Scope
    u.InputType = temp.InputType
    u.Enum = temp.Enum
    return nil
}

// tempUpdateMetafield is a temporary struct used for validating the fields of UpdateMetafield.
type tempUpdateMetafield  struct {
    CurrentName *string         `json:"current_name,omitempty"`
    Name        *string         `json:"name,omitempty"`
    Scope       *MetafieldScope `json:"scope,omitempty"`
    InputType   *MetafieldInput `json:"input_type,omitempty"`
    Enum        []string        `json:"enum,omitempty"`
}
