
# Payment Method Nested Data

A nested data structure detailing the method of payment

## Structure

`PaymentMethodNestedData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | `*string` | Optional | **Default**: `"apple_pay"` |
| `MaskedAccountNumber` | `*string` | Optional | - |
| `MaskedRoutingNumber` | `*string` | Optional | - |
| `CardBrand` | `*string` | Optional | - |
| `CardExpiration` | `*string` | Optional | - |
| `LastFour` | `Optional[string]` | Optional | - |
| `MaskedCardNumber` | `*string` | Optional | - |
| `Details` | `*string` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `Email` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "type": "apple_pay",
  "masked_account_number": "masked_account_number8",
  "masked_routing_number": "masked_routing_number8",
  "card_brand": "card_brand4",
  "card_expiration": "card_expiration2"
}
```

