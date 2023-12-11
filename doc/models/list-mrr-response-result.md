
# List MRR Response Result

## Structure

`ListMRRResponseResult`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Page` | `*int` | Optional | - |
| `PerPage` | `*int` | Optional | - |
| `TotalPages` | `*int` | Optional | - |
| `TotalEntries` | `*int` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `CurrencySymbol` | `*string` | Optional | - |
| `Movements` | [`[]models.Movement`](movement.md) | Optional | - |

## Example (as JSON)

```json
{
  "page": 150,
  "per_page": 238,
  "total_pages": 16,
  "total_entries": 112,
  "currency": "currency8"
}
```

