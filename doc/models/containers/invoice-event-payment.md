
# Invoice Event Payment

A nested data structure detailing the method of payment

## Class Name

`InvoiceEventPayment`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.PaymentMethodApplePay`](../../../doc/models/payment-method-apple-pay.md) | models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay paymentMethodApplePay) |
| [`models.PaymentMethodBankAccount`](../../../doc/models/payment-method-bank-account.md) | models.InvoiceEventPaymentContainer.FromPaymentMethodBankAccount(models.PaymentMethodBankAccount paymentMethodBankAccount) |
| [`models.PaymentMethodCreditCard`](../../../doc/models/payment-method-credit-card.md) | models.InvoiceEventPaymentContainer.FromPaymentMethodCreditCard(models.PaymentMethodCreditCard paymentMethodCreditCard) |
| [`models.PaymentMethodExternal`](../../../doc/models/payment-method-external.md) | models.InvoiceEventPaymentContainer.FromPaymentMethodExternal(models.PaymentMethodExternal paymentMethodExternal) |
| [`models.PaymentMethodPaypal`](../../../doc/models/payment-method-paypal.md) | models.InvoiceEventPaymentContainer.FromPaymentMethodPaypal(models.PaymentMethodPaypal paymentMethodPaypal) |

## models.PaymentMethodApplePay

### Initialization Code

#### Example

```go
value := models.InvoiceEventPaymentContainer.FromPaymentMethodApplePay(models.PaymentMethodApplePay{
    Type:                 models.InvoiceEventPaymentMethod_APPLEPAY,
})
```

## models.PaymentMethodBankAccount

### Initialization Code

#### Example

```go
value := models.InvoiceEventPaymentContainer.FromPaymentMethodBankAccount(models.PaymentMethodBankAccount{
    MaskedAccountNumber:  "masked_account_number2",
    MaskedRoutingNumber:  "masked_routing_number2",
    Type:                 models.InvoiceEventPaymentMethod_BANKACCOUNT,
})
```

## models.PaymentMethodCreditCard

### Initialization Code

#### Example

```go
value := models.InvoiceEventPaymentContainer.FromPaymentMethodCreditCard(models.PaymentMethodCreditCard{
    CardBrand:            "card_brand4",
    MaskedCardNumber:     "masked_card_number0",
    Type:                 models.InvoiceEventPaymentMethod_CREDITCARD,
})
```

## models.PaymentMethodExternal

### Initialization Code

#### Example

```go
value := models.InvoiceEventPaymentContainer.FromPaymentMethodExternal(models.PaymentMethodExternal{
    Details:              models.ToPointer("details4"),
    Kind:                 "kind2",
    Memo:                 models.ToPointer("memo8"),
    Type:                 models.InvoiceEventPaymentMethod_EXTERNAL,
})
```

## models.PaymentMethodPaypal

### Initialization Code

#### Example

```go
value := models.InvoiceEventPaymentContainer.FromPaymentMethodPaypal(models.PaymentMethodPaypal{
    Email:                "email2",
    Type:                 models.InvoiceEventPaymentMethod_PAYPALACCOUNT,
})
```

