
# ACH Agreement

(Optional) If passed, the proof of the authorized ACH agreement terms will be persisted.

## Structure

`ACHAgreement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AgreementTerms` | `*string` | Optional | (Required when providing ACH agreement params) The ACH authorization agreement terms. |
| `AuthorizerFirstName` | `*string` | Optional | (Required when providing ACH agreement params) The first name of the person authorizing the ACH agreement. |
| `AuthorizerLastName` | `*string` | Optional | (Required when providing ACH agreement params) The last name of the person authorizing the ACH agreement. |
| `IpAddress` | `*string` | Optional | (Required when providing ACH agreement params) The IP address of the person authorizing the ACH agreement. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    achAgreement := models.ACHAgreement{
        AgreementTerms:       models.ToPointer("agreement_terms8"),
        AuthorizerFirstName:  models.ToPointer("authorizer_first_name8"),
        AuthorizerLastName:   models.ToPointer("authorizer_last_name6"),
        IpAddress:            models.ToPointer("ip_address6"),
    }

}
```

