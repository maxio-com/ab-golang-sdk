
# Bank Account Attributes

## Structure

`BankAccountAttributes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyToken` | `*string` | Optional | - |
| `BankName` | `*string` | Optional | (Required when creating a subscription with ACH or GoCardless) The name of the bank where the customer’s account resides |
| `BankRoutingNumber` | `*string` | Optional | (Required when creating a subscription with ACH; optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API. |
| `BankAccountNumber` | `*string` | Optional | (Required when creating a subscription with ACH. Required when creating a subscription with GoCardless and bank_iban is blank) The customerʼs bank account number |
| `BankAccountType` | [`*models.BankAccountType`](../../doc/models/bank-account-type.md) | Optional | Defaults to checking |
| `BankBranchCode` | `*string` | Optional | (Optional when creating a subscription with GoCardless) Branch code. Alternatively, an IBAN can be provided. |
| `BankIban` | `*string` | Optional | (Optional when creating a subscription with GoCardless). International Bank Account Number. Alternatively, local bank details can be provided. |
| `BankAccountHolderType` | [`*models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Optional | Defaults to personal |
| `PaymentType` | [`*models.PaymentType`](../../doc/models/payment-type.md) | Optional | - |
| `CurrentVault` | [`*models.BankAccountVault`](../../doc/models/bank-account-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing. |
| `VaultToken` | `*string` | Optional | - |
| `CustomerVaultToken` | `*string` | Optional | (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountAttributes := models.BankAccountAttributes{
        ChargifyToken:         models.ToPointer("chargify_token0"),
        BankName:              models.ToPointer("bank_name2"),
        BankRoutingNumber:     models.ToPointer("bank_routing_number8"),
        BankAccountNumber:     models.ToPointer("bank_account_number4"),
        BankAccountType:       models.ToPointer(models.BankAccountType_CHECKING),
    }

}
```

