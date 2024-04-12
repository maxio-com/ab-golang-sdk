package models

import (
    "encoding/json"
    "errors"
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
    // (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://help.chargify.com/products/product-components.html#general-price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
    Prices                    []Price                   `json:"prices,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge             Optional[CreditType]      `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit           Optional[CreditType]      `json:"downgrade_credit"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    // The amount the customer will be charged per unit when the pricing scheme is “per_unit”. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice                 *EBBComponentUnitPrice    `json:"unit_price,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    // deprecated May 2011 - use unit_price instead
    PriceInCents              *string                   `json:"price_in_cents,omitempty"`
    // The ID of an event based billing metric that will be attached to this component.
    EventBasedBillingMetricId int                       `json:"event_based_billing_metric_id"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                  *int                      `json:"interval,omitempty"`
    // A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit              *IntervalUnit             `json:"interval_unit,omitempty"`
    AdditionalProperties      map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EBBComponent.
// It customizes the JSON marshaling process for EBBComponent objects.
func (e EBBComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EBBComponent object to a map representation for JSON marshaling.
func (e EBBComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
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
    if e.UpgradeCharge.IsValueSet() {
        if e.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = e.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if e.DowngradeCredit.IsValueSet() {
        if e.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = e.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
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
    if e.PriceInCents != nil {
        structMap["price_in_cents"] = e.PriceInCents
    }
    structMap["event_based_billing_metric_id"] = e.EventBasedBillingMetricId
    if e.Interval != nil {
        structMap["interval"] = e.Interval
    }
    if e.IntervalUnit != nil {
        structMap["interval_unit"] = e.IntervalUnit
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBComponent.
// It customizes the JSON unmarshaling process for EBBComponent objects.
func (e *EBBComponent) UnmarshalJSON(input []byte) error {
    var temp ebbComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "upgrade_charge", "downgrade_credit", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "price_in_cents", "event_based_billing_metric_id", "interval", "interval_unit")
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
    e.UpgradeCharge = temp.UpgradeCharge
    e.DowngradeCredit = temp.DowngradeCredit
    e.PricePoints = temp.PricePoints
    e.UnitPrice = temp.UnitPrice
    e.TaxCode = temp.TaxCode
    e.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    e.PriceInCents = temp.PriceInCents
    e.EventBasedBillingMetricId = *temp.EventBasedBillingMetricId
    e.Interval = temp.Interval
    e.IntervalUnit = temp.IntervalUnit
    return nil
}

// TODO
type ebbComponent  struct {
    Name                      *string                   `json:"name"`
    UnitName                  *string                   `json:"unit_name"`
    Description               *string                   `json:"description,omitempty"`
    Handle                    *string                   `json:"handle,omitempty"`
    Taxable                   *bool                     `json:"taxable,omitempty"`
    PricingScheme             *PricingScheme            `json:"pricing_scheme"`
    Prices                    []Price                   `json:"prices,omitempty"`
    UpgradeCharge             Optional[CreditType]      `json:"upgrade_charge"`
    DowngradeCredit           Optional[CreditType]      `json:"downgrade_credit"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    UnitPrice                 *EBBComponentUnitPrice    `json:"unit_price,omitempty"`
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    PriceInCents              *string                   `json:"price_in_cents,omitempty"`
    EventBasedBillingMetricId *int                      `json:"event_based_billing_metric_id"`
    Interval                  *int                      `json:"interval,omitempty"`
    IntervalUnit              *IntervalUnit             `json:"interval_unit,omitempty"`
}

func (e *ebbComponent) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
