
# Dunner Data

## Structure

`DunnerData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `State` | `string` | Required | - |
| `SubscriptionId` | `int` | Required | - |
| `RevenueAtRiskInCents` | `int64` | Required | - |
| `CreatedAt` | `time.Time` | Required | - |
| `Attempts` | `int` | Required | - |
| `LastAttemptedAt` | `time.Time` | Required | - |

## Example (as JSON)

```json
{
  "state": "state4",
  "subscription_id": 126,
  "revenue_at_risk_in_cents": 30,
  "created_at": "2016-03-13T12:52:32.123Z",
  "attempts": 110,
  "last_attempted_at": "2016-03-13T12:52:32.123Z"
}
```

