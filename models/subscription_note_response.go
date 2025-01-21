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

// SubscriptionNoteResponse represents a SubscriptionNoteResponse struct.
type SubscriptionNoteResponse struct {
    Note                 SubscriptionNote       `json:"note"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionNoteResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionNoteResponse) String() string {
    return fmt.Sprintf(
    	"SubscriptionNoteResponse[Note=%v, AdditionalProperties=%v]",
    	s.Note, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNoteResponse.
// It customizes the JSON marshaling process for SubscriptionNoteResponse objects.
func (s SubscriptionNoteResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "note"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNoteResponse object to a map representation for JSON marshaling.
func (s SubscriptionNoteResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["note"] = s.Note.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionNoteResponse.
// It customizes the JSON unmarshaling process for SubscriptionNoteResponse objects.
func (s *SubscriptionNoteResponse) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionNoteResponse
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
    s.AdditionalProperties = additionalProperties
    
    s.Note = *temp.Note
    return nil
}

// tempSubscriptionNoteResponse is a temporary struct used for validating the fields of SubscriptionNoteResponse.
type tempSubscriptionNoteResponse  struct {
    Note *SubscriptionNote `json:"note"`
}

func (s *tempSubscriptionNoteResponse) validate() error {
    var errs []string
    if s.Note == nil {
        errs = append(errs, "required field `note` is missing for type `Subscription Note Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
