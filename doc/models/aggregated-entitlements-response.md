
# Aggregated Entitlements Response

## Structure

`AggregatedEntitlementsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `int` | Required | - |
| `CustomerId` | `int` | Required | - |
| `Status` | `string` | Required | The subscription's current state, e.g. `active`, `trialing`, `canceled`. |
| `Entitlements` | [`[]models.AggregatedEntitlement`](../../doc/models/aggregated-entitlement.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    aggregatedEntitlementsResponse := models.AggregatedEntitlementsResponse{
        SubscriptionId:       114,
        CustomerId:           42,
        Status:               "status6",
        Entitlements:         []models.AggregatedEntitlement{
            models.AggregatedEntitlement{
                FeatureKey:           models.ToPointer("feature.sso"),
                PeriodicityKey:       models.ToPointer("usage.api_calls:1:month"),
                Name:                 models.ToPointer("name6"),
                Type:                 models.ToPointer(models.FeatureKind_SERVICERIGHT),
                Value:                models.ToPointer(models.AggregatedEntitlementValueContainer.FromBoolean(true)),
            },
        },
    }

}
```

