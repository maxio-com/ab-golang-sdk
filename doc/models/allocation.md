
# Allocation

## Structure

`Allocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllocationId` | `*int` | Optional | The allocation unique id |
| `ComponentId` | `*int` | Optional | The integer component ID for the allocation. This references a component that you have created in your Product setup |
| `ComponentHandle` | `models.Optional[string]` | Optional | The handle of the component. This references a component that you have created in your Product setup |
| `SubscriptionId` | `*int` | Optional | The integer subscription ID for the allocation. This references a unique subscription in your Site |
| `Quantity` | [`*models.AllocationQuantity`](../../doc/models/containers/allocation-quantity.md) | Optional | This is a container for one-of cases. |
| `PreviousQuantity` | [`*models.AllocationPreviousQuantity`](../../doc/models/containers/allocation-previous-quantity.md) | Optional | This is a container for one-of cases. |
| `Memo` | `models.Optional[string]` | Optional | The memo passed when the allocation was created |
| `Timestamp` | `*time.Time` | Optional | The time that the allocation was recorded, in format and UTC timezone, i.e. 2012-11-20T22:00:37Z |
| `CreatedAt` | `*time.Time` | Optional | Timestamp indicating when this allocation was created |
| `ProrationUpgradeScheme` | `*string` | Optional | The scheme used if the proration was an upgrade. This is only present when the allocation was created mid-period. |
| `ProrationDowngradeScheme` | `*string` | Optional | The scheme used if the proration was a downgrade. This is only present when the allocation was created mid-period. |
| `PricePointId` | `*int` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `PricePointHandle` | `*string` | Optional | - |
| `Interval` | `*int` | Optional | The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled. |
| `IntervalUnit` | [`models.Optional[models.IntervalUnit]`](../../doc/models/interval-unit.md) | Optional | A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled. |
| `PreviousPricePointId` | `*int` | Optional | - |
| `AccrueCharge` | `*bool` | Optional | If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately. |
| `InitiateDunning` | `*bool` | Optional | If true, if the immediate component payment fails, initiate dunning for the subscription.<br>Otherwise, leave the charges on the subscription to pay for at renewal. |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `Payment` | [`models.Optional[models.PaymentForAllocation]`](../../doc/models/payment-for-allocation.md) | Optional | - |
| `ExpiresAt` | `*time.Time` | Optional | - |
| `UsedQuantity` | `*int64` | Optional | - |
| `ChargeId` | `*int64` | Optional | - |

## Example (as JSON)

```json
{
  "allocation_id": 102,
  "component_id": 144,
  "component_handle": "component_handle0",
  "subscription_id": 144,
  "quantity": 168
}
```

