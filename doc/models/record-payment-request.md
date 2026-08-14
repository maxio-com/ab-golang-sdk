
# Record Payment Request

## Structure

`RecordPaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payment` | [`models.CreatePayment`](../../doc/models/create-payment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    recordPaymentRequest := models.RecordPaymentRequest{
        Payment:              models.CreatePayment{
            Amount:               "amount8",
            Memo:                 "memo0",
            PaymentDetails:       "payment_details6",
            PaymentMethod:        models.InvoicePaymentMethodType_CASH,
        },
    }

}
```

