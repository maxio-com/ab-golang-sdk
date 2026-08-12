
# Upsert Prepaid Configuration Request

## Structure

`UpsertPrepaidConfigurationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidConfiguration` | [`models.UpsertPrepaidConfiguration`](../../doc/models/upsert-prepaid-configuration.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    upsertPrepaidConfigurationRequest := models.UpsertPrepaidConfigurationRequest{
        PrepaidConfiguration: models.UpsertPrepaidConfiguration{
            InitialFundingAmountInCents:     models.ToPointer(int64(74)),
            ReplenishToAmountInCents:        models.ToPointer(int64(76)),
            AutoReplenish:                   models.ToPointer(false),
            ReplenishThresholdAmountInCents: models.ToPointer(int64(20)),
        },
    }

}
```

