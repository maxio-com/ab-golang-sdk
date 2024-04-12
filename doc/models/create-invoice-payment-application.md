
# Create Invoice Payment Application

## Structure

`CreateInvoicePaymentApplication`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InvoiceUid` | `string` | Required | Unique identifier for the invoice. It has the prefix "inv_" followed by alphanumeric characters. |
| `Amount` | `string` | Required | Dollar amount of the invoice payment (eg. "10.50" => $10.50). |

## Example (as JSON)

```json
{
  "invoice_uid": "invoice_uid4",
  "amount": "amount6"
}
```

