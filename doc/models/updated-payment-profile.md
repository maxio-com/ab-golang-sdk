
# Updated Payment Profile

## Structure

`UpdatedPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `CardType` | `*string` | Optional | - |
| `ExpirationMonth` | `*int` | Optional | - |
| `ExpirationYear` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `CurrentVault` | [`*models.CurrentVault`](../../doc/models/current-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `*string` | Optional | - |
| `BillingAddress` | `*string` | Optional | - |
| `BillingAddress2` | `*string` | Optional | - |
| `BillingCity` | `*string` | Optional | - |
| `BillingState` | `*string` | Optional | - |
| `BillingZip` | `*string` | Optional | - |
| `BillingCountry` | `*string` | Optional | - |
| `PaymentType` | `*string` | Optional | - |
| `SiteGatewaySettingId` | `*int` | Optional | - |
| `GatewayHandle` | `Optional[string]` | Optional | - |
| `MaskedCardNumber` | `*string` | Optional | - |
| `CustomerVaultToken` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 232,
  "first_name": "first_name0",
  "last_name": "last_name8",
  "card_type": "card_type4",
  "expiration_month": 150
}
```

