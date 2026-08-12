
# Apple Pay Payment Profile

## Structure

`ApplePayPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the Apple Pay payment profile. |
| `FirstName` | `*string` | Optional | The first name of the Apple Pay account holder |
| `LastName` | `*string` | Optional | The last name of the Apple Pay account holder |
| `CustomerId` | `*int` | Optional | The Chargify-assigned ID for the customer record to which the Apple Pay account belongs |
| `CurrentVault` | [`*models.ApplePayVault`](../../doc/models/apple-pay-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. |
| `VaultToken` | `*string` | Optional | The “token” provided by your vault storage for an already stored payment profile |
| `BillingAddress` | `models.Optional[string]` | Optional | The current billing street address for the Apple Pay account |
| `BillingCity` | `models.Optional[string]` | Optional | The current billing address city for the Apple Pay account |
| `BillingState` | `models.Optional[string]` | Optional | The current billing address state for the Apple Pay account |
| `BillingZip` | `models.Optional[string]` | Optional | The current billing address zip code for the Apple Pay account |
| `BillingCountry` | `models.Optional[string]` | Optional | The current billing address country for the Apple Pay account |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |
| `BillingAddress2` | `models.Optional[string]` | Optional | The current billing street address, second line, for the Apple Pay account |
| `PaymentType` | [`models.PaymentType`](../../doc/models/payment-type.md) | Required | **Default**: `"apple_pay"` |
| `SiteGatewaySettingId` | `models.Optional[int]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was created |
| `UpdatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was last updated |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    applePayPaymentProfile := models.ApplePayPaymentProfile{
        Id:                   models.ToPointer(252),
        FirstName:            models.ToPointer("first_name0"),
        LastName:             models.ToPointer("last_name8"),
        CustomerId:           models.ToPointer(34),
        CurrentVault:         models.ToPointer(models.ApplePayVault_BRAINTREEBLUE),
        PaymentType:          models.PaymentType_APPLEPAY,
    }

}
```

