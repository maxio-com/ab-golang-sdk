
# Prepaid Configuration

## Structure

`PrepaidConfiguration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
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
    prepaidConfiguration := models.PrepaidConfiguration{
        Id:                              models.ToPointer(146),
        InitialFundingAmountInCents:     models.ToPointer(int64(78)),
        ReplenishToAmountInCents:        models.ToPointer(int64(80)),
        AutoReplenish:                   models.ToPointer(false),
        ReplenishThresholdAmountInCents: models.ToPointer(int64(232)),
    }

}
```

