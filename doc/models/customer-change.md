
# Customer Change

## Structure

`CustomerChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Payer` | [`models.Optional[models.CustomerChangePayer]`](../../doc/models/containers/customer-change-payer.md) | Optional | This is a container for one-of cases. |
| `ShippingAddress` | [`models.Optional[models.CustomerChangeShippingAddress]`](../../doc/models/containers/customer-change-shipping-address.md) | Optional | This is a container for one-of cases. |
| `BillingAddress` | [`models.Optional[models.CustomerChangeBillingAddress]`](../../doc/models/containers/customer-change-billing-address.md) | Optional | This is a container for one-of cases. |
| `CustomFields` | [`models.Optional[models.CustomerChangeCustomFields]`](../../doc/models/containers/customer-change-custom-fields.md) | Optional | This is a container for one-of cases. |

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
      }
    ]
  }
}
```

