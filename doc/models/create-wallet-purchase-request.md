
# Create Wallet Purchase Request

Purchases prepaid units for the wallet at a price point, charging the subscription's payment method.

## Structure

`CreateWalletPurchaseRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Purchase` | [`models.Purchase`](../../doc/models/purchase.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createWalletPurchaseRequest := models.CreateWalletPurchaseRequest{
        Purchase:             models.Purchase{
            Quantity:             models.PurchaseQuantityContainer.FromString("String5"),
            PricePointId:         models.ToPointer(20),
            Memo:                 models.NewOptional(models.ToPointer("memo2")),
        },
    }

}
```

