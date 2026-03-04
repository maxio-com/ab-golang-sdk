
# Allocate Components

## Structure

`AllocateComponents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProrationUpgradeScheme` | `*string` | Optional | - |
| `ProrationDowngradeScheme` | `*string` | Optional | - |
| `Allocations` | [`[]models.CreateAllocation`](../../doc/models/create-allocation.md) | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `UpgradeCharge` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `DowngradeCredit` | [`models.Optional[models.CreditType]`](../../doc/models/credit-type.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided. |
| `PaymentCollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | (Optional) If not passed, the allocation(s) will use the payment collection method on the subscription |
| `InitiateDunning` | `*bool` | Optional | If true, if the immediate component payment fails, initiate dunning for the subscription.<br>Otherwise, leave the charges on the subscription to pay for at renewal. |

## Example (as JSON)

```json
{
  "proration_upgrade_scheme": "proration_upgrade_scheme2",
  "proration_downgrade_scheme": "proration_downgrade_scheme0",
  "allocations": [
    {
      "quantity": 26.48,
      "decimal_quantity": "decimal_quantity8",
      "previous_quantity": 55.5,
      "decimal_previous_quantity": "decimal_previous_quantity2",
      "component_id": 242,
      "memo": "memo6"
    },
    {
      "quantity": 26.48,
      "decimal_quantity": "decimal_quantity8",
      "previous_quantity": 55.5,
      "decimal_previous_quantity": "decimal_previous_quantity2",
      "component_id": 242,
      "memo": "memo6"
    }
  ],
  "accrue_charge": false,
  "upgrade_charge": "full"
}
```

