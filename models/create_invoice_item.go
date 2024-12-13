/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CreateInvoiceItem represents a CreateInvoiceItem struct.
type CreateInvoiceItem struct {
    Title                *string                               `json:"title,omitempty"`
    // The quantity can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place.
    Quantity             *CreateInvoiceItemQuantity            `json:"quantity,omitempty"`
    // The unit_price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place.
    UnitPrice            *CreateInvoiceItemUnitPrice           `json:"unit_price,omitempty"`
    // Set to true to automatically calculate taxes. Site must be configured to use and calculate taxes.
    // If using Avalara, a tax_code parameter must also be sent.
    Taxable              *bool                                 `json:"taxable,omitempty"`
    TaxCode              *string                               `json:"tax_code,omitempty"`
    // YYYY-MM-DD
    PeriodRangeStart     *string                               `json:"period_range_start,omitempty"`
    // YYYY-MM-DD
    PeriodRangeEnd       *string                               `json:"period_range_end,omitempty"`
    // Product handle or product id.
    ProductId            *CreateInvoiceItemProductId           `json:"product_id,omitempty"`
    // Component handle or component id.
    ComponentId          *CreateInvoiceItemComponentId         `json:"component_id,omitempty"`
    // Price point handle or id. For component.
    PricePointId         *CreateInvoiceItemPricePointId        `json:"price_point_id,omitempty"`
    ProductPricePointId  *CreateInvoiceItemProductPricePointId `json:"product_price_point_id,omitempty"`
    Description          *string                               `json:"description,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceItem.
// It customizes the JSON marshaling process for CreateInvoiceItem objects.
func (c CreateInvoiceItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "title", "quantity", "unit_price", "taxable", "tax_code", "period_range_start", "period_range_end", "product_id", "component_id", "price_point_id", "product_price_point_id", "description"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceItem object to a map representation for JSON marshaling.
func (c CreateInvoiceItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Title != nil {
        structMap["title"] = c.Title
    }
    if c.Quantity != nil {
        structMap["quantity"] = c.Quantity.toMap()
    }
    if c.UnitPrice != nil {
        structMap["unit_price"] = c.UnitPrice.toMap()
    }
    if c.Taxable != nil {
        structMap["taxable"] = c.Taxable
    }
    if c.TaxCode != nil {
        structMap["tax_code"] = c.TaxCode
    }
    if c.PeriodRangeStart != nil {
        structMap["period_range_start"] = c.PeriodRangeStart
    }
    if c.PeriodRangeEnd != nil {
        structMap["period_range_end"] = c.PeriodRangeEnd
    }
    if c.ProductId != nil {
        structMap["product_id"] = c.ProductId.toMap()
    }
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId.toMap()
    }
    if c.PricePointId != nil {
        structMap["price_point_id"] = c.PricePointId.toMap()
    }
    if c.ProductPricePointId != nil {
        structMap["product_price_point_id"] = c.ProductPricePointId.toMap()
    }
    if c.Description != nil {
        structMap["description"] = c.Description
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceItem.
// It customizes the JSON unmarshaling process for CreateInvoiceItem objects.
func (c *CreateInvoiceItem) UnmarshalJSON(input []byte) error {
    var temp tempCreateInvoiceItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "title", "quantity", "unit_price", "taxable", "tax_code", "period_range_start", "period_range_end", "product_id", "component_id", "price_point_id", "product_price_point_id", "description")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Title = temp.Title
    c.Quantity = temp.Quantity
    c.UnitPrice = temp.UnitPrice
    c.Taxable = temp.Taxable
    c.TaxCode = temp.TaxCode
    c.PeriodRangeStart = temp.PeriodRangeStart
    c.PeriodRangeEnd = temp.PeriodRangeEnd
    c.ProductId = temp.ProductId
    c.ComponentId = temp.ComponentId
    c.PricePointId = temp.PricePointId
    c.ProductPricePointId = temp.ProductPricePointId
    c.Description = temp.Description
    return nil
}

// tempCreateInvoiceItem is a temporary struct used for validating the fields of CreateInvoiceItem.
type tempCreateInvoiceItem  struct {
    Title               *string                               `json:"title,omitempty"`
    Quantity            *CreateInvoiceItemQuantity            `json:"quantity,omitempty"`
    UnitPrice           *CreateInvoiceItemUnitPrice           `json:"unit_price,omitempty"`
    Taxable             *bool                                 `json:"taxable,omitempty"`
    TaxCode             *string                               `json:"tax_code,omitempty"`
    PeriodRangeStart    *string                               `json:"period_range_start,omitempty"`
    PeriodRangeEnd      *string                               `json:"period_range_end,omitempty"`
    ProductId           *CreateInvoiceItemProductId           `json:"product_id,omitempty"`
    ComponentId         *CreateInvoiceItemComponentId         `json:"component_id,omitempty"`
    PricePointId        *CreateInvoiceItemPricePointId        `json:"price_point_id,omitempty"`
    ProductPricePointId *CreateInvoiceItemProductPricePointId `json:"product_price_point_id,omitempty"`
    Description         *string                               `json:"description,omitempty"`
}
