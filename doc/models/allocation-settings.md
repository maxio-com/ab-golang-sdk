
# Allocation Settings

## Structure

`AllocationSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `AccrueCharge` | `*string` | Optional | Either "true" or "false". |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    allocationSettings := models.AllocationSettings{
        UpgradeCharge:        models.NewOptional(models.ToPointer(models.CreditType_PRORATED)),
        DowngradeCredit:      models.NewOptional(models.ToPointer(models.CreditType_PRORATED)),
        AccrueCharge:         models.ToPointer("accrue_charge2"),
    }

}
```

