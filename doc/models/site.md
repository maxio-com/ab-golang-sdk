
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
| `ScheduleSubscriptionCancellationEnabled` | `*bool` | Optional | - |
| `CustomerHierarchyEnabled` | `*bool` | Optional | - |
| `WhopaysEnabled` | `*bool` | Optional | - |
| `WhopaysDefaultPayer` | `*string` | Optional | - |
| `AllocationSettings` | [`*models.AllocationSettings`](../../doc/models/allocation-settings.md) | Optional | - |
| `DefaultPaymentCollectionMethod` | `*string` | Optional | - |
| `OrganizationAddress` | [`*models.OrganizationAddress`](../../doc/models/organization-address.md) | Optional | - |
| `TaxConfiguration` | [`*models.TaxConfiguration`](../../doc/models/tax-configuration.md) | Optional | - |
| `NetTerms` | [`*models.NetTerms`](../../doc/models/net-terms.md) | Optional | - |
| `MultiFrequencyEnabled` | `*bool` | Optional | Whether the site has the multi-frequency billing feature enabled. Only present when relationship invoicing is active. |
| `AutoRenewalsEnabled` | `*bool` | Optional | Whether the auto-renewals feature is enabled for this site. |
| `PortalEnabled` | `*bool` | Optional | Whether the Billing Portal is enabled for this site. |
| `Test` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    site := models.Site{
        Id:                                      models.ToPointer(64),
        Name:                                    models.ToPointer("name4"),
        Subdomain:                               models.ToPointer("subdomain0"),
        Currency:                                models.ToPointer("currency4"),
        SellerId:                                models.ToPointer(228),
    }

}
```

