
# Refund Success

## Structure

`RefundSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RefundId` | `int` | Required | - |
| `GatewayTransactionId` | `int` | Required | - |
| `ProductId` | `int` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    refundSuccess := models.RefundSuccess{
        RefundId:             34,
        GatewayTransactionId: 160,
        ProductId:            190,
    }

}
```

