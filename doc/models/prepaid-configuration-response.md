
# Prepaid Configuration Response

## Structure

`PrepaidConfigurationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidConfiguration` | [`models.PrepaidConfiguration`](../../doc/models/prepaid-configuration.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidConfigurationResponse := models.PrepaidConfigurationResponse{
        PrepaidConfiguration: models.PrepaidConfiguration{
            Id:                              models.ToPointer(142),
            InitialFundingAmountInCents:     models.ToPointer(int64(74)),
            ReplenishToAmountInCents:        models.ToPointer(int64(76)),
            AutoReplenish:                   models.ToPointer(false),
            ReplenishThresholdAmountInCents: models.ToPointer(int64(20)),
        },
    }

}
```

