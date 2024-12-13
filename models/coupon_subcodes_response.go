/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CouponSubcodesResponse represents a CouponSubcodesResponse struct.
type CouponSubcodesResponse struct {
    CreatedCodes         []string               `json:"created_codes,omitempty"`
    DuplicateCodes       []string               `json:"duplicate_codes,omitempty"`
    InvalidCodes         []string               `json:"invalid_codes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CouponSubcodesResponse.
// It customizes the JSON marshaling process for CouponSubcodesResponse objects.
func (c CouponSubcodesResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "created_codes", "duplicate_codes", "invalid_codes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponSubcodesResponse object to a map representation for JSON marshaling.
func (c CouponSubcodesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.CreatedCodes != nil {
        structMap["created_codes"] = c.CreatedCodes
    }
    if c.DuplicateCodes != nil {
        structMap["duplicate_codes"] = c.DuplicateCodes
    }
    if c.InvalidCodes != nil {
        structMap["invalid_codes"] = c.InvalidCodes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponSubcodesResponse.
// It customizes the JSON unmarshaling process for CouponSubcodesResponse objects.
func (c *CouponSubcodesResponse) UnmarshalJSON(input []byte) error {
    var temp tempCouponSubcodesResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_codes", "duplicate_codes", "invalid_codes")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.CreatedCodes = temp.CreatedCodes
    c.DuplicateCodes = temp.DuplicateCodes
    c.InvalidCodes = temp.InvalidCodes
    return nil
}

// tempCouponSubcodesResponse is a temporary struct used for validating the fields of CouponSubcodesResponse.
type tempCouponSubcodesResponse  struct {
    CreatedCodes   []string `json:"created_codes,omitempty"`
    DuplicateCodes []string `json:"duplicate_codes,omitempty"`
    InvalidCodes   []string `json:"invalid_codes,omitempty"`
}
