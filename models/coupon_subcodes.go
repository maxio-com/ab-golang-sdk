package models

import (
    "encoding/json"
)

// CouponSubcodes represents a CouponSubcodes struct.
type CouponSubcodes struct {
    Codes []string `json:"codes,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponSubcodes.
// It customizes the JSON marshaling process for CouponSubcodes objects.
func (c *CouponSubcodes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CouponSubcodes object to a map representation for JSON marshaling.
func (c *CouponSubcodes) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Codes != nil {
        structMap["codes"] = c.Codes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponSubcodes.
// It customizes the JSON unmarshaling process for CouponSubcodes objects.
func (c *CouponSubcodes) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Codes []string `json:"codes,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Codes = temp.Codes
    return nil
}
