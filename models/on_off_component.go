/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OnOffComponent represents a OnOffComponent struct.
type OnOffComponent struct {
    // A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes".
    Name                      string                    `json:"name"`
    // A description for the component that will be displayed to the user on the hosted signup page.
    Description               *string                   `json:"description,omitempty"`
    // A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.
    Handle                    *string                   `json:"handle,omitempty"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                     `json:"taxable,omitempty"`
    // (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261191737101-Price-Points-Components) for an overview of how price brackets work for different pricing schemes.
    Prices                    []Price                   `json:"prices,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge             Optional[CreditType]      `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit           Optional[CreditType]      `json:"downgrade_credit"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    // The amount the customer will be charged per unit when the pricing scheme is “per_unit”. For On/Off Components, this is the amount that the customer will be charged when they turn the component on for the subscription. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice                 *OnOffComponentUnitPrice  `json:"unit_price,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    // deprecated May 2011 - use unit_price instead
    PriceInCents              *string                   `json:"price_in_cents,omitempty"`
    DisplayOnHostedPage       *bool                     `json:"display_on_hosted_page,omitempty"`
    AllowFractionalQuantities *bool                     `json:"allow_fractional_quantities,omitempty"`
    PublicSignupPageIds       []int                     `json:"public_signup_page_ids,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                  *int                      `json:"interval,omitempty"`
    // A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit              Optional[IntervalUnit]    `json:"interval_unit"`
    AdditionalProperties      map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OnOffComponent.
// It customizes the JSON marshaling process for OnOffComponent objects.
func (o OnOffComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OnOffComponent object to a map representation for JSON marshaling.
func (o OnOffComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["name"] = o.Name
    if o.Description != nil {
        structMap["description"] = o.Description
    }
    if o.Handle != nil {
        structMap["handle"] = o.Handle
    }
    if o.Taxable != nil {
        structMap["taxable"] = o.Taxable
    }
    if o.Prices != nil {
        structMap["prices"] = o.Prices
    }
    if o.UpgradeCharge.IsValueSet() {
        if o.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = o.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if o.DowngradeCredit.IsValueSet() {
        if o.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = o.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    if o.PricePoints != nil {
        structMap["price_points"] = o.PricePoints
    }
    if o.UnitPrice != nil {
        structMap["unit_price"] = o.UnitPrice.toMap()
    }
    if o.TaxCode != nil {
        structMap["tax_code"] = o.TaxCode
    }
    if o.HideDateRangeOnInvoice != nil {
        structMap["hide_date_range_on_invoice"] = o.HideDateRangeOnInvoice
    }
    if o.PriceInCents != nil {
        structMap["price_in_cents"] = o.PriceInCents
    }
    if o.DisplayOnHostedPage != nil {
        structMap["display_on_hosted_page"] = o.DisplayOnHostedPage
    }
    if o.AllowFractionalQuantities != nil {
        structMap["allow_fractional_quantities"] = o.AllowFractionalQuantities
    }
    if o.PublicSignupPageIds != nil {
        structMap["public_signup_page_ids"] = o.PublicSignupPageIds
    }
    if o.Interval != nil {
        structMap["interval"] = o.Interval
    }
    if o.IntervalUnit.IsValueSet() {
        if o.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = o.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OnOffComponent.
// It customizes the JSON unmarshaling process for OnOffComponent objects.
func (o *OnOffComponent) UnmarshalJSON(input []byte) error {
    var temp tempOnOffComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "description", "handle", "taxable", "prices", "upgrade_charge", "downgrade_credit", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "price_in_cents", "display_on_hosted_page", "allow_fractional_quantities", "public_signup_page_ids", "interval", "interval_unit")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Name = *temp.Name
    o.Description = temp.Description
    o.Handle = temp.Handle
    o.Taxable = temp.Taxable
    o.Prices = temp.Prices
    o.UpgradeCharge = temp.UpgradeCharge
    o.DowngradeCredit = temp.DowngradeCredit
    o.PricePoints = temp.PricePoints
    o.UnitPrice = temp.UnitPrice
    o.TaxCode = temp.TaxCode
    o.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    o.PriceInCents = temp.PriceInCents
    o.DisplayOnHostedPage = temp.DisplayOnHostedPage
    o.AllowFractionalQuantities = temp.AllowFractionalQuantities
    o.PublicSignupPageIds = temp.PublicSignupPageIds
    o.Interval = temp.Interval
    o.IntervalUnit = temp.IntervalUnit
    return nil
}

// tempOnOffComponent is a temporary struct used for validating the fields of OnOffComponent.
type tempOnOffComponent  struct {
    Name                      *string                   `json:"name"`
    Description               *string                   `json:"description,omitempty"`
    Handle                    *string                   `json:"handle,omitempty"`
    Taxable                   *bool                     `json:"taxable,omitempty"`
    Prices                    []Price                   `json:"prices,omitempty"`
    UpgradeCharge             Optional[CreditType]      `json:"upgrade_charge"`
    DowngradeCredit           Optional[CreditType]      `json:"downgrade_credit"`
    PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
    UnitPrice                 *OnOffComponentUnitPrice  `json:"unit_price,omitempty"`
    TaxCode                   *string                   `json:"tax_code,omitempty"`
    HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
    PriceInCents              *string                   `json:"price_in_cents,omitempty"`
    DisplayOnHostedPage       *bool                     `json:"display_on_hosted_page,omitempty"`
    AllowFractionalQuantities *bool                     `json:"allow_fractional_quantities,omitempty"`
    PublicSignupPageIds       []int                     `json:"public_signup_page_ids,omitempty"`
    Interval                  *int                      `json:"interval,omitempty"`
    IntervalUnit              Optional[IntervalUnit]    `json:"interval_unit"`
}

func (o *tempOnOffComponent) validate() error {
    var errs []string
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `On/Off Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
