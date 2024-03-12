package models

import (
	"encoding/json"
)

// UpdateMetafieldsRequest represents a UpdateMetafieldsRequest struct.
type UpdateMetafieldsRequest struct {
	Metafields *UpdateMetafieldsRequestMetafields `json:"metafields,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetafieldsRequest.
// It customizes the JSON marshaling process for UpdateMetafieldsRequest objects.
func (u *UpdateMetafieldsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetafieldsRequest object to a map representation for JSON marshaling.
func (u *UpdateMetafieldsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.Metafields != nil {
		structMap["metafields"] = u.Metafields.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetafieldsRequest.
// It customizes the JSON unmarshaling process for UpdateMetafieldsRequest objects.
func (u *UpdateMetafieldsRequest) UnmarshalJSON(input []byte) error {
	var temp updateMetafieldsRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	u.Metafields = temp.Metafields
	return nil
}

// TODO
type updateMetafieldsRequest struct {
	Metafields *UpdateMetafieldsRequestMetafields `json:"metafields,omitempty"`
}
