package models

import (
    "encoding/json"
)

// UpdateComponentPricePoint represents a UpdateComponentPricePoint struct.
type UpdateComponentPricePoint struct {
    Name                 *string        `json:"name,omitempty"`
    Handle               *string        `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme `json:"pricing_scheme,omitempty"`
    // Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
    UseSiteExchangeRate  *bool          `json:"use_site_exchange_rate,omitempty"`
    // Whether or not the price point includes tax
    TaxIncluded          *bool          `json:"tax_included,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval             *int           `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit         *IntervalUnit  `json:"interval_unit,omitempty"`
    Prices               []UpdatePrice  `json:"prices,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePoint.
// It customizes the JSON marshaling process for UpdateComponentPricePoint objects.
func (u UpdateComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePoint object to a map representation for JSON marshaling.
func (u UpdateComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Handle != nil {
        structMap["handle"] = u.Handle
    }
    if u.PricingScheme != nil {
        structMap["pricing_scheme"] = u.PricingScheme
    }
    if u.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = u.UseSiteExchangeRate
    }
    if u.TaxIncluded != nil {
        structMap["tax_included"] = u.TaxIncluded
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.IntervalUnit != nil {
        structMap["interval_unit"] = u.IntervalUnit
    }
    if u.Prices != nil {
        structMap["prices"] = u.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePoint.
// It customizes the JSON unmarshaling process for UpdateComponentPricePoint objects.
func (u *UpdateComponentPricePoint) UnmarshalJSON(input []byte) error {
    var temp updateComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "handle", "pricing_scheme", "use_site_exchange_rate", "tax_included", "interval", "interval_unit", "prices")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Name = temp.Name
    u.Handle = temp.Handle
    u.PricingScheme = temp.PricingScheme
    u.UseSiteExchangeRate = temp.UseSiteExchangeRate
    u.TaxIncluded = temp.TaxIncluded
    u.Interval = temp.Interval
    u.IntervalUnit = temp.IntervalUnit
    u.Prices = temp.Prices
    return nil
}

// TODO
type updateComponentPricePoint  struct {
    Name                *string        `json:"name,omitempty"`
    Handle              *string        `json:"handle,omitempty"`
    PricingScheme       *PricingScheme `json:"pricing_scheme,omitempty"`
    UseSiteExchangeRate *bool          `json:"use_site_exchange_rate,omitempty"`
    TaxIncluded         *bool          `json:"tax_included,omitempty"`
    Interval            *int           `json:"interval,omitempty"`
    IntervalUnit        *IntervalUnit  `json:"interval_unit,omitempty"`
    Prices              []UpdatePrice  `json:"prices,omitempty"`
}
