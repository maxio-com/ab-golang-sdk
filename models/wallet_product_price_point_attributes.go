// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WalletProductPricePointAttributes represents a WalletProductPricePointAttributes struct.
// A price point to create for the Wallet Product, in the same call that creates or updates it.
type WalletProductPricePointAttributes struct {
    // The price point's name. Required. Omitting it fails validation with a message that reads `Name: cannot be blank.`, not naming the price point.
    Name                 string                 `json:"name"`
    Handle               *string                `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    // Required when `pricing_scheme` is `per_unit`.
    UnitPrice            Optional[string]       `json:"unit_price"`
    TaxIncluded          *bool                  `json:"tax_included,omitempty"`
    // Requires Multifrequency to be enabled on the site.
    Interval             *int                   `json:"interval,omitempty"`
    IntervalUnit         *IntervalUnit          `json:"interval_unit,omitempty"`
    UseSiteExchangeRate  *bool                  `json:"use_site_exchange_rate,omitempty"`
    // Required for every `pricing_scheme` other than `per_unit`.
    Prices               []Price1               `json:"prices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductPricePointAttributes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductPricePointAttributes) String() string {
    return fmt.Sprintf(
    	"WalletProductPricePointAttributes[Name=%v, Handle=%v, PricingScheme=%v, UnitPrice=%v, TaxIncluded=%v, Interval=%v, IntervalUnit=%v, UseSiteExchangeRate=%v, Prices=%v, AdditionalProperties=%v]",
    	w.Name, w.Handle, w.PricingScheme, w.UnitPrice, w.TaxIncluded, w.Interval, w.IntervalUnit, w.UseSiteExchangeRate, w.Prices, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductPricePointAttributes.
// It customizes the JSON marshaling process for WalletProductPricePointAttributes objects.
func (w WalletProductPricePointAttributes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "name", "handle", "pricing_scheme", "unit_price", "tax_included", "interval", "interval_unit", "use_site_exchange_rate", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductPricePointAttributes object to a map representation for JSON marshaling.
func (w WalletProductPricePointAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["name"] = w.Name
    if w.Handle != nil {
        structMap["handle"] = w.Handle
    }
    if w.PricingScheme != nil {
        structMap["pricing_scheme"] = w.PricingScheme
    }
    if w.UnitPrice.IsValueSet() {
        if w.UnitPrice.Value() != nil {
            structMap["unit_price"] = w.UnitPrice.Value()
        } else {
            structMap["unit_price"] = nil
        }
    }
    if w.TaxIncluded != nil {
        structMap["tax_included"] = w.TaxIncluded
    }
    if w.Interval != nil {
        structMap["interval"] = w.Interval
    }
    if w.IntervalUnit != nil {
        structMap["interval_unit"] = w.IntervalUnit
    }
    if w.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = w.UseSiteExchangeRate
    }
    if w.Prices != nil {
        structMap["prices"] = w.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductPricePointAttributes.
// It customizes the JSON unmarshaling process for WalletProductPricePointAttributes objects.
func (w *WalletProductPricePointAttributes) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductPricePointAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle", "pricing_scheme", "unit_price", "tax_included", "interval", "interval_unit", "use_site_exchange_rate", "prices")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Name = *temp.Name
    w.Handle = temp.Handle
    w.PricingScheme = temp.PricingScheme
    w.UnitPrice = temp.UnitPrice
    w.TaxIncluded = temp.TaxIncluded
    w.Interval = temp.Interval
    w.IntervalUnit = temp.IntervalUnit
    w.UseSiteExchangeRate = temp.UseSiteExchangeRate
    w.Prices = temp.Prices
    return nil
}

// tempWalletProductPricePointAttributes is a temporary struct used for validating the fields of WalletProductPricePointAttributes.
type tempWalletProductPricePointAttributes  struct {
    Name                *string          `json:"name"`
    Handle              *string          `json:"handle,omitempty"`
    PricingScheme       *PricingScheme   `json:"pricing_scheme,omitempty"`
    UnitPrice           Optional[string] `json:"unit_price"`
    TaxIncluded         *bool            `json:"tax_included,omitempty"`
    Interval            *int             `json:"interval,omitempty"`
    IntervalUnit        *IntervalUnit    `json:"interval_unit,omitempty"`
    UseSiteExchangeRate *bool            `json:"use_site_exchange_rate,omitempty"`
    Prices              []Price1         `json:"prices,omitempty"`
}

func (w *tempWalletProductPricePointAttributes) validate() error {
    var errs []string
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Wallet Product Price Point Attributes`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
