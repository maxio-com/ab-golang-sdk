
# Wallet Purchase Preview Response

## Structure

`WalletPurchasePreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletPurchasePreviews` | [`[]models.WalletPurchasePreview`](../../doc/models/wallet-purchase-preview.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletPurchasePreviewResponse := models.WalletPurchasePreviewResponse{
        WalletPurchasePreviews: []models.WalletPurchasePreview{
            models.WalletPurchasePreview{
                PricePointId:         126,
                Currency:             "USD",
                Quantity:             "quantity6",
                Amount:               "amount2",
            },
        },
    }

}
```

