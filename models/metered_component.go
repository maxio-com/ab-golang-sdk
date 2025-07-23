// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// MeteredComponent represents a MeteredComponent struct.
type MeteredComponent struct {
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
	// (Not required for ‘per_unit’ pricing schemes) One or more price brackets. See [Price Bracket Rules](https://maxio.zendesk.com/hc/en-us/articles/24261149166733-Component-Pricing-Schemes#price-bracket-rules) for an overview of how price brackets work for different pricing schemes.
	Prices      []Price                   `json:"prices,omitempty"`
	PricePoints []ComponentPricePointItem `json:"price_points,omitempty"`
	// The amount the customer will be charged per unit when the pricing scheme is “per_unit”. For On/Off Components, this is the amount that the customer will be charged when they turn the component on for the subscription. The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
	UnitPrice *MeteredComponentUnitPrice `json:"unit_price,omitempty"`
	// A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
	TaxCode *string `json:"tax_code,omitempty"`
	// (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
	HideDateRangeOnInvoice    *bool `json:"hide_date_range_on_invoice,omitempty"`
	DisplayOnHostedPage       *bool `json:"display_on_hosted_page,omitempty"`
	AllowFractionalQuantities *bool `json:"allow_fractional_quantities,omitempty"`
	PublicSignupPageIds       []int `json:"public_signup_page_ids,omitempty"`
	// The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
	Interval *int `json:"interval,omitempty"`
	// A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
	IntervalUnit         Optional[IntervalUnit] `json:"interval_unit"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MeteredComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MeteredComponent) String() string {
	return fmt.Sprintf(
		"MeteredComponent[Name=%v, UnitName=%v, Description=%v, Handle=%v, Taxable=%v, PricingScheme=%v, Prices=%v, PricePoints=%v, UnitPrice=%v, TaxCode=%v, HideDateRangeOnInvoice=%v, DisplayOnHostedPage=%v, AllowFractionalQuantities=%v, PublicSignupPageIds=%v, Interval=%v, IntervalUnit=%v, AdditionalProperties=%v]",
		m.Name, m.UnitName, m.Description, m.Handle, m.Taxable, m.PricingScheme, m.Prices, m.PricePoints, m.UnitPrice, m.TaxCode, m.HideDateRangeOnInvoice, m.DisplayOnHostedPage, m.AllowFractionalQuantities, m.PublicSignupPageIds, m.Interval, m.IntervalUnit, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MeteredComponent.
// It customizes the JSON marshaling process for MeteredComponent objects.
func (m MeteredComponent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "display_on_hosted_page", "allow_fractional_quantities", "public_signup_page_ids", "interval", "interval_unit"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MeteredComponent object to a map representation for JSON marshaling.
func (m MeteredComponent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	structMap["name"] = m.Name
	structMap["unit_name"] = m.UnitName
	if m.Description != nil {
		structMap["description"] = m.Description
	}
	if m.Handle != nil {
		structMap["handle"] = m.Handle
	}
	if m.Taxable != nil {
		structMap["taxable"] = m.Taxable
	}
	structMap["pricing_scheme"] = m.PricingScheme
	if m.Prices != nil {
		structMap["prices"] = m.Prices
	}
	if m.PricePoints != nil {
		structMap["price_points"] = m.PricePoints
	}
	if m.UnitPrice != nil {
		structMap["unit_price"] = m.UnitPrice.toMap()
	}
	if m.TaxCode != nil {
		structMap["tax_code"] = m.TaxCode
	}
	if m.HideDateRangeOnInvoice != nil {
		structMap["hide_date_range_on_invoice"] = m.HideDateRangeOnInvoice
	}
	if m.DisplayOnHostedPage != nil {
		structMap["display_on_hosted_page"] = m.DisplayOnHostedPage
	}
	if m.AllowFractionalQuantities != nil {
		structMap["allow_fractional_quantities"] = m.AllowFractionalQuantities
	}
	if m.PublicSignupPageIds != nil {
		structMap["public_signup_page_ids"] = m.PublicSignupPageIds
	}
	if m.Interval != nil {
		structMap["interval"] = m.Interval
	}
	if m.IntervalUnit.IsValueSet() {
		if m.IntervalUnit.Value() != nil {
			structMap["interval_unit"] = m.IntervalUnit.Value()
		} else {
			structMap["interval_unit"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MeteredComponent.
// It customizes the JSON unmarshaling process for MeteredComponent objects.
func (m *MeteredComponent) UnmarshalJSON(input []byte) error {
	var temp tempMeteredComponent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "unit_name", "description", "handle", "taxable", "pricing_scheme", "prices", "price_points", "unit_price", "tax_code", "hide_date_range_on_invoice", "display_on_hosted_page", "allow_fractional_quantities", "public_signup_page_ids", "interval", "interval_unit")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Name = *temp.Name
	m.UnitName = *temp.UnitName
	m.Description = temp.Description
	m.Handle = temp.Handle
	m.Taxable = temp.Taxable
	m.PricingScheme = *temp.PricingScheme
	m.Prices = temp.Prices
	m.PricePoints = temp.PricePoints
	m.UnitPrice = temp.UnitPrice
	m.TaxCode = temp.TaxCode
	m.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
	m.DisplayOnHostedPage = temp.DisplayOnHostedPage
	m.AllowFractionalQuantities = temp.AllowFractionalQuantities
	m.PublicSignupPageIds = temp.PublicSignupPageIds
	m.Interval = temp.Interval
	m.IntervalUnit = temp.IntervalUnit
	return nil
}

// tempMeteredComponent is a temporary struct used for validating the fields of MeteredComponent.
type tempMeteredComponent struct {
	Name                      *string                    `json:"name"`
	UnitName                  *string                    `json:"unit_name"`
	Description               *string                    `json:"description,omitempty"`
	Handle                    *string                    `json:"handle,omitempty"`
	Taxable                   *bool                      `json:"taxable,omitempty"`
	PricingScheme             *PricingScheme             `json:"pricing_scheme"`
	Prices                    []Price                    `json:"prices,omitempty"`
	PricePoints               []ComponentPricePointItem  `json:"price_points,omitempty"`
	UnitPrice                 *MeteredComponentUnitPrice `json:"unit_price,omitempty"`
	TaxCode                   *string                    `json:"tax_code,omitempty"`
	HideDateRangeOnInvoice    *bool                      `json:"hide_date_range_on_invoice,omitempty"`
	DisplayOnHostedPage       *bool                      `json:"display_on_hosted_page,omitempty"`
	AllowFractionalQuantities *bool                      `json:"allow_fractional_quantities,omitempty"`
	PublicSignupPageIds       []int                      `json:"public_signup_page_ids,omitempty"`
	Interval                  *int                       `json:"interval,omitempty"`
	IntervalUnit              Optional[IntervalUnit]     `json:"interval_unit"`
}

func (m *tempMeteredComponent) validate() error {
	var errs []string
	if m.Name == nil {
		errs = append(errs, "required field `name` is missing for type `Metered Component`")
	}
	if m.UnitName == nil {
		errs = append(errs, "required field `unit_name` is missing for type `Metered Component`")
	}
	if m.PricingScheme == nil {
		errs = append(errs, "required field `pricing_scheme` is missing for type `Metered Component`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
