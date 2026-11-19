
# Aggregated Entitlement

One entitlement in a subscriber's aggregated entitlements list. Entries are aggregated per feature key and periodicity window, not per feature key alone. A `usage_limit` feature granted with two different periodicities yields two entries sharing one `feature_key`. Use `periodicity_key` to identify an entry uniquely.

## Structure

`AggregatedEntitlement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FeatureKey` | `*string` | Optional | The feature's key, prefixed by kind: `feature.*` for `access_right`, `usage.*` for `usage_limit`, `service.*` for `service_right`. |
| `PeriodicityKey` | `*string` | Optional | Uniquely identifies this aggregated entry: the prefixed feature key, suffixed with `:{interval}:{unit}` when the entitlement has a periodicity window. Equal to `feature_key` when `periodicity` is `null`. |
| `Name` | `*string` | Optional | - |
| `Type` | [`*models.FeatureKind`](../../doc/models/feature-kind.md) | Optional | The behavior of a feature:<br><br>- `access_right`: a boolean entitlement. A subscriber either has access or does not.<br>- `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").<br>- `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit. |
| `Value` | [`*models.AggregatedEntitlementValue`](../../doc/models/containers/aggregated-entitlement-value.md) | Optional | This is a container for one-of cases. |
| `Enabled` | `*bool` | Optional | `true` only when the aggregated value is truthy for this feature's kind, and the subscription is in a live state (`active`, `trialing`, `assessing`, `past_due`, or `soft_failure`). `false` otherwise, including for `awaiting_signup`, canceled, expired, and on-hold subscriptions. Entitlements deliberately stay enabled through dunning. |
| `Periodicity` | [`models.Optional[models.AggregatedEntitlementPeriodicity]`](../../doc/models/aggregated-entitlement-periodicity.md) | Optional | - |
| `SourceProducts` | `[]string` | Optional | The names of the products/components contributing to this entitlement. For `access_right` features, only contributors that granted `true` are listed. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    aggregatedEntitlement := models.AggregatedEntitlement{
        FeatureKey:           models.ToPointer("feature.sso"),
        PeriodicityKey:       models.ToPointer("usage.api_calls:1:month"),
        Name:                 models.ToPointer("name2"),
        Type:                 models.ToPointer(models.FeatureKind_ACCESSRIGHT),
        Value:                models.ToPointer(models.AggregatedEntitlementValueContainer.FromBoolean(true)),
    }

}
```

