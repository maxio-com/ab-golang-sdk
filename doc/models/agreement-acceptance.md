
# Agreement Acceptance

Required when creating a subscription with Maxio Payments.

## Structure

`AgreementAcceptance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IpAddress` | `*string` | Optional | Required when providing agreement acceptance params. |
| `TermsUrl` | `*string` | Optional | Required when creating a subscription with Maxio Payments. Either terms_url or privacy_policy_url is required when providing agreement_acceptance params. |
| `PrivacyPolicyUrl` | `*string` | Optional | - |
| `ReturnRefundPolicyUrl` | `*string` | Optional | - |
| `DeliveryPolicyUrl` | `*string` | Optional | - |
| `SecureCheckoutPolicyUrl` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    agreementAcceptance := models.AgreementAcceptance{
        IpAddress:               models.ToPointer("ip_address8"),
        TermsUrl:                models.ToPointer("terms_url6"),
        PrivacyPolicyUrl:        models.ToPointer("privacy_policy_url4"),
        ReturnRefundPolicyUrl:   models.ToPointer("return_refund_policy_url0"),
        DeliveryPolicyUrl:       models.ToPointer("delivery_policy_url4"),
    }

}
```

