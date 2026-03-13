
# Create Allocation

## Structure

`CreateAllocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Quantity` | `float64` | Required | The allocated quantity to which to set the line-items allocated quantity. By default, this is an integer. If decimal allocations are enabled for the component, it will be a decimal number. For On/Off components, use 1for on and 0 for off. |
| `DecimalQuantity` | `*string` | Optional | Decimal representation of the allocated quantity. Only valid when decimal<br>allocations are enabled for the component. |
| `PreviousQuantity` | `*float64` | Optional | The quantity that was in effect before this allocation. Responses always<br>include this value; it may be supplied on preview requests to ensure the<br>expected change is evaluated. |
| `DecimalPreviousQuantity` | `*string` | Optional | Decimal representation of `previous_quantity`. Only valid when decimal<br>allocations are enabled for the component. |
| `ComponentId` | `*int` | Optional | (required for the multiple allocations endpoint) The id associated with the component for which the allocation is being made. |
| `Memo` | `*string` | Optional | A memo to record along with the allocation. |
| `ProrationDowngradeScheme` | `*string` | Optional | The scheme used if the proration is a downgrade. Defaults to the site setting if one is not provided. |
| `ProrationUpgradeScheme` | `*string` | Optional | The scheme used if the proration is an upgrade. Defaults to the site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.DowngradeCreditCreditType]`](../../doc/models/downgrade-credit-credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. Values are:<br><br>`full` -  A full price credit is added for the amount owed.<br><br>`prorated` - A prorated credit is added for the amount owed.<br><br>`none` - No charge is added. |
| `UpgradeCharge` | [`models.Optional[models.UpgradeChargeCreditType]`](../../doc/models/upgrade-charge-credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. Values are:<br><br>`full` - A charge is added for the full price of the component.<br><br>`prorated` - A charge is added for the prorated price of the component change.<br><br>`none` - No charge is added. |
| `AccrueCharge` | `*bool` | Optional | "If the change in cost is an upgrade, this determines if the charge should accrue to the next renewal or if capture should be attempted immediately.<br><br>`true` - Attempt to charge the customer at the next renewal.<br><br>`false` - Attempt to charge the customer right away. If it fails, the charge will be accrued until the next renewal.<br><br>Defaults to the site setting if unspecified in the request. |
| `InitiateDunning` | `*bool` | Optional | If set to true, if the immediate component payment fails, initiate dunning for the subscription.<br>Otherwise, leave the charges on the subscription to pay for at renewal. Defaults to false. |
| `PricePointId` | [`models.Optional[models.CreateAllocationPricePointId]`](../../doc/models/containers/create-allocation-price-point-id.md) | Optional | This is a container for one-of cases. |
| `BillingSchedule` | [`*models.BillingSchedule`](../../doc/models/billing-schedule.md) | Optional | This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. This only works for site with Multifrequency enabled. |
| `CustomPrice` | [`*models.ComponentCustomPrice`](../../doc/models/component-custom-price.md) | Optional | Create or update custom pricing unique to the subscription. Used in place of `price_point_id`. |

## Example (as JSON)

```json
{
  "quantity": 8.06,
  "decimal_quantity": "decimal_quantity4",
  "previous_quantity": 218.92,
  "decimal_previous_quantity": "decimal_previous_quantity0",
  "component_id": 192,
  "memo": "memo4"
}
```

