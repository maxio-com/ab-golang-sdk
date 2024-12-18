/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ReferralValidationResponse represents a ReferralValidationResponse struct.
type ReferralValidationResponse struct {
    ReferralCode         *ReferralCode          `json:"referral_code,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReferralValidationResponse.
// It customizes the JSON marshaling process for ReferralValidationResponse objects.
func (r ReferralValidationResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "referral_code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReferralValidationResponse object to a map representation for JSON marshaling.
func (r ReferralValidationResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ReferralCode != nil {
        structMap["referral_code"] = r.ReferralCode.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReferralValidationResponse.
// It customizes the JSON unmarshaling process for ReferralValidationResponse objects.
func (r *ReferralValidationResponse) UnmarshalJSON(input []byte) error {
    var temp tempReferralValidationResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "referral_code")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ReferralCode = temp.ReferralCode
    return nil
}

// tempReferralValidationResponse is a temporary struct used for validating the fields of ReferralValidationResponse.
type tempReferralValidationResponse  struct {
    ReferralCode *ReferralCode `json:"referral_code,omitempty"`
}
