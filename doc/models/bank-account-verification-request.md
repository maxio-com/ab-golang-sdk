
# Bank Account Verification Request

## Structure

`BankAccountVerificationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BankAccountVerification` | [`models.BankAccountVerification`](../../doc/models/bank-account-verification.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountVerificationRequest := models.BankAccountVerificationRequest{
        BankAccountVerification: models.BankAccountVerification{
            Deposit1InCents:      models.ToPointer(int64(244)),
            Deposit2InCents:      models.ToPointer(int64(6)),
        },
    }

}
```

