/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateAllocationRequest represents a CreateAllocationRequest struct.
type CreateAllocationRequest struct {
    Allocation           CreateAllocation `json:"allocation"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAllocationRequest.
// It customizes the JSON marshaling process for CreateAllocationRequest objects.
func (c CreateAllocationRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateAllocationRequest object to a map representation for JSON marshaling.
func (c CreateAllocationRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["allocation"] = c.Allocation.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAllocationRequest.
// It customizes the JSON unmarshaling process for CreateAllocationRequest objects.
func (c *CreateAllocationRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateAllocationRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Allocation = *temp.Allocation
    return nil
}

// tempCreateAllocationRequest is a temporary struct used for validating the fields of CreateAllocationRequest.
type tempCreateAllocationRequest  struct {
    Allocation *CreateAllocation `json:"allocation"`
}

func (c *tempCreateAllocationRequest) validate() error {
    var errs []string
    if c.Allocation == nil {
        errs = append(errs, "required field `allocation` is missing for type `Create Allocation Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
