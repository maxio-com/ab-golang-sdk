
# Allocation Preview Item

## Structure

`AllocationPreviewItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Quantity` | [`*models.AllocationPreviewItemQuantity`](../../doc/models/containers/allocation-preview-item-quantity.md) | Optional | This is a container for one-of cases. |
| `PreviousQuantity` | [`*models.AllocationPreviewItemPreviousQuantity`](../../doc/models/containers/allocation-preview-item-previous-quantity.md) | Optional | This is a container for one-of cases. |
| `Memo` | `models.Optional[string]` | Optional | - |
| `Timestamp` | `models.Optional[string]` | Optional | - |
| `ProrationUpgradeScheme` | `*string` | Optional | - |
| `ProrationDowngradeScheme` | `*string` | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `PricePointId` | `*int` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `PreviousPricePointId` | `*int` | Optional | - |
| `PricePointHandle` | `*string` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `ComponentHandle` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "component_id": 54,
  "subscription_id": 54,
  "quantity": 78,
  "previous_quantity": 192,
  "memo": "memo6"
}
```

