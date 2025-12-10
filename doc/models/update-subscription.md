
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
| `SnapDay` | [`models.Optional[models.UpdateSubscriptionSnapDay]`](../../doc/models/containers/update-subscription-snap-day.md) | Optional | This is a container for one-of cases. |
| `InitialBillingAt` | `*time.Time` | Optional | (Optional) Set this attribute to a future date/time to update a subscription in the Awaiting Signup Date state, to Awaiting Signup. In the Awaiting Signup state, a subscription behaves like any other. It can be canceled, allocated to, or have its billing date changed. etc. When the `initial_billing_at` date hits, the subscription will transition to the expected state. If the product has a trial, the subscription will enter a trial, otherwise it will go active. Setup fees will be respected either before or after the trial, as configured on the price point. If the payment is due at the initial_billing_at and it fails the subscription will be immediately canceled. You can omit the initial_billing_at date to activate the subscription immediately. See the [subscription import](https://maxio.zendesk.com/hc/en-us/articles/24251489107213-Advanced-Billing-Subscription-Imports#date-format) documentation for more information about Date/Time formats. |
| `DeferSignup` | `*bool` | Optional | (Optional) Set this attribute to true to move the subscription from Awaiting Signup, to Awaiting Signup Date. Use this when you want to update a subscription that has an unknown initial billing date. When the first billing date is known, update a subscription to set the `initial_billing_at` date. The subscription moves to the awaiting signup with a scheduled initial billing date. You can omit the initial_billing_at date to activate the subscription immediately. See [Subscription States](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404222005773-Subscription-States) for more information.<br><br>**Default**: `false` |
| `NextBillingAt` | `*time.Time` | Optional | - |
| `ExpiresAt` | `*time.Time` | Optional | Timestamp giving the expiration date of this subscription (if any). You may manually change the expiration date at any point during a subscription period. |
| `PaymentCollectionMethod` | `*string` | Optional | - |
| `ReceivesInvoiceEmails` | `*bool` | Optional | - |
| `NetTerms` | [`*models.UpdateSubscriptionNetTerms`](../../doc/models/containers/update-subscription-net-terms.md) | Optional | This is a container for one-of cases. |
| `StoredCredentialTransactionId` | `*int` | Optional | - |
| `Reference` | `*string` | Optional | - |
| `CustomPrice` | [`*models.SubscriptionCustomPrice`](../../doc/models/subscription-custom-price.md) | Optional | (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription |
| `Components` | [`[]models.UpdateSubscriptionComponent`](../../doc/models/update-subscription-component.md) | Optional | (Optional) An array of component ids and custom prices to be added to the subscription. |
| `DunningCommunicationDelayEnabled` | `*bool` | Optional | Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute. |
| `DunningCommunicationDelayTimeZone` | `models.Optional[string]` | Optional | Time zone for the Dunning Communication Delay feature. |
| `ProductPricePointId` | `*int` | Optional | Set to change the current product's price point. |
| `ProductPricePointHandle` | `*string` | Optional | Set to change the current product's price point. |

## Example (as JSON)

```json
{
  "defer_signup": false,
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

