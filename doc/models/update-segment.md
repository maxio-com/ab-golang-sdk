
# Update Segment

## Structure

`UpdateSegment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricingScheme` | [`models.PricingScheme`](../../doc/models/pricing-scheme.md) | Required | The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes. |
| `Prices` | [`[]models.CreateOrUpdateSegmentPrice`](../../doc/models/create-or-update-segment-price.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateSegment := models.UpdateSegment{
        PricingScheme:        models.PricingScheme_STAIRSTEP,
        Prices:               []models.CreateOrUpdateSegmentPrice{
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
            models.CreateOrUpdateSegmentPrice{
                StartingQuantity:     models.ToPointer(64),
                EndingQuantity:       models.ToPointer(38),
                UnitPrice:            models.CreateOrUpdateSegmentPriceUnitPriceContainer.FromString("String3"),
            },
        },
    }

}
```

