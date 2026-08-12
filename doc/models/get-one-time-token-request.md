
# Get One Time Token Request

## Structure

`GetOneTimeTokenRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.GetOneTimeTokenRequestPaymentProfile`](../../doc/models/containers/get-one-time-token-request-payment-profile.md) | Required | This is a container for any-of cases. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    getOneTimeTokenRequest := models.GetOneTimeTokenRequest{
        PaymentProfile:       models.GetOneTimeTokenRequestPaymentProfileContainer.FromGetOneTimeTokenPaymentProfile(models.GetOneTimeTokenPaymentProfile{
            Id:                   models.NewOptional(models.ToPointer("id2")),
            FirstName:            "first_name2",
            LastName:             "last_name0",
            MaskedCardNumber:     "masked_card_number0",
            CardType:             models.CardType_ROUTEX,
            ExpirationMonth:      float64(187.78),
            ExpirationYear:       float64(164.44),
            CustomerId:           models.NewOptional(models.ToPointer("customer_id0")),
            CurrentVault:         models.CreditCardVault_BRAINTREEBLUE,
            VaultToken:           "vault_token4",
            BillingAddress:       "billing_address4",
            BillingAddress2:      models.ToPointer("billing_address_24"),
            BillingCity:          "billing_city0",
            BillingCountry:       "billing_country6",
            BillingState:         "billing_state6",
            BillingZip:           "billing_zip0",
            PaymentType:          "payment_type2",
            Disabled:             false,
            SiteGatewaySettingId: 232,
            CustomerVaultToken:   models.NewOptional(models.ToPointer("customer_vault_token0")),
            GatewayHandle:        models.NewOptional(models.ToPointer("gateway_handle4")),
        }),
    }

}
```

