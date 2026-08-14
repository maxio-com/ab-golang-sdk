// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// UpdateInvoiceItem represents a UpdateInvoiceItem struct.
// A line item change for a draft ad hoc invoice. Supports the same attributes as line items on invoice creation, plus `uid` and `_destroy` for updating or removing existing line items.
type UpdateInvoiceItem struct {
    Title                *string                               `json:"title,omitempty"`
    // The quantity can contain up to 8 decimal places. e.g., 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place.
    Quantity             *UpdateInvoiceItemQuantity            `json:"quantity,omitempty"`
    // The unit_price can contain up to 8 decimal places. e.g., 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place.
    UnitPrice            *UpdateInvoiceItemUnitPrice           `json:"unit_price,omitempty"`
    // Set to true to automatically calculate taxes. Site must be configured to use and calculate taxes. If using AvaTax, a tax_code parameter must also be sent.
    Taxable              *bool                                 `json:"taxable,omitempty"`
    // A string representing the tax code related to the product type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters.
    TaxCode              *string                               `json:"tax_code,omitempty"`
    // YYYY-MM-DD
    PeriodRangeStart     *string                               `json:"period_range_start,omitempty"`
    // YYYY-MM-DD
    PeriodRangeEnd       *string                               `json:"period_range_end,omitempty"`
    // Product handle or product id.
    ProductId            *UpdateInvoiceItemProductId           `json:"product_id,omitempty"`
    // Component handle or component id.
    ComponentId          *UpdateInvoiceItemComponentId         `json:"component_id,omitempty"`
    // Price point handle or id. For component.
    PricePointId         *UpdateInvoiceItemPricePointId        `json:"price_point_id,omitempty"`
    ProductPricePointId  *UpdateInvoiceItemProductPricePointId `json:"product_price_point_id,omitempty"`
    Description          *string                               `json:"description,omitempty"`
    // Unique identifier of an existing line item on the invoice. When provided, the matching line item is updated with the submitted attributes. When omitted, a new line item is added to the invoice.
    Uid                  *string                               `json:"uid,omitempty"`
    // Set to `true` together with `uid` to remove the matching line item from the invoice. Line items not referenced in the request remain unchanged.
    Destroy              *bool                                 `json:"_destroy,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateInvoiceItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceItem) String() string {
    return fmt.Sprintf(
    	"UpdateInvoiceItem[Title=%v, Quantity=%v, UnitPrice=%v, Taxable=%v, TaxCode=%v, PeriodRangeStart=%v, PeriodRangeEnd=%v, ProductId=%v, ComponentId=%v, PricePointId=%v, ProductPricePointId=%v, Description=%v, Uid=%v, Destroy=%v, AdditionalProperties=%v]",
    	u.Title, u.Quantity, u.UnitPrice, u.Taxable, u.TaxCode, u.PeriodRangeStart, u.PeriodRangeEnd, u.ProductId, u.ComponentId, u.PricePointId, u.ProductPricePointId, u.Description, u.Uid, u.Destroy, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceItem.
// It customizes the JSON marshaling process for UpdateInvoiceItem objects.
func (u UpdateInvoiceItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "title", "quantity", "unit_price", "taxable", "tax_code", "period_range_start", "period_range_end", "product_id", "component_id", "price_point_id", "product_price_point_id", "description", "uid", "_destroy"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceItem object to a map representation for JSON marshaling.
func (u UpdateInvoiceItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Title != nil {
        structMap["title"] = u.Title
    }
    if u.Quantity != nil {
        structMap["quantity"] = u.Quantity.toMap()
    }
    if u.UnitPrice != nil {
        structMap["unit_price"] = u.UnitPrice.toMap()
    }
    if u.Taxable != nil {
        structMap["taxable"] = u.Taxable
    }
    if u.TaxCode != nil {
        structMap["tax_code"] = u.TaxCode
    }
    if u.PeriodRangeStart != nil {
        structMap["period_range_start"] = u.PeriodRangeStart
    }
    if u.PeriodRangeEnd != nil {
        structMap["period_range_end"] = u.PeriodRangeEnd
    }
    if u.ProductId != nil {
        structMap["product_id"] = u.ProductId.toMap()
    }
    if u.ComponentId != nil {
        structMap["component_id"] = u.ComponentId.toMap()
    }
    if u.PricePointId != nil {
        structMap["price_point_id"] = u.PricePointId.toMap()
    }
    if u.ProductPricePointId != nil {
        structMap["product_price_point_id"] = u.ProductPricePointId.toMap()
    }
    if u.Description != nil {
        structMap["description"] = u.Description
    }
    if u.Uid != nil {
        structMap["uid"] = u.Uid
    }
    if u.Destroy != nil {
        structMap["_destroy"] = u.Destroy
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceItem.
// It customizes the JSON unmarshaling process for UpdateInvoiceItem objects.
func (u *UpdateInvoiceItem) UnmarshalJSON(input []byte) error {
    var temp tempUpdateInvoiceItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "title", "quantity", "unit_price", "taxable", "tax_code", "period_range_start", "period_range_end", "product_id", "component_id", "price_point_id", "product_price_point_id", "description", "uid", "_destroy")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Title = temp.Title
    u.Quantity = temp.Quantity
    u.UnitPrice = temp.UnitPrice
    u.Taxable = temp.Taxable
    u.TaxCode = temp.TaxCode
    u.PeriodRangeStart = temp.PeriodRangeStart
    u.PeriodRangeEnd = temp.PeriodRangeEnd
    u.ProductId = temp.ProductId
    u.ComponentId = temp.ComponentId
    u.PricePointId = temp.PricePointId
    u.ProductPricePointId = temp.ProductPricePointId
    u.Description = temp.Description
    u.Uid = temp.Uid
    u.Destroy = temp.Destroy
    return nil
}

// tempUpdateInvoiceItem is a temporary struct used for validating the fields of UpdateInvoiceItem.
type tempUpdateInvoiceItem  struct {
    Title               *string                               `json:"title,omitempty"`
    Quantity            *UpdateInvoiceItemQuantity            `json:"quantity,omitempty"`
    UnitPrice           *UpdateInvoiceItemUnitPrice           `json:"unit_price,omitempty"`
    Taxable             *bool                                 `json:"taxable,omitempty"`
    TaxCode             *string                               `json:"tax_code,omitempty"`
    PeriodRangeStart    *string                               `json:"period_range_start,omitempty"`
    PeriodRangeEnd      *string                               `json:"period_range_end,omitempty"`
    ProductId           *UpdateInvoiceItemProductId           `json:"product_id,omitempty"`
    ComponentId         *UpdateInvoiceItemComponentId         `json:"component_id,omitempty"`
    PricePointId        *UpdateInvoiceItemPricePointId        `json:"price_point_id,omitempty"`
    ProductPricePointId *UpdateInvoiceItemProductPricePointId `json:"product_price_point_id,omitempty"`
    Description         *string                               `json:"description,omitempty"`
    Uid                 *string                               `json:"uid,omitempty"`
    Destroy             *bool                                 `json:"_destroy,omitempty"`
}
