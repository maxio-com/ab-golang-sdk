
# Tokenized Payment Profile

## Structure

`TokenizedPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | - |
| `VaultToken` | `*string` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    tokenizedPaymentProfile := models.TokenizedPaymentProfile{
        Id:                   116,
        VaultToken:           models.ToPointer("vault_token0"),
        GatewayHandle:        models.NewOptional(models.ToPointer("gateway_handle0")),
        CustomerVaultToken:   models.NewOptional(models.ToPointer("customer_vault_token6")),
    }

}
```

