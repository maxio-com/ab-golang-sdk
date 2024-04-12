
# Subscription Group Signup Success Data

## Structure

`SubscriptionGroupSignupSuccessData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `string` | Required | - |
| `Scheme` | `int` | Required | - |
| `CustomerId` | `int` | Required | - |
| `PaymentProfileId` | `int` | Required | - |
| `SubscriptionIds` | `[]int` | Required | - |
| `PrimarySubscriptionId` | `int` | Required | - |
| `NextAssessmentAt` | `time.Time` | Required | - |
| `State` | `string` | Required | - |
| `CancelAtEndOfPeriod` | `bool` | Required | - |

## Example (as JSON)

```json
{
  "uid": "uid8",
  "scheme": 154,
  "customer_id": 174,
  "payment_profile_id": 82,
  "subscription_ids": [
    28,
    29
  ],
  "primary_subscription_id": 102,
  "next_assessment_at": "2016-03-13T12:52:32.123Z",
  "state": "state4",
  "cancel_at_end_of_period": false
}
```

