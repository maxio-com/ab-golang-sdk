
# Allocation

## Structure

`Allocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | The integer component ID for the allocation. This references a component that you have created in your Product setup |
| `SubscriptionId` | `*int` | Optional | The integer subscription ID for the allocation. This references a unique subscription in your Site |
| `Quantity` | `*int` | Optional | The allocated quantity set in to effect by the allocation |
| `PreviousQuantity` | `*int` | Optional | The allocated quantity that was in effect before this allocation was created |
| `Memo` | `Optional[string]` | Optional | The memo passed when the allocation was created |
| `Timestamp` | `*string` | Optional | The time that the allocation was recorded, in  format and UTC timezone, i.e. 2012-11-20T22:00:37Z |
| `ProrationUpgradeScheme` | `*string` | Optional | The scheme used if the proration was an upgrade. This is only present when the allocation was created mid-period. |
| `ProrationDowngradeScheme` | `*string` | Optional | The scheme used if the proration was a downgrade. This is only present when the allocation was created mid-period. |
| `PricePointId` | `*int` | Optional | - |
| `PricePointName` | `*string` | Optional | - |
| `PricePointHandle` | `*string` | Optional | - |
| `PreviousPricePointId` | `*int` | Optional | - |
| `AccrueCharge` | `*bool` | Optional | If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately. |
| `UpgradeCharge` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `Payment` | [`Optional[models.AllocationPayment]`](allocation-payment.md) | Optional | - |

## Example (as JSON)

```json
{
  "component_id": 144,
  "subscription_id": 144,
  "quantity": 246,
  "previous_quantity": 180,
  "memo": "memo4"
}
```

