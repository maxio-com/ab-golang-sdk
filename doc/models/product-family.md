
# Product Family

## Structure

`ProductFamily`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | Timestamp indicating when this product family was archived. `null` if the product family is not archived. |

## Example (as JSON)

```json
{
  "id": 194,
  "name": "name2",
  "handle": "handle8",
  "accounting_code": "accounting_code8",
  "description": "description8"
}
```

