
# Custom Field Value Change

## Structure

`CustomFieldValueChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EventType` | `string` | Required | - |
| `MetafieldName` | `string` | Required | - |
| `MetafieldId` | `int` | Required | - |
| `OldValue` | `*string` | Required | - |
| `NewValue` | `*string` | Required | - |
| `ResourceType` | `string` | Required | - |
| `ResourceId` | `int` | Required | - |

## Example (as JSON)

```json
{
  "event_type": "event_type2",
  "metafield_name": "metafield_name6",
  "metafield_id": 78,
  "old_value": "old_value2",
  "new_value": "new_value8",
  "resource_type": "resource_type2",
  "resource_id": 74
}
```

