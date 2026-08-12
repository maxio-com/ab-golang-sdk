
# MRR Response

## Structure

`MRRResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mrr` | [`models.MRR`](../../doc/models/mrr.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    mrrResponse := models.MRRResponse{
        Mrr:                  models.MRR{
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
        },
    }

}
```

