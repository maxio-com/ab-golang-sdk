
# Site

## Structure

`Site`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Subdomain` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `SellerId` | `*int` | Optional | - |
| `NonPrimaryCurrencies` | `[]string` | Optional | - |
| `RelationshipInvoicingEnabled` | `*bool` | Optional | - |
| `CustomerHierarchyEnabled` | `*bool` | Optional | - |
| `WhopaysEnabled` | `*bool` | Optional | - |
| `WhopaysDefaultPayer` | `*string` | Optional | - |
| `AllocationSettings` | [`*models.AllocationSettings`](allocation-settings.md) | Optional | - |
| `DefaultPaymentCollectionMethod` | `*string` | Optional | - |
| `OrganizationAddress` | [`*models.OrganizationAddress`](organization-address.md) | Optional | - |
| `TaxConfiguration` | [`*models.TaxConfiguration`](tax-configuration.md) | Optional | - |
| `NetTerms` | [`*models.NetTerms`](net-terms.md) | Optional | - |
| `Test` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "id": 34,
  "name": "name0",
  "subdomain": "subdomain4",
  "currency": "currency0",
  "seller_id": 198
}
```

