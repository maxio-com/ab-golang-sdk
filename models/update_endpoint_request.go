package models

import (
	"encoding/json"
)

// UpdateEndpointRequest represents a UpdateEndpointRequest struct.
// Used to Create or Update Endpoint
type UpdateEndpointRequest struct {
	// Used to Create or Update Endpoint
	Endpoint UpdateEndpoint `json:"endpoint"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateEndpointRequest.
// It customizes the JSON marshaling process for UpdateEndpointRequest objects.
func (u *UpdateEndpointRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateEndpointRequest object to a map representation for JSON marshaling.
func (u *UpdateEndpointRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["endpoint"] = u.Endpoint
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateEndpointRequest.
// It customizes the JSON unmarshaling process for UpdateEndpointRequest objects.
func (u *UpdateEndpointRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Endpoint UpdateEndpoint `json:"endpoint"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Endpoint = temp.Endpoint
	return nil
}
