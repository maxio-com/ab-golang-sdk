package models

import (
	"encoding/json"
	"log"
	"time"
)

// Component represents a Component struct.
type Component struct {
	// The unique ID assigned to the component by Chargify. This ID can be used to fetch the component from the API.
	Id *int `json:"id,omitempty"`
	// The name of the Component, suitable for display on statements. i.e. Text Messages.
	Name *string `json:"name,omitempty"`
	// The component API handle
	Handle        Optional[string]                 `json:"handle"`
	PricingScheme Optional[ComponentPricingScheme] `json:"pricing_scheme"`
	// The name of the unit that the component’s usage is measured in. i.e. message
	UnitName *string `json:"unit_name,omitempty"`
	// The amount the customer will be charged per unit. This field is only populated for ‘per_unit’ pricing schemes, otherwise it may be null.
	UnitPrice Optional[string] `json:"unit_price"`
	// The id of the Product Family to which the Component belongs
	ProductFamilyId *int `json:"product_family_id,omitempty"`
	// The name of the Product Family to which the Component belongs
	ProductFamilyName *string `json:"product_family_name,omitempty"`
	// deprecated - use unit_price instead
	PricePerUnitInCents Optional[int64] `json:"price_per_unit_in_cents"`
	// A handle for the component type
	Kind *ComponentKind `json:"kind,omitempty"`
	// Boolean flag describing whether a component is archived or not.
	Archived *bool `json:"archived,omitempty"`
	// Boolean flag describing whether a component is taxable or not.
	Taxable *bool `json:"taxable,omitempty"`
	// The description of the component.
	Description         Optional[string] `json:"description"`
	DefaultPricePointId Optional[int]    `json:"default_price_point_id"`
	// An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty.
	OveragePrices Optional[[]ComponentPrice] `json:"overage_prices"`
	// An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty.
	Prices Optional[[]ComponentPrice] `json:"prices"`
	// Count for the number of price points associated with the component
	PricePointCount *int `json:"price_point_count,omitempty"`
	// URL that points to the location to read the existing price points via GET request
	PricePointsUrl        *string `json:"price_points_url,omitempty"`
	DefaultPricePointName *string `json:"default_price_point_name,omitempty"`
	// A string representing the tax code related to the component type. This is especially important when using the Avalara service to tax based on locale. This attribute has a max length of 10 characters.
	TaxCode   Optional[string] `json:"tax_code"`
	Recurring *bool            `json:"recurring,omitempty"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	UpgradeCharge Optional[CreditType] `json:"upgrade_charge"`
	// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
	// Available values: `full`, `prorated`, `none`.
	DowngradeCredit Optional[CreditType] `json:"downgrade_credit"`
	// Timestamp indicating when this component was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp indicating when this component was updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Timestamp indicating when this component was archived
	ArchivedAt Optional[time.Time] `json:"archived_at"`
	// (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
	HideDateRangeOnInvoice    *bool `json:"hide_date_range_on_invoice,omitempty"`
	AllowFractionalQuantities *bool `json:"allow_fractional_quantities,omitempty"`
	// One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
	ItemCategory        Optional[ItemCategory] `json:"item_category"`
	UseSiteExchangeRate Optional[bool]         `json:"use_site_exchange_rate"`
	// E.g. Internal ID or SKU Number
	AccountingCode Optional[string] `json:"accounting_code"`
	// (Only for Event Based Components) This is an ID of a metric attached to the component. This metric is used to bill upon collected events.
	EventBasedBillingMetricId *int `json:"event_based_billing_metric_id,omitempty"`
	// The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
	Interval *int `json:"interval,omitempty"`
	// A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
	IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Component.
// It customizes the JSON marshaling process for Component objects.
func (c *Component) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the Component object to a map representation for JSON marshaling.
func (c *Component) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.Handle.IsValueSet() {
		if c.Handle.Value() != nil {
			structMap["handle"] = c.Handle.Value()
		} else {
			structMap["handle"] = nil
		}
	}
	if c.PricingScheme.IsValueSet() {
		if c.PricingScheme.Value() != nil {
			structMap["pricing_scheme"] = c.PricingScheme.Value().toMap()
		} else {
			structMap["pricing_scheme"] = nil
		}
	}
	if c.UnitName != nil {
		structMap["unit_name"] = c.UnitName
	}
	if c.UnitPrice.IsValueSet() {
		if c.UnitPrice.Value() != nil {
			structMap["unit_price"] = c.UnitPrice.Value()
		} else {
			structMap["unit_price"] = nil
		}
	}
	if c.ProductFamilyId != nil {
		structMap["product_family_id"] = c.ProductFamilyId
	}
	if c.ProductFamilyName != nil {
		structMap["product_family_name"] = c.ProductFamilyName
	}
	if c.PricePerUnitInCents.IsValueSet() {
		if c.PricePerUnitInCents.Value() != nil {
			structMap["price_per_unit_in_cents"] = c.PricePerUnitInCents.Value()
		} else {
			structMap["price_per_unit_in_cents"] = nil
		}
	}
	if c.Kind != nil {
		structMap["kind"] = c.Kind
	}
	if c.Archived != nil {
		structMap["archived"] = c.Archived
	}
	if c.Taxable != nil {
		structMap["taxable"] = c.Taxable
	}
	if c.Description.IsValueSet() {
		if c.Description.Value() != nil {
			structMap["description"] = c.Description.Value()
		} else {
			structMap["description"] = nil
		}
	}
	if c.DefaultPricePointId.IsValueSet() {
		if c.DefaultPricePointId.Value() != nil {
			structMap["default_price_point_id"] = c.DefaultPricePointId.Value()
		} else {
			structMap["default_price_point_id"] = nil
		}
	}
	if c.OveragePrices.IsValueSet() {
		if c.OveragePrices.Value() != nil {
			structMap["overage_prices"] = c.OveragePrices.Value()
		} else {
			structMap["overage_prices"] = nil
		}
	}
	if c.Prices.IsValueSet() {
		if c.Prices.Value() != nil {
			structMap["prices"] = c.Prices.Value()
		} else {
			structMap["prices"] = nil
		}
	}
	if c.PricePointCount != nil {
		structMap["price_point_count"] = c.PricePointCount
	}
	if c.PricePointsUrl != nil {
		structMap["price_points_url"] = c.PricePointsUrl
	}
	if c.DefaultPricePointName != nil {
		structMap["default_price_point_name"] = c.DefaultPricePointName
	}
	if c.TaxCode.IsValueSet() {
		if c.TaxCode.Value() != nil {
			structMap["tax_code"] = c.TaxCode.Value()
		} else {
			structMap["tax_code"] = nil
		}
	}
	if c.Recurring != nil {
		structMap["recurring"] = c.Recurring
	}
	if c.UpgradeCharge.IsValueSet() {
		if c.UpgradeCharge.Value() != nil {
			structMap["upgrade_charge"] = c.UpgradeCharge.Value()
		} else {
			structMap["upgrade_charge"] = nil
		}
	}
	if c.DowngradeCredit.IsValueSet() {
		if c.DowngradeCredit.Value() != nil {
			structMap["downgrade_credit"] = c.DowngradeCredit.Value()
		} else {
			structMap["downgrade_credit"] = nil
		}
	}
	if c.CreatedAt != nil {
		structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
	}
	if c.UpdatedAt != nil {
		structMap["updated_at"] = c.UpdatedAt.Format(time.RFC3339)
	}
	if c.ArchivedAt.IsValueSet() {
		var ArchivedAtVal *string = nil
		if c.ArchivedAt.Value() != nil {
			val := c.ArchivedAt.Value().Format(time.RFC3339)
			ArchivedAtVal = &val
		}
		if c.ArchivedAt.Value() != nil {
			structMap["archived_at"] = ArchivedAtVal
		} else {
			structMap["archived_at"] = nil
		}
	}
	if c.HideDateRangeOnInvoice != nil {
		structMap["hide_date_range_on_invoice"] = c.HideDateRangeOnInvoice
	}
	if c.AllowFractionalQuantities != nil {
		structMap["allow_fractional_quantities"] = c.AllowFractionalQuantities
	}
	if c.ItemCategory.IsValueSet() {
		if c.ItemCategory.Value() != nil {
			structMap["item_category"] = c.ItemCategory.Value()
		} else {
			structMap["item_category"] = nil
		}
	}
	if c.UseSiteExchangeRate.IsValueSet() {
		if c.UseSiteExchangeRate.Value() != nil {
			structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate.Value()
		} else {
			structMap["use_site_exchange_rate"] = nil
		}
	}
	if c.AccountingCode.IsValueSet() {
		if c.AccountingCode.Value() != nil {
			structMap["accounting_code"] = c.AccountingCode.Value()
		} else {
			structMap["accounting_code"] = nil
		}
	}
	if c.EventBasedBillingMetricId != nil {
		structMap["event_based_billing_metric_id"] = c.EventBasedBillingMetricId
	}
	if c.Interval != nil {
		structMap["interval"] = c.Interval
	}
	if c.IntervalUnit != nil {
		structMap["interval_unit"] = c.IntervalUnit
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Component.
// It customizes the JSON unmarshaling process for Component objects.
func (c *Component) UnmarshalJSON(input []byte) error {
	var temp component
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	c.Id = temp.Id
	c.Name = temp.Name
	c.Handle = temp.Handle
	c.PricingScheme = temp.PricingScheme
	c.UnitName = temp.UnitName
	c.UnitPrice = temp.UnitPrice
	c.ProductFamilyId = temp.ProductFamilyId
	c.ProductFamilyName = temp.ProductFamilyName
	c.PricePerUnitInCents = temp.PricePerUnitInCents
	c.Kind = temp.Kind
	c.Archived = temp.Archived
	c.Taxable = temp.Taxable
	c.Description = temp.Description
	c.DefaultPricePointId = temp.DefaultPricePointId
	c.OveragePrices = temp.OveragePrices
	c.Prices = temp.Prices
	c.PricePointCount = temp.PricePointCount
	c.PricePointsUrl = temp.PricePointsUrl
	c.DefaultPricePointName = temp.DefaultPricePointName
	c.TaxCode = temp.TaxCode
	c.Recurring = temp.Recurring
	c.UpgradeCharge = temp.UpgradeCharge
	c.DowngradeCredit = temp.DowngradeCredit
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		c.CreatedAt = &CreatedAtVal
	}
	if temp.UpdatedAt != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		c.UpdatedAt = &UpdatedAtVal
	}
	c.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
	if temp.ArchivedAt.Value() != nil {
		ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
		}
		c.ArchivedAt.SetValue(&ArchivedAtVal)
	}
	c.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
	c.AllowFractionalQuantities = temp.AllowFractionalQuantities
	c.ItemCategory = temp.ItemCategory
	c.UseSiteExchangeRate = temp.UseSiteExchangeRate
	c.AccountingCode = temp.AccountingCode
	c.EventBasedBillingMetricId = temp.EventBasedBillingMetricId
	c.Interval = temp.Interval
	c.IntervalUnit = temp.IntervalUnit
	return nil
}

// TODO
type component struct {
	Id                        *int                             `json:"id,omitempty"`
	Name                      *string                          `json:"name,omitempty"`
	Handle                    Optional[string]                 `json:"handle"`
	PricingScheme             Optional[ComponentPricingScheme] `json:"pricing_scheme"`
	UnitName                  *string                          `json:"unit_name,omitempty"`
	UnitPrice                 Optional[string]                 `json:"unit_price"`
	ProductFamilyId           *int                             `json:"product_family_id,omitempty"`
	ProductFamilyName         *string                          `json:"product_family_name,omitempty"`
	PricePerUnitInCents       Optional[int64]                  `json:"price_per_unit_in_cents"`
	Kind                      *ComponentKind                   `json:"kind,omitempty"`
	Archived                  *bool                            `json:"archived,omitempty"`
	Taxable                   *bool                            `json:"taxable,omitempty"`
	Description               Optional[string]                 `json:"description"`
	DefaultPricePointId       Optional[int]                    `json:"default_price_point_id"`
	OveragePrices             Optional[[]ComponentPrice]       `json:"overage_prices"`
	Prices                    Optional[[]ComponentPrice]       `json:"prices"`
	PricePointCount           *int                             `json:"price_point_count,omitempty"`
	PricePointsUrl            *string                          `json:"price_points_url,omitempty"`
	DefaultPricePointName     *string                          `json:"default_price_point_name,omitempty"`
	TaxCode                   Optional[string]                 `json:"tax_code"`
	Recurring                 *bool                            `json:"recurring,omitempty"`
	UpgradeCharge             Optional[CreditType]             `json:"upgrade_charge"`
	DowngradeCredit           Optional[CreditType]             `json:"downgrade_credit"`
	CreatedAt                 *string                          `json:"created_at,omitempty"`
	UpdatedAt                 *string                          `json:"updated_at,omitempty"`
	ArchivedAt                Optional[string]                 `json:"archived_at"`
	HideDateRangeOnInvoice    *bool                            `json:"hide_date_range_on_invoice,omitempty"`
	AllowFractionalQuantities *bool                            `json:"allow_fractional_quantities,omitempty"`
	ItemCategory              Optional[ItemCategory]           `json:"item_category"`
	UseSiteExchangeRate       Optional[bool]                   `json:"use_site_exchange_rate"`
	AccountingCode            Optional[string]                 `json:"accounting_code"`
	EventBasedBillingMetricId *int                             `json:"event_based_billing_metric_id,omitempty"`
	Interval                  *int                             `json:"interval,omitempty"`
	IntervalUnit              *IntervalUnit                    `json:"interval_unit,omitempty"`
}
