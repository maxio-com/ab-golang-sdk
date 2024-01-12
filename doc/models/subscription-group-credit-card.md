
# Subscription Group Credit Card

## Structure

`SubscriptionGroupCreditCard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullNumber` | `*interface{}` | Optional | - |
| `ExpirationMonth` | `*interface{}` | Optional | - |
| `ExpirationYear` | `*interface{}` | Optional | - |
| `ChargifyToken` | `*string` | Optional | - |
| `VaultToken` | `*string` | Optional | - |
| `CurrentVault` | [`*models.CurrentVault`](../../doc/models/current-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `GatewayHandle` | `*string` | Optional | - |
| `FirstName` | `*string` | Optional | - |
| `LastName` | `*string` | Optional | - |
| `BillingAddress` | `*string` | Optional | - |
| `BillingAddress2` | `*string` | Optional | - |
| `BillingCity` | `*string` | Optional | - |
| `BillingState` | `*string` | Optional | - |
| `BillingZip` | `*string` | Optional | - |
| `BillingCountry` | `*string` | Optional | - |
| `LastFour` | `*string` | Optional | - |
| `CardType` | `*string` | Optional | - |
| `CustomerVaultToken` | `*string` | Optional | - |
| `Cvv` | `*string` | Optional | - |
| `PaymentType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "full_number": {
    "key1": "val1",
    "key2": "val2"
  },
  "chargify_token": "tok_592nf92ng0sjd4300p",
  "card_type": "visa",
  "expiration_month": {
    "key1": "val1",
    "key2": "val2"
  },
  "expiration_year": {
    "key1": "val1",
    "key2": "val2"
  },
  "vault_token": "vault_token6"
}
```

