
# Create or Update Endpoint

Used to Create or Update Endpoint

## Structure

`CreateOrUpdateEndpoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `string` | Required | - |
| `WebhookSubscriptions` | [`[]models.WebhookSubscription`](../../doc/models/webhook-subscription.md) | Required | - |

## Example (as JSON)

```json
{
  "url": "url8",
  "webhook_subscriptions": [
    "payment_success"
  ]
}
```

