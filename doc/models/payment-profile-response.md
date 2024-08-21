
# Payment Profile Response

## Structure

`PaymentProfileResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.PaymentProfile`](../../doc/models/containers/payment-profile.md) | Required | - |

## Example (as JSON)

```json
{
  "payment_profile": {
    "payment_type": "apple_pay",
    "id": 60,
    "first_name": "first_name2",
    "last_name": "last_name0",
    "customer_id": 98,
    "current_vault": "braintree_blue"
  }
}
```

