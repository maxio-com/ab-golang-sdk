package models

import (
	"encoding/json"
)

// Metadata represents a Metadata struct.
type Metadata struct {
	Id          *int             `json:"id,omitempty"`
	Value       *string          `json:"value,omitempty"`
	ResourceId  *int             `json:"resource_id,omitempty"`
	Name        *string          `json:"name,omitempty"`
	DeletedAt   Optional[string] `json:"deleted_at"`
	MetafieldId *int             `json:"metafield_id,omitempty"`
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
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.Value != nil {
		structMap["value"] = m.Value
	}
	if m.ResourceId != nil {
		structMap["resource_id"] = m.ResourceId
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.DeletedAt.IsValueSet() {
		structMap["deleted_at"] = m.DeletedAt.Value()
	}
	if m.MetafieldId != nil {
		structMap["metafield_id"] = m.MetafieldId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Metadata.
// It customizes the JSON unmarshaling process for Metadata objects.
func (m *Metadata) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          *int             `json:"id,omitempty"`
		Value       *string          `json:"value,omitempty"`
		ResourceId  *int             `json:"resource_id,omitempty"`
		Name        *string          `json:"name,omitempty"`
		DeletedAt   Optional[string] `json:"deleted_at"`
		MetafieldId *int             `json:"metafield_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	m.Id = temp.Id
	m.Value = temp.Value
	m.ResourceId = temp.ResourceId
	m.Name = temp.Name
	m.DeletedAt = temp.DeletedAt
	m.MetafieldId = temp.MetafieldId
	return nil
}
