
# Offer

## Structure

`Offer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `ProductFamilyId` | `*int` | Optional | - |
| `ProductId` | `*int` | Optional | - |
| `ProductPricePointId` | `*int` | Optional | - |
| `ProductRevisableNumber` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | - |
| `OfferItems` | [`[]models.OfferItem`](../../doc/models/offer-item.md) | Optional | - |
| `OfferDiscounts` | [`[]models.OfferDiscount`](../../doc/models/offer-discount.md) | Optional | - |
| `ProductFamilyName` | `*string` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `ProductPricePointName` | `*string` | Optional | - |
| `ProductPriceInCents` | `*int64` | Optional | - |
| `OfferSignupPages` | [`[]models.OfferSignupPage`](../../doc/models/offer-signup-page.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    offer := models.Offer{
        Id:                     models.ToPointer(28),
        SiteId:                 models.ToPointer(210),
        ProductFamilyId:        models.ToPointer(224),
        ProductId:              models.ToPointer(30),
        ProductPricePointId:    models.ToPointer(150),
    }

}
```

