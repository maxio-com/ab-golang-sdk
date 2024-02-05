package models

import (
    "encoding/json"
)

// CreateOrUpdateCoupon represents a CreateOrUpdateCoupon struct.
type CreateOrUpdateCoupon struct {
    Coupon               *interface{}    `json:"coupon,omitempty"`
    // An object where the keys are product_ids and the values are booleans indicating if the coupon should be applicable to the product
    RestrictedProducts   map[string]bool `json:"restricted_products,omitempty"`
    // An object where the keys are component_ids and the values are booleans indicating if the coupon should be applicable to the component
    RestrictedComponents map[string]bool `json:"restricted_components,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateCoupon.
// It customizes the JSON marshaling process for CreateOrUpdateCoupon objects.
func (c *CreateOrUpdateCoupon) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateCoupon object to a map representation for JSON marshaling.
func (c *CreateOrUpdateCoupon) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Coupon != nil {
        structMap["coupon"] = c.Coupon
    }
    if c.RestrictedProducts != nil {
        structMap["restricted_products"] = c.RestrictedProducts
    }
    if c.RestrictedComponents != nil {
        structMap["restricted_components"] = c.RestrictedComponents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateCoupon.
// It customizes the JSON unmarshaling process for CreateOrUpdateCoupon objects.
func (c *CreateOrUpdateCoupon) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Coupon               *interface{}    `json:"coupon,omitempty"`
        RestrictedProducts   map[string]bool `json:"restricted_products,omitempty"`
        RestrictedComponents map[string]bool `json:"restricted_components,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Coupon = temp.Coupon
    c.RestrictedProducts = temp.RestrictedProducts
    c.RestrictedComponents = temp.RestrictedComponents
    return nil
}
