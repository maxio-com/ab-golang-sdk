package models

import (
    "encoding/json"
)

// ReferralValidationResponse represents a ReferralValidationResponse struct.
type ReferralValidationResponse struct {
    ReferralCode         *ReferralCode  `json:"referral_code,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReferralValidationResponse.
// It customizes the JSON marshaling process for ReferralValidationResponse objects.
func (r ReferralValidationResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReferralValidationResponse object to a map representation for JSON marshaling.
func (r ReferralValidationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ReferralCode != nil {
        structMap["referral_code"] = r.ReferralCode.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReferralValidationResponse.
// It customizes the JSON unmarshaling process for ReferralValidationResponse objects.
func (r *ReferralValidationResponse) UnmarshalJSON(input []byte) error {
    var temp referralValidationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "referral_code")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ReferralCode = temp.ReferralCode
    return nil
}

// TODO
type referralValidationResponse  struct {
    ReferralCode *ReferralCode `json:"referral_code,omitempty"`
}
