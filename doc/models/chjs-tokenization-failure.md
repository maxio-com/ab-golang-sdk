
# Chjs Tokenization Failure

## Structure

`ChjsTokenizationFailure`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `string` | Required | - |
| `PaymentProfileParams` | [`*models.PaymentProfileParams`](../../doc/models/payment-profile-params.md) | Optional | PCI-safe cardholder fields only. Full card numbers, CVV, and billing address are never included. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    chjsTokenizationFailure := models.ChjsTokenizationFailure{
        Errors:               "errors2",
        PaymentProfileParams: models.ToPointer(models.PaymentProfileParams{
            FirstName:            models.ToPointer("first_name2"),
            LastName:             models.ToPointer("last_name0"),
            CardType:             models.ToPointer("card_type2"),
        }),
    }

}
```

