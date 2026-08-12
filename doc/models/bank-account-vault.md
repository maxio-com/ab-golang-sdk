
# Bank Account Vault

The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing.

## Enumeration

`BankAccountVault`

## Fields

| Name |
|  --- |
| `AUTHORIZENET` |
| `BLUESNAP` |
| `BOGUS` |
| `FORTE` |
| `GOCARDLESS` |
| `MAXIOPAYMENTS` |
| `MAXP` |
| `STRIPECONNECT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountVault := models.BankAccountVault_MAXP

}
```

