
# Renewal Preview Request

## Structure

`RenewalPreviewRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Components` | [`[]models.RenewalPreviewComponent`](../../doc/models/renewal-preview-component.md) | Optional | An optional array of component definitions to preview. Providing any component definitions here will override the actual components on the subscription (and their quantities), and the billing preview will contain only these components (in addition to any product base fees). |

## Example (as JSON)

```json
{
  "components": [
    {
      "component_id": "String5",
      "quantity": 210,
      "price_point_id": "String3"
    },
    {
      "component_id": "String5",
      "quantity": 210,
      "price_point_id": "String3"
    },
    {
      "component_id": "String5",
      "quantity": 210,
      "price_point_id": "String3"
    }
  ]
}
```

