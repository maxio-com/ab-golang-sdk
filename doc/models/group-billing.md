
# Group Billing

(Optional) Attributes related to billing date and accrual. Note: Only applicable for new subscriptions.

## Structure

`GroupBilling`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accrue` | `*bool` | Optional | A flag indicating whether or not to accrue charges on the new subscription.<br><br>**Default**: `false` |
| `AlignDate` | `*bool` | Optional | A flag indicating whether or not to align the billing date of the new subscription with the billing date of the primary subscription of the hierarchy's default subscription group. Required to be true if prorate is also true.<br><br>**Default**: `false` |
| `Prorate` | `*bool` | Optional | A flag indicating whether or not to prorate billing of the new subscription for the current period. A value of true is ignored unless align_date is also true.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    groupBilling := models.GroupBilling{
        Accrue:               models.ToPointer(false),
        AlignDate:            models.ToPointer(false),
        Prorate:              models.ToPointer(false),
    }

}
```

