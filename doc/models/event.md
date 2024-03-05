
# Event

## Structure

`Event`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | - |
| `Key` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Message` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `SubscriptionId` | `*int` | Required | - |
| `CustomerId` | `*int` | Required | - |
| `CreatedAt` | `time.Time` | Required | - |
| `EventSpecificData` | `interface{}` | Required | - |

## Example (as JSON)

```json
{
  "id": 40,
  "key": "key2",
  "message": "message8",
  "subscription_id": 150,
  "customer_id": 78,
  "created_at": "2016-03-13T12:52:32.123Z",
  "event_specific_data": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

