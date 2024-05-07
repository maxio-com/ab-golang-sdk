package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateSubscriptionNote represents a UpdateSubscriptionNote struct.
// Updatable fields for Subscription Note
type UpdateSubscriptionNote struct {
    Body                 string         `json:"body"`
    Sticky               bool           `json:"sticky"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionNote.
// It customizes the JSON marshaling process for UpdateSubscriptionNote objects.
func (u UpdateSubscriptionNote) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionNote object to a map representation for JSON marshaling.
func (u UpdateSubscriptionNote) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["body"] = u.Body
    structMap["sticky"] = u.Sticky
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionNote.
// It customizes the JSON unmarshaling process for UpdateSubscriptionNote objects.
func (u *UpdateSubscriptionNote) UnmarshalJSON(input []byte) error {
    var temp updateSubscriptionNote
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "body", "sticky")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Body = *temp.Body
    u.Sticky = *temp.Sticky
    return nil
}

// updateSubscriptionNote is a temporary struct used for validating the fields of UpdateSubscriptionNote.
type updateSubscriptionNote  struct {
    Body   *string `json:"body"`
    Sticky *bool   `json:"sticky"`
}

func (u *updateSubscriptionNote) validate() error {
    var errs []string
    if u.Body == nil {
        errs = append(errs, "required field `body` is missing for type `Update Subscription Note`")
    }
    if u.Sticky == nil {
        errs = append(errs, "required field `sticky` is missing for type `Update Subscription Note`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
