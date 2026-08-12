
# Reactivation Billing

These values are only applicable to subscriptions using calendar billing.

## Structure

`ReactivationBilling`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ReactivationCharge` | [`*models.ReactivationCharge`](../../doc/models/reactivation-charge.md) | Optional | You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal.<br><br>**Default**: `"prorated"` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reactivationBilling := models.ReactivationBilling{
        ReactivationCharge:   models.ToPointer(models.ReactivationCharge_PRORATED),
    }

}
```

