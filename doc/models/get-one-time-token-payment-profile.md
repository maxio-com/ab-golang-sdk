
# Get One Time Token Payment Profile

## Structure

`GetOneTimeTokenPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[string]` | Optional | - |
| `FirstName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `LastName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `MaskedCardNumber` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `CardType` | [`models.CardType`](../../doc/models/card-type.md) | Required | The type of card used. |
| `ExpirationMonth` | `float64` | Required | - |
| `ExpirationYear` | `float64` | Required | - |
| `CustomerId` | `models.Optional[string]` | Optional | - |
| `CurrentVault` | [`models.CreditCardVault`](../../doc/models/credit-card-vault.md) | Required | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress2` | `*string` | Optional | - |
| `BillingCity` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingCountry` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingZip` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `PaymentType` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Disabled` | `bool` | Required | - |
| `SiteGatewaySettingId` | `int` | Required | - |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    getOneTimeTokenPaymentProfile := models.GetOneTimeTokenPaymentProfile{
        Id:                   models.NewOptional(models.ToPointer("id4")),
        FirstName:            "first_name4",
        LastName:             "last_name2",
        MaskedCardNumber:     "masked_card_number2",
        CardType:             models.CardType_MAESTRONOLUHN,
        ExpirationMonth:      float64(28.3),
        ExpirationYear:       float64(4.96),
        CustomerId:           models.NewOptional(models.ToPointer("customer_id2")),
        CurrentVault:         models.CreditCardVault_TRUSTCOMMERCE,
        VaultToken:           "vault_token6",
        BillingAddress:       "billing_address6",
        BillingAddress2:      models.ToPointer("billing_address_26"),
        BillingCity:          "billing_city2",
        BillingCountry:       "billing_country8",
        BillingState:         "billing_state2",
        BillingZip:           "billing_zip2",
        PaymentType:          "payment_type6",
        Disabled:             false,
        SiteGatewaySettingId: 156,
        CustomerVaultToken:   models.NewOptional(models.ToPointer("customer_vault_token2")),
        GatewayHandle:        models.NewOptional(models.ToPointer("gateway_handle6")),
    }

}
```

