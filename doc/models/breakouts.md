
# Breakouts

## Structure

`Breakouts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PlanAmountInCents` | `*int64` | Optional | - |
| `PlanAmountFormatted` | `*string` | Optional | - |
| `UsageAmountInCents` | `*int64` | Optional | - |
| `UsageAmountFormatted` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    breakouts := models.Breakouts{
        PlanAmountInCents:    models.ToPointer(int64(254)),
        PlanAmountFormatted:  models.ToPointer("plan_amount_formatted0"),
        UsageAmountInCents:   models.ToPointer(int64(106)),
        UsageAmountFormatted: models.ToPointer("usage_amount_formatted8"),
    }

}
```

