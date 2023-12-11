
# Renewal Preview Request

## Structure

`RenewalPreviewRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Components` | [`[]models.RenewalPreviewComponent`](renewal-preview-component.md) | Optional | An optional array of component definitions to preview. Providing any component definitions here will override the actual components on the subscription (and their quantities), and the billing preview will contain only these components (in addition to any product base fees). |

## Example (as JSON)

```json
{
  "components": [
    {
      "component_id": {
        "key1": "val1",
        "key2": "val2"
      },
      "quantity": 210,
      "price_point_id": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "component_id": {
        "key1": "val1",
        "key2": "val2"
      },
      "quantity": 210,
      "price_point_id": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "component_id": {
        "key1": "val1",
        "key2": "val2"
      },
      "quantity": 210,
      "price_point_id": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ]
}
```

