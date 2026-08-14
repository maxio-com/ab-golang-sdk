
# Bank Account Payment Profile

## Structure

`BankAccountPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the stored bank account. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer. |
| `FirstName` | `*string` | Optional | The first name of the bank account holder |
| `LastName` | `*string` | Optional | The last name of the bank account holder |
| `CustomerId` | `*int` | Optional | The Chargify-assigned ID for the customer record to which the bank account belongs |
| `CurrentVault` | [`*models.BankAccountVault`](../../doc/models/bank-account-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing. |
| `VaultToken` | `*string` | Optional | The "token" provided by your vault storage for an already stored payment profile |
| `BillingAddress` | `models.Optional[string]` | Optional | The current billing street address for the bank account |
| `BillingCity` | `models.Optional[string]` | Optional | The current billing address city for the bank account |
| `BillingState` | `models.Optional[string]` | Optional | The current billing address state for the bank account |
| `BillingZip` | `models.Optional[string]` | Optional | The current billing address zip code for the bank account |
| `BillingCountry` | `models.Optional[string]` | Optional | The current billing address country for the bank account |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token. |
| `BillingAddress2` | `models.Optional[string]` | Optional | The current billing street address, second line, for the bank account |
| `BankName` | `*string` | Optional | The bank where the account resides |
| `MaskedBankRoutingNumber` | `models.Optional[string]` | Optional | A string representation of the stored bank routing number with all but the last 4 digits marked with X's (i.e. 'XXXXXXX1111'). payment_type will be bank_account. |
| `MaskedBankAccountNumber` | `models.Optional[string]` | Optional | A string representation of the stored bank account number with all but the last 4 digits marked with X's (i.e. 'XXXXXXX1111'). |
| `BankAccountType` | [`*models.BankAccountType`](../../doc/models/bank-account-type.md) | Optional | Defaults to checking |
| `BankAccountHolderType` | [`*models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Optional | Defaults to personal |
| `PaymentType` | [`models.PaymentType`](../../doc/models/payment-type.md) | Required | **Default**: `"bank_account"` |
| `Verified` | `*bool` | Optional | Denotes whether a bank account has been verified by providing the amounts of two small deposits made into the account.<br><br>**Default**: `false` |
| `SiteGatewaySettingId` | `models.Optional[int]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was created |
| `UpdatedAt` | `*time.Time` | Optional | A timestamp indicating when this payment profile was last updated |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    bankAccountPaymentProfile := models.BankAccountPaymentProfile{
        Id:                      models.ToPointer(16),
        FirstName:               models.ToPointer("first_name0"),
        LastName:                models.ToPointer("last_name8"),
        CustomerId:              models.ToPointer(54),
        CurrentVault:            models.ToPointer(models.BankAccountVault_GOCARDLESS),
        PaymentType:             models.PaymentType_BANKACCOUNT,
        Verified:                models.ToPointer(false),
    }

}
```

