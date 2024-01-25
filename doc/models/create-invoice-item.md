
# Create Invoice Item

## Structure

`CreateInvoiceItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Title` | `*string` | Optional | - |
| `Quantity` | `*interface{}` | Optional | The quantity can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place. |
| `UnitPrice` | `*interface{}` | Optional | The unit_price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065. If you submit a value with more than 8 decimal places, we will round it down to the 8th decimal place. |
| `Taxable` | `*bool` | Optional | Set to true to automatically calculate taxes. Site must be configured to use and calculate taxes.<br><br>If using Avalara, a tax_code parameter must also be sent. |
| `TaxCode` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | YYYY-MM-DD |
| `PeriodRangeEnd` | `*string` | Optional | YYYY-MM-DD |
| `ProductId` | `*interface{}` | Optional | Product handle or product id. |
| `ComponentId` | `*interface{}` | Optional | Component handle or component id. |
| `PricePointId` | `*interface{}` | Optional | Price point handle or id. For component. |
| `ProductPricePointId` | `*interface{}` | Optional | - |
| `Description` | `*string` | Optional | **Constraints**: *Maximum Length*: `255` |

## Example (as JSON)

```json
{
  "title": "title2",
  "quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "unit_price": {
    "key1": "val1",
    "key2": "val2"
  },
  "taxable": false,
  "tax_code": "tax_code4"
}
```

