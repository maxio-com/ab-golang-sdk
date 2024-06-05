
# Historic Usage

Optional for Event Based Components. If the `include=historic_usages` query param is provided, the last ten billing periods will be returned.

## Structure

`HistoricUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalUsageQuantity` | `*float64` | Optional | Total usage of a component for billing period |
| `BillingPeriodStartsAt` | `*time.Time` | Optional | Start date of billing period |
| `BillingPeriodEndsAt` | `*time.Time` | Optional | End date of billing period |

## Example (as JSON)

```json
{
  "total_usage_quantity": 26.6,
  "billing_period_starts_at": "2016-03-13T12:52:32.123Z",
  "billing_period_ends_at": "2016-03-13T12:52:32.123Z"
}
```

