
# Refund Prepayment

## Structure

`RefundPrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmountInCents` | `*int64` | Required | `amount` is not required if you pass `amount_in_cents`. |
| `Amount` | [`models.RefundPrepaymentAmount`](../../doc/models/containers/refund-prepayment-amount.md) | Required | This is a container for one-of cases. |
| `Memo` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `External` | `*bool` | Optional | Specify the type of refund you wish to initiate. When the prepayment is external, the `external` flag is optional. But if the prepayment was made through a payment profile, the `external` flag is required. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundPrepayment := models.RefundPrepayment{
        AmountInCents:        models.ToPointer(int64(22)),
        Amount:               models.RefundPrepaymentAmountContainer.FromString("String1"),
        Memo:                 "memo2",
        External:             models.ToPointer(false),
    }

}
```

