
# Bank Account Verification

## Structure

`BankAccountVerification`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Deposit1InCents` | `*int64` | Optional | - |
| `Deposit2InCents` | `*int64` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountVerification := models.BankAccountVerification{
        Deposit1InCents:      models.ToPointer(int64(248)),
        Deposit2InCents:      models.ToPointer(int64(10)),
    }

}
```

