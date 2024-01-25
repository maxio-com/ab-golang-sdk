package models

import (
	"encoding/json"
)

// ReplayWebhooksRequest represents a ReplayWebhooksRequest struct.
type ReplayWebhooksRequest struct {
	Ids []int `json:"ids"`
}

// MarshalJSON implements the json.Marshaler interface for ReplayWebhooksRequest.
// It customizes the JSON marshaling process for ReplayWebhooksRequest objects.
func (r *ReplayWebhooksRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ReplayWebhooksRequest object to a map representation for JSON marshaling.
func (r *ReplayWebhooksRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["ids"] = r.Ids
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplayWebhooksRequest.
// It customizes the JSON unmarshaling process for ReplayWebhooksRequest objects.
func (r *ReplayWebhooksRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Ids []int `json:"ids"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.Ids = temp.Ids
	return nil
}
