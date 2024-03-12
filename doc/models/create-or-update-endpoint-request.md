
# Create or Update Endpoint Request

Used to Create or Update Endpoint

## Structure

`CreateOrUpdateEndpointRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Endpoint` | [`models.CreateOrUpdateEndpoint`](../../doc/models/create-or-update-endpoint.md) | Required | Used to Create or Update Endpoint |

## Example (as JSON)

```json
{
  "endpoint": {
    "url": "url2",
    "webhook_subscriptions": [
      "dunning_step_reached"
    ]
  }
}
```

