
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
| `Description` | `Optional[string]` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `UpdatedAt` | `*string` | Optional | - |
| `ArchivedAt` | `Optional[string]` | Optional | - |
| `OfferItems` | [`[]models.OfferItem`](offer-item.md) | Optional | - |
| `OfferDiscounts` | [`[]models.OfferDiscount`](offer-discount.md) | Optional | - |
| `ProductFamilyName` | `*string` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `ProductPricePointName` | `*string` | Optional | - |
| `ProductPriceInCents` | `*int64` | Optional | - |
| `OfferSignupPages` | [`[]models.OfferSignupPage`](offer-signup-page.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": 154,
  "site_id": 80,
  "product_family_id": 158,
  "product_id": 96,
  "product_price_point_id": 20
}
```

