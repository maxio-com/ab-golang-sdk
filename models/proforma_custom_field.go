package models

import (
	"encoding/json"
)

// ProformaCustomField represents a ProformaCustomField struct.
type ProformaCustomField struct {
	OwnerId     *int    `json:"owner_id,omitempty"`
	OwnerType   *string `json:"owner_type,omitempty"`
	Name        *string `json:"name,omitempty"`
	Value       *string `json:"value,omitempty"`
	MetadatumId *int    `json:"metadatum_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaCustomField.
// It customizes the JSON marshaling process for ProformaCustomField objects.
func (p *ProformaCustomField) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the ProformaCustomField object to a map representation for JSON marshaling.
func (p *ProformaCustomField) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.OwnerId != nil {
		structMap["owner_id"] = p.OwnerId
	}
	if p.OwnerType != nil {
		structMap["owner_type"] = p.OwnerType
	}
	if p.Name != nil {
		structMap["name"] = p.Name
	}
	if p.Value != nil {
		structMap["value"] = p.Value
	}
	if p.MetadatumId != nil {
		structMap["metadatum_id"] = p.MetadatumId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaCustomField.
// It customizes the JSON unmarshaling process for ProformaCustomField objects.
func (p *ProformaCustomField) UnmarshalJSON(input []byte) error {
	temp := &struct {
		OwnerId     *int    `json:"owner_id,omitempty"`
		OwnerType   *string `json:"owner_type,omitempty"`
		Name        *string `json:"name,omitempty"`
		Value       *string `json:"value,omitempty"`
		MetadatumId *int    `json:"metadatum_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.OwnerId = temp.OwnerId
	p.OwnerType = temp.OwnerType
	p.Name = temp.Name
	p.Value = temp.Value
	p.MetadatumId = temp.MetadatumId
	return nil
}
