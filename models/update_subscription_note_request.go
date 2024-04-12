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
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionNoteRequest objects.
func (u UpdateSubscriptionNoteRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNoteRequest object to a map representation for JSON marshaling.
func (u UpdateSubscriptionNoteRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["note"] = u.Note.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNoteRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNoteRequest objects.
func (u *UpdateSubscriptionNoteRequest) UnmarshalJSON(input []byte) error {
    var temp updateSubscriptionNoteRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "note")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Note = *temp.Note
    return nil
}

// TODO
type updateSubscriptionNoteRequest  struct {
    Note *UpdateSubscriptionNote `json:"note"`
}

func (u *updateSubscriptionNoteRequest) validate() error {
    var errs []string
    if u.Note == nil {
        errs = append(errs, "required field `note` is missing for type `Update Subscription Note Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
