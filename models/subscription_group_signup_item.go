package models

import (
	"encoding/json"
)

// SubscriptionGroupSignupItem represents a SubscriptionGroupSignupItem struct.
type SubscriptionGroupSignupItem struct {
	// The API Handle of the product for which you are creating a subscription. Required, unless a `product_id` is given instead.
	ProductHandle *string `json:"product_handle,omitempty"`
	// The Product ID of the product for which you are creating a subscription. You can pass either `product_id` or `product_handle`.
	ProductId *int `json:"product_id,omitempty"`
	// The ID of the particular price point on the product.
	ProductPricePointId *int `json:"product_price_point_id,omitempty"`
	// The user-friendly API handle of a product's particular price point.
	ProductPricePointHandle *string `json:"product_price_point_handle,omitempty"`
	// Use in place of passing product and component information to set up the subscription with an existing offer. May be either the Chargify ID of the offer or its handle prefixed with `handle:`
	OfferId *int `json:"offer_id,omitempty"`
	// The reference value (provided by your app) for the subscription itelf.
	Reference *string `json:"reference,omitempty"`
	// One of the subscriptions must be marked as primary in the group.
	Primary *bool `json:"primary,omitempty"`
	// (Optional) If Multi-Currency is enabled and the currency is configured in Chargify, pass it at signup to create a subscription on a non-default currency. Note that you cannot update the currency of an existing subscription.
	Currency *string `json:"currency,omitempty"`
	// An array for all the coupons attached to the subscription.
	CouponCodes []string                           `json:"coupon_codes,omitempty"`
	Components  []SubscriptionGroupSignupComponent `json:"components,omitempty"`
	// (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
	CustomPrice *SubscriptionCustomPrice `json:"custom_price,omitempty"`
	// (Optional). Cannot be used when also specifying next_billing_at
	CalendarBilling *CalendarBilling `json:"calendar_billing,omitempty"`
	// (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet.
	Metafields map[string]string `json:"metafields,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupSignupItem.
// It customizes the JSON marshaling process for SubscriptionGroupSignupItem objects.
func (s *SubscriptionGroupSignupItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupSignupItem object to a map representation for JSON marshaling.
func (s *SubscriptionGroupSignupItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.ProductHandle != nil {
		structMap["product_handle"] = s.ProductHandle
	}
	if s.ProductId != nil {
		structMap["product_id"] = s.ProductId
	}
	if s.ProductPricePointId != nil {
		structMap["product_price_point_id"] = s.ProductPricePointId
	}
	if s.ProductPricePointHandle != nil {
		structMap["product_price_point_handle"] = s.ProductPricePointHandle
	}
	if s.OfferId != nil {
		structMap["offer_id"] = s.OfferId
	}
	if s.Reference != nil {
		structMap["reference"] = s.Reference
	}
	if s.Primary != nil {
		structMap["primary"] = s.Primary
	}
	if s.Currency != nil {
		structMap["currency"] = s.Currency
	}
	if s.CouponCodes != nil {
		structMap["coupon_codes"] = s.CouponCodes
	}
	if s.Components != nil {
		structMap["components"] = s.Components
	}
	if s.CustomPrice != nil {
		structMap["custom_price"] = s.CustomPrice.toMap()
	}
	if s.CalendarBilling != nil {
		structMap["calendar_billing"] = s.CalendarBilling.toMap()
	}
	if s.Metafields != nil {
		structMap["metafields"] = s.Metafields
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupSignupItem.
// It customizes the JSON unmarshaling process for SubscriptionGroupSignupItem objects.
func (s *SubscriptionGroupSignupItem) UnmarshalJSON(input []byte) error {
	var temp subscriptionGroupSignupItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	s.ProductHandle = temp.ProductHandle
	s.ProductId = temp.ProductId
	s.ProductPricePointId = temp.ProductPricePointId
	s.ProductPricePointHandle = temp.ProductPricePointHandle
	s.OfferId = temp.OfferId
	s.Reference = temp.Reference
	s.Primary = temp.Primary
	s.Currency = temp.Currency
	s.CouponCodes = temp.CouponCodes
	s.Components = temp.Components
	s.CustomPrice = temp.CustomPrice
	s.CalendarBilling = temp.CalendarBilling
	s.Metafields = temp.Metafields
	return nil
}

// TODO
type subscriptionGroupSignupItem struct {
	ProductHandle           *string                            `json:"product_handle,omitempty"`
	ProductId               *int                               `json:"product_id,omitempty"`
	ProductPricePointId     *int                               `json:"product_price_point_id,omitempty"`
	ProductPricePointHandle *string                            `json:"product_price_point_handle,omitempty"`
	OfferId                 *int                               `json:"offer_id,omitempty"`
	Reference               *string                            `json:"reference,omitempty"`
	Primary                 *bool                              `json:"primary,omitempty"`
	Currency                *string                            `json:"currency,omitempty"`
	CouponCodes             []string                           `json:"coupon_codes,omitempty"`
	Components              []SubscriptionGroupSignupComponent `json:"components,omitempty"`
	CustomPrice             *SubscriptionCustomPrice           `json:"custom_price,omitempty"`
	CalendarBilling         *CalendarBilling                   `json:"calendar_billing,omitempty"`
	Metafields              map[string]string                  `json:"metafields,omitempty"`
}
