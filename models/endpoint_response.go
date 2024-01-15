package models

import (
    "encoding/json"
)

// EndpointResponse represents a EndpointResponse struct.
type EndpointResponse struct {
    Endpoint *Endpoint `json:"endpoint,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for EndpointResponse.
// It customizes the JSON marshaling process for EndpointResponse objects.
func (e *EndpointResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EndpointResponse object to a map representation for JSON marshaling.
func (e *EndpointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if e.Endpoint != nil {
        structMap["endpoint"] = e.Endpoint
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EndpointResponse.
// It customizes the JSON unmarshaling process for EndpointResponse objects.
func (e *EndpointResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Endpoint *Endpoint `json:"endpoint,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    e.Endpoint = temp.Endpoint
    return nil
}
