/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// AppliedCreditNoteData represents a AppliedCreditNoteData struct.
type AppliedCreditNoteData struct {
    // The UID of the credit note
    Uid                  *string                `json:"uid,omitempty"`
    // The number of the credit note
    Number               *string                `json:"number,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AppliedCreditNoteData.
// It customizes the JSON marshaling process for AppliedCreditNoteData objects.
func (a AppliedCreditNoteData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "uid", "number"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AppliedCreditNoteData object to a map representation for JSON marshaling.
func (a AppliedCreditNoteData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Uid != nil {
        structMap["uid"] = a.Uid
    }
    if a.Number != nil {
        structMap["number"] = a.Number
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AppliedCreditNoteData.
// It customizes the JSON unmarshaling process for AppliedCreditNoteData objects.
func (a *AppliedCreditNoteData) UnmarshalJSON(input []byte) error {
    var temp tempAppliedCreditNoteData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "number")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Uid = temp.Uid
    a.Number = temp.Number
    return nil
}

// tempAppliedCreditNoteData is a temporary struct used for validating the fields of AppliedCreditNoteData.
type tempAppliedCreditNoteData  struct {
    Uid    *string `json:"uid,omitempty"`
    Number *string `json:"number,omitempty"`
}
