package models

import (
    "encoding/json"
)

// CouponResponse represents a CouponResponse struct.
type CouponResponse struct {
    Coupon               *Coupon        `json:"coupon,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CouponResponse.
// It customizes the JSON marshaling process for CouponResponse objects.
func (c CouponResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CouponResponse object to a map representation for JSON marshaling.
func (c CouponResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Coupon != nil {
        structMap["coupon"] = c.Coupon.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponResponse.
// It customizes the JSON unmarshaling process for CouponResponse objects.
func (c *CouponResponse) UnmarshalJSON(input []byte) error {
    var temp couponResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "coupon")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Coupon = temp.Coupon
    return nil
}

// TODO
type couponResponse  struct {
    Coupon *Coupon `json:"coupon,omitempty"`
}
