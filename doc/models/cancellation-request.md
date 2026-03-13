
# Cancellation Request

## Structure

`CancellationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`models.CancellationOptions`](../../doc/models/cancellation-options.md) | Required | - |

## Example (as JSON)

```json
{
  "subscription": {
    "cancellation_message": "cancellation_message2",
    "reason_code": "reason_code8",
    "cancel_at_end_of_period": false,
    "scheduled_cancellation_at": "2016-03-13T12:52:32.123Z",
    "refund_prepayment_account_balance": false
  }
}
```

