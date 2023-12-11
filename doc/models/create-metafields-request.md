
# Create Metafields Request

## Structure

`CreateMetafieldsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metafields` | [`models.Metafields`](metafields.md) | Required | - |

## Example (as JSON)

```json
{
  "metafields": {
    "input_type": "text",
    "id": 22,
    "name": "name2",
    "scope": {
      "csv": "0",
      "invoices": "0",
      "statements": "0",
      "portal": "0",
      "public_show": "0"
    },
    "enum": [
      "enum8",
      "enum9",
      "enum0"
    ]
  }
}
```

