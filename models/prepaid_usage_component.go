package models

import (
    "encoding/json"
)

// PrepaidUsageComponent represents a PrepaidUsageComponent struct.
type PrepaidUsageComponent struct {
    // A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes".
    Name                      string                       `json:"name"`
    // The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. i.e. “message”, which may then be shown as “5 messages” on a subscription’s component line-item
    UnitName                  *string                      `json:"unit_name,omitempty"`
    // A description for the component that will be displayed to the user on the hosted signup page.
    Description               *string                      `json:"description,omitempty"`
    // A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.
    Handle                    *string                      `json:"handle,omitempty"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                        `json:"taxable,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme             *PricingScheme               `json:"pricing_scheme,omitempty"`
    // (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#general-price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
    Prices                    []Price                      `json:"prices,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    UpgradeCharge             Optional[CreditType]         `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    // Available values: `full`, `prorated`, `none`.
    DowngradeCredit           Optional[CreditType]         `json:"downgrade_credit"`
    PricePoints               []PrepaidComponentPricePoint `json:"price_points,omitempty"`
    // The amount the customer will be charged per unit when the pricing scheme is “per_unit”. For On/Off Components, this is the amount that the customer will be charged when they turn the component on for the subscription. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice                 *interface{}                 `json:"unit_price,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode                   *string                      `json:"tax_code,omitempty"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                        `json:"hide_date_range_on_invoice,omitempty"`
    // deprecated May 2011 - use unit_price instead
    PriceInCents              *string                      `json:"price_in_cents,omitempty"`
    OveragePricing            *OveragePricing              `json:"overage_pricing,omitempty"`
    // Boolean which controls whether or not remaining units should be rolled over to the next period
    RolloverPrepaidRemainder  *bool                        `json:"rollover_prepaid_remainder,omitempty"`
    // Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period
    RenewPrepaidAllocation    *bool                        `json:"renew_prepaid_allocation,omitempty"`
    // (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire
    ExpirationInterval        *float64                     `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit    *IntervalUnit                `json:"expiration_interval_unit,omitempty"`
    DisplayOnHostedPage       *bool                        `json:"display_on_hosted_page,omitempty"`
    AllowFractionalQuantities *bool                        `json:"allow_fractional_quantities,omitempty"`
    PublicSignupPageIds       []int                        `json:"public_signup_page_ids,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidUsageComponent.
// It customizes the JSON marshaling process for PrepaidUsageComponent objects.
func (p *PrepaidUsageComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidUsageComponent object to a map representation for JSON marshaling.
func (p *PrepaidUsageComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["name"] = p.Name
    if p.UnitName != nil {
        structMap["unit_name"] = p.UnitName
    }
    if p.Description != nil {
        structMap["description"] = p.Description
    }
    if p.Handle != nil {
        structMap["handle"] = p.Handle
    }
    if p.Taxable != nil {
        structMap["taxable"] = p.Taxable
    }
    if p.PricingScheme != nil {
        structMap["pricing_scheme"] = p.PricingScheme
    }
    if p.Prices != nil {
        structMap["prices"] = p.Prices
    }
    if p.UpgradeCharge.IsValueSet() {
        structMap["upgrade_charge"] = p.UpgradeCharge.Value()
    }
    if p.DowngradeCredit.IsValueSet() {
        structMap["downgrade_credit"] = p.DowngradeCredit.Value()
    }
    if p.PricePoints != nil {
        structMap["price_points"] = p.PricePoints
    }
    if p.UnitPrice != nil {
        structMap["unit_price"] = p.UnitPrice
    }
    if p.TaxCode != nil {
        structMap["tax_code"] = p.TaxCode
    }
    if p.HideDateRangeOnInvoice != nil {
        structMap["hide_date_range_on_invoice"] = p.HideDateRangeOnInvoice
    }
    if p.PriceInCents != nil {
        structMap["price_in_cents"] = p.PriceInCents
    }
    if p.OveragePricing != nil {
        structMap["overage_pricing"] = p.OveragePricing.toMap()
    }
    if p.RolloverPrepaidRemainder != nil {
        structMap["rollover_prepaid_remainder"] = p.RolloverPrepaidRemainder
    }
    if p.RenewPrepaidAllocation != nil {
        structMap["renew_prepaid_allocation"] = p.RenewPrepaidAllocation
    }
    if p.ExpirationInterval != nil {
        structMap["expiration_interval"] = p.ExpirationInterval
    }
    if p.ExpirationIntervalUnit != nil {
        structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit
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
    temp := &struct {
        Name                      string                       `json:"name"`
        UnitName                  *string                      `json:"unit_name,omitempty"`
        Description               *string                      `json:"description,omitempty"`
        Handle                    *string                      `json:"handle,omitempty"`
        Taxable                   *bool                        `json:"taxable,omitempty"`
        PricingScheme             *PricingScheme               `json:"pricing_scheme,omitempty"`
        Prices                    []Price                      `json:"prices,omitempty"`
        UpgradeCharge             Optional[CreditType]         `json:"upgrade_charge"`
        DowngradeCredit           Optional[CreditType]         `json:"downgrade_credit"`
        PricePoints               []PrepaidComponentPricePoint `json:"price_points,omitempty"`
        UnitPrice                 *interface{}                 `json:"unit_price,omitempty"`
        TaxCode                   *string                      `json:"tax_code,omitempty"`
        HideDateRangeOnInvoice    *bool                        `json:"hide_date_range_on_invoice,omitempty"`
        PriceInCents              *string                      `json:"price_in_cents,omitempty"`
        OveragePricing            *OveragePricing              `json:"overage_pricing,omitempty"`
        RolloverPrepaidRemainder  *bool                        `json:"rollover_prepaid_remainder,omitempty"`
        RenewPrepaidAllocation    *bool                        `json:"renew_prepaid_allocation,omitempty"`
        ExpirationInterval        *float64                     `json:"expiration_interval,omitempty"`
        ExpirationIntervalUnit    *IntervalUnit                `json:"expiration_interval_unit,omitempty"`
        DisplayOnHostedPage       *bool                        `json:"display_on_hosted_page,omitempty"`
        AllowFractionalQuantities *bool                        `json:"allow_fractional_quantities,omitempty"`
        PublicSignupPageIds       []int                        `json:"public_signup_page_ids,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Name = temp.Name
    p.UnitName = temp.UnitName
    p.Description = temp.Description
    p.Handle = temp.Handle
    p.Taxable = temp.Taxable
    p.PricingScheme = temp.PricingScheme
    p.Prices = temp.Prices
    p.UpgradeCharge = temp.UpgradeCharge
    p.DowngradeCredit = temp.DowngradeCredit
    p.PricePoints = temp.PricePoints
    p.UnitPrice = temp.UnitPrice
    p.TaxCode = temp.TaxCode
    p.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    p.PriceInCents = temp.PriceInCents
    p.OveragePricing = temp.OveragePricing
    p.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
    p.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
    p.ExpirationInterval = temp.ExpirationInterval
    p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    p.DisplayOnHostedPage = temp.DisplayOnHostedPage
    p.AllowFractionalQuantities = temp.AllowFractionalQuantities
    p.PublicSignupPageIds = temp.PublicSignupPageIds
    return nil
}
