
# Create Prepayment

## Structure

`CreatePrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `float64` | Required | - |
| `Details` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `Method` | [`models.CreatePrepaymentMethod`](../../doc/models/create-prepayment-method.md) | Required | When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions. |
| `PaymentProfileId` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPrepayment := models.CreatePrepayment{
        Amount:               float64(73.78),
        Details:              "details6",
        Memo:                 "memo0",
        Method:               models.CreatePrepaymentMethod_PAYPALACCOUNT,
        PaymentProfileId:     models.ToPointer(58),
    }

}
```

