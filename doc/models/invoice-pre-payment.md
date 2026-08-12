
# Invoice Pre Payment

## Structure

`InvoicePrePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `*int` | Optional | The subscription id for the prepayment account |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the prepayment that was created as a result of this payment. |
| `EndingBalanceInCents` | `*int64` | Optional | The total balance of the prepayment account for this subscription including any prior prepayments |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoicePrePayment := models.InvoicePrePayment{
        SubscriptionId:       models.ToPointer(252),
        AmountInCents:        models.ToPointer(int64(28)),
        EndingBalanceInCents: models.ToPointer(int64(244)),
    }

}
```

