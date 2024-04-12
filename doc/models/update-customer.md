
# Update Customer

## Structure

`UpdateCustomer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `Email` | `*string` | Optional | - |
| `CcEmails` | `*string` | Optional | - |
| `Organization` | `*string` | Optional | - |
| `Reference` | `*string` | Optional | - |
| `Address` | `*string` | Optional | - |
| `Address2` | `*string` | Optional | - |
| `City` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `Zip` | `*string` | Optional | - |
| `Country` | `*string` | Optional | - |
| `Phone` | `*string` | Optional | - |
| `Locale` | `*string` | Optional | Set a specific language on a customer record. |
| `VatNumber` | `*string` | Optional | - |
| `TaxExempt` | `*bool` | Optional | - |
| `TaxExemptReason` | `*string` | Optional | - |
| `ParentId` | `models.Optional[int]` | Optional | - |
| `Verified` | `models.Optional[bool]` | Optional | Is the customer verified to use ACH as a payment method. Available only on Authorize.Net gateway |

## Example (as JSON)

```json
{
  "first_name": "first_name2",
  "last_name": "last_name0",
  "email": "email4",
  "cc_emails": "cc_emails8",
  "organization": "organization6"
}
```

