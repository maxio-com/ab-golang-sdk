
# Upsert Prepaid Configuration

## Structure

`UpsertPrepaidConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InitialFundingAmountInCents` | `*int64` | Optional | - |
| `ReplenishToAmountInCents` | `*int64` | Optional | - |
| `AutoReplenish` | `*bool` | Optional | - |
| `ReplenishThresholdAmountInCents` | `*int64` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    upsertPrepaidConfiguration := models.UpsertPrepaidConfiguration{
        InitialFundingAmountInCents:     models.ToPointer(int64(244)),
        ReplenishToAmountInCents:        models.ToPointer(int64(246)),
        AutoReplenish:                   models.ToPointer(false),
        ReplenishThresholdAmountInCents: models.ToPointer(int64(190)),
    }

}
```

