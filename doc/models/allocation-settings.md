
# Allocation Settings

## Structure

`AllocationSettings`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `UpgradeCharge` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `DowngradeCredit` | [`Optional[models.CreditTypeEnum]`](credit-type-enum.md) | Optional | The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.<br>Available values: `full`, `prorated`, `none`. |
| `AccrueCharge` | `*bool` | Optional | - |

## Example (as JSON)

```json
{
  "upgrade_charge": "none",
  "downgrade_credit": "prorated",
  "accrue_charge": false
}
```

