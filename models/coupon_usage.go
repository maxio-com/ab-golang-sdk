package models

import (
	"encoding/json"
)

// CouponUsage represents a CouponUsage struct.
type CouponUsage struct {
	// The Chargify id of the product
	Id *int `json:"id,omitempty"`
	// Name of the product
	Name *string `json:"name,omitempty"`
	// Number of times the coupon has been applied
	Signups *int `json:"signups,omitempty"`
	// Dollar amount of customer savings as a result of the coupon.
	Savings Optional[int] `json:"savings"`
	// Dollar amount of customer savings as a result of the coupon.
	SavingsInCents Optional[int64] `json:"savings_in_cents"`
	// Total revenue of the all subscriptions that have received a discount from this coupon.
	Revenue Optional[int] `json:"revenue"`
	// Total revenue of the all subscriptions that have received a discount from this coupon.
	RevenueInCents *int64 `json:"revenue_in_cents,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponUsage.
// It customizes the JSON marshaling process for CouponUsage objects.
func (c *CouponUsage) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponUsage object to a map representation for JSON marshaling.
func (c *CouponUsage) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["savings"] = c.Savings.Value()
	}
	if c.SavingsInCents.IsValueSet() {
		structMap["savings_in_cents"] = c.SavingsInCents.Value()
	}
	if c.Revenue.IsValueSet() {
		structMap["revenue"] = c.Revenue.Value()
	}
	if c.RevenueInCents != nil {
		structMap["revenue_in_cents"] = c.RevenueInCents
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponUsage.
// It customizes the JSON unmarshaling process for CouponUsage objects.
func (c *CouponUsage) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id             *int            `json:"id,omitempty"`
		Name           *string         `json:"name,omitempty"`
		Signups        *int            `json:"signups,omitempty"`
		Savings        Optional[int]   `json:"savings"`
		SavingsInCents Optional[int64] `json:"savings_in_cents"`
		Revenue        Optional[int]   `json:"revenue"`
		RevenueInCents *int64          `json:"revenue_in_cents,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.Name = temp.Name
	c.Signups = temp.Signups
	c.Savings = temp.Savings
	c.SavingsInCents = temp.SavingsInCents
	c.Revenue = temp.Revenue
	c.RevenueInCents = temp.RevenueInCents
	return nil
}
