
# Create Invoice Item

## Structure

`CreateInvoiceItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Title` | `*string` | Optional | - |
| `Quantity` | [`*models.CreateInvoiceItemQuantity`](../../doc/models/containers/create-invoice-item-quantity.md) | Optional | This is a container for one-of cases. |
| `UnitPrice` | [`*models.CreateInvoiceItemUnitPrice`](../../doc/models/containers/create-invoice-item-unit-price.md) | Optional | This is a container for one-of cases. |
| `Taxable` | `*bool` | Optional | Set to true to automatically calculate taxes. Site must be configured to use and calculate taxes. If using AvaTax, a tax_code parameter must also be sent. |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the product type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `PeriodRangeStart` | `*string` | Optional | YYYY-MM-DD |
| `PeriodRangeEnd` | `*string` | Optional | YYYY-MM-DD |
| `ProductId` | [`*models.CreateInvoiceItemProductId`](../../doc/models/containers/create-invoice-item-product-id.md) | Optional | This is a container for one-of cases. |
| `ComponentId` | [`*models.CreateInvoiceItemComponentId`](../../doc/models/containers/create-invoice-item-component-id.md) | Optional | This is a container for one-of cases. |
| `PricePointId` | [`*models.CreateInvoiceItemPricePointId`](../../doc/models/containers/create-invoice-item-price-point-id.md) | Optional | This is a container for one-of cases. |
| `ProductPricePointId` | [`*models.CreateInvoiceItemProductPricePointId`](../../doc/models/containers/create-invoice-item-product-price-point-id.md) | Optional | This is a container for one-of cases. |
| `Description` | `*string` | Optional | **Constraints**: *Maximum Length*: `255` |

## Example (as JSON)

```json
{
  "title": "title2",
  "quantity": 154.86,
  "unit_price": 138.08,
  "taxable": false,
  "tax_code": "tax_code4"
}
```

