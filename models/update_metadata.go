// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// UpdateMetadata represents a UpdateMetadata struct.
type UpdateMetadata struct {
    CurrentName          *string                `json:"current_name,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Value                *string                `json:"value,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateMetadata,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateMetadata) String() string {
    return fmt.Sprintf(
    	"UpdateMetadata[CurrentName=%v, Name=%v, Value=%v, AdditionalProperties=%v]",
    	u.CurrentName, u.Name, u.Value, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetadata.
// It customizes the JSON marshaling process for UpdateMetadata objects.
func (u UpdateMetadata) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "current_name", "name", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetadata object to a map representation for JSON marshaling.
func (u UpdateMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CurrentName != nil {
        structMap["current_name"] = u.CurrentName
    }
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Value != nil {
        structMap["value"] = u.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetadata.
// It customizes the JSON unmarshaling process for UpdateMetadata objects.
func (u *UpdateMetadata) UnmarshalJSON(input []byte) error {
    var temp tempUpdateMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "current_name", "name", "value")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.CurrentName = temp.CurrentName
    u.Name = temp.Name
    u.Value = temp.Value
    return nil
}

// tempUpdateMetadata is a temporary struct used for validating the fields of UpdateMetadata.
type tempUpdateMetadata  struct {
    CurrentName *string `json:"current_name,omitempty"`
    Name        *string `json:"name,omitempty"`
    Value       *string `json:"value,omitempty"`
}
