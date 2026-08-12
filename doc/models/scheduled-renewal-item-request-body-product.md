
# Scheduled Renewal Item Request Body Product

## Structure

`ScheduledRenewalItemRequestBodyProduct`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ItemType` | `string` | Required, Constant | Item type to add. Either Product or Component.<br><br>**Value**: `"Product"` |
| `ItemId` | `int` | Required | Product or component identifier. |
| `PricePointId` | `*int` | Optional | Price point identifier. |
| `Quantity` | `*int` | Optional | (Optional) Quantity for the item. |
| `CustomPrice` | [`*models.ScheduledRenewalProductPricePoint`](../../doc/models/scheduled-renewal-product-price-point.md) | Optional | Custom pricing for a product within a scheduled renewal. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    scheduledRenewalItemRequestBodyProduct := models.ScheduledRenewalItemRequestBodyProduct{
        ItemType:             "Product",
        ItemId:               154,
        PricePointId:         models.ToPointer(168),
        Quantity:             models.ToPointer(166),
        CustomPrice:          models.ToPointer(models.ScheduledRenewalProductPricePoint{
            Name:                   models.ToPointer("name4"),
            Handle:                 models.ToPointer("handle0"),
            PriceInCents:           models.ScheduledRenewalProductPricePointPriceInCentsContainer.FromString("String3"),
            Interval:               models.ScheduledRenewalProductPricePointIntervalContainer.FromString("String3"),
            IntervalUnit:           models.ToPointer(models.IntervalUnit_DAY),
            TaxIncluded:            models.ToPointer(false),
            InitialChargeInCents:   models.ToPointer(int64(30)),
            ExpirationInterval:     models.ToPointer(52),
        }),
    }

}
```

