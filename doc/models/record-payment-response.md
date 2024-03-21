
# Record Payment Response

## Structure

`RecordPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaidInvoices` | [`[]models.PaidInvoice`](../../doc/models/paid-invoice.md) | Optional | - |
| `Prepayment` | [`models.Optional[models.RecordPaymentResponsePrepayment]`](../../doc/models/containers/record-payment-response-prepayment.md) | Optional | This is a container for one-of cases. |

## Example (as JSON)

```json
{
  "paid_invoices": [
    {
      "invoice_id": "invoice_id8",
      "status": "draft",
      "due_amount": "due_amount0",
      "paid_amount": "paid_amount0"
    },
    {
      "invoice_id": "invoice_id8",
      "status": "draft",
      "due_amount": "due_amount0",
      "paid_amount": "paid_amount0"
    },
    {
      "invoice_id": "invoice_id8",
      "status": "draft",
      "due_amount": "due_amount0",
      "paid_amount": "paid_amount0"
    }
  ],
  "prepayment": {
    "subscription_id": 180,
    "amount_in_cents": 100,
    "ending_balance_in_cents": 60
  }
}
```

