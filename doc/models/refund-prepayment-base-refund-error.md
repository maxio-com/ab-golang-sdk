
# Refund Prepayment Base Refund Error

## Structure

`RefundPrepaymentBaseRefundError`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Refund` | [`*models.BaseRefundError`](../../doc/models/base-refund-error.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundPrepaymentBaseRefundError := models.RefundPrepaymentBaseRefundError{
        Refund:               models.ToPointer(models.BaseRefundError{
            Base:                 []interface{}{
                interface{}("[key1, val1][key2, val2]"),
            },
        }),
    }

}
```

