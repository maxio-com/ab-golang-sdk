/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Metadata represents a Metadata struct.
type Metadata struct {
    Id                   Optional[int]          `json:"id"`
    Value                Optional[string]       `json:"value"`
    ResourceId           Optional[int]          `json:"resource_id"`
    Name                 *string                `json:"name,omitempty"`
    DeletedAt            Optional[time.Time]    `json:"deleted_at"`
    MetafieldId          Optional[int]          `json:"metafield_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Metadata,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Metadata) String() string {
    return fmt.Sprintf(
    	"Metadata[Id=%v, Value=%v, ResourceId=%v, Name=%v, DeletedAt=%v, MetafieldId=%v, AdditionalProperties=%v]",
    	m.Id, m.Value, m.ResourceId, m.Name, m.DeletedAt, m.MetafieldId, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Metadata.
// It customizes the JSON marshaling process for Metadata objects.
func (m Metadata) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "id", "value", "resource_id", "name", "deleted_at", "metafield_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Metadata object to a map representation for JSON marshaling.
func (m Metadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Id.IsValueSet() {
        if m.Id.Value() != nil {
            structMap["id"] = m.Id.Value()
        } else {
            structMap["id"] = nil
        }
    }
    if m.Value.IsValueSet() {
        if m.Value.Value() != nil {
            structMap["value"] = m.Value.Value()
        } else {
            structMap["value"] = nil
        }
    }
    if m.ResourceId.IsValueSet() {
        if m.ResourceId.Value() != nil {
            structMap["resource_id"] = m.ResourceId.Value()
        } else {
            structMap["resource_id"] = nil
        }
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.DeletedAt.IsValueSet() {
        var DeletedAtVal *string = nil
        if m.DeletedAt.Value() != nil {
            val := m.DeletedAt.Value().Format(time.RFC3339)
            DeletedAtVal = &val
        }
        if m.DeletedAt.Value() != nil {
            structMap["deleted_at"] = DeletedAtVal
        } else {
            structMap["deleted_at"] = nil
        }
    }
    if m.MetafieldId.IsValueSet() {
        if m.MetafieldId.Value() != nil {
            structMap["metafield_id"] = m.MetafieldId.Value()
        } else {
            structMap["metafield_id"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata.
// It customizes the JSON unmarshaling process for Metadata objects.
func (m *Metadata) UnmarshalJSON(input []byte) error {
    var temp tempMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "value", "resource_id", "name", "deleted_at", "metafield_id")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Id = temp.Id
    m.Value = temp.Value
    m.ResourceId = temp.ResourceId
    m.Name = temp.Name
    m.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
    if temp.DeletedAt.Value() != nil {
        DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
        }
        m.DeletedAt.SetValue(&DeletedAtVal)
    }
    m.MetafieldId = temp.MetafieldId
    return nil
}

// tempMetadata is a temporary struct used for validating the fields of Metadata.
type tempMetadata  struct {
    Id          Optional[int]    `json:"id"`
    Value       Optional[string] `json:"value"`
    ResourceId  Optional[int]    `json:"resource_id"`
    Name        *string          `json:"name,omitempty"`
    DeletedAt   Optional[string] `json:"deleted_at"`
    MetafieldId Optional[int]    `json:"metafield_id"`
}
