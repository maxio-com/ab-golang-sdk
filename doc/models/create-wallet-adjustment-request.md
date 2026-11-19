
# Create Wallet Adjustment Request

A signed correction to a specific grant's balance, or to an overdraft entry's usage, intended for support and ops use.

## Structure

`CreateWalletAdjustmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Adjustment` | [`models.Adjustment`](../../doc/models/adjustment.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createWalletAdjustmentRequest := models.CreateWalletAdjustmentRequest{
        Adjustment:           models.Adjustment{
            Amount:                    models.AdjustmentAmountContainer.FromString("String7"),
            Reason:                    "reason0",
            GrantId:                   models.NewOptional(models.ToPointer(104)),
            SourceTransactionId:       models.NewOptional(models.ToPointer(126)),
            ExpectedRemaining:         models.NewOptional(models.ToPointer("expected_remaining6")),
            ExpectedSourceCorrections: models.NewOptional(models.ToPointer("expected_source_corrections8")),
        },
    }

}
```

