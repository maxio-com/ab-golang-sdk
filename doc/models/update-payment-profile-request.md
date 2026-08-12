
# Update Payment Profile Request

## Structure

`UpdatePaymentProfileRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.UpdatePaymentProfile`](../../doc/models/update-payment-profile.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updatePaymentProfileRequest := models.UpdatePaymentProfileRequest{
        PaymentProfile:       models.UpdatePaymentProfile{
            FirstName:            models.ToPointer("first_name4"),
            LastName:             models.ToPointer("last_name2"),
            FullNumber:           models.ToPointer("5424000000000015"),
            CardType:             models.ToPointer(models.CardType_BOGUS),
            ExpirationMonth:      models.ToPointer("expiration_month0"),
        },
    }

}
```

