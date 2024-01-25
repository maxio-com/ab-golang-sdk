
# Refund Prepayment

## Structure

`RefundPrepayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmountInCents` | `int64` | Required | `amount` is not required if you pass `amount_in_cents`. |
| `Amount` | `interface{}` | Required | `amount_in_cents` is not required if you pass `amount`. |
| `Memo` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `External` | `*bool` | Optional | Specify the type of refund you wish to initiate. When the prepayment is external, the `external` flag is optional. But if the prepayment was made through a payment profile, the `external` flag is required. |

## Example (as JSON)

```json
{
  "amount_in_cents": 110,
  "amount": {
    "key1": "val1",
    "key2": "val2"
  },
  "memo": "memo4",
  "external": false
}
```

