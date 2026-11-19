
# Update Wallet Request

## Structure

`UpdateWalletRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Wallet` | [`models.Wallet2`](../../doc/models/wallet-2.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateWalletRequest := models.UpdateWalletRequest{
        Wallet:               models.Wallet2{
            Name:                 models.NewOptional(models.ToPointer("name6")),
        },
    }

}
```

