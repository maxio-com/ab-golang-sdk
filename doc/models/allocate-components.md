
# Allocate Components

## Structure

`AllocateComponents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProrationUpgradeScheme` | `*string` | Optional | - |
| `ProrationDowngradeScheme` | `*string` | Optional | - |
| `Allocations` | [`[]models.CreateAllocation`](../../doc/models/create-allocation.md) | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `PaymentCollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | (Optional) If not passed, the allocation(s) will use the payment collection method on the subscription. |
| `InitiateDunning` | `*bool` | Optional | If true, if the immediate component payment fails, initiate dunning for the subscription.<br>Otherwise, leave the charges on the subscription to pay for at renewal. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    allocateComponents := models.AllocateComponents{
        ProrationUpgradeScheme:   models.ToPointer("proration_upgrade_scheme8"),
        ProrationDowngradeScheme: models.ToPointer("proration_downgrade_scheme6"),
        Allocations:              []models.CreateAllocation{
            models.CreateAllocation{
                Quantity:                 float64(26.48),
                DecimalQuantity:          models.ToPointer("decimal_quantity8"),
                PreviousQuantity:         models.ToPointer(float64(55.5)),
                DecimalPreviousQuantity:  models.ToPointer("decimal_previous_quantity2"),
                ComponentId:              models.ToPointer(242),
                Memo:                     models.ToPointer("memo6"),
            },
            models.CreateAllocation{
                Quantity:                 float64(26.48),
                DecimalQuantity:          models.ToPointer("decimal_quantity8"),
                PreviousQuantity:         models.ToPointer(float64(55.5)),
                DecimalPreviousQuantity:  models.ToPointer("decimal_previous_quantity2"),
                ComponentId:              models.ToPointer(242),
                Memo:                     models.ToPointer("memo6"),
            },
            models.CreateAllocation{
                Quantity:                 float64(26.48),
                DecimalQuantity:          models.ToPointer("decimal_quantity8"),
                PreviousQuantity:         models.ToPointer(float64(55.5)),
                DecimalPreviousQuantity:  models.ToPointer("decimal_previous_quantity2"),
                ComponentId:              models.ToPointer(242),
                Memo:                     models.ToPointer("memo6"),
            },
        },
        AccrueCharge:             models.ToPointer(false),
        UpgradeCharge:            models.NewOptional(models.ToPointer(models.CreditType_PRORATED)),
    }

}
```

