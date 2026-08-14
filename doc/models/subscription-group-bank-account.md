
# Subscription Group Bank Account

## Structure

`SubscriptionGroupBankAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BankName` | `*string` | Optional | (Required when creating a subscription with ACH or GoCardless) The name of the bank where the customer’s account resides |
| `BankAccountNumber` | `*string` | Optional | (Required when creating a subscription with ACH. Required when creating a subscription with GoCardless and bank_iban is blank) The customerʼs bank account number |
| `BankRoutingNumber` | `*string` | Optional | (Required when creating a subscription with ACH. Optional when creating a subscription with GoCardless.) The routing number of the bank. It becomes bank_code while passing via GoCardless API. |
| `BankIban` | `*string` | Optional | (Optional when creating a subscription with GoCardless). International Bank Account Number. Alternatively, local bank details can be provided. |
| `BankBranchCode` | `*string` | Optional | (Optional when creating a subscription with GoCardless) Branch code. Alternatively, an IBAN can be provided. |
| `BankAccountType` | [`*models.BankAccountType`](../../doc/models/bank-account-type.md) | Optional | Defaults to checking |
| `BankAccountHolderType` | [`*models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Optional | Defaults to personal |
| `PaymentType` | [`*models.PaymentType`](../../doc/models/payment-type.md) | Optional | - |
| `BillingAddress` | `*string` | Optional | - |
| `BillingCity` | `*string` | Optional | - |
| `BillingState` | `*string` | Optional | - |
| `BillingZip` | `*string` | Optional | - |
| `BillingCountry` | `*string` | Optional | - |
| `ChargifyToken` | `*string` | Optional | - |
| `CurrentVault` | [`*models.BankAccountVault`](../../doc/models/bank-account-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing. |
| `GatewayHandle` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionGroupBankAccount := models.SubscriptionGroupBankAccount{
        BankName:              models.ToPointer("bank_name2"),
        BankAccountNumber:     models.ToPointer("bank_account_number4"),
        BankRoutingNumber:     models.ToPointer("bank_routing_number8"),
        BankIban:              models.ToPointer("bank_iban6"),
        BankBranchCode:        models.ToPointer("bank_branch_code6"),
    }

}
```

