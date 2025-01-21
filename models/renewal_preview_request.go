/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// RenewalPreviewRequest represents a RenewalPreviewRequest struct.
type RenewalPreviewRequest struct {
    // An optional array of component definitions to preview. Providing any component definitions here will override the actual components on the subscription (and their quantities), and the billing preview will contain only these components (in addition to any product base fees).
    Components           []RenewalPreviewComponent `json:"components,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for RenewalPreviewRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RenewalPreviewRequest) String() string {
    return fmt.Sprintf(
    	"RenewalPreviewRequest[Components=%v, AdditionalProperties=%v]",
    	r.Components, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewRequest.
// It customizes the JSON marshaling process for RenewalPreviewRequest objects.
func (r RenewalPreviewRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "components"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewRequest object to a map representation for JSON marshaling.
func (r RenewalPreviewRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Components != nil {
        structMap["components"] = r.Components
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewRequest.
// It customizes the JSON unmarshaling process for RenewalPreviewRequest objects.
func (r *RenewalPreviewRequest) UnmarshalJSON(input []byte) error {
    var temp tempRenewalPreviewRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "components")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Components = temp.Components
    return nil
}

// tempRenewalPreviewRequest is a temporary struct used for validating the fields of RenewalPreviewRequest.
type tempRenewalPreviewRequest  struct {
    Components []RenewalPreviewComponent `json:"components,omitempty"`
}
