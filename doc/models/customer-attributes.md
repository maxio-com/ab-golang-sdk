
# Customer Attributes

## Structure

`CustomerAttributes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | The first name of the customer. Required when creating a customer via attributes. |
| `LastName` | `*string` | Optional | The last name of the customer. Required when creating a customer via attributes. |
| `Email` | `*string` | Optional | The email address of the customer. Required when creating a customer via attributes. |
| `CcEmails` | `*string` | Optional | (Optional) A list of emails that should be cc’d on all customer communications. |
| `Organization` | `*string` | Optional | (Optional) The organization/company of the customer. |
| `Reference` | `*string` | Optional | (Optional) A customer “reference”, or unique identifier from your app, stored in Chargify. Can be used so that you may reference your customer’s within Chargify using the same unique value you use in your application. |
| `Address` | `*string` | Optional | (Optional) The customer’s shipping street address (e.g., “123 Main St.”). |
| `Address2` | `models.Optional[string]` | Optional | (Optional) Second line of the customer’s shipping address e.g., “Apt. 100” |
| `City` | `*string` | Optional | (Optional) The customer’s shipping address city (e.g., “Boston”). |
| `State` | `*string` | Optional | “(Optional) The customer’s shipping address state (e.g., “MA”). This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes.” |
| `Zip` | `*string` | Optional | (Optional) The customer’s shipping address zip code (e.g., “12345”). |
| `Country` | `*string` | Optional | “(Optional) The customer shipping address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (e.g., “US”).” |
| `Phone` | `*string` | Optional | (Optional) The phone number of the customer. |
| `Verified` | `*bool` | Optional | - |
| `TaxExempt` | `*bool` | Optional | (Optional) The tax_exempt status of the customer. Acceptable values are true or 1 for true and false or 0 for false. |
| `Surcharging` | `*bool` | Optional | (Optional) Whether surcharging is enabled for the customer. Defaults to `true` when omitted. Only applied on sites where surcharging control is enabled. |
| `VatNumber` | `*string` | Optional | (Optional) Supplying the VAT number allows EU customers to opt-out of the Value Added Tax assuming the merchant address and customer billing address are not within the same EU country. It’s important to omit the country code from the VAT number upon entry. Otherwise, taxes will be assessed upon the purchase. |
| `Metafields` | `map[string]string` | Optional | (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet. |
| `ParentId` | `models.Optional[int]` | Optional | The parent ID in Chargify if applicable. Parent is another Customer object. |
| `SalesforceId` | `models.Optional[string]` | Optional | (Optional) The Salesforce ID of the customer. |
| `DefaultAutoRenewalProfileId` | `models.Optional[int]` | Optional | (Optional) The default auto-renewal profile ID for the customer |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerAttributes := models.CustomerAttributes{
        FirstName:                   models.ToPointer("first_name2"),
        LastName:                    models.ToPointer("last_name0"),
        Email:                       models.ToPointer("email4"),
        CcEmails:                    models.ToPointer("cc_emails2"),
        Organization:                models.ToPointer("organization6"),
        Metafields:                  map[string]string{
            "custom_field_name_1": "custom_field_value_1",
            "custom_field_name_2": "custom_field_value_2",
        },
    }

}
```

