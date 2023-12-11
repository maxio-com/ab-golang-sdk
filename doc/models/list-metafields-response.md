
# List Metafields Response

## Structure

`ListMetafieldsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TotalCount` | `*int` | Optional | - |
| `CurrentPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `Metafields` | [`[]models.Metafield`](metafield.md) | Optional | - |

## Example (as JSON)

```json
{
  "total_count": 210,
  "current_page": 186,
  "total_pages": 198,
  "per_page": 92,
  "metafields": [
    {
      "id": 22,
      "name": "name2",
      "scope": {
        "csv": "0",
        "invoices": "0",
        "statements": "0",
        "portal": "0",
        "public_show": "0"
      },
      "data_count": 10,
      "input_type": "input_type4"
    }
  ]
}
```

