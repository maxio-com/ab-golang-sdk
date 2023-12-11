
# Allocate Components

## Structure

`AllocateComponents`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProrationUpgradeScheme` | `*string` | Optional | **Default**: `"no-prorate"` |
| `ProrationDowngradeScheme` | `*string` | Optional | **Default**: `"no-prorate"` |
| `Allocations` | [`[]models.CreateAllocationRequest`](create-allocation-request.md) | Optional | - |
| `AccrueCharge` | `*bool` | Optional | - |
| `UpgradeCharge` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `PaymentCollectionMethod` | [`*models.PaymentCollectionMethod1Enum`](payment-collection-method-1-enum.md) | Optional | (Optional) If not passed, the allocation(s) will use the payment collection method on the subscription<br>**Default**: `"automatic"` |

## Example (as JSON)

```json
{
  "proration_upgrade_scheme": "no-prorate",
  "proration_downgrade_scheme": "no-prorate",
  "payment_collection_method": "automatic",
  "allocations": [
    {
      "allocation": {
        "quantity": 228.94,
        "component_id": 8,
        "memo": "memo2",
        "proration_downgrade_scheme": "proration_downgrade_scheme4",
        "proration_upgrade_scheme": "proration_upgrade_scheme6",
        "accrue_charge": false
      }
    },
    {
      "allocation": {
        "quantity": 228.94,
        "component_id": 8,
        "memo": "memo2",
        "proration_downgrade_scheme": "proration_downgrade_scheme4",
        "proration_upgrade_scheme": "proration_upgrade_scheme6",
        "accrue_charge": false
      }
    }
  ],
  "accrue_charge": false,
  "upgrade_charge": "full"
}
```

