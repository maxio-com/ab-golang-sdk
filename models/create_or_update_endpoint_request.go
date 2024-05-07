package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateOrUpdateEndpointRequest represents a CreateOrUpdateEndpointRequest struct.
// Used to Create or Update Endpoint
type CreateOrUpdateEndpointRequest struct {
    // Used to Create or Update Endpoint
    Endpoint             CreateOrUpdateEndpoint `json:"endpoint"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON marshaling process for CreateOrUpdateEndpointRequest objects.
func (c CreateOrUpdateEndpointRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateEndpointRequest object to a map representation for JSON marshaling.
func (c CreateOrUpdateEndpointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["endpoint"] = c.Endpoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateEndpointRequest objects.
func (c *CreateOrUpdateEndpointRequest) UnmarshalJSON(input []byte) error {
    var temp createOrUpdateEndpointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "endpoint")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Endpoint = *temp.Endpoint
    return nil
}

// createOrUpdateEndpointRequest is a temporary struct used for validating the fields of CreateOrUpdateEndpointRequest.
type createOrUpdateEndpointRequest  struct {
    Endpoint *CreateOrUpdateEndpoint `json:"endpoint"`
}

func (c *createOrUpdateEndpointRequest) validate() error {
    var errs []string
    if c.Endpoint == nil {
        errs = append(errs, "required field `endpoint` is missing for type `Create or Update Endpoint Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
