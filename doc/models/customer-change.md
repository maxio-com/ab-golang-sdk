
# Customer Change

## Structure

`CustomerChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payer` | [`models.Optional[models.CustomerPayerChange]`](../../doc/models/customer-payer-change.md) | Optional | - |
| `ShippingAddress` | [`models.Optional[models.AddressChange]`](../../doc/models/address-change.md) | Optional | - |
| `BillingAddress` | [`models.Optional[models.AddressChange]`](../../doc/models/address-change.md) | Optional | - |
| `CustomFields` | [`models.Optional[models.CustomerCustomFieldsChange]`](../../doc/models/customer-custom-fields-change.md) | Optional | - |

## Example (as JSON)

```json
{
  "payer": {
    "before": {
      "first_name": "first_name0",
      "last_name": "last_name8",
      "organization": "organization4",
      "email": "email6"
    },
    "after": {
      "first_name": "first_name2",
      "last_name": "last_name0",
      "organization": "organization4",
      "email": "email4"
    }
  },
  "shipping_address": {
    "before": {
      "street": "street0",
      "line2": "line24",
      "city": "city0",
      "state": "state6",
      "zip": "zip4"
    },
    "after": {
      "street": "street2",
      "line2": "line26",
      "city": "city8",
      "state": "state2",
      "zip": "zip4"
    }
  },
  "billing_address": {
    "before": {
      "street": "street0",
      "line2": "line24",
      "city": "city0",
      "state": "state6",
      "zip": "zip4"
    },
    "after": {
      "street": "street2",
      "line2": "line26",
      "city": "city8",
      "state": "state2",
      "zip": "zip4"
    }
  },
  "custom_fields": {
    "before": [
      {
        "owner_id": 26,
        "owner_type": "Customer",
        "name": "name0",
        "value": "value2",
        "metadatum_id": 26
      },
      {
        "owner_id": 26,
        "owner_type": "Customer",
        "name": "name0",
        "value": "value2",
        "metadatum_id": 26
      }
    ],
    "after": [
      {
        "owner_id": 130,
        "owner_type": "Customer",
        "name": "name2",
        "value": "value4",
        "metadatum_id": 130
      },
      {
        "owner_id": 130,
        "owner_type": "Customer",
        "name": "name2",
        "value": "value4",
        "metadatum_id": 130
      },
      {
        "owner_id": 130,
        "owner_type": "Customer",
        "name": "name2",
        "value": "value4",
        "metadatum_id": 130
      }
    ]
  }
}
```

