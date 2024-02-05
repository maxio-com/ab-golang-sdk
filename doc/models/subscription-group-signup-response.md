
# Subscription Group Signup Response

## Structure

`SubscriptionGroupSignupResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `Scheme` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `PaymentProfileId` | `*int` | Optional | - |
| `SubscriptionIds` | `[]int` | Optional | - |
| `PrimarySubscriptionId` | `*int` | Optional | - |
| `NextAssessmentAt` | `*string` | Optional | - |
| `State` | `*string` | Optional | - |
| `CancelAtEndOfPeriod` | `*bool` | Optional | - |
| `Subscriptions` | [`[]models.SubscriptionGroupItem`](../../doc/models/subscription-group-item.md) | Optional | - |
| `PaymentCollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.<br>**Default**: `"automatic"` |

## Example (as JSON)

```json
{
  "payment_collection_method": "automatic",
  "uid": "uid8",
  "scheme": 28,
  "customer_id": 48,
  "payment_profile_id": 44,
  "subscription_ids": [
    158,
    159,
    160
  ]
}
```

