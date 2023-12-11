
# Get One Time Token Payment Profile

## Structure

`GetOneTimeTokenPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `FirstName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `LastName` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `MaskedCardNumber` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `CardType` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `ExpirationMonth` | `float64` | Required | - |
| `ExpirationYear` | `float64` | Required | - |
| `CustomerId` | `Optional[string]` | Optional | - |
| `CurrentVault` | [`models.CurrentVaultEnum`](current-vault-enum.md) | Required | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress2` | `*string` | Optional | - |
| `BillingCity` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingCountry` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `BillingZip` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `PaymentType` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Disabled` | `bool` | Required | - |
| `SiteGatewaySettingId` | `int` | Required | - |
| `CustomerVaultToken` | `Optional[string]` | Optional | - |
| `GatewayHandle` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id2",
  "first_name": "first_name2",
  "last_name": "last_name0",
  "masked_card_number": "masked_card_number0",
  "card_type": "card_type8",
  "expiration_month": 187.78,
  "expiration_year": 164.44,
  "customer_id": "customer_id0",
  "current_vault": "firstdata",
  "vault_token": "vault_token4",
  "billing_address": "billing_address4",
  "billing_address_2": "billing_address_24",
  "billing_city": "billing_city0",
  "billing_country": "billing_country6",
  "billing_state": "billing_state6",
  "billing_zip": "billing_zip0",
  "payment_type": "payment_type2",
  "disabled": false,
  "site_gateway_setting_id": 232,
  "customer_vault_token": "customer_vault_token0",
  "gateway_handle": "gateway_handle4"
}
```

