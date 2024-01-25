package models

import (
	"encoding/json"
)

// QuantityBasedComponent represents a QuantityBasedComponent struct.
type QuantityBasedComponent struct {
	// A name for this component that is suitable for showing customers and displaying on billing statements, ie. "Minutes".
	Name string `json:"name"`
	// The name of the unit of measurement for the component. It should be singular since it will be automatically pluralized when necessary. i.e. “message”, which may then be shown as “5 messages” on a subscription’s component line-item
	UnitName string `json:"unit_name"`
	// A description for the component that will be displayed to the user on the hosted signup page.
	Description *string `json:"description,omitempty"`
	// A unique identifier for your use that can be used to retrieve this component is subsequent requests.  Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.
	Handle *string `json:"handle,omitempty"`
	// Boolean flag describing whether a component is taxable or not.
	Taxable *bool `json:"taxable,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme PricingScheme `json:"pricing_scheme"`
	// (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://chargify.zendesk.com/hc/en-us/articles/4407755865883#price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
	Prices []Price `json:"prices,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditType] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit Optional[CreditType]      `json:"downgrade_credit"`
	PricePoints     []ComponentPricePointItem `json:"price_points,omitempty"`
	// The amount the customer will be charged per unit when the pricing scheme is “per_unit”. For On/Off Components, this is the amount that the customer will be charged when they turn the component on for the subscription. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
	UnitPrice *interface{} `json:"unit_price,omitempty"`
	// A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
	TaxCode *string `json:"tax_code,omitempty"`
	// (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
	HideDateRangeOnInvoice *bool `json:"hide_date_range_on_invoice,omitempty"`
	// deprecated May 2011 - use unit_price instead
	PriceInCents              *string `json:"price_in_cents,omitempty"`
	Recurring                 *bool   `json:"recurring,omitempty"`
	DisplayOnHostedPage       *bool   `json:"display_on_hosted_page,omitempty"`
	AllowFractionalQuantities *bool   `json:"allow_fractional_quantities,omitempty"`
	PublicSignupPageIds       []int   `json:"public_signup_page_ids,omitempty"`
	// The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
	Interval *int `json:"interval,omitempty"`
	// A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
	IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for QuantityBasedComponent.
// It customizes the JSON marshaling process for QuantityBasedComponent objects.
func (q *QuantityBasedComponent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(q.toMap())
}

// toMap converts the QuantityBasedComponent object to a map representation for JSON marshaling.
func (q *QuantityBasedComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = q.Name
	structMap["unit_name"] = q.UnitName
	if q.Description != nil {
		structMap["description"] = q.Description
	}
	if q.Handle != nil {
		structMap["handle"] = q.Handle
	}
	if q.Taxable != nil {
		structMap["taxable"] = q.Taxable
	}
	structMap["pricing_scheme"] = q.PricingScheme
	if q.Prices != nil {
		structMap["prices"] = q.Prices
	}
	if q.UpgradeCharge.IsValueSet() {
		structMap["upgrade_charge"] = q.UpgradeCharge.Value()
	}
	if q.DowngradeCredit.IsValueSet() {
		structMap["downgrade_credit"] = q.DowngradeCredit.Value()
	}
	if q.PricePoints != nil {
		structMap["price_points"] = q.PricePoints
	}
	if q.UnitPrice != nil {
		structMap["unit_price"] = q.UnitPrice
	}
	if q.TaxCode != nil {
		structMap["tax_code"] = q.TaxCode
	}
	if q.HideDateRangeOnInvoice != nil {
		structMap["hide_date_range_on_invoice"] = q.HideDateRangeOnInvoice
	}
	if q.PriceInCents != nil {
		structMap["price_in_cents"] = q.PriceInCents
	}
	if q.Recurring != nil {
		structMap["recurring"] = q.Recurring
	}
	if q.DisplayOnHostedPage != nil {
		structMap["display_on_hosted_page"] = q.DisplayOnHostedPage
	}
	if q.AllowFractionalQuantities != nil {
		structMap["allow_fractional_quantities"] = q.AllowFractionalQuantities
	}
	if q.PublicSignupPageIds != nil {
		structMap["public_signup_page_ids"] = q.PublicSignupPageIds
	}
	if q.Interval != nil {
		structMap["interval"] = q.Interval
	}
	if q.IntervalUnit != nil {
		structMap["interval_unit"] = q.IntervalUnit
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for QuantityBasedComponent.
// It customizes the JSON unmarshaling process for QuantityBasedComponent objects.
func (q *QuantityBasedComponent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                      string                    `json:"name"`
		UnitName                  string                    `json:"unit_name"`
		Description               *string                   `json:"description,omitempty"`
		Handle                    *string                   `json:"handle,omitempty"`
		Taxable                   *bool                     `json:"taxable,omitempty"`
		PricingScheme             PricingScheme             `json:"pricing_scheme"`
		Prices                    []Price                   `json:"prices,omitempty"`
		UpgradeCharge             Optional[CreditType]      `json:"upgrade_charge"`
		DowngradeCredit           Optional[CreditType]      `json:"downgrade_credit"`
		PricePoints               []ComponentPricePointItem `json:"price_points,omitempty"`
		UnitPrice                 *interface{}              `json:"unit_price,omitempty"`
		TaxCode                   *string                   `json:"tax_code,omitempty"`
		HideDateRangeOnInvoice    *bool                     `json:"hide_date_range_on_invoice,omitempty"`
		PriceInCents              *string                   `json:"price_in_cents,omitempty"`
		Recurring                 *bool                     `json:"recurring,omitempty"`
		DisplayOnHostedPage       *bool                     `json:"display_on_hosted_page,omitempty"`
		AllowFractionalQuantities *bool                     `json:"allow_fractional_quantities,omitempty"`
		PublicSignupPageIds       []int                     `json:"public_signup_page_ids,omitempty"`
		Interval                  *int                      `json:"interval,omitempty"`
		IntervalUnit              *IntervalUnit             `json:"interval_unit,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	q.Name = temp.Name
	q.UnitName = temp.UnitName
	q.Description = temp.Description
	q.Handle = temp.Handle
	q.Taxable = temp.Taxable
	q.PricingScheme = temp.PricingScheme
	q.Prices = temp.Prices
	q.UpgradeCharge = temp.UpgradeCharge
	q.DowngradeCredit = temp.DowngradeCredit
	q.PricePoints = temp.PricePoints
	q.UnitPrice = temp.UnitPrice
	q.TaxCode = temp.TaxCode
	q.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
	q.PriceInCents = temp.PriceInCents
	q.Recurring = temp.Recurring
	q.DisplayOnHostedPage = temp.DisplayOnHostedPage
	q.AllowFractionalQuantities = temp.AllowFractionalQuantities
	q.PublicSignupPageIds = temp.PublicSignupPageIds
	q.Interval = temp.Interval
	q.IntervalUnit = temp.IntervalUnit
	return nil
}
