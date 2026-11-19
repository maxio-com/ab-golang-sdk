
# Wallet Product Price Point Summary

A summary of the Wallet Product's default price point, for rendering a catalog listing without a separate price points call.

## Structure

`WalletProductPricePointSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ArchivedCount` | `*int` | Optional | Number of archived price points on this Wallet Product. |
| `PricingScheme` | `models.Optional[string]` | Optional | The default price point's pricing scheme. `null` if the Wallet Product has no default price point. |
| `TierCount` | `*int` | Optional | Number of price brackets on the default price point. A `per_unit` scheme always stores exactly one synthesized bracket, so this is `1` for `per_unit`, not `0`. |
| `EntryPrice` | `models.Optional[string]` | Optional | The unit price of the default price point's first bracket, as a decimal string. `null` if there are no brackets. |
| `FinalPrice` | `models.Optional[string]` | Optional | The unit price of the default price point's last bracket, as a decimal string. `null` if there are no brackets. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    walletProductPricePointSummary := models.WalletProductPricePointSummary{
        ArchivedCount:        models.ToPointer(244),
        PricingScheme:        models.NewOptional(models.ToPointer("pricing_scheme8")),
        TierCount:            models.ToPointer(30),
        EntryPrice:           models.NewOptional(models.ToPointer("entry_price8")),
        FinalPrice:           models.NewOptional(models.ToPointer("final_price4")),
    }

}
```

