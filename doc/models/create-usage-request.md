
# Create Usage Request

## Structure

`CreateUsageRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | [`models.CreateUsage`](../../doc/models/create-usage.md) | Required | - |

## Example (as JSON)

```json
{
  "usage": {
    "quantity": 162.34,
    "price_point_id": "price_point_id0",
    "memo": "memo2",
    "billing_schedule": {
      "initial_billing_at": "2016-03-13T12:52:32.123Z"
    },
    "custom_price": {
      "tax_included": false,
      "pricing_scheme": "stairstep",
      "interval": 66,
      "interval_unit": "day",
      "prices": [
        {
          "starting_quantity": 242,
          "ending_quantity": 40,
          "unit_price": 23.26
        },
        {
          "starting_quantity": 242,
          "ending_quantity": 40,
          "unit_price": 23.26
        }
      ],
      "renew_prepaid_allocation": false
    }
  }
}
```

