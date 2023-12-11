
# Full Subscription Group Response

## Structure

`FullSubscriptionGroupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Scheme` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `PaymentProfileId` | `*int` | Optional | - |
| `SubscriptionIds` | `[]int` | Optional | - |
| `PrimarySubscriptionId` | `*int` | Optional | - |
| `NextAssessmentAt` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `CancelAtEndOfPeriod` | `*bool` | Optional | - |
| `CurrentBillingAmountInCents` | `*int64` | Optional | - |
| `Customer` | [`*models.SubscriptionGroupCustomer`](subscription-group-customer.md) | Optional | - |
| `AccountBalances` | [`*models.SubscriptionGroupBalances`](subscription-group-balances.md) | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid8",
  "scheme": 90,
  "customer_id": 110,
  "payment_profile_id": 18,
  "subscription_ids": [
    220,
    221,
    222
  ]
}
```

