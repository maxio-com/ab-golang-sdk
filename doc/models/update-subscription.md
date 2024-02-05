
# Update Subscription

## Structure

`UpdateSubscription`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditCardAttributes` | [`*models.CreditCardAttributes`](../../doc/models/credit-card-attributes.md) | Optional | - |
| `ProductHandle` | `*string` | Optional | Set to the handle of a different product to change the subscription's product |
| `ProductId` | `*int` | Optional | Set to the id of a different product to change the subscription's product |
| `ProductChangeDelayed` | `*bool` | Optional | - |
| `NextProductId` | `*string` | Optional | Set to an empty string to cancel a delayed product change. |
| `NextProductPricePointId` | `*string` | Optional | - |
| `SnapDay` | `*interface{}` | Optional | Use for subscriptions with product eligible for calendar billing only. Value can be 1-28 or 'end'. |
| `NextBillingAt` | `*string` | Optional | - |
| `PaymentCollectionMethod` | `*string` | Optional | - |
| `ReceivesInvoiceEmails` | `*bool` | Optional | - |
| `NetTerms` | `*interface{}` | Optional | - |
| `StoredCredentialTransactionId` | `*int` | Optional | - |
| `Reference` | `*string` | Optional | - |
| `CustomPrice` | [`*models.SubscriptionCustomPrice`](../../doc/models/subscription-custom-price.md) | Optional | (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription |
| `Components` | [`[]models.UpdateSubscriptionComponent`](../../doc/models/update-subscription-component.md) | Optional | (Optional) An array of component ids and custom prices to be added to the subscription. |
| `DunningCommunicationDelayEnabled` | `Optional[bool]` | Optional | Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute. |
| `DunningCommunicationDelayTimeZone` | `Optional[string]` | Optional | Time zone for the Dunning Communication Delay feature. |

## Example (as JSON)

```json
{
  "dunning_communication_delay_time_zone": "\"Eastern Time (US & Canada)\"",
  "credit_card_attributes": {
    "full_number": "full_number2",
    "expiration_month": "expiration_month6",
    "expiration_year": "expiration_year2"
  },
  "product_handle": "product_handle2",
  "product_id": 114,
  "product_change_delayed": false,
  "next_product_id": "next_product_id8"
}
```

