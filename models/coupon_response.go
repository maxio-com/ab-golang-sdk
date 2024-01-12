package models

import (
    "encoding/json"
)

// CouponResponse represents a CouponResponse struct.
type CouponResponse struct {
    Coupon *Coupon `json:"coupon,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponResponse.
// It customizes the JSON marshaling process for CouponResponse objects.
func (c *CouponResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CouponResponse object to a map representation for JSON marshaling.
func (c *CouponResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Coupon != nil {
        structMap["coupon"] = c.Coupon
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponResponse.
// It customizes the JSON unmarshaling process for CouponResponse objects.
func (c *CouponResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Coupon *Coupon `json:"coupon,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Coupon = temp.Coupon
    return nil
}
