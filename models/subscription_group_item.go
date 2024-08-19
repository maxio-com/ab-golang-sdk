/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// SubscriptionGroupItem represents a SubscriptionGroupItem struct.
type SubscriptionGroupItem struct {
    Id                      *int             `json:"id,omitempty"`
    Reference               Optional[string] `json:"reference"`
    ProductId               *int             `json:"product_id,omitempty"`
    ProductHandle           Optional[string] `json:"product_handle"`
    ProductPricePointId     *int             `json:"product_price_point_id,omitempty"`
    ProductPricePointHandle *string          `json:"product_price_point_handle,omitempty"`
    Currency                *string          `json:"currency,omitempty"`
    CouponCode              Optional[string] `json:"coupon_code"`
    TotalRevenueInCents     *int64           `json:"total_revenue_in_cents,omitempty"`
    BalanceInCents          *int64           `json:"balance_in_cents,omitempty"`
    AdditionalProperties    map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupItem.
// It customizes the JSON marshaling process for SubscriptionGroupItem objects.
func (s SubscriptionGroupItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupItem object to a map representation for JSON marshaling.
func (s SubscriptionGroupItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.Reference.IsValueSet() {
        if s.Reference.Value() != nil {
            structMap["reference"] = s.Reference.Value()
        } else {
            structMap["reference"] = nil
        }
    }
    if s.ProductId != nil {
        structMap["product_id"] = s.ProductId
    }
    if s.ProductHandle.IsValueSet() {
        if s.ProductHandle.Value() != nil {
            structMap["product_handle"] = s.ProductHandle.Value()
        } else {
            structMap["product_handle"] = nil
        }
    }
    if s.ProductPricePointId != nil {
        structMap["product_price_point_id"] = s.ProductPricePointId
    }
    if s.ProductPricePointHandle != nil {
        structMap["product_price_point_handle"] = s.ProductPricePointHandle
    }
    if s.Currency != nil {
        structMap["currency"] = s.Currency
    }
    if s.CouponCode.IsValueSet() {
        if s.CouponCode.Value() != nil {
            structMap["coupon_code"] = s.CouponCode.Value()
        } else {
            structMap["coupon_code"] = nil
        }
    }
    if s.TotalRevenueInCents != nil {
        structMap["total_revenue_in_cents"] = s.TotalRevenueInCents
    }
    if s.BalanceInCents != nil {
        structMap["balance_in_cents"] = s.BalanceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupItem.
// It customizes the JSON unmarshaling process for SubscriptionGroupItem objects.
func (s *SubscriptionGroupItem) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionGroupItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "reference", "product_id", "product_handle", "product_price_point_id", "product_price_point_handle", "currency", "coupon_code", "total_revenue_in_cents", "balance_in_cents")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.Reference = temp.Reference
    s.ProductId = temp.ProductId
    s.ProductHandle = temp.ProductHandle
    s.ProductPricePointId = temp.ProductPricePointId
    s.ProductPricePointHandle = temp.ProductPricePointHandle
    s.Currency = temp.Currency
    s.CouponCode = temp.CouponCode
    s.TotalRevenueInCents = temp.TotalRevenueInCents
    s.BalanceInCents = temp.BalanceInCents
    return nil
}

// tempSubscriptionGroupItem is a temporary struct used for validating the fields of SubscriptionGroupItem.
type tempSubscriptionGroupItem  struct {
    Id                      *int             `json:"id,omitempty"`
    Reference               Optional[string] `json:"reference"`
    ProductId               *int             `json:"product_id,omitempty"`
    ProductHandle           Optional[string] `json:"product_handle"`
    ProductPricePointId     *int             `json:"product_price_point_id,omitempty"`
    ProductPricePointHandle *string          `json:"product_price_point_handle,omitempty"`
    Currency                *string          `json:"currency,omitempty"`
    CouponCode              Optional[string] `json:"coupon_code"`
    TotalRevenueInCents     *int64           `json:"total_revenue_in_cents,omitempty"`
    BalanceInCents          *int64           `json:"balance_in_cents,omitempty"`
}
