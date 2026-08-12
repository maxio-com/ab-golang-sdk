
# Invoice Line Item Event Data

## Structure

`InvoiceLineItemEventData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Title` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Quantity` | `*int` | Optional | - |
| `QuantityDelta` | `models.Optional[int]` | Optional | - |
| `UnitPrice` | `*string` | Optional | - |
| `PeriodRangeStart` | `*string` | Optional | - |
| `PeriodRangeEnd` | `*string` | Optional | - |
| `Amount` | `*string` | Optional | - |
| `LineReferences` | `*string` | Optional | - |
| `PricingDetailsIndex` | `models.Optional[int]` | Optional | - |
| `PricingDetails` | [`[]models.InvoiceLineItemPricingDetail`](../../doc/models/invoice-line-item-pricing-detail.md) | Optional | - |
| `TaxCode` | `models.Optional[string]` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductPricePointId` | `models.Optional[int]` | Optional | - |
| `PricePointId` | `models.Optional[int]` | Optional | - |
| `ComponentId` | `models.Optional[int]` | Optional | - |
| `BillingScheduleItemId` | `models.Optional[int]` | Optional | - |
| `CustomItem` | `models.Optional[bool]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceLineItemEventData := models.InvoiceLineItemEventData{
        Uid:                   models.ToPointer("uid4"),
        Title:                 models.ToPointer("title0"),
        Description:           models.ToPointer("description4"),
        Quantity:              models.ToPointer(190),
        QuantityDelta:         models.NewOptional(models.ToPointer(36)),
    }

}
```

