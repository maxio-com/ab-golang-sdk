// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SendEmail represents a SendEmail struct.
type SendEmail struct {
    CanExecute           bool                   `json:"can_execute"`
    Url                  string                 `json:"url"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SendEmail,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SendEmail) String() string {
    return fmt.Sprintf(
    	"SendEmail[CanExecute=%v, Url=%v, AdditionalProperties=%v]",
    	s.CanExecute, s.Url, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SendEmail.
// It customizes the JSON marshaling process for SendEmail objects.
func (s SendEmail) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "can_execute", "url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SendEmail object to a map representation for JSON marshaling.
func (s SendEmail) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["can_execute"] = s.CanExecute
    structMap["url"] = s.Url
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SendEmail.
// It customizes the JSON unmarshaling process for SendEmail objects.
func (s *SendEmail) UnmarshalJSON(input []byte) error {
    var temp tempSendEmail
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "can_execute", "url")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CanExecute = *temp.CanExecute
    s.Url = *temp.Url
    return nil
}

// tempSendEmail is a temporary struct used for validating the fields of SendEmail.
type tempSendEmail  struct {
    CanExecute *bool   `json:"can_execute"`
    Url        *string `json:"url"`
}

func (s *tempSendEmail) validate() error {
    var errs []string
    if s.CanExecute == nil {
        errs = append(errs, "required field `can_execute` is missing for type `SendEmail`")
    }
    if s.Url == nil {
        errs = append(errs, "required field `url` is missing for type `SendEmail`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
