
# Subscription Group Signup Component

## Structure

`SubscriptionGroupSignupComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*interface{}` | Optional | Required if passing any component to `components` attribute. |
| `AllocatedQuantity` | `*interface{}` | Optional | - |
| `UnitBalance` | `*interface{}` | Optional | - |
| `PricePointId` | `*interface{}` | Optional | - |
| `CustomPrice` | [`*models.SubscriptionGroupComponentCustomPrice`](subscription-group-component-custom-price.md) | Optional | Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`. |

## Example (as JSON)

```json
{
  "component_id": {
    "key1": "val1",
    "key2": "val2"
  },
  "allocated_quantity": {
    "key1": "val1",
    "key2": "val2"
  },
  "unit_balance": {
    "key1": "val1",
    "key2": "val2"
  },
  "price_point_id": {
    "key1": "val1",
    "key2": "val2"
  },
  "custom_price": {
    "pricing_scheme": "stairstep",
    "prices": [
      {
        "starting_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "ending_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      },
      {
        "starting_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "ending_quantity": {
          "key1": "val1",
          "key2": "val2"
        },
        "unit_price": {
          "key1": "val1",
          "key2": "val2"
        }
      }
    ],
    "overage_pricing": [
      {
        "pricing_scheme": "stairstep",
        "prices": [
          {
            "starting_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "ending_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "unit_price": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        ]
      },
      {
        "pricing_scheme": "stairstep",
        "prices": [
          {
            "starting_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "ending_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "unit_price": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        ]
      },
      {
        "pricing_scheme": "stairstep",
        "prices": [
          {
            "starting_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "ending_quantity": {
              "key1": "val1",
              "key2": "val2"
            },
            "unit_price": {
              "key1": "val1",
              "key2": "val2"
            }
          }
        ]
      }
    ]
  }
}
```

