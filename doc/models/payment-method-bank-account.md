
# Payment Method Bank Account

## Structure

`PaymentMethodBankAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MaskedAccountNumber` | `string` | Required | - |
| `MaskedRoutingNumber` | `string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentMethodBankAccount := models.PaymentMethodBankAccount{
        MaskedAccountNumber:  "masked_account_number4",
        MaskedRoutingNumber:  "masked_routing_number4",
        Type:                 models.InvoiceEventPaymentMethod_BANKACCOUNT,
    }

}
```

