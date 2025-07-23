// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// IssueAdvanceInvoiceRequest represents a IssueAdvanceInvoiceRequest struct.
type IssueAdvanceInvoiceRequest struct {
    Force                *bool                  `json:"force,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IssueAdvanceInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssueAdvanceInvoiceRequest) String() string {
    return fmt.Sprintf(
    	"IssueAdvanceInvoiceRequest[Force=%v, AdditionalProperties=%v]",
    	i.Force, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssueAdvanceInvoiceRequest.
// It customizes the JSON marshaling process for IssueAdvanceInvoiceRequest objects.
func (i IssueAdvanceInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "force"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssueAdvanceInvoiceRequest object to a map representation for JSON marshaling.
func (i IssueAdvanceInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Force != nil {
        structMap["force"] = i.Force
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueAdvanceInvoiceRequest.
// It customizes the JSON unmarshaling process for IssueAdvanceInvoiceRequest objects.
func (i *IssueAdvanceInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempIssueAdvanceInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "force")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Force = temp.Force
    return nil
}

// tempIssueAdvanceInvoiceRequest is a temporary struct used for validating the fields of IssueAdvanceInvoiceRequest.
type tempIssueAdvanceInvoiceRequest  struct {
    Force *bool `json:"force,omitempty"`
}
