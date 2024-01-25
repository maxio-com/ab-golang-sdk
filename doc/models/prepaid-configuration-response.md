
# Prepaid Configuration Response

## Structure

`PrepaidConfigurationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PrepaidConfiguration` | [`models.PrepaidConfiguration`](../../doc/models/prepaid-configuration.md) | Required | - |

## Example (as JSON)

```json
{
  "prepaid_configuration": {
    "id": 142,
    "initial_funding_amount_in_cents": 74,
    "replenish_to_amount_in_cents": 76,
    "auto_replenish": false,
    "replenish_threshold_amount_in_cents": 20
  }
}
```

