
# EBB Event

## Structure

`EBBEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Chargify` | [`*models.ChargifyEBB`](../../doc/models/chargify-ebb.md) | Optional | - |

## Example (as JSON)

```json
{
  "chargify": {
    "subscription_id": 1,
    "timestamp": "2020-02-27T17:45:50-05:00",
    "id": "id6",
    "created_at": "created_at4",
    "uniqueness_token": "uniqueness_token2"
  },
  "messages": 150,
  "country": "US",
  "customer": {
    "name": "John",
    "lastName": "Doe",
    "address": {
      "street": "Maple Street",
      "zip": 4888,
      "state": "MA"
    }
  }
}
```

