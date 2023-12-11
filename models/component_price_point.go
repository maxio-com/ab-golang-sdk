package models

import (
	"encoding/json"
)

// ComponentPricePoint represents a ComponentPricePoint struct.
type ComponentPricePoint struct {
	Id *int `json:"id,omitempty"`
	// Price point type. We expose the following types:
	// 1. **default**: a price point that is marked as a default price for a certain product.
	// 2. **custom**: a custom price point.
	// 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
	Type *PricePointTypeEnum `json:"type,omitempty"`
	// Note: Refer to type attribute instead
	Default *bool   `json:"default,omitempty"` // Deprecated
	Name    *string `json:"name,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme *PricingSchemeEnum         `json:"pricing_scheme,omitempty"`
	ComponentId   *int                       `json:"component_id,omitempty"`
	Handle        *string                    `json:"handle,omitempty"`
	ArchivedAt    Optional[string]           `json:"archived_at"`
	CreatedAt     *string                    `json:"created_at,omitempty"`
	UpdatedAt     *string                    `json:"updated_at,omitempty"`
	Prices        []ComponentPricePointPrice `json:"prices,omitempty"`
	// Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
	UseSiteExchangeRate *bool `json:"use_site_exchange_rate,omitempty"`
	// (only used for Custom Pricing - ie. when the price point's type is `custom`) The id of the subscription that the custom price point is for.
	SubscriptionId *int  `json:"subscription_id,omitempty"`
	TaxIncluded    *bool `json:"tax_included,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePoint.
// It customizes the JSON marshaling process for ComponentPricePoint objects.
func (c *ComponentPricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePoint object to a map representation for JSON marshaling.
func (c *ComponentPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.Type != nil {
		structMap["type"] = c.Type
	}
	if c.Default != nil {
		structMap["default"] = c.Default
	}
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.PricingScheme != nil {
		structMap["pricing_scheme"] = c.PricingScheme
	}
	if c.ComponentId != nil {
		structMap["component_id"] = c.ComponentId
	}
	if c.Handle != nil {
		structMap["handle"] = c.Handle
	}
	if c.ArchivedAt.IsValueSet() {
		structMap["archived_at"] = c.ArchivedAt.Value()
	}
	if c.CreatedAt != nil {
		structMap["created_at"] = c.CreatedAt
	}
	if c.UpdatedAt != nil {
		structMap["updated_at"] = c.UpdatedAt
	}
	if c.Prices != nil {
		structMap["prices"] = c.Prices
	}
	if c.UseSiteExchangeRate != nil {
		structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
	}
	if c.SubscriptionId != nil {
		structMap["subscription_id"] = c.SubscriptionId
	}
	if c.TaxIncluded != nil {
		structMap["tax_included"] = c.TaxIncluded
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePoint.
// It customizes the JSON unmarshaling process for ComponentPricePoint objects.
func (c *ComponentPricePoint) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                  *int                       `json:"id,omitempty"`
		Type                *PricePointTypeEnum        `json:"type,omitempty"`
		Default             *bool                      `json:"default,omitempty"`
		Name                *string                    `json:"name,omitempty"`
		PricingScheme       *PricingSchemeEnum         `json:"pricing_scheme,omitempty"`
		ComponentId         *int                       `json:"component_id,omitempty"`
		Handle              *string                    `json:"handle,omitempty"`
		ArchivedAt          Optional[string]           `json:"archived_at"`
		CreatedAt           *string                    `json:"created_at,omitempty"`
		UpdatedAt           *string                    `json:"updated_at,omitempty"`
		Prices              []ComponentPricePointPrice `json:"prices,omitempty"`
		UseSiteExchangeRate *bool                      `json:"use_site_exchange_rate,omitempty"`
		SubscriptionId      *int                       `json:"subscription_id,omitempty"`
		TaxIncluded         *bool                      `json:"tax_included,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.Type = temp.Type
	c.Default = temp.Default
	c.Name = temp.Name
	c.PricingScheme = temp.PricingScheme
	c.ComponentId = temp.ComponentId
	c.Handle = temp.Handle
	c.ArchivedAt = temp.ArchivedAt
	c.CreatedAt = temp.CreatedAt
	c.UpdatedAt = temp.UpdatedAt
	c.Prices = temp.Prices
	c.UseSiteExchangeRate = temp.UseSiteExchangeRate
	c.SubscriptionId = temp.SubscriptionId
	c.TaxIncluded = temp.TaxIncluded
	return nil
}
