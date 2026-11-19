
# Wallet Purchase Preview

A price quote for a purchase, without charging anything.

## Structure

`WalletPurchasePreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePointId` | `int` | Required | - |
| `Currency` | `string` | Required | The three-letter currency code the `amount` is denominated in, matching the subscription's currency. |
| `Quantity` | `string` | Required | The number of units the quote is for, as a decimal string. |
| `Amount` | `string` | Required | The price for `quantity` units at this price point, as a decimal string. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletPurchasePreview := models.WalletPurchasePreview{
        PricePointId:         98,
        Currency:             "USD",
        Quantity:             "quantity4",
        Amount:               "amount0",
    }

}
```

