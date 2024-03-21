
# Batch Job

## Structure

`BatchJob`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FinishedAt` | `models.Optional[time.Time]` | Optional | - |
| `RowCount` | `models.Optional[int]` | Optional | - |
| `CreatedAt` | `models.Optional[time.Time]` | Optional | - |
| `Completed` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "id": 246,
  "finished_at": "2016-03-13T12:52:32.123Z",
  "row_count": 254,
  "created_at": "2016-03-13T12:52:32.123Z",
  "completed": "completed0"
}
```

