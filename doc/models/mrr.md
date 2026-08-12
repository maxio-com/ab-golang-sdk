
# MRR

## Structure

`MRR`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmountInCents` | `*int64` | Optional | - |
| `AmountFormatted` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `CurrencySymbol` | `*string` | Optional | - |
| `Breakouts` | [`*models.Breakouts`](../../doc/models/breakouts.md) | Optional | - |
| `AtTime` | `*time.Time` | Optional | ISO8601 timestamp |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    mrr := models.MRR{
        AmountInCents:        models.ToPointer(int64(198)),
        AmountFormatted:      models.ToPointer("amount_formatted6"),
        Currency:             models.ToPointer("currency4"),
        CurrencySymbol:       models.ToPointer("currency_symbol2"),
        Breakouts:            models.ToPointer(models.Breakouts{
            PlanAmountInCents:    models.ToPointer(int64(254)),
            PlanAmountFormatted:  models.ToPointer("plan_amount_formatted0"),
            UsageAmountInCents:   models.ToPointer(int64(106)),
            UsageAmountFormatted: models.ToPointer("usage_amount_formatted8"),
        }),
    }

}
```

