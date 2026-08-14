
# Create Offer

## Structure

`CreateOffer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | - |
| `Handle` | `string` | Required | - |
| `Description` | `*string` | Optional | - |
| `ProductId` | `int` | Required | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `Components` | [`[]models.CreateOfferComponent`](../../doc/models/create-offer-component.md) | Optional | - |
| `Coupons` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOffer := models.CreateOffer{
        Name:                 "name6",
        Handle:               "handle2",
        Description:          models.ToPointer("description6"),
        ProductId:            66,
        ProductPricePointId:  models.ToPointer(246),
        Components:           []models.CreateOfferComponent{
            models.CreateOfferComponent{
                ComponentId:          models.ToPointer(108),
                PricePointId:         models.ToPointer(124),
                StartingQuantity:     models.ToPointer(84),
            },
            models.CreateOfferComponent{
                ComponentId:          models.ToPointer(108),
                PricePointId:         models.ToPointer(124),
                StartingQuantity:     models.ToPointer(84),
            },
            models.CreateOfferComponent{
                ComponentId:          models.ToPointer(108),
                PricePointId:         models.ToPointer(124),
                StartingQuantity:     models.ToPointer(84),
            },
        },
        Coupons:              []string{
            "coupons6",
            "coupons5",
            "coupons4",
        },
    }

}
```

