
# Bank Account Response

## Structure

`BankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.BankAccountPaymentProfile`](../../doc/models/bank-account-payment-profile.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountResponse := models.BankAccountResponse{
        PaymentProfile:       models.BankAccountPaymentProfile{
            Id:                      models.ToPointer(44),
            FirstName:               models.ToPointer("first_name4"),
            LastName:                models.ToPointer("last_name2"),
            CustomerId:              models.ToPointer(82),
            CurrentVault:            models.ToPointer(models.BankAccountVault_AUTHORIZENET),
            PaymentType:             models.PaymentType_BANKACCOUNT,
            Verified:                models.ToPointer(false),
        },
    }

}
```

