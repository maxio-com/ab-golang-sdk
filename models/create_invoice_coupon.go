// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CreateInvoiceCoupon represents a CreateInvoiceCoupon struct.
type CreateInvoiceCoupon struct {
    Code                 *string                             `json:"code,omitempty"`
    Percentage           *CreateInvoiceCouponPercentage      `json:"percentage,omitempty"`
    Amount               *CreateInvoiceCouponAmount          `json:"amount,omitempty"`
    Description          *string                             `json:"description,omitempty"`
    ProductFamilyId      *CreateInvoiceCouponProductFamilyId `json:"product_family_id,omitempty"`
    // Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.
    CompoundingStrategy  *CompoundingStrategy                `json:"compounding_strategy,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for CreateInvoiceCoupon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateInvoiceCoupon) String() string {
    return fmt.Sprintf(
    	"CreateInvoiceCoupon[Code=%v, Percentage=%v, Amount=%v, Description=%v, ProductFamilyId=%v, CompoundingStrategy=%v, AdditionalProperties=%v]",
    	c.Code, c.Percentage, c.Amount, c.Description, c.ProductFamilyId, c.CompoundingStrategy, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceCoupon.
// It customizes the JSON marshaling process for CreateInvoiceCoupon objects.
func (c CreateInvoiceCoupon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "code", "percentage", "amount", "description", "product_family_id", "compounding_strategy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceCoupon object to a map representation for JSON marshaling.
func (c CreateInvoiceCoupon) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Code != nil {
        structMap["code"] = c.Code
    }
    if c.Percentage != nil {
        structMap["percentage"] = c.Percentage.toMap()
    }
    if c.Amount != nil {
        structMap["amount"] = c.Amount.toMap()
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.ProductFamilyId != nil {
        structMap["product_family_id"] = c.ProductFamilyId.toMap()
    }
    if c.CompoundingStrategy != nil {
        structMap["compounding_strategy"] = c.CompoundingStrategy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceCoupon.
// It customizes the JSON unmarshaling process for CreateInvoiceCoupon objects.
func (c *CreateInvoiceCoupon) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoiceCoupon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "code", "percentage", "amount", "description", "product_family_id", "compounding_strategy")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Code = temp.Code
    c.Percentage = temp.Percentage
    c.Amount = temp.Amount
    c.Description = temp.Description
    c.ProductFamilyId = temp.ProductFamilyId
    c.CompoundingStrategy = temp.CompoundingStrategy
    return nil
}

// tempCreateInvoiceCoupon is a temporary struct used for validating the fields of CreateInvoiceCoupon.
type tempCreateInvoiceCoupon  struct {
    Code                *string                             `json:"code,omitempty"`
    Percentage          *CreateInvoiceCouponPercentage      `json:"percentage,omitempty"`
    Amount              *CreateInvoiceCouponAmount          `json:"amount,omitempty"`
    Description         *string                             `json:"description,omitempty"`
    ProductFamilyId     *CreateInvoiceCouponProductFamilyId `json:"product_family_id,omitempty"`
    CompoundingStrategy *CompoundingStrategy                `json:"compounding_strategy,omitempty"`
}
