
# Wallet Grant

A funding tranche embedded on the wallet. `original` is the amount the grant was created with; `posted_remaining` is what debits and adjustments written to the ledger have left; `remaining` is what can still be spent. Debits drain grants earliest-expiration-first (grants with no expiration are drained last).

## Structure

`WalletGrant`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Advanced Billing id of the grant. |
| `Name` | `*string` | Optional | The memo the grant was created with, or `Grant |
| `Original` | `*string` | Optional | The amount the grant was created with, encoded as a decimal string. Never changes after creation. |
| `Remaining` | `*string` | Optional | What the grant can still spend, encoded as a decimal string. On [Read Wallet](../../doc/controllers/wallets.md#read-wallet) for a metered wallet whose `balance` is the live figure, this is `posted_remaining` less the pending usage drawn from this grant in drain order: usage past a grant spills to the next one, and usage past every grant is an overdraft that belongs to no grant. It never goes below zero, and the grants' `remaining` add up to the wallet's `balance` whenever `balance` is zero or more. Everywhere else it equals `posted_remaining`. |
| `PostedRemaining` | `*string` | Optional | The grant's remainder from written ledger entries alone, encoded as a decimal string. Always present. This is the figure an adjustment checks `expected_remaining` and its bounds against, so send it back, not `remaining`, when correcting a grant. |
| `ExpiresAt` | `models.Optional[time.Time]` | Optional | When this grant's unconsumed remainder is automatically written off. `null` for an evergreen grant. |
| `ExpiresOn` | `models.Optional[time.Time]` | Optional | The calendar day `expires_at` falls on, already resolved in the site's time zone. |
| `ExpiredAt` | `models.Optional[time.Time]` | Optional | When the expiry sweep actually wrote this grant's remainder off. Usually falls shortly after `expires_at`; use `expires_on` for the date the grant expired, not this field. `null` if the grant was never written off (including a grant simply spent to zero before its `expires_at`). |
| `ExpiredOn` | `models.Optional[time.Time]` | Optional | The calendar day `expired_at` falls on, resolved in the site's time zone. |
| `Origin` | [`*models.Origin`](../../doc/models/origin.md) | Optional | How the grant was created. `recurring` comes from the wallet's automatic funding schedule (see the wallet's `recurring_funding` array). `manual` comes from a one-time grant posted to [Create Wallet Transaction](../../doc/controllers/wallet-transactions.md#create-wallet-transaction). `initial_balance` comes from a non-zero `initial_balance` at wallet creation. `allowance_increase` comes from raising a funding entitlement's allowance while it's already funding a wallet. `purchase` comes from [Create Wallet Purchase](../../doc/controllers/wallet-purchases.md#create-wallet-purchase). |
| `GrantedForPeriodStart` | `models.Optional[time.Time]` | Optional | For a `recurring` grant, the start of the funding cadence period this grant covers. `null` for every other origin. After a gap in funding (see `skipped_periods`), this is the period currently in progress, not the oldest period that was missed. |
| `InvoiceUid` | `models.Optional[string]` | Optional | The uid of the invoice that paid for this grant, for a grant with `origin` of `purchase`. `null` for every other origin. |
| `CreatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletGrant := models.WalletGrant{
        Id:                    models.ToPointer(8),
        Name:                  models.ToPointer("name4"),
        Original:              models.ToPointer("100.0"),
        Remaining:             models.ToPointer("50.0"),
        PostedRemaining:       models.ToPointer("50.0"),
    }

}
```

