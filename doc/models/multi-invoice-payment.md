
# Multi Invoice Payment

## Structure

`MultiInvoicePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `*int` | Optional | The numeric ID of the transaction. |
| `TotalAmount` | `*string` | Optional | Dollar amount of the sum of the paid invoices. |
| `CurrencyCode` | `*string` | Optional | The ISO 4217 currency code (3 character string) representing the currency of invoice transaction. |
| `Applications` | [`[]models.InvoicePaymentApplication`](invoice-payment-application.md) | Optional | - |

## Example (as JSON)

```json
{
  "transaction_id": 144,
  "total_amount": "total_amount2",
  "currency_code": "currency_code2",
  "applications": [
    {
      "invoice_uid": "invoice_uid8",
      "application_uid": "application_uid8",
      "applied_amount": "applied_amount0"
    }
  ]
}
```

