
# Customer Changes Preview Response

## Structure

`CustomerChangesPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Changes` | [`models.CustomerChange`](../../doc/models/customer-change.md) | Required | - |

## Example (as JSON)

```json
{
  "changes": {
    "payer": {
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "after": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "shipping_address": {
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "after": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "billing_address": {
      "before": {
        "key1": "val1",
        "key2": "val2"
      },
      "after": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "custom_fields": {
      "before": [
        {
          "owner_id": 26,
          "owner_type": "owner_type2",
          "name": "name0",
          "value": "value2",
          "metadatum_id": 26
        },
        {
          "owner_id": 26,
          "owner_type": "owner_type2",
          "name": "name0",
          "value": "value2",
          "metadatum_id": 26
        }
      ],
      "after": [
        {
          "owner_id": 130,
          "owner_type": "owner_type4",
          "name": "name2",
          "value": "value4",
          "metadatum_id": 130
        },
        {
          "owner_id": 130,
          "owner_type": "owner_type4",
          "name": "name2",
          "value": "value4",
          "metadatum_id": 130
        },
        {
          "owner_id": 130,
          "owner_type": "owner_type4",
          "name": "name2",
          "value": "value4",
          "metadatum_id": 130
        }
      ]
    }
  }
}
```

