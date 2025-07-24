// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ComponentPricePoint represents a ComponentPricePoint struct.
type ComponentPricePoint struct {
    Id                       *int                             `json:"id,omitempty"`
    // Price point type. We expose the following types:
    // 1. **default**: a price point that is marked as a default price for a certain product.
    // 2. **custom**: a custom price point.
    // 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
    Type                     *PricePointType                  `json:"type,omitempty"`
    // Note: Refer to type attribute instead
    Default                  *bool                            `json:"default,omitempty"`                    // Deprecated
    Name                     *string                          `json:"name,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme            *PricingScheme                   `json:"pricing_scheme,omitempty"`
    ComponentId              *int                             `json:"component_id,omitempty"`
    Handle                   Optional[string]                 `json:"handle"`
    ArchivedAt               Optional[time.Time]              `json:"archived_at"`
    CreatedAt                *time.Time                       `json:"created_at,omitempty"`
    UpdatedAt                *time.Time                       `json:"updated_at,omitempty"`
    Prices                   []ComponentPrice                 `json:"prices,omitempty"`
    // Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site. Defaults to true during creation.
    UseSiteExchangeRate      *bool                            `json:"use_site_exchange_rate,omitempty"`
    // (only used for Custom Pricing - ie. when the price point's type is `custom`) The id of the subscription that the custom price point is for.
    SubscriptionId           *int                             `json:"subscription_id,omitempty"`
    TaxIncluded              *bool                            `json:"tax_included,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                 Optional[int]                    `json:"interval"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit             Optional[IntervalUnit]           `json:"interval_unit"`
    // An array of currency pricing data is available when multiple currencies are defined for the site. It varies based on the use_site_exchange_rate setting for the price point. This parameter is present only in the response of read endpoints, after including the appropriate query parameter.
    CurrencyPrices           []ComponentCurrencyPrice         `json:"currency_prices,omitempty"`
    // Applicable only to prepaid usage components. An array of overage price brackets.
    OveragePrices            []ComponentPrice                 `json:"overage_prices,omitempty"`
    // Applicable only to prepaid usage components. Pricing scheme for overage pricing.
    OveragePricingScheme     *PricingScheme                   `json:"overage_pricing_scheme,omitempty"`
    // Applicable only to prepaid usage components. Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period.
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    // Applicable only to prepaid usage components. Boolean which controls whether or not remaining units should be rolled over to the next period.
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    // Applicable only to prepaid usage components where rollover_prepaid_remainder is true. The number of `expiration_interval_unit`s after which rollover amounts should expire.
    ExpirationInterval       Optional[int]                    `json:"expiration_interval"`
    // Applicable only to prepaid usage components where rollover_prepaid_remainder is true. A string representing the expiration interval unit for this component, either month or day.
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
    AdditionalProperties     map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentPricePoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPricePoint) String() string {
    return fmt.Sprintf(
    	"ComponentPricePoint[Id=%v, Type=%v, Default=%v, Name=%v, PricingScheme=%v, ComponentId=%v, Handle=%v, ArchivedAt=%v, CreatedAt=%v, UpdatedAt=%v, Prices=%v, UseSiteExchangeRate=%v, SubscriptionId=%v, TaxIncluded=%v, Interval=%v, IntervalUnit=%v, CurrencyPrices=%v, OveragePrices=%v, OveragePricingScheme=%v, RenewPrepaidAllocation=%v, RolloverPrepaidRemainder=%v, ExpirationInterval=%v, ExpirationIntervalUnit=%v, AdditionalProperties=%v]",
    	c.Id, c.Type, c.Default, c.Name, c.PricingScheme, c.ComponentId, c.Handle, c.ArchivedAt, c.CreatedAt, c.UpdatedAt, c.Prices, c.UseSiteExchangeRate, c.SubscriptionId, c.TaxIncluded, c.Interval, c.IntervalUnit, c.CurrencyPrices, c.OveragePrices, c.OveragePricingScheme, c.RenewPrepaidAllocation, c.RolloverPrepaidRemainder, c.ExpirationInterval, c.ExpirationIntervalUnit, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePoint.
// It customizes the JSON marshaling process for ComponentPricePoint objects.
func (c ComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "type", "default", "name", "pricing_scheme", "component_id", "handle", "archived_at", "created_at", "updated_at", "prices", "use_site_exchange_rate", "subscription_id", "tax_included", "interval", "interval_unit", "currency_prices", "overage_prices", "overage_pricing_scheme", "renew_prepaid_allocation", "rollover_prepaid_remainder", "expiration_interval", "expiration_interval_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePoint object to a map representation for JSON marshaling.
func (c ComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    if c.Handle.IsValueSet() {
        if c.Handle.Value() != nil {
            structMap["handle"] = c.Handle.Value()
        } else {
            structMap["handle"] = nil
        }
    }
    if c.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if c.ArchivedAt.Value() != nil {
            val := c.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        if c.ArchivedAt.Value() != nil {
            structMap["archived_at"] = ArchivedAtVal
        } else {
            structMap["archived_at"] = nil
        }
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
        if c.Interval.Value() != nil {
            structMap["interval"] = c.Interval.Value()
        } else {
            structMap["interval"] = nil
        }
    }
    if c.IntervalUnit.IsValueSet() {
        if c.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = c.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    if c.CurrencyPrices != nil {
        structMap["currency_prices"] = c.CurrencyPrices
    }
    if c.OveragePrices != nil {
        structMap["overage_prices"] = c.OveragePrices
    }
    if c.OveragePricingScheme != nil {
        structMap["overage_pricing_scheme"] = c.OveragePricingScheme
    }
    if c.RenewPrepaidAllocation != nil {
        structMap["renew_prepaid_allocation"] = c.RenewPrepaidAllocation
    }
    if c.RolloverPrepaidRemainder != nil {
        structMap["rollover_prepaid_remainder"] = c.RolloverPrepaidRemainder
    }
    if c.ExpirationInterval.IsValueSet() {
        if c.ExpirationInterval.Value() != nil {
            structMap["expiration_interval"] = c.ExpirationInterval.Value()
        } else {
            structMap["expiration_interval"] = nil
        }
    }
    if c.ExpirationIntervalUnit.IsValueSet() {
        if c.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = c.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePoint.
// It customizes the JSON unmarshaling process for ComponentPricePoint objects.
func (c *ComponentPricePoint) UnmarshalJSON(input []byte) error {
    var temp tempComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "type", "default", "name", "pricing_scheme", "component_id", "handle", "archived_at", "created_at", "updated_at", "prices", "use_site_exchange_rate", "subscription_id", "tax_included", "interval", "interval_unit", "currency_prices", "overage_prices", "overage_pricing_scheme", "renew_prepaid_allocation", "rollover_prepaid_remainder", "expiration_interval", "expiration_interval_unit")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
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
    c.OveragePrices = temp.OveragePrices
    c.OveragePricingScheme = temp.OveragePricingScheme
    c.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
    c.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
    c.ExpirationInterval = temp.ExpirationInterval
    c.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    return nil
}

// tempComponentPricePoint is a temporary struct used for validating the fields of ComponentPricePoint.
type tempComponentPricePoint  struct {
    Id                       *int                             `json:"id,omitempty"`
    Type                     *PricePointType                  `json:"type,omitempty"`
    Default                  *bool                            `json:"default,omitempty"`
    Name                     *string                          `json:"name,omitempty"`
    PricingScheme            *PricingScheme                   `json:"pricing_scheme,omitempty"`
    ComponentId              *int                             `json:"component_id,omitempty"`
    Handle                   Optional[string]                 `json:"handle"`
    ArchivedAt               Optional[string]                 `json:"archived_at"`
    CreatedAt                *string                          `json:"created_at,omitempty"`
    UpdatedAt                *string                          `json:"updated_at,omitempty"`
    Prices                   []ComponentPrice                 `json:"prices,omitempty"`
    UseSiteExchangeRate      *bool                            `json:"use_site_exchange_rate,omitempty"`
    SubscriptionId           *int                             `json:"subscription_id,omitempty"`
    TaxIncluded              *bool                            `json:"tax_included,omitempty"`
    Interval                 Optional[int]                    `json:"interval"`
    IntervalUnit             Optional[IntervalUnit]           `json:"interval_unit"`
    CurrencyPrices           []ComponentCurrencyPrice         `json:"currency_prices,omitempty"`
    OveragePrices            []ComponentPrice                 `json:"overage_prices,omitempty"`
    OveragePricingScheme     *PricingScheme                   `json:"overage_pricing_scheme,omitempty"`
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    ExpirationInterval       Optional[int]                    `json:"expiration_interval"`
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
}
