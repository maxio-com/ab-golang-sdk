
# Get One Time Token Request Payment Profile

## Class Name

`GetOneTimeTokenRequestPaymentProfile`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.GetOneTimeTokenPaymentProfile`](../../../doc/models/get-one-time-token-payment-profile.md) | models.GetOneTimeTokenRequestPaymentProfileContainer.FromGetOneTimeTokenPaymentProfile(models.GetOneTimeTokenPaymentProfile getOneTimeTokenPaymentProfile) |
| [`models.GetOneTimeTokenBankAccountPaymentProfile`](../../../doc/models/get-one-time-token-bank-account-payment-profile.md) | models.GetOneTimeTokenRequestPaymentProfileContainer.FromGetOneTimeTokenBankAccountPaymentProfile(models.GetOneTimeTokenBankAccountPaymentProfile getOneTimeTokenBankAccountPaymentProfile) |

## models.GetOneTimeTokenPaymentProfile

### Initialization Code

#### Example

```go
value := models.GetOneTimeTokenRequestPaymentProfileContainer.FromGetOneTimeTokenPaymentProfile(models.GetOneTimeTokenPaymentProfile{
    FirstName:            "first_name2",
    LastName:             "last_name0",
    MaskedCardNumber:     "masked_card_number0",
    CardType:             models.CardType_ROUTEX,
    ExpirationMonth:      float64(187.78),
    ExpirationYear:       float64(164.44),
    CurrentVault:         models.CreditCardVault_BRAINTREEBLUE,
    VaultToken:           "vault_token4",
    BillingAddress:       "billing_address4",
    BillingCity:          "billing_city0",
    BillingCountry:       "billing_country6",
    BillingState:         "billing_state6",
    BillingZip:           "billing_zip0",
    PaymentType:          "payment_type2",
    Disabled:             false,
    SiteGatewaySettingId: 232,
})
```

## models.GetOneTimeTokenBankAccountPaymentProfile

### Initialization Code

#### Example

```go
value := models.GetOneTimeTokenRequestPaymentProfileContainer.FromGetOneTimeTokenBankAccountPaymentProfile(models.GetOneTimeTokenBankAccountPaymentProfile{
    FirstName:               "first_name8",
    LastName:                "last_name6",
    CurrentVault:            models.BankAccountVault_MAXP,
    VaultToken:              "vault_token0",
    BillingAddress:          "billing_address0",
    BillingCity:             "billing_city4",
    BillingCountry:          "billing_country2",
    BillingState:            "billing_state8",
    BillingZip:              "billing_zip6",
    BankName:                "bank_name6",
    MaskedBankRoutingNumber: "masked_bank_routing_number6",
    MaskedBankAccountNumber: "masked_bank_account_number0",
    BankAccountType:         models.BankAccountType_CHECKING,
    BankAccountHolderType:   models.BankAccountHolderType_PERSONAL,
    PaymentType:             "payment_type2",
    Disabled:                false,
    SiteGatewaySettingId:    254,
})
```

