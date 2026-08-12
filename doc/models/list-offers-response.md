
# List Offers Response

## Structure

`ListOffersResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Offers` | [`[]models.Offer`](../../doc/models/offer.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listOffersResponse := models.ListOffersResponse{
        Offers:               []models.Offer{
            models.Offer{
                Id:                     models.ToPointer(12),
                SiteId:                 models.ToPointer(194),
                ProductFamilyId:        models.ToPointer(16),
                ProductId:              models.ToPointer(210),
                ProductPricePointId:    models.ToPointer(134),
            },
        },
    }

}
```

