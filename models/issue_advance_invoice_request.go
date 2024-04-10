package models

import (
    "encoding/json"
)

// IssueAdvanceInvoiceRequest represents a IssueAdvanceInvoiceRequest struct.
type IssueAdvanceInvoiceRequest struct {
    Force                *bool          `json:"force,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueAdvanceInvoiceRequest.
// It customizes the JSON marshaling process for IssueAdvanceInvoiceRequest objects.
func (i IssueAdvanceInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssueAdvanceInvoiceRequest object to a map representation for JSON marshaling.
func (i IssueAdvanceInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Force != nil {
        structMap["force"] = i.Force
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueAdvanceInvoiceRequest.
// It customizes the JSON unmarshaling process for IssueAdvanceInvoiceRequest objects.
func (i *IssueAdvanceInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp issueAdvanceInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "force")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Force = temp.Force
    return nil
}

// TODO
type issueAdvanceInvoiceRequest  struct {
    Force *bool `json:"force,omitempty"`
}
