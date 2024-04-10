
# Subscription Included Coupon

## Structure

`SubscriptionIncludedCoupon`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `*string` | Optional | - |
| `UseCount` | `*int` | Optional | - |
| `UsesAllowed` | `*int` | Optional | - |
| `ExpiresAt` | `models.Optional[string]` | Optional | - |
| `Recurring` | `*bool` | Optional | - |
| `AmountInCents` | `models.Optional[int64]` | Optional | **Constraints**: `>= 0` |
| `Percentage` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "code": "\"ABCD_10\"",
  "use_count": 2,
  "uses_allowed": 10,
  "expires_at": "\"2023-07-13T05:18:58-04:00\"",
  "amount_in_cents": 1000,
  "percentage": "\"15.0\"",
  "recurring": false
}
```

