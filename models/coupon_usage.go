/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CouponUsage represents a CouponUsage struct.
type CouponUsage struct {
    // The Chargify id of the product
    Id                   *int                   `json:"id,omitempty"`
    // Name of the product
    Name                 *string                `json:"name,omitempty"`
    // Number of times the coupon has been applied
    Signups              *int                   `json:"signups,omitempty"`
    // Dollar amount of customer savings as a result of the coupon.
    Savings              Optional[int]          `json:"savings"`
    // Dollar amount of customer savings as a result of the coupon.
    SavingsInCents       Optional[int64]        `json:"savings_in_cents"`
    // Total revenue of the all subscriptions that have received a discount from this coupon.
    Revenue              Optional[int]          `json:"revenue"`
    // Total revenue of the all subscriptions that have received a discount from this coupon.
    RevenueInCents       *int64                 `json:"revenue_in_cents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponUsage) String() string {
    return fmt.Sprintf(
    	"CouponUsage[Id=%v, Name=%v, Signups=%v, Savings=%v, SavingsInCents=%v, Revenue=%v, RevenueInCents=%v, AdditionalProperties=%v]",
    	c.Id, c.Name, c.Signups, c.Savings, c.SavingsInCents, c.Revenue, c.RevenueInCents, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponUsage.
// It customizes the JSON marshaling process for CouponUsage objects.
func (c CouponUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "name", "signups", "savings", "savings_in_cents", "revenue", "revenue_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponUsage object to a map representation for JSON marshaling.
func (c CouponUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Signups != nil {
        structMap["signups"] = c.Signups
    }
    if c.Savings.IsValueSet() {
        if c.Savings.Value() != nil {
            structMap["savings"] = c.Savings.Value()
        } else {
            structMap["savings"] = nil
        }
    }
    if c.SavingsInCents.IsValueSet() {
        if c.SavingsInCents.Value() != nil {
            structMap["savings_in_cents"] = c.SavingsInCents.Value()
        } else {
            structMap["savings_in_cents"] = nil
        }
    }
    if c.Revenue.IsValueSet() {
        if c.Revenue.Value() != nil {
            structMap["revenue"] = c.Revenue.Value()
        } else {
            structMap["revenue"] = nil
        }
    }
    if c.RevenueInCents != nil {
        structMap["revenue_in_cents"] = c.RevenueInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponUsage.
// It customizes the JSON unmarshaling process for CouponUsage objects.
func (c *CouponUsage) UnmarshalJSON(input []byte) error {
    var temp tempCouponUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "signups", "savings", "savings_in_cents", "revenue", "revenue_in_cents")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.Name = temp.Name
    c.Signups = temp.Signups
    c.Savings = temp.Savings
    c.SavingsInCents = temp.SavingsInCents
    c.Revenue = temp.Revenue
    c.RevenueInCents = temp.RevenueInCents
    return nil
}

// tempCouponUsage is a temporary struct used for validating the fields of CouponUsage.
type tempCouponUsage  struct {
    Id             *int            `json:"id,omitempty"`
    Name           *string         `json:"name,omitempty"`
    Signups        *int            `json:"signups,omitempty"`
    Savings        Optional[int]   `json:"savings"`
    SavingsInCents Optional[int64] `json:"savings_in_cents"`
    Revenue        Optional[int]   `json:"revenue"`
    RevenueInCents *int64          `json:"revenue_in_cents,omitempty"`
}
