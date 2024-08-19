
# Subscription Group Credit Card

## Structure

`SubscriptionGroupCreditCard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullNumber` | [`*models.SubscriptionGroupCreditCardFullNumber`](../../doc/models/containers/subscription-group-credit-card-full-number.md) | Optional | This is a container for one-of cases. |
| `ExpirationMonth` | [`*models.SubscriptionGroupCreditCardExpirationMonth`](../../doc/models/containers/subscription-group-credit-card-expiration-month.md) | Optional | This is a container for one-of cases. |
| `ExpirationYear` | [`*models.SubscriptionGroupCreditCardExpirationYear`](../../doc/models/containers/subscription-group-credit-card-expiration-year.md) | Optional | This is a container for one-of cases. |
| `ChargifyToken` | `*string` | Optional | - |
| `VaultToken` | `*string` | Optional | - |
| `CurrentVault` | [`*models.CreditCardVault`](../../doc/models/credit-card-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
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
| `CardType` | [`*models.CardType`](../../doc/models/card-type.md) | Optional | The type of card used. |
| `CustomerVaultToken` | `*string` | Optional | - |
| `Cvv` | `*string` | Optional | - |
| `PaymentType` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "full_number": 4111111111111111,
  "chargify_token": "tok_592nf92ng0sjd4300p",
  "expiration_month": "String1",
  "expiration_year": "String5",
  "vault_token": "vault_token6"
}
```

