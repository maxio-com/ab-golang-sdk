
# Prepayment

## Structure

`Prepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | - |
| `SubscriptionId` | `int` | Required | - |
| `AmountInCents` | `int64` | Required | - |
| `RemainingAmountInCents` | `int64` | Required | - |
| `RefundedAmountInCents` | `*int64` | Optional | - |
| `Details` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `External` | `bool` | Required | - |
| `Memo` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `PaymentType` | [`*models.PrepaymentMethod`](../../doc/models/prepayment-method.md) | Optional | The payment type of the prepayment. |
| `CreatedAt` | `time.Time` | Required | - |

## Example (as JSON)

```json
{
  "id": 50,
  "subscription_id": 160,
  "amount_in_cents": 120,
  "remaining_amount_in_cents": 194,
  "refunded_amount_in_cents": 144,
  "details": "details4",
  "external": false,
  "memo": "memo8",
  "payment_type": "cash",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

