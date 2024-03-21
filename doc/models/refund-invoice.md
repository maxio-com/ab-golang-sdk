
# Refund Invoice

Refund an invoice or a segment of a consolidated invoice.

## Structure

`RefundInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `string` | Required | The amount to be refunded in decimal format as a string. Example: "10.50". Must not exceed the remaining refundable balance of the payment. |
| `Memo` | `string` | Required | A description that will be attached to the refund |
| `PaymentId` | `int` | Required | The ID of the payment to be refunded |
| `External` | `*bool` | Optional | Flag that marks refund as external (no money is returned to the customer). Defaults to `false`. |
| `ApplyCredit` | `*bool` | Optional | If set to true, creates credit and applies it to an invoice. Defaults to `false`. |
| `VoidInvoice` | `*bool` | Optional | If `apply_credit` set to false and refunding full amount, if `void_invoice` set to true, invoice will be voided after refund. Defaults to `false`. |

## Example (as JSON)

```json
{
  "amount": "amount8",
  "memo": "memo0",
  "payment_id": 0,
  "external": false,
  "apply_credit": false,
  "void_invoice": false
}
```

