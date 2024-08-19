
# Credit Card Payment Profile

## Structure

`CreditCardPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | The Chargify-assigned ID of the stored card. This value can be used as an input to payment_profile_id when creating a subscription, in order to re-use a stored payment profile for the same customer. |
| `FirstName` | `*string` | Optional | The first name of the card holder. |
| `LastName` | `*string` | Optional | The last name of the card holder. |
| `MaskedCardNumber` | `*string` | Optional | A string representation of the credit card number with all but the last 4 digits masked with X’s (i.e. ‘XXXX-XXXX-XXXX-1234’). |
| `CardType` | [`*models.CardType`](../../doc/models/card-type.md) | Optional | The type of card used. |
| `ExpirationMonth` | `*int` | Optional | An integer representing the expiration month of the card(1 – 12). |
| `ExpirationYear` | `*int` | Optional | An integer representing the 4-digit expiration year of the card(i.e. ‘2012’). |
| `CustomerId` | `*int` | Optional | The Chargify-assigned id for the customer record to which the card belongs. |
| `CurrentVault` | [`*models.CreditCardVault`](../../doc/models/credit-card-vault.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `models.Optional[string]` | Optional | The “token” provided by your vault storage for an already stored payment profile. |
| `BillingAddress` | `models.Optional[string]` | Optional | The current billing street address for the card. |
| `BillingCity` | `models.Optional[string]` | Optional | The current billing address city for the card. |
| `BillingState` | `models.Optional[string]` | Optional | The current billing address state for the card. |
| `BillingZip` | `models.Optional[string]` | Optional | The current billing address zip code for the card. |
| `BillingCountry` | `models.Optional[string]` | Optional | The current billing address country for the card. |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | (only for Authorize.Net CIM storage): the customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token. |
| `BillingAddress2` | `models.Optional[string]` | Optional | The current billing street address, second line, for the card. |
| `PaymentType` | [`models.PaymentType`](../../doc/models/payment-type.md) | Required | **Default**: `"credit_card"` |
| `Disabled` | `*bool` | Optional | - |
| `ChargifyToken` | `*string` | Optional | Token received after sending billing information using chargify.js. This token will only be received if passed as a sole attribute of credit_card_attributes (i.e. tok_9g6hw85pnpt6knmskpwp4ttt) |
| `SiteGatewaySettingId` | `models.Optional[int]` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | An identifier of connected gateway. |

## Example (as JSON)

```json
{
  "id": 10088716,
  "first_name": "Test",
  "last_name": "Subscription",
  "masked_card_number": "XXXX-XXXX-XXXX-1",
  "card_type": "bogus",
  "expiration_month": 1,
  "expiration_year": 2022,
  "customer_id": 14543792,
  "current_vault": "bogus",
  "vault_token": "1",
  "billing_address": "123 Montana Way",
  "billing_city": "Billings",
  "billing_state": "MT",
  "billing_zip": "59101",
  "billing_country": "US",
  "customer_vault_token": null,
  "billing_address_2": "",
  "payment_type": "credit_card",
  "site_gateway_setting_id": 1,
  "gateway_handle": null
}
```

