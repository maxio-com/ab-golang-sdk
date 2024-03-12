package models

import (
	"encoding/json"
)

// ResumeOptions represents a ResumeOptions struct.
type ResumeOptions struct {
	// Chargify will only attempt to resume the subscription's billing period. If not resumable, the subscription will be left in it's current state.
	RequireResume *bool `json:"require_resume,omitempty"`
	// Indicates whether or not Chargify should clear the subscription's existing balance before attempting to resume the subscription. If subscription cannot be resumed, the balance will remain as it was before the attempt to resume was made.
	ForgiveBalance *bool `json:"forgive_balance,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ResumeOptions.
// It customizes the JSON marshaling process for ResumeOptions objects.
func (r *ResumeOptions) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResumeOptions object to a map representation for JSON marshaling.
func (r *ResumeOptions) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.RequireResume != nil {
		structMap["require_resume"] = r.RequireResume
	}
	if r.ForgiveBalance != nil {
		structMap["forgive_balance"] = r.ForgiveBalance
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResumeOptions.
// It customizes the JSON unmarshaling process for ResumeOptions objects.
func (r *ResumeOptions) UnmarshalJSON(input []byte) error {
	var temp resumeOptions
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	r.RequireResume = temp.RequireResume
	r.ForgiveBalance = temp.ForgiveBalance
	return nil
}

// TODO
type resumeOptions struct {
	RequireResume  *bool `json:"require_resume,omitempty"`
	ForgiveBalance *bool `json:"forgive_balance,omitempty"`
}
