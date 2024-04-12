
# Record Payment Response

## Structure

`RecordPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaidInvoices` | [`[]models.PaidInvoice`](../../doc/models/paid-invoice.md) | Optional | - |
| `Prepayment` | [`models.Optional[models.InvoicePrePayment]`](../../doc/models/invoice-pre-payment.md) | Optional | - |

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
    "subscription_id": 148,
    "amount_in_cents": 124,
    "ending_balance_in_cents": 164
  }
}
```

