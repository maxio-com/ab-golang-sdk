
# Subscription Bank Account

## Structure

`SubscriptionBankAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BankAccountHolderType` | `*string` | Optional | Defaults to personal |
| `BankAccountType` | `*string` | Optional | Defaults to checking |
| `BankName` | `*string` | Optional | The bank where the account resides |
| `BillingAddress` | `*string` | Optional | The current billing street address for the bank account |
| `BillingAddress2` | `*string` | Optional | The current billing street address, second line, for the bank account |
| `BillingCity` | `*string` | Optional | The current billing address city for the bank account |
| `BillingState` | `*string` | Optional | The current billing address state for the bank account |
| `BillingZip` | `*string` | Optional | The current billing address zip code for the bank account |
| `BillingCountry` | `*string` | Optional | The current billing address country for the bank account |
| `CurrentVault` | [`*models.BankAccountVaultEnum`](bank-account-vault-enum.md) | Optional | The vault that stores the payment profile with the provided vault_token. |
| `CustomerId` | `*int` | Optional | The Chargify-assigned id for the customer record to which the bank account belongs |
| `CustomerVaultToken` | `*string` | Optional | (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token |
| `FirstName` | `*string` | Optional | The first name of the bank account holder |
| `LastName` | `*string` | Optional | The last name of the bank account holder |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the stored bank account. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer |
| `MaskedBankAccountNumber` | `*string` | Optional | A string representation of the stored bank account number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’) |
| `MaskedBankRoutingNumber` | `*string` | Optional | A string representation of the stored bank routing number with all but the last 4 digits marked with X’s (i.e. ‘XXXXXXX1111’). payment_type will be bank_account |
| `VaultToken` | `*string` | Optional | The “token” provided by your vault storage for an already stored payment profile |
| `ChargifyToken` | `*string` | Optional | Token received after sending billing informations using chargify.js. This token will only be received if passed as a sole attribute of credit_card_attributes (i.e. tok_9g6hw85pnpt6knmskpwp4ttt) |
| `SiteGatewaySettingId` | `*int` | Optional | - |
| `GatewayHandle` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "bank_account_holder_type": "bank_account_holder_type4",
  "bank_account_type": "bank_account_type4",
  "bank_name": "bank_name8",
  "billing_address": "billing_address8",
  "billing_address_2": "billing_address_28"
}
```

