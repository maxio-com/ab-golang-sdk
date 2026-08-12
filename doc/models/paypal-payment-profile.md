
# Paypal Payment Profile

## Structure

`PaypalPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the stored PayPal payment profile. |
| `FirstName` | `*string` | Optional | The first name of the PayPal account holder |
| `LastName` | `*string` | Optional | The last name of the PayPal account holder |
| `CustomerId` | `*int` | Optional | The Chargify-assigned id for the customer record to which the PayPal account belongs |
| `CurrentVault` | [`*models.PayPalVault`](../../doc/models/pay-pal-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. |
| `VaultToken` | `*string` | Optional | The “token” provided by your vault storage for an already stored payment profile |
| `BillingAddress` | `models.Optional[string]` | Optional | The current billing street address for the PayPal account |
| `BillingCity` | `models.Optional[string]` | Optional | The current billing address city for the PayPal account |
| `BillingState` | `models.Optional[string]` | Optional | The current billing address state for the PayPal account |
| `BillingZip` | `models.Optional[string]` | Optional | The current billing address zip code for the PayPal account |
| `BillingCountry` | `models.Optional[string]` | Optional | The current billing address country for the PayPal account |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |
| `BillingAddress2` | `models.Optional[string]` | Optional | The current billing street address, second line, for the PayPal account |
| `PaymentType` | [`models.PaymentType`](../../doc/models/payment-type.md) | Required | **Default**: `"paypal_account"` |
| `SiteGatewaySettingId` | `models.Optional[int]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `PaypalEmail` | `*string` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was created |
| `UpdatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was last updated |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paypalPaymentProfile := models.PaypalPaymentProfile{
        Id:                   models.ToPointer(214),
        FirstName:            models.ToPointer("first_name0"),
        LastName:             models.ToPointer("last_name8"),
        CustomerId:           models.ToPointer(252),
        CurrentVault:         models.ToPointer(models.PayPalVault_MODUSLINK),
        PaymentType:          models.PaymentType_PAYPALACCOUNT,
    }

}
```

