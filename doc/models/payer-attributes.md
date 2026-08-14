
# Payer Attributes

## Structure

`PayerAttributes`

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
| `Locale` | `*string` | Optional | - |
| `VatNumber` | `*string` | Optional | - |
| `TaxExempt` | `*bool` | Optional | - |
| `TaxExemptReason` | `*string` | Optional | - |
| `Metafields` | `map[string]string` | Optional | (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    payerAttributes := models.PayerAttributes{
        FirstName:            models.ToPointer("first_name2"),
        LastName:             models.ToPointer("last_name0"),
        Email:                models.ToPointer("email4"),
        CcEmails:             models.ToPointer("cc_emails8"),
        Organization:         models.ToPointer("organization4"),
        Metafields:           map[string]string{
            "custom_field_name_1": "custom_field_value_1",
            "custom_field_name_2": "custom_field_value_2",
        },
    }

}
```

