
# Payment Method Credit Card

## Structure

`PaymentMethodCreditCard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CardBrand` | `string` | Required | - |
| `CardExpiration` | `*string` | Optional | - |
| `LastFour` | `models.Optional[string]` | Optional | - |
| `MaskedCardNumber` | `string` | Required | - |
| `Type` | [`models.InvoiceEventPaymentMethod`](../../doc/models/invoice-event-payment-method.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentMethodCreditCard := models.PaymentMethodCreditCard{
        CardBrand:            "card_brand4",
        CardExpiration:       models.ToPointer("card_expiration2"),
        LastFour:             models.NewOptional(models.ToPointer("last_four6")),
        MaskedCardNumber:     "masked_card_number0",
        Type:                 models.InvoiceEventPaymentMethod_CREDITCARD,
    }

}
```

