
# Batch Job

## Structure

`BatchJob`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `FinishedAt` | `Optional[time.Time]` | Optional | - |
| `RowCount` | `Optional[int]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
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

