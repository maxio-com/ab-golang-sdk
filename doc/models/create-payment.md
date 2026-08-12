
# Create Payment

## Structure

`CreatePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `PaymentDetails` | `string` | Required | - |
| `PaymentMethod` | [`models.InvoicePaymentMethodType`](../../doc/models/invoice-payment-method-type.md) | Required | The type of payment method used. Defaults to other. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPayment := models.CreatePayment{
        Amount:               "amount6",
        Memo:                 "memo8",
        PaymentDetails:       "payment_details4",
        PaymentMethod:        models.InvoicePaymentMethodType_CREDITCARD,
    }

}
```

