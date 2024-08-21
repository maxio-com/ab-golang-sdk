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

// RenewalPreviewResponse represents a RenewalPreviewResponse struct.
type RenewalPreviewResponse struct {
    RenewalPreview       RenewalPreview `json:"renewal_preview"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewResponse.
// It customizes the JSON marshaling process for RenewalPreviewResponse objects.
func (r RenewalPreviewResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewResponse object to a map representation for JSON marshaling.
func (r RenewalPreviewResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["renewal_preview"] = r.RenewalPreview.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewResponse.
// It customizes the JSON unmarshaling process for RenewalPreviewResponse objects.
func (r *RenewalPreviewResponse) UnmarshalJSON(input []byte) error {
    var temp tempRenewalPreviewResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "renewal_preview")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.RenewalPreview = *temp.RenewalPreview
    return nil
}

// tempRenewalPreviewResponse is a temporary struct used for validating the fields of RenewalPreviewResponse.
type tempRenewalPreviewResponse  struct {
    RenewalPreview *RenewalPreview `json:"renewal_preview"`
}

func (r *tempRenewalPreviewResponse) validate() error {
    var errs []string
    if r.RenewalPreview == nil {
        errs = append(errs, "required field `renewal_preview` is missing for type `Renewal Preview Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
