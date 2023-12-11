package models

import (
	"encoding/json"
)

// ReferralValidationResponse represents a ReferralValidationResponse struct.
type ReferralValidationResponse struct {
	ReferralCode *ReferralCode `json:"referral_code,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ReferralValidationResponse.
// It customizes the JSON marshaling process for ReferralValidationResponse objects.
func (r *ReferralValidationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ReferralValidationResponse object to a map representation for JSON marshaling.
func (r *ReferralValidationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.ReferralCode != nil {
		structMap["referral_code"] = r.ReferralCode
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReferralValidationResponse.
// It customizes the JSON unmarshaling process for ReferralValidationResponse objects.
func (r *ReferralValidationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ReferralCode *ReferralCode `json:"referral_code,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	r.ReferralCode = temp.ReferralCode
	return nil
}
