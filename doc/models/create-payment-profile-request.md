
# Create Payment Profile Request

## Structure

`CreatePaymentProfileRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.CreatePaymentProfile`](../../doc/models/create-payment-profile.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPaymentProfileRequest := models.CreatePaymentProfileRequest{
        PaymentProfile:       models.CreatePaymentProfile{
            ChargifyToken:         models.ToPointer("tok_9g6hw85pnpt6knmskpwp4ttt"),
            Id:                    models.ToPointer(44),
            PaymentType:           models.ToPointer(models.PaymentType_CREDITCARD),
            FirstName:             models.ToPointer("first_name4"),
            LastName:              models.ToPointer("last_name2"),
            FullNumber:            models.ToPointer("5424000000000015"),
        },
    }

}
```

