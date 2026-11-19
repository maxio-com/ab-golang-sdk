
# Create Customer

## Structure

`CreateCustomer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `string` | Required | - |
| `LastName` | `string` | Required | - |
| `Email` | `string` | Required | - |
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
| `VatCountry` | `*string` | Optional | The two-letter ISO 3166-1 country code that qualifies the customer's tax ID. Required when `entity_identifier_kind` is `vat_eu` or `national_tax`, and used to derive the kind when only the legacy `vat_number` is sent. |
| `EntityIdentifierKind` | [`*models.EntityIdentifierKind`](../../doc/models/entity-identifier-kind.md) | Optional | The kind of tax or business identifier held by the customer:<br><br>- `vat_eu`: an EU VAT number. Requires `vat_country` to be an EU member state code or `GB`.<br>- `national_tax`: a national tax ID registered outside the EU. Requires `vat_country` to be one of `AL`, `AM`, `AR`, `AU`, `BR`, `CA`, `CH`, `DZ`, `IN`, `MX`, `NO`, `NZ`, or `ZA`.<br>- `company_reg`: a company registration number, such as a French SIREN. No `vat_country` is required.<br>- `gln`: a Global Location Number. The value must be 13 digits.<br>- `duns`: a D-U-N-S Number. The value must be 9 digits.<br>- `lei`: a Legal Entity Identifier. The value must be 20 characters: 18 letters or digits followed by 2 digits.<br><br>A customer holds one identifier at a time. Saving an identifier of a different kind replaces the existing one. |
| `EntityIdentifierValue` | `*string` | Optional | The customer's tax or business identifier, sent together with `entity_identifier_kind`. Advanced Billing trims surrounding whitespace and stores the value in uppercase. |
| `TaxExempt` | `*bool` | Optional | - |
| `Surcharging` | `*bool` | Optional | Whether surcharging is enabled for the customer. Defaults to `true` when omitted. Only applied on sites where surcharging control is enabled. |
| `TaxExemptReason` | `*string` | Optional | - |
| `ParentId` | `models.Optional[int]` | Optional | The parent ID in Chargify if applicable. Parent is another Customer object. |
| `SalesforceId` | `models.Optional[string]` | Optional | The Salesforce ID of the customer |
| `BrandingThemeId` | `models.Optional[int]` | Optional | The ID of the Branding Theme assigned to this customer as the customer's default Branding Theme. This customer-level Branding Theme is used when a subscription does not have its own subscription-level Branding Theme. Available only when Branding Themes are enabled for the site. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createCustomer := models.CreateCustomer{
        FirstName:             "first_name0",
        LastName:              "last_name8",
        Email:                 "email6",
        CcEmails:              models.ToPointer("cc_emails0"),
        Organization:          models.ToPointer("organization6"),
        Reference:             models.ToPointer("reference4"),
        Address:               models.ToPointer("address6"),
        Address2:              models.ToPointer("address_24"),
    }

}
```

