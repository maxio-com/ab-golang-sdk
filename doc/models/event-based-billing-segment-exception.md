
# Event Based Billing Segment Exception

## Structure

`EventBasedBillingSegmentException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.EventBasedBillingSegmentError`](../../doc/models/event-based-billing-segment-error.md) | Required | - |

## Example (as JSON)

```json
{
  "errors": {
    "segments": {
      "key0": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  }
}
```

