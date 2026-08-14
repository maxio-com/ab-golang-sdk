
# Update Invoice Item

A line item change for a draft ad hoc invoice. Supports the same attributes as line items on invoice creation, plus `uid` and `_destroy` for updating or removing existing line items.

## Structure

`UpdateInvoiceItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Title` | `*string` | Optional | - |
| `Quantity` | [`*models.UpdateInvoiceItemQuantity`](../../doc/models/containers/update-invoice-item-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitPrice` | [`*models.UpdateInvoiceItemUnitPrice`](../../doc/models/containers/update-invoice-item-unit-price.md) | Optional | This is a container for one-of cases. |
| `Taxable` | `*bool` | Optional | Set to true to automatically calculate taxes. Site must be configured to use and calculate taxes. If using AvaTax, a tax_code parameter must also be sent. |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the product type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `PeriodRangeStart` | `*string` | Optional | YYYY-MM-DD |
| `PeriodRangeEnd` | `*string` | Optional | YYYY-MM-DD |
| `ProductId` | [`*models.UpdateInvoiceItemProductId`](../../doc/models/containers/update-invoice-item-product-id.md) | Optional | This is a container for one-of cases. |
| `ComponentId` | [`*models.UpdateInvoiceItemComponentId`](../../doc/models/containers/update-invoice-item-component-id.md) | Optional | This is a container for one-of cases. |
| `PricePointId` | [`*models.UpdateInvoiceItemPricePointId`](../../doc/models/containers/update-invoice-item-price-point-id.md) | Optional | This is a container for one-of cases. |
| `ProductPricePointId` | [`*models.UpdateInvoiceItemProductPricePointId`](../../doc/models/containers/update-invoice-item-product-price-point-id.md) | Optional | This is a container for one-of cases. |
| `Description` | `*string` | Optional | **Constraints**: *Maximum Length*: `255` |
| `Uid` | `*string` | Optional | Unique identifier of an existing line item on the invoice. When provided, the matching line item is updated with the submitted attributes. When omitted, a new line item is added to the invoice. |
| `Destroy` | `*bool` | Optional | Set to `true` together with `uid` to remove the matching line item from the invoice. Line items not referenced in the request remain unchanged. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateInvoiceItem := models.UpdateInvoiceItem{
        Title:                models.ToPointer("title8"),
        Quantity:             models.ToPointer(models.UpdateInvoiceItemQuantityContainer.FromPrecision(float64(94.82))),
        UnitPrice:            models.ToPointer(models.UpdateInvoiceItemUnitPriceContainer.FromPrecision(float64(78.04))),
        Taxable:              models.ToPointer(false),
        TaxCode:              models.ToPointer("tax_code0"),
    }

}
```

