package models

import (
	"encoding/json"
	"log"
	"time"
)

// AutoResume represents a AutoResume struct.
type AutoResume struct {
	AutomaticallyResumeAt Optional[time.Time] `json:"automatically_resume_at"`
}

// MarshalJSON implements the json.Marshaler interface for AutoResume.
// It customizes the JSON marshaling process for AutoResume objects.
func (a *AutoResume) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AutoResume object to a map representation for JSON marshaling.
func (a *AutoResume) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.AutomaticallyResumeAt.IsValueSet() {
		var AutomaticallyResumeAtVal *string = nil
		if a.AutomaticallyResumeAt.Value() != nil {
			val := a.AutomaticallyResumeAt.Value().Format(time.RFC3339)
			AutomaticallyResumeAtVal = &val
		}
		structMap["automatically_resume_at"] = AutomaticallyResumeAtVal
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoResume.
// It customizes the JSON unmarshaling process for AutoResume objects.
func (a *AutoResume) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AutomaticallyResumeAt Optional[string] `json:"automatically_resume_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

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
