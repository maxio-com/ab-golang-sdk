
# Subscription Group Payment Profile

## Structure

`SubscriptionGroupPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `MaskedCardNumber` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupPaymentProfile := models.SubscriptionGroupPaymentProfile{
        Id:                   models.ToPointer(246),
        FirstName:            models.ToPointer("first_name6"),
        LastName:             models.ToPointer("last_name4"),
        MaskedCardNumber:     models.ToPointer("masked_card_number4"),
    }

}
```

