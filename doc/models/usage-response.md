
# Usage Response

## Structure

`UsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Usage` | [`models.Usage`](usage.md) | Required | - |

## Example (as JSON)

```json
{
  "usage": {
    "id": 150,
    "memo": "memo2",
    "created_at": "created_at6",
    "price_point_id": 28,
    "quantity": {
      "key1": "val1",
      "key2": "val2"
    }
  }
}
```

