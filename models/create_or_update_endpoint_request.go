package models

import (
    "encoding/json"
)

// CreateOrUpdateEndpointRequest represents a CreateOrUpdateEndpointRequest struct.
// Used to Create or Update Endpoint
type CreateOrUpdateEndpointRequest struct {
    // Used to Create or Update Endpoint
    Endpoint CreateOrUpdateEndpoint `json:"endpoint"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON marshaling process for CreateOrUpdateEndpointRequest objects.
func (c *CreateOrUpdateEndpointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateEndpointRequest object to a map representation for JSON marshaling.
func (c *CreateOrUpdateEndpointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["endpoint"] = c.Endpoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateEndpointRequest objects.
func (c *CreateOrUpdateEndpointRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Endpoint CreateOrUpdateEndpoint `json:"endpoint"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Endpoint = temp.Endpoint
    return nil
}
