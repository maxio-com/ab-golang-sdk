
# Bank Account

## Structure

`BankAccount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `CurrentVault` | [`*models.BankAccountVaultEnum`](bank-account-vault-enum.md) | Optional | The vault that stores the payment profile with the provided vault_token. |
| `VaultToken` | `*string` | Optional | - |
| `BillingAddress` | `*string` | Optional | - |
| `BillingCity` | `*string` | Optional | - |
| `BillingState` | `*string` | Optional | - |
| `BillingZip` | `*string` | Optional | - |
| `BillingCountry` | `*string` | Optional | - |
| `CustomerVaultToken` | `Optional[string]` | Optional | - |
| `BillingAddress2` | `*string` | Optional | - |
| `BankName` | `*string` | Optional | - |
| `MaskedBankRoutingNumber` | `*string` | Optional | - |
| `MaskedBankAccountNumber` | `*string` | Optional | - |
| `BankAccountType` | `*string` | Optional | - |
| `BankAccountHolderType` | `*string` | Optional | - |
| `PaymentType` | `*string` | Optional | - |
| `Verified` | `*bool` | Optional | denotes whether a bank account has been verified by providing the amounts of two small deposits made into the account<br>**Default**: `false` |
| `SiteGatewaySettingId` | `*int` | Optional | - |
| `GatewayHandle` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "verified": false,
  "id": 190,
  "first_name": "first_name2",
  "last_name": "last_name0",
  "customer_id": 228,
  "current_vault": "stripe_connect"
}
```

