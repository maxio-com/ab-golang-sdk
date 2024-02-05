package models

import (
    "encoding/json"
)

// CreateOrUpdateProduct represents a CreateOrUpdateProduct struct.
type CreateOrUpdateProduct struct {
    // The product name
    Name                 string       `json:"name"`
    // The product API handle
    Handle               *string      `json:"handle,omitempty"`
    // The product description
    Description          string       `json:"description"`
    // E.g. Internal ID or SKU Number
    AccountingCode       *string      `json:"accounting_code,omitempty"`
    // Deprecated value that can be ignored unless you have legacy hosted pages. For Public Signup Page users, please read this attribute from under the signup page.
    RequireCreditCard    *bool        `json:"require_credit_card,omitempty"`
    // The product price, in integer cents
    PriceInCents         int64        `json:"price_in_cents"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product would renew every 30 days
    Interval             int          `json:"interval"`
    // A string representing the interval unit for this product, either month or day
    IntervalUnit         IntervalUnit `json:"interval_unit"`
    AutoCreateSignupPage *bool        `json:"auto_create_signup_page,omitempty"`
    // A string representing the tax code related to the product type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
    TaxCode              *string      `json:"tax_code,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateProduct.
// It customizes the JSON marshaling process for CreateOrUpdateProduct objects.
func (c *CreateOrUpdateProduct) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateProduct object to a map representation for JSON marshaling.
func (c *CreateOrUpdateProduct) toMap() map[string]any {
    structMap := make(map[string]any)
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
    temp := &struct {
        Name                 string       `json:"name"`
        Handle               *string      `json:"handle,omitempty"`
        Description          string       `json:"description"`
        AccountingCode       *string      `json:"accounting_code,omitempty"`
        RequireCreditCard    *bool        `json:"require_credit_card,omitempty"`
        PriceInCents         int64        `json:"price_in_cents"`
        Interval             int          `json:"interval"`
        IntervalUnit         IntervalUnit `json:"interval_unit"`
        AutoCreateSignupPage *bool        `json:"auto_create_signup_page,omitempty"`
        TaxCode              *string      `json:"tax_code,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Name = temp.Name
    c.Handle = temp.Handle
    c.Description = temp.Description
    c.AccountingCode = temp.AccountingCode
    c.RequireCreditCard = temp.RequireCreditCard
    c.PriceInCents = temp.PriceInCents
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.AutoCreateSignupPage = temp.AutoCreateSignupPage
    c.TaxCode = temp.TaxCode
    return nil
}
