
# Invoice Event Payment

A nested data structure detailing the method of payment

## Structure

`InvoiceEventPayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | `*string` | Optional | - |
| `MaskedAccountNumber` | `*string` | Optional | - |
| `MaskedRoutingNumber` | `*string` | Optional | - |
| `CardBrand` | `*string` | Optional | - |
| `CardExpiration` | `*string` | Optional | - |
| `LastFour` | `models.Optional[string]` | Optional | - |
| `MaskedCardNumber` | `*string` | Optional | - |
| `Details` | `models.Optional[string]` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `Memo` | `models.Optional[string]` | Optional | - |
| `Email` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "Invoice Event Payment",
  "masked_account_number": "masked_account_number8",
  "masked_routing_number": "masked_routing_number8",
  "card_brand": "card_brand4",
  "card_expiration": "card_expiration2"
}
```

