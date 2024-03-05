package models

import (
    "encoding/json"
    "log"
    "time"
)

// Metadata represents a Metadata struct.
type Metadata struct {
    Id          Optional[int]       `json:"id"`
    Value       Optional[string]    `json:"value"`
    ResourceId  Optional[int]       `json:"resource_id"`
    Name        *string             `json:"name,omitempty"`
    DeletedAt   Optional[time.Time] `json:"deleted_at"`
    MetafieldId Optional[int]       `json:"metafield_id"`
}

// MarshalJSON implements the json.Marshaler interface for Metadata.
// It customizes the JSON marshaling process for Metadata objects.
func (m *Metadata) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the Metadata object to a map representation for JSON marshaling.
func (m *Metadata) toMap() map[string]any {
    structMap := make(map[string]any)
    if m.Id.IsValueSet() {
        structMap["id"] = m.Id.Value()
    }
    if m.Value.IsValueSet() {
        structMap["value"] = m.Value.Value()
    }
    if m.ResourceId.IsValueSet() {
        structMap["resource_id"] = m.ResourceId.Value()
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
        structMap["deleted_at"] = DeletedAtVal
    }
    if m.MetafieldId.IsValueSet() {
        structMap["metafield_id"] = m.MetafieldId.Value()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata.
// It customizes the JSON unmarshaling process for Metadata objects.
func (m *Metadata) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id          Optional[int]    `json:"id"`
        Value       Optional[string] `json:"value"`
        ResourceId  Optional[int]    `json:"resource_id"`
        Name        *string          `json:"name,omitempty"`
        DeletedAt   Optional[string] `json:"deleted_at"`
        MetafieldId Optional[int]    `json:"metafield_id"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
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
