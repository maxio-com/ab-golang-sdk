
# Payment Profile

## Class Name

`PaymentProfile`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ApplePayPaymentProfile`](../../../doc/models/apple-pay-payment-profile.md) | models.PaymentProfileContainer.FromApplePayPaymentProfile(models.ApplePayPaymentProfile applePayPaymentProfile) |
| [`models.BankAccountPaymentProfile`](../../../doc/models/bank-account-payment-profile.md) | models.PaymentProfileContainer.FromBankAccountPaymentProfile(models.BankAccountPaymentProfile bankAccountPaymentProfile) |
| [`models.CreditCardPaymentProfile`](../../../doc/models/credit-card-payment-profile.md) | models.PaymentProfileContainer.FromCreditCardPaymentProfile(models.CreditCardPaymentProfile creditCardPaymentProfile) |
| [`models.PaypalPaymentProfile`](../../../doc/models/paypal-payment-profile.md) | models.PaymentProfileContainer.FromPaypalPaymentProfile(models.PaypalPaymentProfile paypalPaymentProfile) |

## models.ApplePayPaymentProfile

### Initialization Code

#### Example

```go
value := models.PaymentProfileContainer.FromApplePayPaymentProfile(models.ApplePayPaymentProfile{
    PaymentType:          models.PaymentType_APPLEPAY,
})
```

## models.BankAccountPaymentProfile

### Initialization Code

#### Example

```go
value := models.PaymentProfileContainer.FromBankAccountPaymentProfile(models.BankAccountPaymentProfile{
    PaymentType:             models.PaymentType_BANKACCOUNT,
    Verified:                models.ToPointer(false),
})
```

## models.CreditCardPaymentProfile

### Initialization Code

#### Example

```go
value := models.PaymentProfileContainer.FromCreditCardPaymentProfile(models.CreditCardPaymentProfile{
    Id:                   models.ToPointer(10088716),
    FirstName:            models.ToPointer("Test"),
    LastName:             models.ToPointer("Subscription"),
    MaskedCardNumber:     models.ToPointer("XXXX-XXXX-XXXX-1"),
    CardType:             models.ToPointer(models.CardType_BOGUS),
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
})
```

## models.PaypalPaymentProfile

### Initialization Code

#### Example

```go
value := models.PaymentProfileContainer.FromPaypalPaymentProfile(models.PaypalPaymentProfile{
    PaymentType:          models.PaymentType_PAYPALACCOUNT,
})
```

