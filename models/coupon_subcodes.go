/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CouponSubcodes represents a CouponSubcodes struct.
type CouponSubcodes struct {
    Codes                []string               `json:"codes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponSubcodes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponSubcodes) String() string {
    return fmt.Sprintf(
    	"CouponSubcodes[Codes=%v, AdditionalProperties=%v]",
    	c.Codes, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponSubcodes.
// It customizes the JSON marshaling process for CouponSubcodes objects.
func (c CouponSubcodes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "codes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponSubcodes object to a map representation for JSON marshaling.
func (c CouponSubcodes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Codes != nil {
        structMap["codes"] = c.Codes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponSubcodes.
// It customizes the JSON unmarshaling process for CouponSubcodes objects.
func (c *CouponSubcodes) UnmarshalJSON(input []byte) error {
    var temp tempCouponSubcodes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "codes")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Codes = temp.Codes
    return nil
}

// tempCouponSubcodes is a temporary struct used for validating the fields of CouponSubcodes.
type tempCouponSubcodes  struct {
    Codes []string `json:"codes,omitempty"`
}
