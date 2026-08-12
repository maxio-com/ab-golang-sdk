
# Payment Method External

## Structure

`PaymentMethodExternal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Details` | `*string` | Required | - |
| `Kind` | `string` | Required | - |
| `Memo` | `*string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentMethodExternal := models.PaymentMethodExternal{
        Details:              models.ToPointer("details0"),
        Kind:                 "kind8",
        Memo:                 models.ToPointer("memo4"),
        Type:                 models.InvoiceEventPaymentMethod_EXTERNAL,
    }

}
```

