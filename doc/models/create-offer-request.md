
# Create Offer Request

## Structure

`CreateOfferRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Offer` | [`models.CreateOffer`](../../doc/models/create-offer.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOfferRequest := models.CreateOfferRequest{
        Offer:                models.CreateOffer{
            Name:                 "name4",
            Handle:               "handle0",
            Description:          models.ToPointer("description6"),
            ProductId:            30,
            ProductPricePointId:  models.ToPointer(150),
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
            },
            Coupons:              []string{
                "coupons6",
            },
        },
    }

}
```

