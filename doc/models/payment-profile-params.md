
# Payment Profile Params

PCI-safe cardholder fields only. Full card numbers, CVV, and billing address are never included.

## Structure

`PaymentProfileParams`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `CardType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "first_name": "first_name2",
  "last_name": "last_name0",
  "card_type": "card_type8"
}
```

