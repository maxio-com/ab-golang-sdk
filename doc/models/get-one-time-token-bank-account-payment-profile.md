
# Get One Time Token Bank Account Payment Profile

## Structure

`GetOneTimeTokenBankAccountPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[string]` | Optional | - |
| `FirstName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `LastName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `CustomerId` | `models.Optional[string]` | Optional | - |
| `CurrentVault` | [`models.BankAccountVault`](../../doc/models/bank-account-vault.md) | Required | The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing. |
| `VaultToken` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress2` | `*string` | Optional | - |
| `BillingCity` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingCountry` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingZip` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BankName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `MaskedBankRoutingNumber` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `MaskedBankAccountNumber` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BankAccountType` | [`models.BankAccountType`](../../doc/models/bank-account-type.md) | Required | Defaults to checking |
| `BankAccountHolderType` | [`models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Required | Defaults to personal |
| `PaymentType` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Disabled` | `bool` | Required | - |
| `SiteGatewaySettingId` | `int` | Required | - |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `Verified` | `models.Optional[bool]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    getOneTimeTokenBankAccountPaymentProfile := models.GetOneTimeTokenBankAccountPaymentProfile{
        Id:                      models.NewOptional(models.ToPointer("id0")),
        FirstName:               "first_name0",
        LastName:                "last_name8",
        CustomerId:              models.NewOptional(models.ToPointer("customer_id8")),
        CurrentVault:            models.BankAccountVault_MAXP,
        VaultToken:              "vault_token2",
        BillingAddress:          "billing_address2",
        BillingAddress2:         models.ToPointer("billing_address_22"),
        BillingCity:             "billing_city8",
        BillingCountry:          "billing_country4",
        BillingState:            "billing_state6",
        BillingZip:              "billing_zip8",
        BankName:                "bank_name4",
        MaskedBankRoutingNumber: "masked_bank_routing_number4",
        MaskedBankAccountNumber: "masked_bank_account_number2",
        BankAccountType:         models.BankAccountType_CHECKING,
        BankAccountHolderType:   models.BankAccountHolderType_PERSONAL,
        PaymentType:             "payment_type0",
        Disabled:                false,
        SiteGatewaySettingId:    246,
        CustomerVaultToken:      models.NewOptional(models.ToPointer("customer_vault_token8")),
        GatewayHandle:           models.NewOptional(models.ToPointer("gateway_handle2")),
    }

}
```

