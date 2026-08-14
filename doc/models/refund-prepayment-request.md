
# Refund Prepayment Request

## Structure

`RefundPrepaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`models.RefundPrepayment`](../../doc/models/refund-prepayment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundPrepaymentRequest := models.RefundPrepaymentRequest{
        Refund:               models.RefundPrepayment{
            AmountInCents:        models.ToPointer(int64(132)),
            Amount:               models.RefundPrepaymentAmountContainer.FromString("String1"),
            Memo:                 "memo2",
            External:             models.ToPointer(false),
        },
    }

}
```

