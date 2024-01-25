
# Create Prepayment

## Structure

`CreatePrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `float64` | Required | - |
| `Details` | `string` | Required | - |
| `Memo` | `string` | Required | - |
| `Method` | [`models.PrepaymentMethod`](../../doc/models/prepayment-method.md) | Required | :- When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions. |
| `PaymentProfileId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "amount": 23.92,
  "details": "details6",
  "memo": "memo0",
  "method": "paypal_account",
  "payment_profile_id": 240
}
```

