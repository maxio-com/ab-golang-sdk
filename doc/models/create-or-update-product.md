
# Create or Update Product

## Structure

`CreateOrUpdateProduct`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | The product name |
| `Handle` | `*string` | Optional | The product API handle |
| `Description` | `string` | Required | The product description |
| `AccountingCode` | `*string` | Optional | E.g. Internal ID or SKU Number |
| `RequireCreditCard` | `*bool` | Optional | Deprecated value that can be ignored unless you have legacy hosted pages. For Public Signup Page users, read this attribute from under the signup page. |
| `PriceInCents` | `int64` | Required | The product price, in integer cents |
| `Interval` | `int` | Required | The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this product would renew every 30 days. |
| `IntervalUnit` | [`models.IntervalUnit`](../../doc/models/interval-unit.md) | Required | A string representing the interval unit for this product, either month or day |
| `TrialPriceInCents` | `*int64` | Optional | The product trial price, in integer cents |
| `TrialInterval` | `*int` | Optional | The numerical trial interval. e.g., an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product trial would last 30 days. |
| `TrialIntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the trial interval unit for this product, either month or day |
| `TrialType` | [`models.Optional[models.TrialType]`](../../doc/models/trial-type.md) | Optional | Indicates how a trial is handled when the trial period ends and there is no credit card on file. For `no_obligation`, the subscription transitions to a Trial Ended state. Maxio will not send any emails or statements. For `payment_expected`, the subscription transitions to a Past Due state. Maxio will send normal dunning emails and statements according to your other settings. |
| `ExpirationInterval` | `*int` | Optional | The numerical expiration interval. e.g., an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product would expire after 30 days. |
| `ExpirationIntervalUnit` | [`models.Optional[models.ExpirationIntervalUnit]`](../../doc/models/expiration-interval-unit.md) | Optional | A string representing the expiration interval unit for this product, either month, day or never |
| `AutoCreateSignupPage` | `*bool` | Optional | - |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the product type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `UnspscCode` | `models.Optional[string]` | Optional | (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this product instead of the default derived from item_category. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOrUpdateProduct := models.CreateOrUpdateProduct{
        Name:                   "name6",
        Handle:                 models.ToPointer("handle2"),
        Description:            "description4",
        AccountingCode:         models.ToPointer("accounting_code2"),
        RequireCreditCard:      models.ToPointer(false),
        PriceInCents:           int64(138),
        Interval:               154,
        IntervalUnit:           models.IntervalUnit_DAY,
        TrialPriceInCents:      models.ToPointer(int64(50)),
        TrialInterval:          models.ToPointer(252),
    }

}
```

