
# Chjs Tokenization Success

## Structure

`ChjsTokenizationSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.TokenizedPaymentProfile`](../../doc/models/tokenized-payment-profile.md) | Required | - |
| `GatewayCustomerId` | `models.Optional[int]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    chjsTokenizationSuccess := models.ChjsTokenizationSuccess{
        PaymentProfile:       models.TokenizedPaymentProfile{
            Id:                   44,
            VaultToken:           models.ToPointer("vault_token6"),
            GatewayHandle:        models.NewOptional(models.ToPointer("gateway_handle4")),
            CustomerVaultToken:   models.NewOptional(models.ToPointer("customer_vault_token2")),
        },
        GatewayCustomerId:    models.NewOptional(models.ToPointer(228)),
    }

}
```

