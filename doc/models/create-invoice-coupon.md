
# Create Invoice Coupon

## Structure

`CreateInvoiceCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | - |
| `Subcode` | `*string` | Optional | - |
| `Percentage` | [`*models.CreateInvoiceCouponPercentage`](../../doc/models/containers/create-invoice-coupon-percentage.md) | Optional | This is a container for one-of cases. |
| `Amount` | [`*models.CreateInvoiceCouponAmount`](../../doc/models/containers/create-invoice-coupon-amount.md) | Optional | This is a container for one-of cases. |
| `Description` | `*string` | Optional | **Constraints**: *Maximum Length*: `255` |
| `ProductFamilyId` | [`*models.CreateInvoiceCouponProductFamilyId`](../../doc/models/containers/create-invoice-coupon-product-family-id.md) | Optional | This is a container for one-of cases. |
| `CompoundingStrategy` | [`*models.CompoundingStrategy`](../../doc/models/compounding-strategy.md) | Optional | Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createInvoiceCoupon := models.CreateInvoiceCoupon{
        Code:                 models.ToPointer("code8"),
        Subcode:              models.ToPointer("subcode4"),
        Percentage:           models.ToPointer(models.CreateInvoiceCouponPercentageContainer.FromPrecision(float64(50))),
        Amount:               models.ToPointer(models.CreateInvoiceCouponAmountContainer.FromString("String9")),
        Description:          models.ToPointer("description0"),
    }

}
```

