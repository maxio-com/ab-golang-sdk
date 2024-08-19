/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// EndpointResponse represents a EndpointResponse struct.
type EndpointResponse struct {
    Endpoint             *Endpoint      `json:"endpoint,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EndpointResponse.
// It customizes the JSON marshaling process for EndpointResponse objects.
func (e EndpointResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EndpointResponse object to a map representation for JSON marshaling.
func (e EndpointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Endpoint != nil {
        structMap["endpoint"] = e.Endpoint.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EndpointResponse.
// It customizes the JSON unmarshaling process for EndpointResponse objects.
func (e *EndpointResponse) UnmarshalJSON(input []byte) error {
    var temp tempEndpointResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "endpoint")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Endpoint = temp.Endpoint
    return nil
}

// tempEndpointResponse is a temporary struct used for validating the fields of EndpointResponse.
type tempEndpointResponse  struct {
    Endpoint *Endpoint `json:"endpoint,omitempty"`
}
