
# Bulk Update Segments Item

## Structure

`BulkUpdateSegmentsItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | The ID of the segment you want to update. |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bulkUpdateSegmentsItem := models.BulkUpdateSegmentsItem{
        Id:                   220,
        PricingScheme:        models.PricingScheme_PERUNIT,
        Prices:               []models.CreateOrUpdateSegmentPrice{
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
        },
    }

}
```

