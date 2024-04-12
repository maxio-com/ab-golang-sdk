
# Invoice Customer

Information about the customer who is owner or recipient the invoiced subscription.

## Structure

`InvoiceCustomer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyId` | `models.Optional[int]` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Organization` | `models.Optional[string]` | Optional | - |
| `Email` | `*string` | Optional | - |
| `VatNumber` | `models.Optional[string]` | Optional | - |
| `Reference` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "chargify_id": 236,
  "first_name": "first_name0",
  "last_name": "last_name8",
  "organization": "organization4",
  "email": "email6"
}
```

