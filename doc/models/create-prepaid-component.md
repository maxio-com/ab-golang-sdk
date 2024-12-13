
# Create Prepaid Component

## Structure

`CreatePrepaidComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidUsageComponent` | [`models.PrepaidUsageComponent`](../../doc/models/prepaid-usage-component.md) | Required | - |

## Example (as JSON)

```json
{
  "prepaid_usage_component": {
    "name": "name2",
    "unit_name": "unit_name4",
    "description": "description2",
    "handle": "handle8",
    "taxable": false,
    "pricing_scheme": "per_unit",
    "prices": [
      {
        "starting_quantity": 242,
        "ending_quantity": 40,
        "unit_price": 23.26
      }
    ],
    "upgrade_charge": "full",
    "overage_pricing": {
      "pricing_scheme": "stairstep",
      "prices": [
        {
          "starting_quantity": 242,
          "ending_quantity": 40,
          "unit_price": 23.26
        }
      ]
    }
  }
}
```

