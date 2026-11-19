// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// WalletProductListItem represents a WalletProductListItem struct.
// A Wallet Product as it appears in [List Wallet Products]($e/Wallet%20Products/listWalletProducts): the full Component field set, plus a usage count and a price point summary that only the list response reports.
type WalletProductListItem struct {
    // The unique ID assigned to the component by Chargify. This ID can be used to fetch the component from the API.
    Id                        *int                            `json:"id,omitempty"`
    // The name of the Component, suitable for display on statements. e.g., Text Messages.
    Name                      *string                         `json:"name,omitempty"`
    // The component API handle
    Handle                    Optional[string]                `json:"handle"`
    PricingScheme             Optional[PricingScheme]         `json:"pricing_scheme"`
    // The name of the unit that the component’s usage is measured in. e.g., message
    UnitName                  *string                         `json:"unit_name,omitempty"`
    // The amount the customer will be charged per unit. This field is only populated for ‘per_unit’ pricing schemes, otherwise it may be null.
    UnitPrice                 Optional[string]                `json:"unit_price"`
    // The id of the Product Family to which the Component belongs. `null` for a Wallet Product, which has no Product Family.
    ProductFamilyId           Optional[int]                   `json:"product_family_id"`
    // The name of the Product Family to which the Component belongs. `null` for a Wallet Product, which has no Product Family.
    ProductFamilyName         Optional[string]                `json:"product_family_name"`
    // The handle of the Product Family to which the Component belongs. `null` for a Wallet Product, which has no Product Family.
    ProductFamilyHandle       Optional[string]                `json:"product_family_handle"`
    // deprecated - use unit_price instead.
    PricePerUnitInCents       Optional[int64]                 `json:"price_per_unit_in_cents"`
    // A handle for the component type
    Kind                      *ComponentKind                  `json:"kind,omitempty"`
    // Boolean flag describing whether a component is archived or not.
    Archived                  *bool                           `json:"archived,omitempty"`
    // The description of the component.
    Description               Optional[string]                `json:"description"`
    DefaultPricePointId       Optional[int]                   `json:"default_price_point_id"`
    // Applicable only to prepaid usage components. An array of overage price brackets.
    OveragePrices             Optional[[]ComponentPrice]      `json:"overage_prices"`
    // An array of price brackets. If the component uses the ‘per_unit’ pricing scheme, this array will be empty.
    Prices                    Optional[[]ComponentPrice]      `json:"prices"`
    // Count for the number of price points associated with the component
    PricePointCount           *int                            `json:"price_point_count,omitempty"`
    // URL that points to the location to read the existing price points via GET request
    PricePointsUrl            Optional[string]                `json:"price_points_url"`
    DefaultPricePointName     *string                         `json:"default_price_point_name,omitempty"`
    // Boolean flag describing whether a component is taxable or not.
    Taxable                   *bool                           `json:"taxable,omitempty"`
    // A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters.
    TaxCode                   Optional[string]                `json:"tax_code"`
    Recurring                 *bool                           `json:"recurring,omitempty"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    UpgradeCharge             Optional[CreditType]            `json:"upgrade_charge"`
    // The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
    DowngradeCredit           Optional[CreditType]            `json:"downgrade_credit"`
    // Timestamp indicating when this component was created
    CreatedAt                 *time.Time                      `json:"created_at,omitempty"`
    // Timestamp indicating when this component was updated
    UpdatedAt                 *time.Time                      `json:"updated_at,omitempty"`
    // Timestamp indicating when this component was archived
    ArchivedAt                Optional[time.Time]             `json:"archived_at"`
    // (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices.
    HideDateRangeOnInvoice    *bool                           `json:"hide_date_range_on_invoice,omitempty"`
    AllowFractionalQuantities *bool                           `json:"allow_fractional_quantities,omitempty"`
    // One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
    ItemCategory              Optional[ItemCategory]          `json:"item_category"`
    UseSiteExchangeRate       Optional[bool]                  `json:"use_site_exchange_rate"`
    // E.g. Internal ID or SKU Number
    AccountingCode            Optional[string]                `json:"accounting_code"`
    // The ID of an attached metric from legacy Events Based Billing. Omitted if no metric is attached. If the metric is linked to a Maxio Metering meter, the response can include both event_based_billing_metric_id and metering_meter_id.
    EventBasedBillingMetricId *int                            `json:"event_based_billing_metric_id,omitempty"`
    // The ID of an attached single Maxio Metering meter. Omitted if no single meter is attached, including when the component uses only an Advanced Formula. If a legacy Events Based Billing metric is linked to the meter, the response can include both metering_meter_id and event_based_billing_metric_id.
    MeteringMeterId           *int                            `json:"metering_meter_id,omitempty"`
    // The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this component’s default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                  *int                            `json:"interval,omitempty"`
    // A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit              Optional[IntervalUnit]          `json:"interval_unit"`
    // (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this component instead of the default derived from item_category.
    UnspscCode                Optional[string]                `json:"unspsc_code"`
    // The active feature catalog items attached to this component. Present only when the request includes `include_features=true`.
    Features                  Optional[[]FeatureCatalogItem]  `json:"features"`
    // A rollup of what's using this Wallet Product, for the catalog listing.
    UsedBy                    *WalletProductUsageCount        `json:"used_by,omitempty"`
    // A summary of the Wallet Product's default price point, for rendering a catalog listing without a separate price points call.
    PricePointSummary         *WalletProductPricePointSummary `json:"price_point_summary,omitempty"`
    AdditionalProperties      map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for WalletProductListItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WalletProductListItem) String() string {
    return fmt.Sprintf(
    	"WalletProductListItem[Id=%v, Name=%v, Handle=%v, PricingScheme=%v, UnitName=%v, UnitPrice=%v, ProductFamilyId=%v, ProductFamilyName=%v, ProductFamilyHandle=%v, PricePerUnitInCents=%v, Kind=%v, Archived=%v, Description=%v, DefaultPricePointId=%v, OveragePrices=%v, Prices=%v, PricePointCount=%v, PricePointsUrl=%v, DefaultPricePointName=%v, Taxable=%v, TaxCode=%v, Recurring=%v, UpgradeCharge=%v, DowngradeCredit=%v, CreatedAt=%v, UpdatedAt=%v, ArchivedAt=%v, HideDateRangeOnInvoice=%v, AllowFractionalQuantities=%v, ItemCategory=%v, UseSiteExchangeRate=%v, AccountingCode=%v, EventBasedBillingMetricId=%v, MeteringMeterId=%v, Interval=%v, IntervalUnit=%v, UnspscCode=%v, Features=%v, UsedBy=%v, PricePointSummary=%v, AdditionalProperties=%v]",
    	w.Id, w.Name, w.Handle, w.PricingScheme, w.UnitName, w.UnitPrice, w.ProductFamilyId, w.ProductFamilyName, w.ProductFamilyHandle, w.PricePerUnitInCents, w.Kind, w.Archived, w.Description, w.DefaultPricePointId, w.OveragePrices, w.Prices, w.PricePointCount, w.PricePointsUrl, w.DefaultPricePointName, w.Taxable, w.TaxCode, w.Recurring, w.UpgradeCharge, w.DowngradeCredit, w.CreatedAt, w.UpdatedAt, w.ArchivedAt, w.HideDateRangeOnInvoice, w.AllowFractionalQuantities, w.ItemCategory, w.UseSiteExchangeRate, w.AccountingCode, w.EventBasedBillingMetricId, w.MeteringMeterId, w.Interval, w.IntervalUnit, w.UnspscCode, w.Features, w.UsedBy, w.PricePointSummary, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WalletProductListItem.
// It customizes the JSON marshaling process for WalletProductListItem objects.
func (w WalletProductListItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "name", "handle", "pricing_scheme", "unit_name", "unit_price", "product_family_id", "product_family_name", "product_family_handle", "price_per_unit_in_cents", "kind", "archived", "description", "default_price_point_id", "overage_prices", "prices", "price_point_count", "price_points_url", "default_price_point_name", "taxable", "tax_code", "recurring", "upgrade_charge", "downgrade_credit", "created_at", "updated_at", "archived_at", "hide_date_range_on_invoice", "allow_fractional_quantities", "item_category", "use_site_exchange_rate", "accounting_code", "event_based_billing_metric_id", "metering_meter_id", "interval", "interval_unit", "unspsc_code", "features", "used_by", "price_point_summary"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WalletProductListItem object to a map representation for JSON marshaling.
func (w WalletProductListItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    if w.Handle.IsValueSet() {
        if w.Handle.Value() != nil {
            structMap["handle"] = w.Handle.Value()
        } else {
            structMap["handle"] = nil
        }
    }
    if w.PricingScheme.IsValueSet() {
        if w.PricingScheme.Value() != nil {
            structMap["pricing_scheme"] = w.PricingScheme.Value()
        } else {
            structMap["pricing_scheme"] = nil
        }
    }
    if w.UnitName != nil {
        structMap["unit_name"] = w.UnitName
    }
    if w.UnitPrice.IsValueSet() {
        if w.UnitPrice.Value() != nil {
            structMap["unit_price"] = w.UnitPrice.Value()
        } else {
            structMap["unit_price"] = nil
        }
    }
    if w.ProductFamilyId.IsValueSet() {
        if w.ProductFamilyId.Value() != nil {
            structMap["product_family_id"] = w.ProductFamilyId.Value()
        } else {
            structMap["product_family_id"] = nil
        }
    }
    if w.ProductFamilyName.IsValueSet() {
        if w.ProductFamilyName.Value() != nil {
            structMap["product_family_name"] = w.ProductFamilyName.Value()
        } else {
            structMap["product_family_name"] = nil
        }
    }
    if w.ProductFamilyHandle.IsValueSet() {
        if w.ProductFamilyHandle.Value() != nil {
            structMap["product_family_handle"] = w.ProductFamilyHandle.Value()
        } else {
            structMap["product_family_handle"] = nil
        }
    }
    if w.PricePerUnitInCents.IsValueSet() {
        if w.PricePerUnitInCents.Value() != nil {
            structMap["price_per_unit_in_cents"] = w.PricePerUnitInCents.Value()
        } else {
            structMap["price_per_unit_in_cents"] = nil
        }
    }
    if w.Kind != nil {
        structMap["kind"] = w.Kind
    }
    if w.Archived != nil {
        structMap["archived"] = w.Archived
    }
    if w.Description.IsValueSet() {
        if w.Description.Value() != nil {
            structMap["description"] = w.Description.Value()
        } else {
            structMap["description"] = nil
        }
    }
    if w.DefaultPricePointId.IsValueSet() {
        if w.DefaultPricePointId.Value() != nil {
            structMap["default_price_point_id"] = w.DefaultPricePointId.Value()
        } else {
            structMap["default_price_point_id"] = nil
        }
    }
    if w.OveragePrices.IsValueSet() {
        if w.OveragePrices.Value() != nil {
            structMap["overage_prices"] = w.OveragePrices.Value()
        } else {
            structMap["overage_prices"] = nil
        }
    }
    if w.Prices.IsValueSet() {
        if w.Prices.Value() != nil {
            structMap["prices"] = w.Prices.Value()
        } else {
            structMap["prices"] = nil
        }
    }
    if w.PricePointCount != nil {
        structMap["price_point_count"] = w.PricePointCount
    }
    if w.PricePointsUrl.IsValueSet() {
        if w.PricePointsUrl.Value() != nil {
            structMap["price_points_url"] = w.PricePointsUrl.Value()
        } else {
            structMap["price_points_url"] = nil
        }
    }
    if w.DefaultPricePointName != nil {
        structMap["default_price_point_name"] = w.DefaultPricePointName
    }
    if w.Taxable != nil {
        structMap["taxable"] = w.Taxable
    }
    if w.TaxCode.IsValueSet() {
        if w.TaxCode.Value() != nil {
            structMap["tax_code"] = w.TaxCode.Value()
        } else {
            structMap["tax_code"] = nil
        }
    }
    if w.Recurring != nil {
        structMap["recurring"] = w.Recurring
    }
    if w.UpgradeCharge.IsValueSet() {
        if w.UpgradeCharge.Value() != nil {
            structMap["upgrade_charge"] = w.UpgradeCharge.Value()
        } else {
            structMap["upgrade_charge"] = nil
        }
    }
    if w.DowngradeCredit.IsValueSet() {
        if w.DowngradeCredit.Value() != nil {
            structMap["downgrade_credit"] = w.DowngradeCredit.Value()
        } else {
            structMap["downgrade_credit"] = nil
        }
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt.Format(time.RFC3339)
    }
    if w.UpdatedAt != nil {
        structMap["updated_at"] = w.UpdatedAt.Format(time.RFC3339)
    }
    if w.ArchivedAt.IsValueSet() {
        var ArchivedAtVal *string = nil
        if w.ArchivedAt.Value() != nil {
            val := w.ArchivedAt.Value().Format(time.RFC3339)
            ArchivedAtVal = &val
        }
        if w.ArchivedAt.Value() != nil {
            structMap["archived_at"] = ArchivedAtVal
        } else {
            structMap["archived_at"] = nil
        }
    }
    if w.HideDateRangeOnInvoice != nil {
        structMap["hide_date_range_on_invoice"] = w.HideDateRangeOnInvoice
    }
    if w.AllowFractionalQuantities != nil {
        structMap["allow_fractional_quantities"] = w.AllowFractionalQuantities
    }
    if w.ItemCategory.IsValueSet() {
        if w.ItemCategory.Value() != nil {
            structMap["item_category"] = w.ItemCategory.Value()
        } else {
            structMap["item_category"] = nil
        }
    }
    if w.UseSiteExchangeRate.IsValueSet() {
        if w.UseSiteExchangeRate.Value() != nil {
            structMap["use_site_exchange_rate"] = w.UseSiteExchangeRate.Value()
        } else {
            structMap["use_site_exchange_rate"] = nil
        }
    }
    if w.AccountingCode.IsValueSet() {
        if w.AccountingCode.Value() != nil {
            structMap["accounting_code"] = w.AccountingCode.Value()
        } else {
            structMap["accounting_code"] = nil
        }
    }
    if w.EventBasedBillingMetricId != nil {
        structMap["event_based_billing_metric_id"] = w.EventBasedBillingMetricId
    }
    if w.MeteringMeterId != nil {
        structMap["metering_meter_id"] = w.MeteringMeterId
    }
    if w.Interval != nil {
        structMap["interval"] = w.Interval
    }
    if w.IntervalUnit.IsValueSet() {
        if w.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = w.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    if w.UnspscCode.IsValueSet() {
        if w.UnspscCode.Value() != nil {
            structMap["unspsc_code"] = w.UnspscCode.Value()
        } else {
            structMap["unspsc_code"] = nil
        }
    }
    if w.Features.IsValueSet() {
        if w.Features.Value() != nil {
            structMap["features"] = w.Features.Value()
        } else {
            structMap["features"] = nil
        }
    }
    if w.UsedBy != nil {
        structMap["used_by"] = w.UsedBy.toMap()
    }
    if w.PricePointSummary != nil {
        structMap["price_point_summary"] = w.PricePointSummary.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WalletProductListItem.
// It customizes the JSON unmarshaling process for WalletProductListItem objects.
func (w *WalletProductListItem) UnmarshalJSON(input []byte) error {
    var temp tempWalletProductListItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "handle", "pricing_scheme", "unit_name", "unit_price", "product_family_id", "product_family_name", "product_family_handle", "price_per_unit_in_cents", "kind", "archived", "description", "default_price_point_id", "overage_prices", "prices", "price_point_count", "price_points_url", "default_price_point_name", "taxable", "tax_code", "recurring", "upgrade_charge", "downgrade_credit", "created_at", "updated_at", "archived_at", "hide_date_range_on_invoice", "allow_fractional_quantities", "item_category", "use_site_exchange_rate", "accounting_code", "event_based_billing_metric_id", "metering_meter_id", "interval", "interval_unit", "unspsc_code", "features", "used_by", "price_point_summary")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.Name = temp.Name
    w.Handle = temp.Handle
    w.PricingScheme = temp.PricingScheme
    w.UnitName = temp.UnitName
    w.UnitPrice = temp.UnitPrice
    w.ProductFamilyId = temp.ProductFamilyId
    w.ProductFamilyName = temp.ProductFamilyName
    w.ProductFamilyHandle = temp.ProductFamilyHandle
    w.PricePerUnitInCents = temp.PricePerUnitInCents
    w.Kind = temp.Kind
    w.Archived = temp.Archived
    w.Description = temp.Description
    w.DefaultPricePointId = temp.DefaultPricePointId
    w.OveragePrices = temp.OveragePrices
    w.Prices = temp.Prices
    w.PricePointCount = temp.PricePointCount
    w.PricePointsUrl = temp.PricePointsUrl
    w.DefaultPricePointName = temp.DefaultPricePointName
    w.Taxable = temp.Taxable
    w.TaxCode = temp.TaxCode
    w.Recurring = temp.Recurring
    w.UpgradeCharge = temp.UpgradeCharge
    w.DowngradeCredit = temp.DowngradeCredit
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        w.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        w.UpdatedAt = &UpdatedAtVal
    }
    w.ArchivedAt.ShouldSetValue(temp.ArchivedAt.IsValueSet())
    if temp.ArchivedAt.Value() != nil {
        ArchivedAtVal, err := time.Parse(time.RFC3339, (*temp.ArchivedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse archived_at as % s format.", time.RFC3339)
        }
        w.ArchivedAt.SetValue(&ArchivedAtVal)
    }
    w.HideDateRangeOnInvoice = temp.HideDateRangeOnInvoice
    w.AllowFractionalQuantities = temp.AllowFractionalQuantities
    w.ItemCategory = temp.ItemCategory
    w.UseSiteExchangeRate = temp.UseSiteExchangeRate
    w.AccountingCode = temp.AccountingCode
    w.EventBasedBillingMetricId = temp.EventBasedBillingMetricId
    w.MeteringMeterId = temp.MeteringMeterId
    w.Interval = temp.Interval
    w.IntervalUnit = temp.IntervalUnit
    w.UnspscCode = temp.UnspscCode
    w.Features = temp.Features
    w.UsedBy = temp.UsedBy
    w.PricePointSummary = temp.PricePointSummary
    return nil
}

// tempWalletProductListItem is a temporary struct used for validating the fields of WalletProductListItem.
type tempWalletProductListItem  struct {
    Id                        *int                            `json:"id,omitempty"`
    Name                      *string                         `json:"name,omitempty"`
    Handle                    Optional[string]                `json:"handle"`
    PricingScheme             Optional[PricingScheme]         `json:"pricing_scheme"`
    UnitName                  *string                         `json:"unit_name,omitempty"`
    UnitPrice                 Optional[string]                `json:"unit_price"`
    ProductFamilyId           Optional[int]                   `json:"product_family_id"`
    ProductFamilyName         Optional[string]                `json:"product_family_name"`
    ProductFamilyHandle       Optional[string]                `json:"product_family_handle"`
    PricePerUnitInCents       Optional[int64]                 `json:"price_per_unit_in_cents"`
    Kind                      *ComponentKind                  `json:"kind,omitempty"`
    Archived                  *bool                           `json:"archived,omitempty"`
    Description               Optional[string]                `json:"description"`
    DefaultPricePointId       Optional[int]                   `json:"default_price_point_id"`
    OveragePrices             Optional[[]ComponentPrice]      `json:"overage_prices"`
    Prices                    Optional[[]ComponentPrice]      `json:"prices"`
    PricePointCount           *int                            `json:"price_point_count,omitempty"`
    PricePointsUrl            Optional[string]                `json:"price_points_url"`
    DefaultPricePointName     *string                         `json:"default_price_point_name,omitempty"`
    Taxable                   *bool                           `json:"taxable,omitempty"`
    TaxCode                   Optional[string]                `json:"tax_code"`
    Recurring                 *bool                           `json:"recurring,omitempty"`
    UpgradeCharge             Optional[CreditType]            `json:"upgrade_charge"`
    DowngradeCredit           Optional[CreditType]            `json:"downgrade_credit"`
    CreatedAt                 *string                         `json:"created_at,omitempty"`
    UpdatedAt                 *string                         `json:"updated_at,omitempty"`
    ArchivedAt                Optional[string]                `json:"archived_at"`
    HideDateRangeOnInvoice    *bool                           `json:"hide_date_range_on_invoice,omitempty"`
    AllowFractionalQuantities *bool                           `json:"allow_fractional_quantities,omitempty"`
    ItemCategory              Optional[ItemCategory]          `json:"item_category"`
    UseSiteExchangeRate       Optional[bool]                  `json:"use_site_exchange_rate"`
    AccountingCode            Optional[string]                `json:"accounting_code"`
    EventBasedBillingMetricId *int                            `json:"event_based_billing_metric_id,omitempty"`
    MeteringMeterId           *int                            `json:"metering_meter_id,omitempty"`
    Interval                  *int                            `json:"interval,omitempty"`
    IntervalUnit              Optional[IntervalUnit]          `json:"interval_unit"`
    UnspscCode                Optional[string]                `json:"unspsc_code"`
    Features                  Optional[[]FeatureCatalogItem]  `json:"features"`
    UsedBy                    *WalletProductUsageCount        `json:"used_by,omitempty"`
    PricePointSummary         *WalletProductPricePointSummary `json:"price_point_summary,omitempty"`
}
