
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
| `Surcharging` | `*bool` | Optional | Whether surcharging is enabled for the customer. Only applied on sites where surcharging control is enabled. |
| `TaxExemptReason` | `*string` | Optional | - |
| `ParentId` | `models.Optional[int]` | Optional | - |
| `Verified` | `models.Optional[bool]` | Optional | Is the customer verified to use ACH as a payment method. Available only on the Authorize.Net gateway. |
| `SalesforceId` | `models.Optional[string]` | Optional | The Salesforce ID of the customer |
| `BrandingThemeId` | `models.Optional[int]` | Optional | The ID of the Branding Theme assigned to this customer as the customer's default Branding Theme. This customer-level Branding Theme is used when a subscription does not have its own subscription-level Branding Theme. Available only when Branding Themes are enabled for the site. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateCustomer := models.UpdateCustomer{
        FirstName:            models.ToPointer("first_name2"),
        LastName:             models.ToPointer("last_name0"),
        Email:                models.ToPointer("email4"),
        CcEmails:             models.ToPointer("cc_emails8"),
        Organization:         models.ToPointer("organization6"),
    }

}
```

