
# Offer Response

## Structure

`OfferResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Offer` | [`*models.Offer`](../../doc/models/offer.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    offerResponse := models.OfferResponse{
        Offer:                models.ToPointer(models.Offer{
            Id:                     models.ToPointer(28),
            SiteId:                 models.ToPointer(210),
            ProductFamilyId:        models.ToPointer(224),
            ProductId:              models.ToPointer(30),
            ProductPricePointId:    models.ToPointer(150),
        }),
    }

}
```

