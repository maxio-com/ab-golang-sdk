
# Update Metafields Request

## Structure

`UpdateMetafieldsRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Metafields` | [`*models.Metafields1`](metafields-1.md) | Optional | - |

## Example (as JSON)

```json
{
  "metafields": {
    "current_name": "current_name6",
    "name": "name2",
    "scope": {
      "csv": "0",
      "invoices": "0",
      "statements": "0",
      "portal": "0",
      "public_show": "0"
    },
    "input_type": "balance_tracker",
    "enum": [
      "enum8",
      "enum9",
      "enum0"
    ]
  }
}
```

