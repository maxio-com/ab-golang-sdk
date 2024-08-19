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

// CreateOrUpdateProduct represents a CreateOrUpdateProduct struct.
type CreateOrUpdateProduct struct {
    // The product name
    Name                   string                           `json:"name"`
    // The product API handle
    Handle                 *string                          `json:"handle,omitempty"`
    // The product description
    Description            string                           `json:"description"`
    // E.g. Internal ID or SKU Number
    AccountingCode         *string                          `json:"accounting_code,omitempty"`
    // Deprecated value that can be ignored unless you have legacy hosted pages. For Public Signup Page users, please read this attribute from under the signup page.
    RequireCreditCard      *bool                            `json:"require_credit_card,omitempty"`
    // The product price, in integer cents
    PriceInCents           int64                            `json:"price_in_cents"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product would renew every 30 days
    Interval               int                              `json:"interval"`
    // A string representing the interval unit for this product, either month or day
    IntervalUnit           IntervalUnit                     `json:"interval_unit"`
    // The product trial price, in integer cents
    TrialPriceInCents      *int64                           `json:"trial_price_in_cents,omitempty"`
    // The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product trial would last 30 days.
    TrialInterval          *int                             `json:"trial_interval,omitempty"`
    // A string representing the trial interval unit for this product, either month or day
    TrialIntervalUnit      Optional[IntervalUnit]           `json:"trial_interval_unit"`
    TrialType              *string                          `json:"trial_type,omitempty"`
    // The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product would expire after 30 days.
    ExpirationInterval     *int                             `json:"expiration_interval,omitempty"`
    // A string representing the expiration interval unit for this product, either month, day or never
    ExpirationIntervalUnit Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
    AutoCreateSignupPage   *bool                            `json:"auto_create_signup_page,omitempty"`
    // A string representing the tax code related to the product type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode                *string                          `json:"tax_code,omitempty"`
    AdditionalProperties   map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateProduct.
// It customizes the JSON marshaling process for CreateOrUpdateProduct objects.
func (c CreateOrUpdateProduct) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateProduct object to a map representation for JSON marshaling.
func (c CreateOrUpdateProduct) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    if c.Handle != nil {
        structMap["handle"] = c.Handle
    }
    structMap["description"] = c.Description
    if c.AccountingCode != nil {
        structMap["accounting_code"] = c.AccountingCode
    }
    if c.RequireCreditCard != nil {
        structMap["require_credit_card"] = c.RequireCreditCard
    }
    structMap["price_in_cents"] = c.PriceInCents
    structMap["interval"] = c.Interval
    structMap["interval_unit"] = c.IntervalUnit
    if c.TrialPriceInCents != nil {
        structMap["trial_price_in_cents"] = c.TrialPriceInCents
    }
    if c.TrialInterval != nil {
        structMap["trial_interval"] = c.TrialInterval
    }
    if c.TrialIntervalUnit.IsValueSet() {
        if c.TrialIntervalUnit.Value() != nil {
            structMap["trial_interval_unit"] = c.TrialIntervalUnit.Value()
        } else {
            structMap["trial_interval_unit"] = nil
        }
    }
    if c.TrialType != nil {
        structMap["trial_type"] = c.TrialType
    }
    if c.ExpirationInterval != nil {
        structMap["expiration_interval"] = c.ExpirationInterval
    }
    if c.ExpirationIntervalUnit.IsValueSet() {
        if c.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = c.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    if c.AutoCreateSignupPage != nil {
        structMap["auto_create_signup_page"] = c.AutoCreateSignupPage
    }
    if c.TaxCode != nil {
        structMap["tax_code"] = c.TaxCode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateProduct.
// It customizes the JSON unmarshaling process for CreateOrUpdateProduct objects.
func (c *CreateOrUpdateProduct) UnmarshalJSON(input []byte) error {
    var temp tempCreateOrUpdateProduct
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "handle", "description", "accounting_code", "require_credit_card", "price_in_cents", "interval", "interval_unit", "trial_price_in_cents", "trial_interval", "trial_interval_unit", "trial_type", "expiration_interval", "expiration_interval_unit", "auto_create_signup_page", "tax_code")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = *temp.Name
    c.Handle = temp.Handle
    c.Description = *temp.Description
    c.AccountingCode = temp.AccountingCode
    c.RequireCreditCard = temp.RequireCreditCard
    c.PriceInCents = *temp.PriceInCents
    c.Interval = *temp.Interval
    c.IntervalUnit = *temp.IntervalUnit
    c.TrialPriceInCents = temp.TrialPriceInCents
    c.TrialInterval = temp.TrialInterval
    c.TrialIntervalUnit = temp.TrialIntervalUnit
    c.TrialType = temp.TrialType
    c.ExpirationInterval = temp.ExpirationInterval
    c.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    c.AutoCreateSignupPage = temp.AutoCreateSignupPage
    c.TaxCode = temp.TaxCode
    return nil
}

// tempCreateOrUpdateProduct is a temporary struct used for validating the fields of CreateOrUpdateProduct.
type tempCreateOrUpdateProduct  struct {
    Name                   *string                          `json:"name"`
    Handle                 *string                          `json:"handle,omitempty"`
    Description            *string                          `json:"description"`
    AccountingCode         *string                          `json:"accounting_code,omitempty"`
    RequireCreditCard      *bool                            `json:"require_credit_card,omitempty"`
    PriceInCents           *int64                           `json:"price_in_cents"`
    Interval               *int                             `json:"interval"`
    IntervalUnit           *IntervalUnit                    `json:"interval_unit"`
    TrialPriceInCents      *int64                           `json:"trial_price_in_cents,omitempty"`
    TrialInterval          *int                             `json:"trial_interval,omitempty"`
    TrialIntervalUnit      Optional[IntervalUnit]           `json:"trial_interval_unit"`
    TrialType              *string                          `json:"trial_type,omitempty"`
    ExpirationInterval     *int                             `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
    AutoCreateSignupPage   *bool                            `json:"auto_create_signup_page,omitempty"`
    TaxCode                *string                          `json:"tax_code,omitempty"`
}

func (c *tempCreateOrUpdateProduct) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Create or Update Product`")
    }
    if c.Description == nil {
        errs = append(errs, "required field `description` is missing for type `Create or Update Product`")
    }
    if c.PriceInCents == nil {
        errs = append(errs, "required field `price_in_cents` is missing for type `Create or Update Product`")
    }
    if c.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `Create or Update Product`")
    }
    if c.IntervalUnit == nil {
        errs = append(errs, "required field `interval_unit` is missing for type `Create or Update Product`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
