/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// EndpointResponse represents a EndpointResponse struct.
type EndpointResponse struct {
    Endpoint             *Endpoint              `json:"endpoint,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EndpointResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EndpointResponse) String() string {
    return fmt.Sprintf(
    	"EndpointResponse[Endpoint=%v, AdditionalProperties=%v]",
    	e.Endpoint, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EndpointResponse.
// It customizes the JSON marshaling process for EndpointResponse objects.
func (e EndpointResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "endpoint"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EndpointResponse object to a map representation for JSON marshaling.
func (e EndpointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "endpoint")
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
