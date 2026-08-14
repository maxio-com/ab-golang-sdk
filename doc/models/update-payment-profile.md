
# Update Payment Profile

## Structure

`UpdatePaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FirstName` | `*string` | Optional | The first name of the card holder. |
| `LastName` | `*string` | Optional | The last name of the card holder. |
| `FullNumber` | `*string` | Optional | The full credit card number |
| `CardType` | [`*models.CardType`](../../doc/models/card-type.md) | Optional | The type of card used. |
| `ExpirationMonth` | `*string` | Optional | (Optional when performing an Import via vault_token, required otherwise) The 1- or 2-digit credit card expiration month, as an integer or string, e.g., 5 |
| `ExpirationYear` | `*string` | Optional | (Optional when performing an Import via vault_token, required otherwise) The 4-digit credit card expiration year, as an integer or string, e.g., 2012 |
| `CurrentVault` | [`*models.AllVaults`](../../doc/models/all-vaults.md) | Optional | The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing. |
| `BillingAddress` | `*string` | Optional | The credit card or bank account billing street address (e.g., 123 Main St.). This value is merely passed through to the payment gateway. |
| `BillingCity` | `*string` | Optional | The credit card or bank account billing address city (e.g., “Boston”). This value is merely passed through to the payment gateway. |
| `BillingState` | `*string` | Optional | The credit card or bank account billing address state (e.g., MA). This value is merely passed through to the payment gateway. This must conform to the [ISO_3166-1](https://en.wikipedia.org/wiki/ISO_3166-1#Current_codes) in order to be valid for tax locale purposes. |
| `BillingZip` | `*string` | Optional | The credit card or bank account billing address zip code (e.g., 12345). This value is merely passed through to the payment gateway. |
| `BillingCountry` | `*string` | Optional | The credit card or bank account billing address country, required in [ISO_3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) format (e.g., “US”). This value is merely passed through to the payment gateway. Some gateways require country codes in a specific format. Check your gateway’s documentation. If creating an ACH subscription, only US is supported at this time. |
| `BillingAddress2` | `models.Optional[string]` | Optional | Second line of the customer’s billing address, e.g., Apt. 100 |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updatePaymentProfile := models.UpdatePaymentProfile{
        FirstName:            models.ToPointer("first_name6"),
        LastName:             models.ToPointer("last_name4"),
        FullNumber:           models.ToPointer("5424000000000015"),
        CardType:             models.ToPointer(models.CardType_AMERICANEXPRESS),
        ExpirationMonth:      models.ToPointer("expiration_month8"),
    }

}
```

