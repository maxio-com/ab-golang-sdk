
# Payment Method Paypal

## Structure

`PaymentMethodPaypal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Email` | `string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentMethodPaypal := models.PaymentMethodPaypal{
        Email:                "email8",
        Type:                 models.InvoiceEventPaymentMethod_PAYPALACCOUNT,
    }

}
```

