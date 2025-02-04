/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceLineItemEventData represents a InvoiceLineItemEventData struct.
type InvoiceLineItemEventData struct {
    Uid                   *string                        `json:"uid,omitempty"`
    Title                 *string                        `json:"title,omitempty"`
    Description           *string                        `json:"description,omitempty"`
    Quantity              *int                           `json:"quantity,omitempty"`
    QuantityDelta         Optional[int]                  `json:"quantity_delta"`
    UnitPrice             *string                        `json:"unit_price,omitempty"`
    PeriodRangeStart      *string                        `json:"period_range_start,omitempty"`
    PeriodRangeEnd        *string                        `json:"period_range_end,omitempty"`
    Amount                *string                        `json:"amount,omitempty"`
    LineReferences        *string                        `json:"line_references,omitempty"`
    PricingDetailsIndex   Optional[int]                  `json:"pricing_details_index"`
    PricingDetails        []InvoiceLineItemPricingDetail `json:"pricing_details,omitempty"`
    TaxCode               Optional[string]               `json:"tax_code"`
    TaxAmount             *string                        `json:"tax_amount,omitempty"`
    ProductId             *int                           `json:"product_id,omitempty"`
    ProductPricePointId   Optional[int]                  `json:"product_price_point_id"`
    PricePointId          Optional[int]                  `json:"price_point_id"`
    ComponentId           Optional[int]                  `json:"component_id"`
    BillingScheduleItemId Optional[int]                  `json:"billing_schedule_item_id"`
    CustomItem            Optional[bool]                 `json:"custom_item"`
    AdditionalProperties  map[string]interface{}         `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceLineItemEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceLineItemEventData) String() string {
    return fmt.Sprintf(
    	"InvoiceLineItemEventData[Uid=%v, Title=%v, Description=%v, Quantity=%v, QuantityDelta=%v, UnitPrice=%v, PeriodRangeStart=%v, PeriodRangeEnd=%v, Amount=%v, LineReferences=%v, PricingDetailsIndex=%v, PricingDetails=%v, TaxCode=%v, TaxAmount=%v, ProductId=%v, ProductPricePointId=%v, PricePointId=%v, ComponentId=%v, BillingScheduleItemId=%v, CustomItem=%v, AdditionalProperties=%v]",
    	i.Uid, i.Title, i.Description, i.Quantity, i.QuantityDelta, i.UnitPrice, i.PeriodRangeStart, i.PeriodRangeEnd, i.Amount, i.LineReferences, i.PricingDetailsIndex, i.PricingDetails, i.TaxCode, i.TaxAmount, i.ProductId, i.ProductPricePointId, i.PricePointId, i.ComponentId, i.BillingScheduleItemId, i.CustomItem, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemEventData.
// It customizes the JSON marshaling process for InvoiceLineItemEventData objects.
func (i InvoiceLineItemEventData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "title", "description", "quantity", "quantity_delta", "unit_price", "period_range_start", "period_range_end", "amount", "line_references", "pricing_details_index", "pricing_details", "tax_code", "tax_amount", "product_id", "product_price_point_id", "price_point_id", "component_id", "billing_schedule_item_id", "custom_item"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemEventData object to a map representation for JSON marshaling.
func (i InvoiceLineItemEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    if i.QuantityDelta.IsValueSet() {
        if i.QuantityDelta.Value() != nil {
            structMap["quantity_delta"] = i.QuantityDelta.Value()
        } else {
            structMap["quantity_delta"] = nil
        }
    }
    if i.UnitPrice != nil {
        structMap["unit_price"] = i.UnitPrice
    }
    if i.PeriodRangeStart != nil {
        structMap["period_range_start"] = i.PeriodRangeStart
    }
    if i.PeriodRangeEnd != nil {
        structMap["period_range_end"] = i.PeriodRangeEnd
    }
    if i.Amount != nil {
        structMap["amount"] = i.Amount
    }
    if i.LineReferences != nil {
        structMap["line_references"] = i.LineReferences
    }
    if i.PricingDetailsIndex.IsValueSet() {
        if i.PricingDetailsIndex.Value() != nil {
            structMap["pricing_details_index"] = i.PricingDetailsIndex.Value()
        } else {
            structMap["pricing_details_index"] = nil
        }
    }
    if i.PricingDetails != nil {
        structMap["pricing_details"] = i.PricingDetails
    }
    if i.TaxCode.IsValueSet() {
        if i.TaxCode.Value() != nil {
            structMap["tax_code"] = i.TaxCode.Value()
        } else {
            structMap["tax_code"] = nil
        }
    }
    if i.TaxAmount != nil {
        structMap["tax_amount"] = i.TaxAmount
    }
    if i.ProductId != nil {
        structMap["product_id"] = i.ProductId
    }
    if i.ProductPricePointId.IsValueSet() {
        if i.ProductPricePointId.Value() != nil {
            structMap["product_price_point_id"] = i.ProductPricePointId.Value()
        } else {
            structMap["product_price_point_id"] = nil
        }
    }
    if i.PricePointId.IsValueSet() {
        if i.PricePointId.Value() != nil {
            structMap["price_point_id"] = i.PricePointId.Value()
        } else {
            structMap["price_point_id"] = nil
        }
    }
    if i.ComponentId.IsValueSet() {
        if i.ComponentId.Value() != nil {
            structMap["component_id"] = i.ComponentId.Value()
        } else {
            structMap["component_id"] = nil
        }
    }
    if i.BillingScheduleItemId.IsValueSet() {
        if i.BillingScheduleItemId.Value() != nil {
            structMap["billing_schedule_item_id"] = i.BillingScheduleItemId.Value()
        } else {
            structMap["billing_schedule_item_id"] = nil
        }
    }
    if i.CustomItem.IsValueSet() {
        if i.CustomItem.Value() != nil {
            structMap["custom_item"] = i.CustomItem.Value()
        } else {
            structMap["custom_item"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItemEventData.
// It customizes the JSON unmarshaling process for InvoiceLineItemEventData objects.
func (i *InvoiceLineItemEventData) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceLineItemEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "title", "description", "quantity", "quantity_delta", "unit_price", "period_range_start", "period_range_end", "amount", "line_references", "pricing_details_index", "pricing_details", "tax_code", "tax_amount", "product_id", "product_price_point_id", "price_point_id", "component_id", "billing_schedule_item_id", "custom_item")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.Title = temp.Title
    i.Description = temp.Description
    i.Quantity = temp.Quantity
    i.QuantityDelta = temp.QuantityDelta
    i.UnitPrice = temp.UnitPrice
    i.PeriodRangeStart = temp.PeriodRangeStart
    i.PeriodRangeEnd = temp.PeriodRangeEnd
    i.Amount = temp.Amount
    i.LineReferences = temp.LineReferences
    i.PricingDetailsIndex = temp.PricingDetailsIndex
    i.PricingDetails = temp.PricingDetails
    i.TaxCode = temp.TaxCode
    i.TaxAmount = temp.TaxAmount
    i.ProductId = temp.ProductId
    i.ProductPricePointId = temp.ProductPricePointId
    i.PricePointId = temp.PricePointId
    i.ComponentId = temp.ComponentId
    i.BillingScheduleItemId = temp.BillingScheduleItemId
    i.CustomItem = temp.CustomItem
    return nil
}

// tempInvoiceLineItemEventData is a temporary struct used for validating the fields of InvoiceLineItemEventData.
type tempInvoiceLineItemEventData  struct {
    Uid                   *string                        `json:"uid,omitempty"`
    Title                 *string                        `json:"title,omitempty"`
    Description           *string                        `json:"description,omitempty"`
    Quantity              *int                           `json:"quantity,omitempty"`
    QuantityDelta         Optional[int]                  `json:"quantity_delta"`
    UnitPrice             *string                        `json:"unit_price,omitempty"`
    PeriodRangeStart      *string                        `json:"period_range_start,omitempty"`
    PeriodRangeEnd        *string                        `json:"period_range_end,omitempty"`
    Amount                *string                        `json:"amount,omitempty"`
    LineReferences        *string                        `json:"line_references,omitempty"`
    PricingDetailsIndex   Optional[int]                  `json:"pricing_details_index"`
    PricingDetails        []InvoiceLineItemPricingDetail `json:"pricing_details,omitempty"`
    TaxCode               Optional[string]               `json:"tax_code"`
    TaxAmount             *string                        `json:"tax_amount,omitempty"`
    ProductId             *int                           `json:"product_id,omitempty"`
    ProductPricePointId   Optional[int]                  `json:"product_price_point_id"`
    PricePointId          Optional[int]                  `json:"price_point_id"`
    ComponentId           Optional[int]                  `json:"component_id"`
    BillingScheduleItemId Optional[int]                  `json:"billing_schedule_item_id"`
    CustomItem            Optional[bool]                 `json:"custom_item"`
}
