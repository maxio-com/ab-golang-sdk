/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// OfferDiscount represents a OfferDiscount struct.
type OfferDiscount struct {
    CouponCode           *string                `json:"coupon_code,omitempty"`
    CouponId             *int                   `json:"coupon_id,omitempty"`
    CouponName           *string                `json:"coupon_name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OfferDiscount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OfferDiscount) String() string {
    return fmt.Sprintf(
    	"OfferDiscount[CouponCode=%v, CouponId=%v, CouponName=%v, AdditionalProperties=%v]",
    	o.CouponCode, o.CouponId, o.CouponName, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OfferDiscount.
// It customizes the JSON marshaling process for OfferDiscount objects.
func (o OfferDiscount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "coupon_code", "coupon_id", "coupon_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OfferDiscount object to a map representation for JSON marshaling.
func (o OfferDiscount) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CouponCode != nil {
        structMap["coupon_code"] = o.CouponCode
    }
    if o.CouponId != nil {
        structMap["coupon_id"] = o.CouponId
    }
    if o.CouponName != nil {
        structMap["coupon_name"] = o.CouponName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferDiscount.
// It customizes the JSON unmarshaling process for OfferDiscount objects.
func (o *OfferDiscount) UnmarshalJSON(input []byte) error {
    var temp tempOfferDiscount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "coupon_code", "coupon_id", "coupon_name")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.CouponCode = temp.CouponCode
    o.CouponId = temp.CouponId
    o.CouponName = temp.CouponName
    return nil
}

// tempOfferDiscount is a temporary struct used for validating the fields of OfferDiscount.
type tempOfferDiscount  struct {
    CouponCode *string `json:"coupon_code,omitempty"`
    CouponId   *int    `json:"coupon_id,omitempty"`
    CouponName *string `json:"coupon_name,omitempty"`
}
