package models

import (
    "encoding/json"
    "log"
    "time"
)

// InvoiceLineItem represents a InvoiceLineItem struct.
type InvoiceLineItem struct {
    // Unique identifier for the line item.  Useful when cross-referencing the line against individual discounts in the `discounts` or `taxes` lists.
    Uid                  *string                                    `json:"uid,omitempty"`
    // A short descriptor for the charge or item represented by this line.
    Title                *string                                    `json:"title,omitempty"`
    // Detailed description for the charge or item represented by this line.  May include proration details in plain text.
    // Note: this string may contain line breaks that are hints for the best display format on the invoice.
    Description          *string                                    `json:"description,omitempty"`
    // The quantity or count of units billed by the line item.
    // This is a decimal number represented as a string. (See "About Decimal Numbers".)
    Quantity             *string                                    `json:"quantity,omitempty"`
    // The price per unit for the line item.
    // When tiered pricing was used (i.e. not every unit was actually priced at the same price) this will be the blended average cost per unit and the `tiered_unit_price` field will be set to `true`.
    UnitPrice            *string                                    `json:"unit_price,omitempty"`
    // The line subtotal, generally calculated as `quantity * unit_price`. This is the canonical amount of record for the line - when rounding differences are in play, `subtotal_amount` takes precedence over the value derived from `quantity * unit_price` (which may not have the proper precision to exactly equal this amount).
    SubtotalAmount       *string                                    `json:"subtotal_amount,omitempty"`
    // The approximate discount applied to just this line.
    // The value is approximated in cases where rounding errors make it difficult to apportion exactly a total discount among many lines. Several lines may have been summed prior to applying the discount to arrive at `discount_amount` for the invoice - backing that out to the discount on a single line may introduce rounding or precision errors.
    DiscountAmount       *string                                    `json:"discount_amount,omitempty"`
    // The approximate tax applied to just this line.
    // The value is approximated in cases where rounding errors make it difficult to apportion exactly a total tax among many lines. Several lines may have been summed prior to applying the tax rate to arrive at `tax_amount` for the invoice - backing that out to the tax on a single line may introduce rounding or precision errors.
    TaxAmount            *string                                    `json:"tax_amount,omitempty"`
    // The non-canonical total amount for the line.
    // `subtotal_amount` is the canonical amount for a line. The invoice `total_amount` is derived from the sum of the line `subtotal_amount`s and discounts or taxes applied thereafter.  Therefore, due to rounding or precision errors, the sum of line `total_amount`s may not equal the invoice `total_amount`.
    TotalAmount          *string                                    `json:"total_amount,omitempty"`
    // When `true`, indicates that the actual pricing scheme for the line was tiered, so the `unit_price` shown is the blended average for all units.
    TieredUnitPrice      *bool                                      `json:"tiered_unit_price,omitempty"`
    // Start date for the period covered by this line. The format is `"YYYY-MM-DD"`.
    // * For periodic charges paid in advance, this date will match the billing date, and the end date will be in the future.
    // * For periodic charges paid in arrears (e.g. metered charges), this date will be the date of the previous billing, and the end date will be the current billing date.
    // * For non-periodic charges, this date and the end date will match.
    PeriodRangeStart     *time.Time                                 `json:"period_range_start,omitempty"`
    // End date for the period covered by this line. The format is `"YYYY-MM-DD"`.
    // * For periodic charges paid in advance, this date will match the next (future) billing date.
    // * For periodic charges paid in arrears (e.g. metered charges), this date will be the date of the current billing date.
    // * For non-periodic charges, this date and the start date will match.
    PeriodRangeEnd       *time.Time                                 `json:"period_range_end,omitempty"`
    TransactionId        *int                                       `json:"transaction_id,omitempty"`
    // The ID of the product subscribed when the charge was made.
    // This may be set even for component charges, so true product-only (non-component) charges will also have a nil `component_id`.
    ProductId            Optional[int]                              `json:"product_id"`
    // The version of the product subscribed when the charge was made.
    ProductVersion       Optional[int]                              `json:"product_version"`
    // The ID of the component being billed. Will be `nil` for non-component charges.
    ComponentId          Optional[int]                              `json:"component_id"`
    // The price point ID of the component being billed. Will be `nil` for non-component charges.
    PricePointId         Optional[int]                              `json:"price_point_id"`
    Hide                 *bool                                      `json:"hide,omitempty"`
    ComponentCostData    Optional[InvoiceLineItemComponentCostData] `json:"component_cost_data"`
    // The price point ID of the line item's product
    ProductPricePointId  Optional[int]                              `json:"product_price_point_id"`
    CustomItem           *bool                                      `json:"custom_item,omitempty"`
    Kind                 *string                                    `json:"kind,omitempty"`
    AdditionalProperties map[string]any                             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItem.
// It customizes the JSON marshaling process for InvoiceLineItem objects.
func (i InvoiceLineItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItem object to a map representation for JSON marshaling.
func (i InvoiceLineItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.Title != nil {
        structMap["title"] = i.Title
    }
    if i.Description != nil {
        structMap["description"] = i.Description
    }
    if i.Quantity != nil {
        structMap["quantity"] = i.Quantity
    }
    if i.UnitPrice != nil {
        structMap["unit_price"] = i.UnitPrice
    }
    if i.SubtotalAmount != nil {
        structMap["subtotal_amount"] = i.SubtotalAmount
    }
    if i.DiscountAmount != nil {
        structMap["discount_amount"] = i.DiscountAmount
    }
    if i.TaxAmount != nil {
        structMap["tax_amount"] = i.TaxAmount
    }
    if i.TotalAmount != nil {
        structMap["total_amount"] = i.TotalAmount
    }
    if i.TieredUnitPrice != nil {
        structMap["tiered_unit_price"] = i.TieredUnitPrice
    }
    if i.PeriodRangeStart != nil {
        structMap["period_range_start"] = i.PeriodRangeStart.Format(DEFAULT_DATE)
    }
    if i.PeriodRangeEnd != nil {
        structMap["period_range_end"] = i.PeriodRangeEnd.Format(DEFAULT_DATE)
    }
    if i.TransactionId != nil {
        structMap["transaction_id"] = i.TransactionId
    }
    if i.ProductId.IsValueSet() {
        if i.ProductId.Value() != nil {
            structMap["product_id"] = i.ProductId.Value()
        } else {
            structMap["product_id"] = nil
        }
    }
    if i.ProductVersion.IsValueSet() {
        if i.ProductVersion.Value() != nil {
            structMap["product_version"] = i.ProductVersion.Value()
        } else {
            structMap["product_version"] = nil
        }
    }
    if i.ComponentId.IsValueSet() {
        if i.ComponentId.Value() != nil {
            structMap["component_id"] = i.ComponentId.Value()
        } else {
            structMap["component_id"] = nil
        }
    }
    if i.PricePointId.IsValueSet() {
        if i.PricePointId.Value() != nil {
            structMap["price_point_id"] = i.PricePointId.Value()
        } else {
            structMap["price_point_id"] = nil
        }
    }
    if i.Hide != nil {
        structMap["hide"] = i.Hide
    }
    if i.ComponentCostData.IsValueSet() {
        if i.ComponentCostData.Value() != nil {
            structMap["component_cost_data"] = i.ComponentCostData.Value().toMap()
        } else {
            structMap["component_cost_data"] = nil
        }
    }
    if i.ProductPricePointId.IsValueSet() {
        if i.ProductPricePointId.Value() != nil {
            structMap["product_price_point_id"] = i.ProductPricePointId.Value()
        } else {
            structMap["product_price_point_id"] = nil
        }
    }
    if i.CustomItem != nil {
        structMap["custom_item"] = i.CustomItem
    }
    if i.Kind != nil {
        structMap["kind"] = i.Kind
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItem.
// It customizes the JSON unmarshaling process for InvoiceLineItem objects.
func (i *InvoiceLineItem) UnmarshalJSON(input []byte) error {
    var temp invoiceLineItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "title", "description", "quantity", "unit_price", "subtotal_amount", "discount_amount", "tax_amount", "total_amount", "tiered_unit_price", "period_range_start", "period_range_end", "transaction_id", "product_id", "product_version", "component_id", "price_point_id", "hide", "component_cost_data", "product_price_point_id", "custom_item", "kind")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Uid = temp.Uid
    i.Title = temp.Title
    i.Description = temp.Description
    i.Quantity = temp.Quantity
    i.UnitPrice = temp.UnitPrice
    i.SubtotalAmount = temp.SubtotalAmount
    i.DiscountAmount = temp.DiscountAmount
    i.TaxAmount = temp.TaxAmount
    i.TotalAmount = temp.TotalAmount
    i.TieredUnitPrice = temp.TieredUnitPrice
    if temp.PeriodRangeStart != nil {
        PeriodRangeStartVal, err := time.Parse(DEFAULT_DATE, *temp.PeriodRangeStart)
        if err != nil {
            log.Fatalf("Cannot Parse period_range_start as % s format.", DEFAULT_DATE)
        }
        i.PeriodRangeStart = &PeriodRangeStartVal
    }
    if temp.PeriodRangeEnd != nil {
        PeriodRangeEndVal, err := time.Parse(DEFAULT_DATE, *temp.PeriodRangeEnd)
        if err != nil {
            log.Fatalf("Cannot Parse period_range_end as % s format.", DEFAULT_DATE)
        }
        i.PeriodRangeEnd = &PeriodRangeEndVal
    }
    i.TransactionId = temp.TransactionId
    i.ProductId = temp.ProductId
    i.ProductVersion = temp.ProductVersion
    i.ComponentId = temp.ComponentId
    i.PricePointId = temp.PricePointId
    i.Hide = temp.Hide
    i.ComponentCostData = temp.ComponentCostData
    i.ProductPricePointId = temp.ProductPricePointId
    i.CustomItem = temp.CustomItem
    i.Kind = temp.Kind
    return nil
}

// TODO
type invoiceLineItem  struct {
    Uid                 *string                                    `json:"uid,omitempty"`
    Title               *string                                    `json:"title,omitempty"`
    Description         *string                                    `json:"description,omitempty"`
    Quantity            *string                                    `json:"quantity,omitempty"`
    UnitPrice           *string                                    `json:"unit_price,omitempty"`
    SubtotalAmount      *string                                    `json:"subtotal_amount,omitempty"`
    DiscountAmount      *string                                    `json:"discount_amount,omitempty"`
    TaxAmount           *string                                    `json:"tax_amount,omitempty"`
    TotalAmount         *string                                    `json:"total_amount,omitempty"`
    TieredUnitPrice     *bool                                      `json:"tiered_unit_price,omitempty"`
    PeriodRangeStart    *string                                    `json:"period_range_start,omitempty"`
    PeriodRangeEnd      *string                                    `json:"period_range_end,omitempty"`
    TransactionId       *int                                       `json:"transaction_id,omitempty"`
    ProductId           Optional[int]                              `json:"product_id"`
    ProductVersion      Optional[int]                              `json:"product_version"`
    ComponentId         Optional[int]                              `json:"component_id"`
    PricePointId        Optional[int]                              `json:"price_point_id"`
    Hide                *bool                                      `json:"hide,omitempty"`
    ComponentCostData   Optional[InvoiceLineItemComponentCostData] `json:"component_cost_data"`
    ProductPricePointId Optional[int]                              `json:"product_price_point_id"`
    CustomItem          *bool                                      `json:"custom_item,omitempty"`
    Kind                *string                                    `json:"kind,omitempty"`
}
