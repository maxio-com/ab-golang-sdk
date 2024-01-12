
# Created Payment Profile

## Structure

`CreatedPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `MaskedCardNumber` | `Optional[string]` | Optional | - |
| `CardType` | `*string` | Optional | - |
| `ExpirationMonth` | `*int` | Optional | - |
| `ExpirationYear` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `CurrentVault` | [`*models.CurrentVault`](../../doc/models/current-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `*string` | Optional | - |
| `BillingAddress` | `*string` | Optional | - |
| `BillingCity` | `*string` | Optional | - |
| `BillingState` | `*string` | Optional | - |
| `BillingZip` | `*string` | Optional | - |
| `BillingCountry` | `*string` | Optional | - |
| `CustomerVaultToken` | `Optional[string]` | Optional | - |
| `BillingAddress2` | `Optional[string]` | Optional | - |
| `PaymentType` | `*string` | Optional | - |
| `BankName` | `*string` | Optional | - |
| `MaskedBankRoutingNumber` | `*string` | Optional | - |
| `MaskedBankAccountNumber` | `*string` | Optional | - |
| `BankAccountType` | `*string` | Optional | - |
| `BankAccountHolderType` | `*string` | Optional | - |
| `Verified` | `*bool` | Optional | - |
| `SiteGatewaySettingId` | `*int` | Optional | - |
| `GatewayHandle` | `*string` | Optional | - |
| `Disabled` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "id": 14,
  "first_name": "first_name0",
  "last_name": "last_name8",
  "masked_card_number": "masked_card_number8",
  "card_type": "card_type4"
}
```

