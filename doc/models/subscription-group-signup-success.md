
# Subscription Group Signup Success

## Structure

`SubscriptionGroupSignupSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionGroup` | [`models.SubscriptionGroupSignupSuccessData`](../../doc/models/subscription-group-signup-success-data.md) | Required | - |
| `Customer` | [`models.Customer`](../../doc/models/customer.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription_group": {
    "uid": "uid8",
    "scheme": 200,
    "customer_id": 220,
    "payment_profile_id": 128,
    "subscription_ids": [
      74,
      75
    ],
    "primary_subscription_id": 148,
    "next_assessment_at": "2016-03-13T12:52:32.123Z",
    "state": "state6",
    "cancel_at_end_of_period": false
  },
  "customer": {
    "first_name": "first_name0",
    "last_name": "last_name8",
    "email": "email6",
    "cc_emails": "cc_emails0",
    "organization": "organization6"
  }
}
```

