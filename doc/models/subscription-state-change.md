
# Subscription State Change

## Structure

`SubscriptionStateChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousSubscriptionState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `NewSubscriptionState` | `string` | Required | **Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "previous_subscription_state": "previous_subscription_state2",
  "new_subscription_state": "new_subscription_state6"
}
```

