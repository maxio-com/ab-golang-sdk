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

// UpdateSubscriptionNoteRequest represents a UpdateSubscriptionNoteRequest struct.
// Updatable fields for Subscription Note
type UpdateSubscriptionNoteRequest struct {
    // Updatable fields for Subscription Note
    Note                 UpdateSubscriptionNote `json:"note"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionNoteRequest objects.
func (u UpdateSubscriptionNoteRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "note"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNoteRequest object to a map representation for JSON marshaling.
func (u UpdateSubscriptionNoteRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["note"] = u.Note.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNoteRequest objects.
func (u *UpdateSubscriptionNoteRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSubscriptionNoteRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "note")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Note = *temp.Note
    return nil
}

// tempUpdateSubscriptionNoteRequest is a temporary struct used for validating the fields of UpdateSubscriptionNoteRequest.
type tempUpdateSubscriptionNoteRequest  struct {
    Note *UpdateSubscriptionNote `json:"note"`
}

func (u *tempUpdateSubscriptionNoteRequest) validate() error {
    var errs []string
    if u.Note == nil {
        errs = append(errs, "required field `note` is missing for type `Update Subscription Note Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
