
# Wallet Purchase

A completed unit purchase.

## Structure

`WalletPurchase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | The id of the grant this purchase created. When the wallet was overdrawn by at least the purchased quantity, the units repay the debt and the grant is still created, with `remaining` of `0` and `original` equal to the purchased quantity. |
| `PricePointId` | `int` | Required | The price point the purchase was priced at. |
| `InvoiceUid` | `string` | Required | The uid of the invoice created to collect payment for this purchase. |
| `Quantity` | `string` | Required | The number of units purchased, as a decimal string. |
| `ExpiresAt` | `*time.Time` | Required | When the purchased units expire, following the funding entitlement's grant expiration policy. `null` when the wallet has no funding entitlement or its entitlements disagree on an expiration policy. |

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
    walletPurchase := models.WalletPurchase{
        Id:                   218,
        PricePointId:         96,
        InvoiceUid:           "invoice_uid0",
        Quantity:             "quantity6",
        ExpiresAt:            models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
    }

}
```

