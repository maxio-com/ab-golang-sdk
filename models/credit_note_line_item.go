package models

import (
    "encoding/json"
    "log"
    "time"
)

// CreditNoteLineItem represents a CreditNoteLineItem struct.
type CreditNoteLineItem struct {
    // Unique identifier for the line item.  Useful when cross-referencing the line against individual discounts in the `discounts` or `taxes` lists.
    Uid                   *string       `json:"uid,omitempty"`
    // A short descriptor for the credit given by this line.
    Title                 *string       `json:"title,omitempty"`
    // Detailed description for the credit given by this line.  May include proration details in plain text.
    // Note: this string may contain line breaks that are hints for the best display format on the credit note.
    Description           *string       `json:"description,omitempty"`
    // The quantity or count of units credited by the line item.
    // This is a decimal number represented as a string. (See "About Decimal Numbers".)
    Quantity              *string       `json:"quantity,omitempty"`
    // The price per unit for the line item.
    // When tiered pricing was used (i.e. not every unit was actually priced at the same price) this will be the blended average cost per unit and the `tiered_unit_price` field will be set to `true`.
    UnitPrice             *string       `json:"unit_price,omitempty"`
    // The line subtotal, generally calculated as `quantity * unit_price`. This is the canonical amount of record for the line - when rounding differences are in play, `subtotal_amount` takes precedence over the value derived from `quantity * unit_price` (which may not have the proper precision to exactly equal this amount).
    SubtotalAmount        *string       `json:"subtotal_amount,omitempty"`
    // The approximate discount of just this line.
    // The value is approximated in cases where rounding errors make it difficult to apportion exactly a total discount among many lines. Several lines may have been summed prior to applying the discount to arrive at `discount_amount` for the invoice - backing that out to the discount on a single line may introduce rounding or precision errors.
    DiscountAmount        *string       `json:"discount_amount,omitempty"`
    // The approximate tax of just this line.
    // The value is approximated in cases where rounding errors make it difficult to apportion exactly a total tax among many lines. Several lines may have been summed prior to applying the tax rate to arrive at `tax_amount` for the invoice - backing that out to the tax on a single line may introduce rounding or precision errors.
    TaxAmount             *string       `json:"tax_amount,omitempty"`
    // The non-canonical total amount for the line.
    // `subtotal_amount` is the canonical amount for a line. The invoice `total_amount` is derived from the sum of the line `subtotal_amount`s and discounts or taxes applied thereafter.  Therefore, due to rounding or precision errors, the sum of line `total_amount`s may not equal the invoice `total_amount`.
    TotalAmount           *string       `json:"total_amount,omitempty"`
    // When `true`, indicates that the actual pricing scheme for the line was tiered, so the `unit_price` shown is the blended average for all units.
    TieredUnitPrice       *bool         `json:"tiered_unit_price,omitempty"`
    // Start date for the period credited by this line. The format is `"YYYY-MM-DD"`.
    PeriodRangeStart      *time.Time    `json:"period_range_start,omitempty"`
    // End date for the period credited by this line. The format is `"YYYY-MM-DD"`.
    PeriodRangeEnd        *time.Time    `json:"period_range_end,omitempty"`
    // The ID of the product being credited.
    // This may be set even for component credits, so true product-only (non-component) credits will also have a nil `component_id`.
    ProductId             *int          `json:"product_id,omitempty"`
    // The version of the product being credited.
    ProductVersion        *int          `json:"product_version,omitempty"`
    // The ID of the component being credited. Will be `nil` for non-component credits.
    ComponentId           Optional[int] `json:"component_id"`
    // The price point ID of the component being credited. Will be `nil` for non-component credits.
    PricePointId          Optional[int] `json:"price_point_id"`
    BillingScheduleItemId Optional[int] `json:"billing_schedule_item_id"`
    CustomItem            *bool         `json:"custom_item,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreditNoteLineItem.
// It customizes the JSON marshaling process for CreditNoteLineItem objects.
func (c *CreditNoteLineItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreditNoteLineItem object to a map representation for JSON marshaling.
func (c *CreditNoteLineItem) toMap() map[string]any {
    structMap := make(map[string]any)
    if c.Uid != nil {
        structMap["uid"] = c.Uid
    }
    if c.Title != nil {
        structMap["title"] = c.Title
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    if c.Quantity != nil {
        structMap["quantity"] = c.Quantity
    }
    if c.UnitPrice != nil {
        structMap["unit_price"] = c.UnitPrice
    }
    if c.SubtotalAmount != nil {
        structMap["subtotal_amount"] = c.SubtotalAmount
    }
    if c.DiscountAmount != nil {
        structMap["discount_amount"] = c.DiscountAmount
    }
    if c.TaxAmount != nil {
        structMap["tax_amount"] = c.TaxAmount
    }
    if c.TotalAmount != nil {
        structMap["total_amount"] = c.TotalAmount
    }
    if c.TieredUnitPrice != nil {
        structMap["tiered_unit_price"] = c.TieredUnitPrice
    }
    if c.PeriodRangeStart != nil {
        structMap["period_range_start"] = c.PeriodRangeStart.Format(DEFAULT_DATE)
    }
    if c.PeriodRangeEnd != nil {
        structMap["period_range_end"] = c.PeriodRangeEnd.Format(DEFAULT_DATE)
    }
    if c.ProductId != nil {
        structMap["product_id"] = c.ProductId
    }
    if c.ProductVersion != nil {
        structMap["product_version"] = c.ProductVersion
    }
    if c.ComponentId.IsValueSet() {
        structMap["component_id"] = c.ComponentId.Value()
    }
    if c.PricePointId.IsValueSet() {
        structMap["price_point_id"] = c.PricePointId.Value()
    }
    if c.BillingScheduleItemId.IsValueSet() {
        structMap["billing_schedule_item_id"] = c.BillingScheduleItemId.Value()
    }
    if c.CustomItem != nil {
        structMap["custom_item"] = c.CustomItem
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditNoteLineItem.
// It customizes the JSON unmarshaling process for CreditNoteLineItem objects.
func (c *CreditNoteLineItem) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid                   *string       `json:"uid,omitempty"`
        Title                 *string       `json:"title,omitempty"`
        Description           *string       `json:"description,omitempty"`
        Quantity              *string       `json:"quantity,omitempty"`
        UnitPrice             *string       `json:"unit_price,omitempty"`
        SubtotalAmount        *string       `json:"subtotal_amount,omitempty"`
        DiscountAmount        *string       `json:"discount_amount,omitempty"`
        TaxAmount             *string       `json:"tax_amount,omitempty"`
        TotalAmount           *string       `json:"total_amount,omitempty"`
        TieredUnitPrice       *bool         `json:"tiered_unit_price,omitempty"`
        PeriodRangeStart      *string       `json:"period_range_start,omitempty"`
        PeriodRangeEnd        *string       `json:"period_range_end,omitempty"`
        ProductId             *int          `json:"product_id,omitempty"`
        ProductVersion        *int          `json:"product_version,omitempty"`
        ComponentId           Optional[int] `json:"component_id"`
        PricePointId          Optional[int] `json:"price_point_id"`
        BillingScheduleItemId Optional[int] `json:"billing_schedule_item_id"`
        CustomItem            *bool         `json:"custom_item,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Uid = temp.Uid
    c.Title = temp.Title
    c.Description = temp.Description
    c.Quantity = temp.Quantity
    c.UnitPrice = temp.UnitPrice
    c.SubtotalAmount = temp.SubtotalAmount
    c.DiscountAmount = temp.DiscountAmount
    c.TaxAmount = temp.TaxAmount
    c.TotalAmount = temp.TotalAmount
    c.TieredUnitPrice = temp.TieredUnitPrice
    if temp.PeriodRangeStart != nil {
        PeriodRangeStartVal, err := time.Parse(DEFAULT_DATE, *temp.PeriodRangeStart)
        if err != nil {
            log.Fatalf("Cannot Parse period_range_start as % s format.", DEFAULT_DATE)
        }
        c.PeriodRangeStart = &PeriodRangeStartVal
    }
    if temp.PeriodRangeEnd != nil {
        PeriodRangeEndVal, err := time.Parse(DEFAULT_DATE, *temp.PeriodRangeEnd)
        if err != nil {
            log.Fatalf("Cannot Parse period_range_end as % s format.", DEFAULT_DATE)
        }
        c.PeriodRangeEnd = &PeriodRangeEndVal
    }
    c.ProductId = temp.ProductId
    c.ProductVersion = temp.ProductVersion
    c.ComponentId = temp.ComponentId
    c.PricePointId = temp.PricePointId
    c.BillingScheduleItemId = temp.BillingScheduleItemId
    c.CustomItem = temp.CustomItem
    return nil
}
