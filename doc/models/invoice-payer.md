
# Invoice Payer

## Structure

`InvoicePayer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyId` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `models.Optional[string]` | Optional | - |
| `Email` | `*string` | Optional | - |
| `VatNumber` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "chargify_id": 46,
  "first_name": "first_name4",
  "last_name": "last_name2",
  "organization": "organization8",
  "email": "email2"
}
```

