
# Payment Related Events

## Structure

`PaymentRelatedEvents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductId` | `int` | Required | - |
| `AccountTransactionId` | `int` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentRelatedEvents := models.PaymentRelatedEvents{
        ProductId:            186,
        AccountTransactionId: 170,
    }

}
```

