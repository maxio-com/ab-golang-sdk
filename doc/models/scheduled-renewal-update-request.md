
# Scheduled Renewal Update Request

## Structure

`ScheduledRenewalUpdateRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalConfigurationItem` | [`models.ScheduledRenewalUpdateRequestRenewalConfigurationItem`](../../doc/models/containers/scheduled-renewal-update-request-renewal-configuration-item.md) | Required | This is a container for one-of cases. |

## Example (as JSON)

```json
{
  "renewal_configuration_item": {
    "item_type": "Component",
    "item_id": 108,
    "price_point_id": 122,
    "quantity": 212,
    "custom_price": {
      "tax_included": false,
      "pricing_scheme": "stairstep",
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
      ]
    }
  }
}
```

