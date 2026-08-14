
# Account Balances

## Structure

`AccountBalances`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OpenInvoices` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | The balance, in cents, of the sum of the subscription's open, payable invoices. |
| `PendingInvoices` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | The balance, in cents, of the sum of the subscription's pending, payable invoices. |
| `PendingDiscounts` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | The balance, in cents, of the subscription's Pending Discount account. |
| `ServiceCredits` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | The balance, in cents, of the subscription's Service Credit account. |
| `Prepayments` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | The balance, in cents, of the subscription's Prepayment account. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    accountBalances := models.AccountBalances{
        OpenInvoices:         models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(40)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(202))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(170))),
        }),
        PendingInvoices:      models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(0)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(242))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(46))),
        }),
        PendingDiscounts:     models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(88)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(154))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(134))),
        }),
        ServiceCredits:       models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(84)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(70))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(38))),
        }),
        Prepayments:          models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(192)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(178))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(146))),
        }),
    }

}
```

