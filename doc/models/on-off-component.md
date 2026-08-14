
# On Off Component

## Structure

`OnOffComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `string` | Required | A name for this component that is suitable for showing customers and displaying on billing statements, e.g., "Minutes". |
| `Description` | `*string` | Optional | A description for the component that will be displayed to the user on the hosted signup page. |
| `Handle` | `*string` | Optional | A unique identifier for your use that can be used to retrieve this component in subsequent requests. Must start with a letter or number and may only contain lowercase letters, numbers, or the characters '.', ':', '-', or '_'.<br><br>**Constraints**: *Pattern*: `^[a-z0-9][a-z0-9\-_:.]*$` |
| `Taxable` | `*bool` | Optional | Boolean flag describing whether a component is taxable or not. |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `PricePoints` | [`[]models.ComponentPricePointItem`](../../doc/models/component-price-point-item.md) | Optional | - |
| `UnitPrice` | [`models.OnOffComponentUnitPrice`](../../doc/models/containers/on-off-component-unit-price.md) | Required | This is a container for one-of cases. |
| `TaxCode` | `*string` | Optional | A string representing the tax code related to the component type. This is especially important when using AvaTax to tax based on locale. This attribute has a max length of 25 characters. |
| `HideDateRangeOnInvoice` | `*bool` | Optional | (Only available on Relationship Invoicing sites) Boolean flag describing if the service date range should show for the component on generated invoices. |
| `DisplayOnHostedPage` | `*bool` | Optional | - |
| `AllowFractionalQuantities` | `*bool` | Optional | - |
| `PublicSignupPageIds` | `[]int` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. e.g., an interval of ‘30’ coupled with an interval_unit of day would mean this component's default price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component's default price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `UnspscCode` | `models.Optional[string]` | Optional | (Optional) Custom UNSPSC commodity code for Level 3/CEDP payment data. When set, this value is sent as the commodity code on invoice line items for this component instead of the default derived from item_category. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    onOffComponent := models.OnOffComponent{
        Name:                      "name0",
        Description:               models.ToPointer("description0"),
        Handle:                    models.ToPointer("handle6"),
        Taxable:                   models.ToPointer(false),
        UpgradeCharge:             models.NewOptional(models.ToPointer(models.CreditType_FULL)),
        DowngradeCredit:           models.NewOptional(models.ToPointer(models.CreditType_FULL)),
        UnitPrice:                 models.OnOffComponentUnitPriceContainer.FromString("String9"),
    }

}
```

