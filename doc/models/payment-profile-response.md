
# Payment Profile Response

## Structure

`PaymentProfileResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.PaymentProfile`](../../doc/models/containers/payment-profile.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentProfileResponse := models.PaymentProfileResponse{
        PaymentProfile:       models.PaymentProfileContainer.FromApplePayPaymentProfile(models.ApplePayPaymentProfile{
            Id:                   models.ToPointer(60),
            FirstName:            models.ToPointer("first_name2"),
            LastName:             models.ToPointer("last_name0"),
            CustomerId:           models.ToPointer(98),
            CurrentVault:         models.ToPointer(models.ApplePayVault_BRAINTREEBLUE),
            PaymentType:          models.PaymentType_APPLEPAY,
        }),
    }

}
```

