// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// PrepaidUsageComponent represents a PrepaidUsageComponent struct.
type PrepaidUsageComponent struct {
    // A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes".
    Name                      string                                  `json:"name"`
    // The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. i.e. “message”, which may then be shown as “5 messages” on a subscription’s component line-item
    UnitName                  string                                  `json:"unit_name"`
    // A description for the component that will be displayed to the user on the hosted signup page.
    Description               *string                                 `json:"description,omitempty"`
    // A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.
    Handle                    *string                                 `json:"handle,omitempty"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                                   `json:"taxable,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme             PricingScheme                           `json:"pricing_scheme"`
    // (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261149166733-Component-Pricing-Schemes#price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
    Prices                    []Price                                 `json:"prices,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge             Optional[CreditType]                    `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit           Optional[CreditType]                    `json:"downgrade_credit"`
    PricePoints               []CreatePrepaidUsageComponentPricePoint `json:"price_points,omitempty"`
    // The amount the customer will be charged per unit when the pricing scheme is “per_unit”. For On/Off Components, this is the amount that the customer will be charged when they turn the component on for the subscription. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice                 *PrepaidUsageComponentUnitPrice         `json:"unit_price,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters.
    TaxCode                   *string                                 `json:"tax_code,omitempty"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                                   `json:"hide_date_range_on_invoice,omitempty"`
    OveragePricing            OveragePricing                          `json:"overage_pricing"`
    // Boolean which controls whether or not remaining units should be rolled over to the next period
    RolloverPrepaidRemainder  *bool                                   `json:"rollover_prepaid_remainder,omitempty"`
    // Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period
    RenewPrepaidAllocation    *bool                                   `json:"renew_prepaid_allocation,omitempty"`
    // (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire
    ExpirationInterval        *float64                                `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit    Optional[ExpirationIntervalUnit]        `json:"expiration_interval_unit"`
    DisplayOnHostedPage       *bool                                   `json:"display_on_hosted_page,omitempty"`
    AllowFractionalQuantities *bool                                   `json:"allow_fractional_quantities,omitempty"`
    PublicSignupPageIds       []int                                   `json:"public_signup_page_ids,omitempty"`
    AdditionalProperties      map[string]interface{}                  `json:"_"`
}

// String implements the fmt.Stringer interface for PrepaidUsageComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PrepaidUsageComponent) String() string {
    return fmt.Sprintf(
    	"PrepaidUsageComponent[Name=%v, UnitName=%v, Description=%v, Handle=%v, Taxable=%v, PricingScheme=%v, Prices=%v, UpgradeCharge=%v, DowngradeCredit=%v, PricePoints=%v, UnitPrice=%v, TaxCode=%v, HideDateRangeOnInvoice=%v, OveragePricing=%v, RolloverPrepaidRemainder=%v, RenewPrepaidAllocation=%v, ExpirationInterval=%v, ExpirationIntervalUnit=%v, DisplayOnHostedPage=%v, AllowFractionalQuantities=%v, PublicSignupPageIds=%v, AdditionalProperties=%v]",
    	p.Name, p.UnitName, p.Description, p.Handle, p.Taxable, p.PricingScheme, p.Prices, p.UpgradeCharge, p.DowngradeCredit, p.PricePoints, p.UnitPrice, p.TaxCode, p.HideDateRangeOnInvoice, p.OveragePricing, p.RolloverPrepaidRemainder, p.RenewPrepaidAllocation, p.ExpirationInterval, p.ExpirationIntervalUnit, p.DisplayOnHostedPage, p.AllowFractionalQuantities, p.PublicSignupPageIds, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageComponent.
// It customizes the JSON marshaling process for PrepaidUsageComponent objects.
func (p PrepaidUsageComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "upgrade_charge", "downgrade_credit", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "overage_pricing", "rollover_prepaid_remainder", "renew_prepaid_allocation", "expiration_interval", "expiration_interval_unit", "display_on_hosted_page", "allow_fractional_quantities", "public_signup_page_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageComponent object to a map representation for JSON marshaling.
func (p PrepaidUsageComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["name"] = p.Name
    structMap["unit_name"] = p.UnitName
    if p.Description != nil {
        structMap["description"] = p.Description
    }
    if p.Handle != nil {
        structMap["handle"] = p.Handle
    }
    if p.Taxable != nil {
        structMap["taxable"] = p.Taxable
    }
    structMap["pricing_scheme"] = p.PricingScheme
    if p.Prices != nil {
        structMap["prices"] = p.Prices
    }
    if p.UpgradeCharge.IsValueSet() {
        if p.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = p.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if p.DowngradeCredit.IsValueSet() {
        if p.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = p.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    if p.PricePoints != nil {
        structMap["price_points"] = p.PricePoints
    }
    if p.UnitPrice != nil {
        structMap["unit_price"] = p.UnitPrice.toMap()
    }
    if p.TaxCode != nil {
        structMap["tax_code"] = p.TaxCode
    }
    if p.HideDateRangeOnInvoice != nil {
        structMap["hide_date_range_on_invoice"] = p.HideDateRangeOnInvoice
    }
    structMap["overage_pricing"] = p.OveragePricing.toMap()
    if p.RolloverPrepaidRemainder != nil {
        structMap["rollover_prepaid_remainder"] = p.RolloverPrepaidRemainder
    }
    if p.RenewPrepaidAllocation != nil {
        structMap["renew_prepaid_allocation"] = p.RenewPrepaidAllocation
    }
    if p.ExpirationInterval != nil {
        structMap["expiration_interval"] = p.ExpirationInterval
    }
    if p.ExpirationIntervalUnit.IsValueSet() {
        if p.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    if p.DisplayOnHostedPage != nil {
        structMap["display_on_hosted_page"] = p.DisplayOnHostedPage
    }
    if p.AllowFractionalQuantities != nil {
        structMap["allow_fractional_quantities"] = p.AllowFractionalQuantities
    }
    if p.PublicSignupPageIds != nil {
        structMap["public_signup_page_ids"] = p.PublicSignupPageIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidUsageComponent.
// It customizes the JSON unmarshaling process for PrepaidUsageComponent objects.
func (p *PrepaidUsageComponent) UnmarshalJSON(input []byte) error {
    var temp tempPrepaidUsageComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "upgrade_charge", "downgrade_credit", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "overage_pricing", "rollover_prepaid_remainder", "renew_prepaid_allocation", "expiration_interval", "expiration_interval_unit", "display_on_hosted_page", "allow_fractional_quantities", "public_signup_page_ids")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Name = *temp.Name
    p.UnitName = *temp.UnitName
    p.Description = temp.Description
    p.Handle = temp.Handle
    p.Taxable = temp.Taxable
    p.PricingScheme = *temp.PricingScheme
    p.Prices = temp.Prices
    p.UpgradeCharge = temp.UpgradeCharge
    p.DowngradeCredit = temp.DowngradeCredit
    p.PricePoints = temp.PricePoints
    p.UnitPrice = temp.UnitPrice
    p.TaxCode = temp.TaxCode
    p.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    p.OveragePricing = *temp.OveragePricing
    p.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
    p.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
    p.ExpirationInterval = temp.ExpirationInterval
    p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    p.DisplayOnHostedPage = temp.DisplayOnHostedPage
    p.AllowFractionalQuantities = temp.AllowFractionalQuantities
    p.PublicSignupPageIds = temp.PublicSignupPageIds
    return nil
}

// tempPrepaidUsageComponent is a temporary struct used for validating the fields of PrepaidUsageComponent.
type tempPrepaidUsageComponent  struct {
    Name                      *string                                 `json:"name"`
    UnitName                  *string                                 `json:"unit_name"`
    Description               *string                                 `json:"description,omitempty"`
    Handle                    *string                                 `json:"handle,omitempty"`
    Taxable                   *bool                                   `json:"taxable,omitempty"`
    PricingScheme             *PricingScheme                          `json:"pricing_scheme"`
    Prices                    []Price                                 `json:"prices,omitempty"`
    UpgradeCharge             Optional[CreditType]                    `json:"upgrade_charge"`
    DowngradeCredit           Optional[CreditType]                    `json:"downgrade_credit"`
    PricePoints               []CreatePrepaidUsageComponentPricePoint `json:"price_points,omitempty"`
    UnitPrice                 *PrepaidUsageComponentUnitPrice         `json:"unit_price,omitempty"`
    TaxCode                   *string                                 `json:"tax_code,omitempty"`
    HideDateRangeOnInvoice    *bool                                   `json:"hide_date_range_on_invoice,omitempty"`
    OveragePricing            *OveragePricing                         `json:"overage_pricing"`
    RolloverPrepaidRemainder  *bool                                   `json:"rollover_prepaid_remainder,omitempty"`
    RenewPrepaidAllocation    *bool                                   `json:"renew_prepaid_allocation,omitempty"`
    ExpirationInterval        *float64                                `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit    Optional[ExpirationIntervalUnit]        `json:"expiration_interval_unit"`
    DisplayOnHostedPage       *bool                                   `json:"display_on_hosted_page,omitempty"`
    AllowFractionalQuantities *bool                                   `json:"allow_fractional_quantities,omitempty"`
    PublicSignupPageIds       []int                                   `json:"public_signup_page_ids,omitempty"`
}

func (p *tempPrepaidUsageComponent) validate() error {
    var errs []string
    if p.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Prepaid Usage Component`")
    }
    if p.UnitName == nil {
        errs = append(errs, "required field `unit_name` is missing for type `Prepaid Usage Component`")
    }
    if p.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Prepaid Usage Component`")
    }
    if p.OveragePricing == nil {
        errs = append(errs, "required field `overage_pricing` is missing for type `Prepaid Usage Component`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
