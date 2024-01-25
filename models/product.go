package models

import (
	"encoding/json"
	"log"
	"time"
)

// Product represents a Product struct.
type Product struct {
	Id *int `json:"id,omitempty"`
	// The product name
	Name *string `json:"name,omitempty"`
	// The product API handle
	Handle Optional[string] `json:"handle"`
	// The product description
	Description Optional[string] `json:"description"`
	// E.g. Internal ID or SKU Number
	AccountingCode Optional[string] `json:"accounting_code"`
	// Deprecated value that can be ignored unless you have legacy hosted pages. For Public Signup Page users, please read this attribute from under the signup page.
	RequestCreditCard *bool `json:"request_credit_card,omitempty"`
	// A numerical interval for the length a subscription to this product will run before it expires. See the description of interval for a description of how this value is coupled with an interval unit to calculate the full interval
	ExpirationInterval Optional[int] `json:"expiration_interval"`
	// A string representing the trial interval unit for this product, either month or day
	ExpirationIntervalUnit Optional[ExtendedIntervalUnit] `json:"expiration_interval_unit"`
	// Timestamp indicating when this product was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp indicating when this product was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The product price, in integer cents
	PriceInCents *int64 `json:"price_in_cents,omitempty"`
	// The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product would renew every 30 days
	Interval *int `json:"interval,omitempty"`
	// A string representing the interval unit for this product, either month or day
	IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
	// The up front charge you have specified.
	InitialChargeInCents Optional[int64] `json:"initial_charge_in_cents"`
	// The price of the trial period for a subscription to this product, in integer cents.
	TrialPriceInCents Optional[int64] `json:"trial_price_in_cents"`
	// A numerical interval for the length of the trial period of a subscription to this product. See the description of interval for a description of how this value is coupled with an interval unit to calculate the full interval
	TrialInterval Optional[int] `json:"trial_interval"`
	// A string representing the trial interval unit for this product, either month or day
	TrialIntervalUnit Optional[IntervalUnit] `json:"trial_interval_unit"`
	// Timestamp indicating when this product was archived
	ArchivedAt Optional[time.Time] `json:"archived_at"`
	// Boolean that controls whether a payment profile is required to be entered for customers wishing to sign up on this product.
	RequireCreditCard *bool            `json:"require_credit_card,omitempty"`
	ReturnParams      Optional[string] `json:"return_params"`
	Taxable           *bool            `json:"taxable,omitempty"`
	// The url to which a customer will be returned after a successful account update
	UpdateReturnUrl         Optional[string] `json:"update_return_url"`
	InitialChargeAfterTrial *bool            `json:"initial_charge_after_trial,omitempty"`
	// The version of the product
	VersionNumber *int `json:"version_number,omitempty"`
	// The parameters will append to the url after a successful account update. See [help documentation](https://help.chargify.com/products/product-editing.html#return-parameters-after-account-update)
	UpdateReturnParams    Optional[string]   `json:"update_return_params"`
	ProductFamily         *ProductFamily     `json:"product_family,omitempty"`
	PublicSignupPages     []PublicSignupPage `json:"public_signup_pages,omitempty"`
	ProductPricePointName *string            `json:"product_price_point_name,omitempty"`
	// A boolean indicating whether to request a billing address on any Self-Service Pages that are used by subscribers of this product.
	RequestBillingAddress *bool `json:"request_billing_address,omitempty"`
	// A boolean indicating whether a billing address is required to add a payment profile, especially at signup.
	RequireBillingAddress *bool `json:"require_billing_address,omitempty"`
	// A boolean indicating whether a shipping address is required for the customer, especially at signup.
	RequireShippingAddress *bool `json:"require_shipping_address,omitempty"`
	// A string representing the tax code related to the product type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
	TaxCode                    Optional[string] `json:"tax_code"`
	DefaultProductPricePointId *int             `json:"default_product_price_point_id,omitempty"`
	UseSiteExchangeRate        Optional[bool]   `json:"use_site_exchange_rate"`
	// One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
	ItemCategory            Optional[string] `json:"item_category"`
	ProductPricePointId     *int             `json:"product_price_point_id,omitempty"`
	ProductPricePointHandle Optional[string] `json:"product_price_point_handle"`
}

// MarshalJSON implements the json.Marshaler interface for Product.
// It customizes the JSON marshaling process for Product objects.
func (p *Product) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the Product object to a map representation for JSON marshaling.
func (p *Product) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Id != nil {
		structMap["id"] = p.Id
	}
	if p.Name != nil {
		structMap["name"] = p.Name
	}
	if p.Handle.IsValueSet() {
		structMap["handle"] = p.Handle.Value()
	}
	if p.Description.IsValueSet() {
		structMap["description"] = p.Description.Value()
	}
	if p.AccountingCode.IsValueSet() {
		structMap["accounting_code"] = p.AccountingCode.Value()
	}
	if p.RequestCreditCard != nil {
		structMap["request_credit_card"] = p.RequestCreditCard
	}
	if p.ExpirationInterval.IsValueSet() {
		structMap["expiration_interval"] = p.ExpirationInterval.Value()
	}
	if p.ExpirationIntervalUnit.IsValueSet() {
		structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit.Value()
	}
	if p.CreatedAt != nil {
		structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
	}
	if p.UpdatedAt != nil {
		structMap["updated_at"] = p.UpdatedAt.Format(time.RFC3339)
	}
	if p.PriceInCents != nil {
		structMap["price_in_cents"] = p.PriceInCents
	}
	if p.Interval != nil {
		structMap["interval"] = p.Interval
	}
	if p.IntervalUnit != nil {
		structMap["interval_unit"] = p.IntervalUnit
	}
	if p.InitialChargeInCents.IsValueSet() {
		structMap["initial_charge_in_cents"] = p.InitialChargeInCents.Value()
	}
	if p.TrialPriceInCents.IsValueSet() {
		structMap["trial_price_in_cents"] = p.TrialPriceInCents.Value()
	}
	if p.TrialInterval.IsValueSet() {
		structMap["trial_interval"] = p.TrialInterval.Value()
	}
	if p.TrialIntervalUnit.IsValueSet() {
		structMap["trial_interval_unit"] = p.TrialIntervalUnit.Value()
	}
	if p.ArchivedAt.IsValueSet() {
		var ArchivedAtVal *string = nil
		if p.ArchivedAt.Value() != nil {
			val := p.ArchivedAt.Value().Format(time.RFC3339)
			ArchivedAtVal = &val
		}
		structMap["archived_at"] = ArchivedAtVal
	}
	if p.RequireCreditCard != nil {
		structMap["require_credit_card"] = p.RequireCreditCard
	}
	if p.ReturnParams.IsValueSet() {
		structMap["return_params"] = p.ReturnParams.Value()
	}
	if p.Taxable != nil {
		structMap["taxable"] = p.Taxable
	}
	if p.UpdateReturnUrl.IsValueSet() {
		structMap["update_return_url"] = p.UpdateReturnUrl.Value()
	}
	if p.InitialChargeAfterTrial != nil {
		structMap["initial_charge_after_trial"] = p.InitialChargeAfterTrial
	}
	if p.VersionNumber != nil {
		structMap["version_number"] = p.VersionNumber
	}
	if p.UpdateReturnParams.IsValueSet() {
		structMap["update_return_params"] = p.UpdateReturnParams.Value()
	}
	if p.ProductFamily != nil {
		structMap["product_family"] = p.ProductFamily
	}
	if p.PublicSignupPages != nil {
		structMap["public_signup_pages"] = p.PublicSignupPages
	}
	if p.ProductPricePointName != nil {
		structMap["product_price_point_name"] = p.ProductPricePointName
	}
	if p.RequestBillingAddress != nil {
		structMap["request_billing_address"] = p.RequestBillingAddress
	}
	if p.RequireBillingAddress != nil {
		structMap["require_billing_address"] = p.RequireBillingAddress
	}
	if p.RequireShippingAddress != nil {
		structMap["require_shipping_address"] = p.RequireShippingAddress
	}
	if p.TaxCode.IsValueSet() {
		structMap["tax_code"] = p.TaxCode.Value()
	}
	if p.DefaultProductPricePointId != nil {
		structMap["default_product_price_point_id"] = p.DefaultProductPricePointId
	}
	if p.UseSiteExchangeRate.IsValueSet() {
		structMap["use_site_exchange_rate"] = p.UseSiteExchangeRate.Value()
	}
	if p.ItemCategory.IsValueSet() {
		structMap["item_category"] = p.ItemCategory.Value()
	}
	if p.ProductPricePointId != nil {
		structMap["product_price_point_id"] = p.ProductPricePointId
	}
	if p.ProductPricePointHandle.IsValueSet() {
		structMap["product_price_point_handle"] = p.ProductPricePointHandle.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Product.
// It customizes the JSON unmarshaling process for Product objects.
func (p *Product) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                         *int                           `json:"id,omitempty"`
		Name                       *string                        `json:"name,omitempty"`
		Handle                     Optional[string]               `json:"handle"`
		Description                Optional[string]               `json:"description"`
		AccountingCode             Optional[string]               `json:"accounting_code"`
		RequestCreditCard          *bool                          `json:"request_credit_card,omitempty"`
		ExpirationInterval         Optional[int]                  `json:"expiration_interval"`
		ExpirationIntervalUnit     Optional[ExtendedIntervalUnit] `json:"expiration_interval_unit"`
		CreatedAt                  *string                        `json:"created_at,omitempty"`
		UpdatedAt                  *string                        `json:"updated_at,omitempty"`
		PriceInCents               *int64                         `json:"price_in_cents,omitempty"`
		Interval                   *int                           `json:"interval,omitempty"`
		IntervalUnit               *IntervalUnit                  `json:"interval_unit,omitempty"`
		InitialChargeInCents       Optional[int64]                `json:"initial_charge_in_cents"`
		TrialPriceInCents          Optional[int64]                `json:"trial_price_in_cents"`
		TrialInterval              Optional[int]                  `json:"trial_interval"`
		TrialIntervalUnit          Optional[IntervalUnit]         `json:"trial_interval_unit"`
		ArchivedAt                 Optional[string]               `json:"archived_at"`
		RequireCreditCard          *bool                          `json:"require_credit_card,omitempty"`
		ReturnParams               Optional[string]               `json:"return_params"`
		Taxable                    *bool                          `json:"taxable,omitempty"`
		UpdateReturnUrl            Optional[string]               `json:"update_return_url"`
		InitialChargeAfterTrial    *bool                          `json:"initial_charge_after_trial,omitempty"`
		VersionNumber              *int                           `json:"version_number,omitempty"`
		UpdateReturnParams         Optional[string]               `json:"update_return_params"`
		ProductFamily              *ProductFamily                 `json:"product_family,omitempty"`
		PublicSignupPages          []PublicSignupPage             `json:"public_signup_pages,omitempty"`
		ProductPricePointName      *string                        `json:"product_price_point_name,omitempty"`
		RequestBillingAddress      *bool                          `json:"request_billing_address,omitempty"`
		RequireBillingAddress      *bool                          `json:"require_billing_address,omitempty"`
		RequireShippingAddress     *bool                          `json:"require_shipping_address,omitempty"`
		TaxCode                    Optional[string]               `json:"tax_code"`
		DefaultProductPricePointId *int                           `json:"default_product_price_point_id,omitempty"`
		UseSiteExchangeRate        Optional[bool]                 `json:"use_site_exchange_rate"`
		ItemCategory               Optional[string]               `json:"item_category"`
		ProductPricePointId        *int                           `json:"product_price_point_id,omitempty"`
		ProductPricePointHandle    Optional[string]               `json:"product_price_point_handle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Id = temp.Id
	p.Name = temp.Name
	p.Handle = temp.Handle
	p.Description = temp.Description
	p.AccountingCode = temp.AccountingCode
	p.RequestCreditCard = temp.RequestCreditCard
	p.ExpirationInterval = temp.ExpirationInterval
	p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		p.CreatedAt = &CreatedAtVal
	}
	if temp.UpdatedAt != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		p.UpdatedAt = &UpdatedAtVal
	}
	p.PriceInCents = temp.PriceInCents
	p.Interval = temp.Interval
	p.IntervalUnit = temp.IntervalUnit
	p.InitialChargeInCents = temp.InitialChargeInCents
	p.TrialPriceInCents = temp.TrialPriceInCents
	p.TrialInterval = temp.TrialInterval
	p.TrialIntervalUnit = temp.TrialIntervalUnit
	p.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
	if temp.ArchivedAt.Value() != nil {
		ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
		}
		p.ArchivedAt.SetValue(&ArchivedAtVal)
	}
	p.RequireCreditCard = temp.RequireCreditCard
	p.ReturnParams = temp.ReturnParams
	p.Taxable = temp.Taxable
	p.UpdateReturnUrl = temp.UpdateReturnUrl
	p.InitialChargeAfterTrial = temp.InitialChargeAfterTrial
	p.VersionNumber = temp.VersionNumber
	p.UpdateReturnParams = temp.UpdateReturnParams
	p.ProductFamily = temp.ProductFamily
	p.PublicSignupPages = temp.PublicSignupPages
	p.ProductPricePointName = temp.ProductPricePointName
	p.RequestBillingAddress = temp.RequestBillingAddress
	p.RequireBillingAddress = temp.RequireBillingAddress
	p.RequireShippingAddress = temp.RequireShippingAddress
	p.TaxCode = temp.TaxCode
	p.DefaultProductPricePointId = temp.DefaultProductPricePointId
	p.UseSiteExchangeRate = temp.UseSiteExchangeRate
	p.ItemCategory = temp.ItemCategory
	p.ProductPricePointId = temp.ProductPricePointId
	p.ProductPricePointHandle = temp.ProductPricePointHandle
	return nil
}
