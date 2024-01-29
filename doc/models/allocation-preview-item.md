
# Allocation Preview Item

## Structure

`AllocationPreviewItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Quantity` | `*interface{}` | Optional | - |
| `PreviousQuantity` | `*interface{}` | Optional | - |
| `Memo` | `Optional[string]` | Optional | - |
| `Timestamp` | `Optional[string]` | Optional | - |
| `ProrationUpgradeScheme` | `*string` | Optional | - |
| `ProrationDowngradeScheme` | `*string` | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `UpgradeCharge` | [`Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `PricePointId` | `*int` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`*models.IntervalUnit`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `PreviousPricePointId` | `*int` | Optional | - |
| `PricePointHandle` | `*string` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `ComponentHandle` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "component_id": 54,
  "subscription_id": 54,
  "quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "previous_quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "memo": "memo6"
}
```

