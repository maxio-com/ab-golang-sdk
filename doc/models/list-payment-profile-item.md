
# List Payment Profile Item

## Structure

`ListPaymentProfileItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `CurrentVault` | [`*models.CurrentVaultEnum`](current-vault-enum.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
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
| `SiteGatewaySettingId` | `*int` | Optional | - |
| `GatewayHandle` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": 56,
  "first_name": "first_name6",
  "last_name": "last_name4",
  "customer_id": 94,
  "current_vault": "bogus"
}
```

