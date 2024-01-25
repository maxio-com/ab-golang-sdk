
# Subscription Group Prepayment Request

## Structure

`SubscriptionGroupPrepaymentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Prepayment` | [`models.SubscriptionGroupPrepayment`](../../doc/models/subscription-group-prepayment.md) | Required | - |

## Example (as JSON)

```json
{
  "prepayment": {
    "amount": 136,
    "details": "details8",
    "memo": "memo2",
    "method": "paypal_account"
  }
}
```

