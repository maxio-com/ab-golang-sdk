
# Create Payment Profile

## Structure

`CreatePaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyToken` | `*string` | Optional | Token received after sending billing information using Maxio.js (formerly Chargify.js). |
| `Id` | `*int` | Optional | - |
| `PaymentType` | [`*models.PaymentType`](../../doc/models/payment-type.md) | Optional | - |
| `FirstName` | `*string` | Optional | First name on card or bank account. If omitted, the first_name from customer attributes will be used. |
| `LastName` | `*string` | Optional | Last name on card or bank account. If omitted, the last_name from customer attributes will be used. |
| `MaskedCardNumber` | `*string` | Optional | - |
| `FullNumber` | `*string` | Optional | The full credit card number |
| `CardType` | [`*models.CardType`](../../doc/models/card-type.md) | Optional | The type of card used. |
| `ExpirationMonth` | [`*models.CreatePaymentProfileExpirationMonth`](../../doc/models/containers/create-payment-profile-expiration-month.md) | Optional | This is a container for one-of cases. |
| `ExpirationYear` | [`*models.CreatePaymentProfileExpirationYear`](../../doc/models/containers/create-payment-profile-expiration-year.md) | Optional | This is a container for one-of cases. |
| `BillingAddress` | `*string` | Optional | The credit card or bank account billing street address (e.g., 123 Main St.). This value is merely passed through to the payment gateway. |
| `BillingAddress2` | `models.Optional[string]` | Optional | Second line of the customer’s billing address e.g., Apt. 100 |
| `BillingCity` | `*string` | Optional | The credit card or bank account billing address city (e.g., “Boston”). This value is merely passed through to the payment gateway. |
| `BillingState` | `*string` | Optional | The credit card or bank account billing address state (e.g., MA). This value is merely passed through to the payment gateway. This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes. |
| `BillingCountry` | `*string` | Optional | “The credit card or bank account billing address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (e.g., “US”). This value is merely passed through to the payment gateway. Some gateways require country codes in a specific format. Check your gateway’s documentation. If creating an ACH subscription, only US is supported at this time.” |
| `BillingZip` | `*string` | Optional | The credit card or bank account billing address zip code (e.g., 12345). This value is merely passed through to the payment gateway. |
| `CurrentVault` | [`*models.AllVaults`](../../doc/models/all-vaults.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `VaultToken` | `*string` | Optional | The “token” provided by your vault storage for an already stored payment profile |
| `CustomerVaultToken` | `*string` | Optional | (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token |
| `CustomerId` | `*int` | Optional | (Required when creating a new payment profile) The Chargify customer id. |
| `PaypalEmail` | `*string` | Optional | used by merchants that implemented BraintreeBlue javaScript libraries on their own. We recommend using Maxio.js (formerly Chargify.js) instead. |
| `PaymentMethodNonce` | `*string` | Optional | used by merchants that implemented BraintreeBlue javaScript libraries on their own. We recommend using Maxio.js (formerly Chargify.js) instead. |
| `GatewayHandle` | `*string` | Optional | This attribute is only available if MultiGateway feature is enabled for your Site. This feature is in the Private Beta currently. gateway_handle is used to directly select a gateway where a payment profile will be stored in. Every connected gateway must have a unique gateway handle specified. Read [Multigateway description](https://chargify.zendesk.com/hc/en-us/articles/4407761759643#connecting-with-multiple-gateways) to learn more about new concepts that MultiGateway introduces and the default behavior when this attribute is not passed. |
| `Cvv` | `*string` | Optional | The 3- or 4-digit Card Verification Value. This value is merely passed through to the payment gateway. |
| `BankName` | `*string` | Optional | (Required when creating with ACH or GoCardless, optional with Stripe Direct Debit). The name of the bank where the customerʼs account resides |
| `BankIban` | `*string` | Optional | (Optional when creating with GoCardless, required with Stripe Direct Debit). International Bank Account Number. Alternatively, local bank details can be provided. |
| `BankRoutingNumber` | `*string` | Optional | (Required when creating with ACH. Optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API. |
| `BankAccountNumber` | `*string` | Optional | (Required when creating with ACH, GoCardless, Stripe BECS or BACS Direct Debit, and bank_iban is blank) The customerʼs bank account number |
| `BankBranchCode` | `*string` | Optional | (Optional when creating with GoCardless, required with Stripe BECS or BACS Direct Debit) Branch/Sort code. Alternatively, an IBAN can be provided. |
| `BankAccountType` | [`*models.BankAccountType`](../../doc/models/bank-account-type.md) | Optional | Defaults to checking |
| `BankAccountHolderType` | [`*models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Optional | Defaults to personal |
| `LastFour` | `*string` | Optional | (Optional) Used for creating subscription with payment profile imported using vault_token, for proper display in Advanced Billing UI |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createPaymentProfile := models.CreatePaymentProfile{
        ChargifyToken:         models.ToPointer("tok_9g6hw85pnpt6knmskpwp4ttt"),
        Id:                    models.ToPointer(124),
        PaymentType:           models.ToPointer(models.PaymentType_CREDITCARD),
        FirstName:             models.ToPointer("first_name6"),
        LastName:              models.ToPointer("last_name4"),
        FullNumber:            models.ToPointer("5424000000000015"),
    }

}
```

