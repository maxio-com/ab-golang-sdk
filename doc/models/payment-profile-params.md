
# Payment Profile Params

PCI-safe cardholder fields only. Full card numbers, CVV, and billing address are never included.

## Structure

`PaymentProfileParams`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `CardType` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentProfileParams := models.PaymentProfileParams{
        FirstName:            models.ToPointer("first_name2"),
        LastName:             models.ToPointer("last_name0"),
        CardType:             models.ToPointer("card_type2"),
    }

}
```

