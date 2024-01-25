
# Usage Response

## Structure

`UsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | [`models.Usage`](../../doc/models/usage.md) | Required | - |

## Example (as JSON)

```json
{
  "usage": {
    "id": 150,
    "memo": "memo2",
    "created_at": "2016-03-13T12:52:32.123Z",
    "price_point_id": 28,
    "quantity": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

