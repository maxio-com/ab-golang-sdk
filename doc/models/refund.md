
# Refund

## Structure

`Refund`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `*string` | Optional | The amount to be refunded in decimal format as a string. Example: "10.50". Must not exceed the remaining refundable balance of the payment. |
| `Memo` | `*string` | Optional | A description that will be attached to the refund |
| `PaymentId` | `*int` | Optional | The ID of the payment to be refunded |
| `External` | `*bool` | Optional | Flag that marks refund as external (no money is returned to the customer). Defaults to `false`. |
| `ApplyCredit` | `*bool` | Optional | If set to true, creates credit and applies it to an invoice. Defaults to `false`. |
| `VoidInvoice` | `*bool` | Optional | If `apply_credit` set to false and refunding full amount, if `void_invoice` set to true, invoice will be voided after refund. Defaults to `false`. |
| `SegmentUids` | `*interface{}` | Optional | An array of segment uids to refund or the string 'all' to indicate that all segments should be refunded |

## Example (as JSON)

```json
{
  "amount": "amount8",
  "memo": "memo0",
  "payment_id": 130,
  "external": false,
  "apply_credit": false
}
```

