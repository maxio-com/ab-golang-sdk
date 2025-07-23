// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// AutoResume represents a AutoResume struct.
type AutoResume struct {
	AutomaticallyResumeAt Optional[time.Time]    `json:"automatically_resume_at"`
	AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AutoResume,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AutoResume) String() string {
	return fmt.Sprintf(
		"AutoResume[AutomaticallyResumeAt=%v, AdditionalProperties=%v]",
		a.AutomaticallyResumeAt, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AutoResume.
// It customizes the JSON marshaling process for AutoResume objects.
func (a AutoResume) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"automatically_resume_at"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AutoResume object to a map representation for JSON marshaling.
func (a AutoResume) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.AutomaticallyResumeAt.IsValueSet() {
		var AutomaticallyResumeAtVal *string = nil
		if a.AutomaticallyResumeAt.Value() != nil {
			val := a.AutomaticallyResumeAt.Value().Format(time.RFC3339)
			AutomaticallyResumeAtVal = &val
		}
		if a.AutomaticallyResumeAt.Value() != nil {
			structMap["automatically_resume_at"] = AutomaticallyResumeAtVal
		} else {
			structMap["automatically_resume_at"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoResume.
// It customizes the JSON unmarshaling process for AutoResume objects.
func (a *AutoResume) UnmarshalJSON(input []byte) error {
	var temp tempAutoResume
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "automatically_resume_at")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.AutomaticallyResumeAt.ShouldSetValue(temp.AutomaticallyResumeAt.IsValueSet())
	if temp.AutomaticallyResumeAt.Value() != nil {
		AutomaticallyResumeAtVal, err := time.Parse(time.RFC3339, (*temp.AutomaticallyResumeAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse automatically_resume_at as % s format.", time.RFC3339)
		}
		a.AutomaticallyResumeAt.SetValue(&AutomaticallyResumeAtVal)
	}
	return nil
}

// tempAutoResume is a temporary struct used for validating the fields of AutoResume.
type tempAutoResume struct {
	AutomaticallyResumeAt Optional[string] `json:"automatically_resume_at"`
}
