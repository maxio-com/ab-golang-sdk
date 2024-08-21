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

// AllocationPreviewResponse represents a AllocationPreviewResponse struct.
type AllocationPreviewResponse struct {
    AllocationPreview    AllocationPreview `json:"allocation_preview"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewResponse.
// It customizes the JSON marshaling process for AllocationPreviewResponse objects.
func (a AllocationPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AllocationPreviewResponse object to a map representation for JSON marshaling.
func (a AllocationPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["allocation_preview"] = a.AllocationPreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewResponse.
// It customizes the JSON unmarshaling process for AllocationPreviewResponse objects.
func (a *AllocationPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempAllocationPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allocation_preview")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AllocationPreview = *temp.AllocationPreview
    return nil
}

// tempAllocationPreviewResponse is a temporary struct used for validating the fields of AllocationPreviewResponse.
type tempAllocationPreviewResponse  struct {
    AllocationPreview *AllocationPreview `json:"allocation_preview"`
}

func (a *tempAllocationPreviewResponse) validate() error {
    var errs []string
    if a.AllocationPreview == nil {
        errs = append(errs, "required field `allocation_preview` is missing for type `Allocation Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
