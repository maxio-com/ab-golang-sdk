
# Cancellation Options

## Structure

`CancellationOptions`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CancellationMessage` | `*string` | Optional | An indication as to why the subscription is being canceled. For your internal use. |
| `ReasonCode` | `*string` | Optional | The reason code associated with the cancellation. Use the [List Reason Codes](../../doc/controllers/reason-codes.md#list-reason-codes) endpoint to retrieve the reason codes associated with your site. |
| `CancelAtEndOfPeriod` | `*bool` | Optional | When true, the subscription is cancelled at the current period end instead of immediately. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site. |
| `ScheduledCancellationAt` | `models.Optional[time.Time]` | Optional | Schedules the cancellation on the provided date. This is option is not applicable for prepaid subscriptions. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site. |
| `RefundPrepaymentAccountBalance` | `*bool` | Optional | Applies to prepaid subscriptions. When true, which is the default, the remaining prepaid balance is refunded as part of cancellation processing. When false, prepaid balance is not refunded as part of cancellation processing. To use this option, the Schedule Subscription Cancellation feature must be enabled on your site. |

## Example (as JSON)

```json
{
  "cancellation_message": "cancellation_message0",
  "reason_code": "reason_code6",
  "cancel_at_end_of_period": false,
  "scheduled_cancellation_at": "2016-03-13T12:52:32.123Z",
  "refund_prepayment_account_balance": false
}
```

