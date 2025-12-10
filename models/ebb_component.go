// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// EBBComponent represents a EBBComponent struct.
type EBBComponent struct {
    // A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes".
    Name                      string                    `json:"name"`
    // The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. i.e. “message”, which may then be shown as “5 messages” on a subscription’s component line-item
    UnitName                  string                    `json:"unit_name"`
    // A description for the component that will be displayed to the user on the hosted signup page.
    Description               *string                   `json:"description,omitempty"`
    // A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.
    Handle                    *string                   `json:"handle,omitempty"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                     `json:"taxable,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme             PricingScheme             `json:"pricing_scheme"`
    // (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261149166733-Component-Pricing-Schemes#price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
    Prices                    []Price                   `json:"prices,omitempty"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    // The amount the customer will be charged per unit when the pricing scheme is “per_unit”. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice                 *EBBComponentUnitPrice    `json:"unit_price,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters.
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    // The ID of an event based billing metric that will be attached to this component.
    EventBasedBillingMetricId int                       `json:"event_based_billing_metric_id"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                  *int                      `json:"interval,omitempty"`
    // A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit              Optional[IntervalUnit]    `json:"interval_unit"`
    AdditionalProperties      map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for EBBComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EBBComponent) String() string {
    return fmt.Sprintf(
    	"EBBComponent[Name=%v, UnitName=%v, Description=%v, Handle=%v, Taxable=%v, PricingScheme=%v, Prices=%v, PricePoints=%v, UnitPrice=%v, TaxCode=%v, HideDateRangeOnInvoice=%v, EventBasedBillingMetricId=%v, Interval=%v, IntervalUnit=%v, AdditionalProperties=%v]",
    	e.Name, e.UnitName, e.Description, e.Handle, e.Taxable, e.PricingScheme, e.Prices, e.PricePoints, e.UnitPrice, e.TaxCode, e.HideDateRangeOnInvoice, e.EventBasedBillingMetricId, e.Interval, e.IntervalUnit, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EBBComponent.
// It customizes the JSON marshaling process for EBBComponent objects.
func (e EBBComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "event_based_billing_metric_id", "interval", "interval_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EBBComponent object to a map representation for JSON marshaling.
func (e EBBComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["name"] = e.Name
    structMap["unit_name"] = e.UnitName
    if e.Description != nil {
        structMap["description"] = e.Description
    }
    if e.Handle != nil {
        structMap["handle"] = e.Handle
    }
    if e.Taxable != nil {
        structMap["taxable"] = e.Taxable
    }
    structMap["pricing_scheme"] = e.PricingScheme
    if e.Prices != nil {
        structMap["prices"] = e.Prices
    }
    if e.PricePoints != nil {
        structMap["price_points"] = e.PricePoints
    }
    if e.UnitPrice != nil {
        structMap["unit_price"] = e.UnitPrice.toMap()
    }
    if e.TaxCode != nil {
        structMap["tax_code"] = e.TaxCode
    }
    if e.HideDateRangeOnInvoice != nil {
        structMap["hide_date_range_on_invoice"] = e.HideDateRangeOnInvoice
    }
    structMap["event_based_billing_metric_id"] = e.EventBasedBillingMetricId
    if e.Interval != nil {
        structMap["interval"] = e.Interval
    }
    if e.IntervalUnit.IsValueSet() {
        if e.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = e.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBComponent.
// It customizes the JSON unmarshaling process for EBBComponent objects.
func (e *EBBComponent) UnmarshalJSON(input []byte) error {
    var temp tempEBBComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "event_based_billing_metric_id", "interval", "interval_unit")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Name = *temp.Name
    e.UnitName = *temp.UnitName
    e.Description = temp.Description
    e.Handle = temp.Handle
    e.Taxable = temp.Taxable
    e.PricingScheme = *temp.PricingScheme
    e.Prices = temp.Prices
    e.PricePoints = temp.PricePoints
    e.UnitPrice = temp.UnitPrice
    e.TaxCode = temp.TaxCode
    e.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    e.EventBasedBillingMetricId = *temp.EventBasedBillingMetricId
    e.Interval = temp.Interval
    e.IntervalUnit = temp.IntervalUnit
    return nil
}

// tempEBBComponent is a temporary struct used for validating the fields of EBBComponent.
type tempEBBComponent  struct {
    Name                      *string                   `json:"name"`
    UnitName                  *string                   `json:"unit_name"`
    Description               *string                   `json:"description,omitempty"`
    Handle                    *string                   `json:"handle,omitempty"`
    Taxable                   *bool                     `json:"taxable,omitempty"`
    PricingScheme             *PricingScheme            `json:"pricing_scheme"`
    Prices                    []Price                   `json:"prices,omitempty"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    UnitPrice                 *EBBComponentUnitPrice    `json:"unit_price,omitempty"`
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    EventBasedBillingMetricId *int                      `json:"event_based_billing_metric_id"`
    Interval                  *int                      `json:"interval,omitempty"`
    IntervalUnit              Optional[IntervalUnit]    `json:"interval_unit"`
}

func (e *tempEBBComponent) validate() error {
    var errs []string
    if e.Name == nil {
        errs = append(errs, "required field `name` is missing for type `EBB Component`")
    }
    if e.UnitName == nil {
        errs = append(errs, "required field `unit_name` is missing for type `EBB Component`")
    }
    if e.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `EBB Component`")
    }
    if e.EventBasedBillingMetricId == nil {
        errs = append(errs, "required field `event_based_billing_metric_id` is missing for type `EBB Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
