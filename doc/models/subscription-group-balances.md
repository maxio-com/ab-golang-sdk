
# Subscription Group Balances

## Structure

`SubscriptionGroupBalances`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayments` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | - |
| `ServiceCredits` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | - |
| `OpenInvoices` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | - |
| `PendingDiscounts` | [`*models.AccountBalance`](../../doc/models/account-balance.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupBalances := models.SubscriptionGroupBalances{
        Prepayments:          models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(192)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(178))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(146))),
        }),
        ServiceCredits:       models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(84)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(70))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(38))),
        }),
        OpenInvoices:         models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(40)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(202))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(170))),
        }),
        PendingDiscounts:     models.ToPointer(models.AccountBalance{
            BalanceInCents:           models.ToPointer(int64(88)),
            AutomaticBalanceInCents:  models.NewOptional(models.ToPointer(int64(154))),
            RemittanceBalanceInCents: models.NewOptional(models.ToPointer(int64(134))),
        }),
    }

}
```

