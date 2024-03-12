
# Coupon Usage

## Structure

`CouponUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify id of the product |
| `Name` | `*string` | Optional | Name of the product |
| `Signups` | `*int` | Optional | Number of times the coupon has been applied |
| `Savings` | `models.Optional[int]` | Optional | Dollar amount of customer savings as a result of the coupon. |
| `SavingsInCents` | `models.Optional[int64]` | Optional | Dollar amount of customer savings as a result of the coupon. |
| `Revenue` | `models.Optional[int]` | Optional | Total revenue of the all subscriptions that have received a discount from this coupon. |
| `RevenueInCents` | `*int64` | Optional | Total revenue of the all subscriptions that have received a discount from this coupon. |

## Example (as JSON)

```json
{
  "id": 14,
  "name": "name0",
  "signups": 34,
  "savings": 52,
  "savings_in_cents": 138
}
```

