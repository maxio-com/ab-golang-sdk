
# Credit Card Payment Profile

## Structure

`CreditCardPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the stored card. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer. |
| `FirstName` | `*string` | Optional | The first name of the card holder. |
| `LastName` | `*string` | Optional | The last name of the card holder. |
| `MaskedCardNumber` | `*string` | Optional | A string representation of the credit card number with all but the last 4 digits masked with X’s (e.g., ‘XXXX-XXXX-XXXX-1234’). |
| `CardType` | [`models.Optional[models.CardType]`](../../doc/models/card-type.md) | Optional | The type of card used. |
| `ExpirationMonth` | `*int` | Optional | An integer representing the expiration month of the card(1 – 12). |
| `ExpirationYear` | `*int` | Optional | An integer representing the 4-digit expiration year of the card(e.g., ‘2012’). |
| `CustomerId` | `*int` | Optional | The Chargify-assigned id for the customer record to which the card belongs. |
| `CurrentVault` | [`*models.CreditCardVault`](../../doc/models/credit-card-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `models.Optional[string]` | Optional | The “token” provided by your vault storage for an already stored payment profile. |
| `BillingAddress` | `models.Optional[string]` | Optional | The current billing street address for the card. |
| `BillingCity` | `models.Optional[string]` | Optional | The current billing address city for the card. |
| `BillingState` | `models.Optional[string]` | Optional | The current billing address state for the card. |
| `BillingZip` | `models.Optional[string]` | Optional | The current billing address zip code for the card. |
| `BillingCountry` | `models.Optional[string]` | Optional | The current billing address country for the card. |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token. |
| `BillingAddress2` | `models.Optional[string]` | Optional | The current billing street address, second line, for the card. |
| `PaymentType` | [`models.PaymentType`](../../doc/models/payment-type.md) | Required | **Default**: `"credit_card"` |
| `Disabled` | `*bool` | Optional | - |
| `ChargifyToken` | `*string` | Optional | Token received after sending billing information using Maxio.js (formerly Chargify.js). This token will only be received if passed as a sole attribute of credit_card_attributes (e.g., tok_9g6hw85pnpt6knmskpwp4ttt). |
| `SiteGatewaySettingId` | `models.Optional[int]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | An identifier of connected gateway. |
| `CreatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was created |
| `UpdatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was last updated |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    creditCardPaymentProfile := models.CreditCardPaymentProfile{
        Id:                   models.ToPointer(10088716),
        FirstName:            models.ToPointer("Test"),
        LastName:             models.ToPointer("Subscription"),
        MaskedCardNumber:     models.ToPointer("XXXX-XXXX-XXXX-1"),
        CardType:             models.NewOptional(models.ToPointer(models.CardType_BOGUS)),
        ExpirationMonth:      models.ToPointer(1),
        ExpirationYear:       models.ToPointer(2022),
        CustomerId:           models.ToPointer(14543792),
        CurrentVault:         models.ToPointer(models.CreditCardVault_BOGUS),
        VaultToken:           models.NewOptional(models.ToPointer("1")),
        BillingAddress:       models.NewOptional(models.ToPointer("123 Montana Way")),
        BillingCity:          models.NewOptional(models.ToPointer("Billings")),
        BillingState:         models.NewOptional(models.ToPointer("MT")),
        BillingZip:           models.NewOptional(models.ToPointer("59101")),
        BillingCountry:       models.NewOptional(models.ToPointer("US")),
        CustomerVaultToken:   models.NewOptional(models.ToPointer("customer_vault_token2")),
        BillingAddress2:      models.NewOptional(models.ToPointer("")),
        PaymentType:          models.PaymentType_CREDITCARD,
        SiteGatewaySettingId: models.NewOptional(models.ToPointer(1)),
        GatewayHandle:        models.NewOptional(models.ToPointer("gateway_handle8")),
    }

}
```

