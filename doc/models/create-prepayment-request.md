
# Create Prepayment Request

## Structure

`CreatePrepaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.CreatePrepayment`](../../doc/models/create-prepayment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPrepaymentRequest := models.CreatePrepaymentRequest{
        Prepayment:           models.CreatePrepayment{
            Amount:               float64(11.6),
            Details:              "details8",
            Memo:                 "memo2",
            Method:               models.CreatePrepaymentMethod_MONEYORDER,
            PaymentProfileId:     models.ToPointer(240),
        },
    }

}
```

