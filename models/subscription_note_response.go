package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SubscriptionNoteResponse represents a SubscriptionNoteResponse struct.
type SubscriptionNoteResponse struct {
    Note                 SubscriptionNote `json:"note"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionNoteResponse.
// It customizes the JSON marshaling process for SubscriptionNoteResponse objects.
func (s SubscriptionNoteResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionNoteResponse object to a map representation for JSON marshaling.
func (s SubscriptionNoteResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["note"] = s.Note.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionNoteResponse.
// It customizes the JSON unmarshaling process for SubscriptionNoteResponse objects.
func (s *SubscriptionNoteResponse) UnmarshalJSON(input []byte) error {
    var temp subscriptionNoteResponse
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
    
    s.AdditionalProperties = additionalProperties
    s.Note = *temp.Note
    return nil
}

// TODO
type subscriptionNoteResponse  struct {
    Note *SubscriptionNote `json:"note"`
}

func (s *subscriptionNoteResponse) validate() error {
    var errs []string
    if s.Note == nil {
        errs = append(errs, "required field `note` is missing for type `Subscription Note Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
