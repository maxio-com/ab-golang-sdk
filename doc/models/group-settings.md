
# Group Settings

## Structure

`GroupSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Target` | [`models.GroupTarget`](../../doc/models/group-target.md) | Required | Attributes of the target customer who will be the responsible payer of the created subscription. Required. |
| `Billing` | [`*models.GroupBilling`](../../doc/models/group-billing.md) | Optional | (Optional) Attributes related to billing date and accrual. Note: Only applicable for new subscriptions. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    groupSettings := models.GroupSettings{
        Target:               models.GroupTarget{
            Type:                 models.GroupTargetType_PARENT,
            Id:                   models.ToPointer(236),
        },
        Billing:              models.ToPointer(models.GroupBilling{
            Accrue:               models.ToPointer(false),
            AlignDate:            models.ToPointer(false),
            Prorate:              models.ToPointer(false),
        }),
    }

}
```

