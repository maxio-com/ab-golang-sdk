
# Wallet Purchase Result Response

## Structure

`WalletPurchaseResultResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `WalletPurchases` | [`[]models.WalletPurchase`](../../doc/models/wallet-purchase.md) | Required | - |
| `Balance` | `string` | Required | The wallet's new balance, as a decimal string. |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    walletPurchaseResultResponse := models.WalletPurchaseResultResponse{
        WalletPurchases:      []models.WalletPurchase{
            models.WalletPurchase{
                Id:                   244,
                PricePointId:         134,
                InvoiceUid:           "invoice_uid0",
                Quantity:             "quantity6",
                ExpiresAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            },
        },
        Balance:              "balance6",
    }

}
```

