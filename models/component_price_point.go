package models

import (
    "encoding/json"
    "log"
    "time"
)

// ComponentPricePoint represents a ComponentPricePoint struct.
type ComponentPricePoint struct {
    Id                  *int                       `json:"id,omitempty"`
    // Price point type. We expose the following types:
    // 1. **default**: a price point that is marked as a default price for a certain product.
    // 2. **custom**: a custom price point.
    // 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
    Type                *PricePointType            `json:"type,omitempty"`
    // Note: Refer to type attribute instead
    Default             *bool                      `json:"default,omitempty"`                // Deprecated
    Name                *string                    `json:"name,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme       *PricingScheme             `json:"pricing_scheme,omitempty"`
    ComponentId         *int                       `json:"component_id,omitempty"`
    Handle              *string                    `json:"handle,omitempty"`
    ArchivedAt          Optional[time.Time]        `json:"archived_at"`
    CreatedAt           *time.Time                 `json:"created_at,omitempty"`
    UpdatedAt           *time.Time                 `json:"updated_at,omitempty"`
    Prices              []ComponentPricePointPrice `json:"prices,omitempty"`
    // Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
    UseSiteExchangeRate *bool                      `json:"use_site_exchange_rate,omitempty"`
    // (only used for Custom Pricing - ie. when the price point's type is `custom`) The id of the subscription that the custom price point is for.
    SubscriptionId      *int                       `json:"subscription_id,omitempty"`
    TaxIncluded         *bool                      `json:"tax_included,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval            Optional[int]              `json:"interval"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit        Optional[IntervalUnit]     `json:"interval_unit"`
    // An array of currency pricing data is available when multiple currencies are defined for the site. It varies based on the use_site_exchange_rate setting for the price point. This parameter is present only in the response of read endpoints, after including the appropriate query parameter.
    CurrencyPrices      []ComponentCurrencyPrice   `json:"currency_prices,omitempty"`
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
        var ArchivedAtVal *string = nil
        if c.ArchivedAt.Value() != nil {
            val := c.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        structMap["archived_at"] = ArchivedAtVal
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.UpdatedAt != nil {
        structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
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
    if c.Interval.IsValueSet() {
        structMap["interval"] = c.Interval.Value()
    }
    if c.IntervalUnit.IsValueSet() {
        structMap["interval_unit"] = c.IntervalUnit.Value()
    }
    if c.CurrencyPrices != nil {
        structMap["currency_prices"] = c.CurrencyPrices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePoint.
// It customizes the JSON unmarshaling process for ComponentPricePoint objects.
func (c *ComponentPricePoint) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Id                  *int                       `json:"id,omitempty"`
        Type                *PricePointType            `json:"type,omitempty"`
        Default             *bool                      `json:"default,omitempty"`
        Name                *string                    `json:"name,omitempty"`
        PricingScheme       *PricingScheme             `json:"pricing_scheme,omitempty"`
        ComponentId         *int                       `json:"component_id,omitempty"`
        Handle              *string                    `json:"handle,omitempty"`
        ArchivedAt          Optional[string]           `json:"archived_at"`
        CreatedAt           *string                    `json:"created_at,omitempty"`
        UpdatedAt           *string                    `json:"updated_at,omitempty"`
        Prices              []ComponentPricePointPrice `json:"prices,omitempty"`
        UseSiteExchangeRate *bool                      `json:"use_site_exchange_rate,omitempty"`
        SubscriptionId      *int                       `json:"subscription_id,omitempty"`
        TaxIncluded         *bool                      `json:"tax_included,omitempty"`
        Interval            Optional[int]              `json:"interval"`
        IntervalUnit        Optional[IntervalUnit]     `json:"interval_unit"`
        CurrencyPrices      []ComponentCurrencyPrice   `json:"currency_prices,omitempty"`
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
    c.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        c.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        c.UpdatedAt = &UpdatedAtVal
    }
    c.Prices = temp.Prices
    c.UseSiteExchangeRate = temp.UseSiteExchangeRate
    c.SubscriptionId = temp.SubscriptionId
    c.TaxIncluded = temp.TaxIncluded
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.CurrencyPrices = temp.CurrencyPrices
    return nil
}
