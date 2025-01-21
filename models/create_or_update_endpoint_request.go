/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateOrUpdateEndpointRequest represents a CreateOrUpdateEndpointRequest struct.
// Used to Create or Update Endpoint
type CreateOrUpdateEndpointRequest struct {
    // Used to Create or Update Endpoint
    Endpoint             CreateOrUpdateEndpoint `json:"endpoint"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateOrUpdateEndpointRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateOrUpdateEndpointRequest) String() string {
    return fmt.Sprintf(
    	"CreateOrUpdateEndpointRequest[Endpoint=%v, AdditionalProperties=%v]",
    	c.Endpoint, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON marshaling process for CreateOrUpdateEndpointRequest objects.
func (c CreateOrUpdateEndpointRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "endpoint"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateEndpointRequest object to a map representation for JSON marshaling.
func (c CreateOrUpdateEndpointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["endpoint"] = c.Endpoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateEndpointRequest.
// It customizes the JSON unmarshaling process for CreateOrUpdateEndpointRequest objects.
func (c *CreateOrUpdateEndpointRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateOrUpdateEndpointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "endpoint")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Endpoint = *temp.Endpoint
    return nil
}

// tempCreateOrUpdateEndpointRequest is a temporary struct used for validating the fields of CreateOrUpdateEndpointRequest.
type tempCreateOrUpdateEndpointRequest  struct {
    Endpoint *CreateOrUpdateEndpoint `json:"endpoint"`
}

func (c *tempCreateOrUpdateEndpointRequest) validate() error {
    var errs []string
    if c.Endpoint == nil {
        errs = append(errs, "required field `endpoint` is missing for type `Create or Update Endpoint Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
