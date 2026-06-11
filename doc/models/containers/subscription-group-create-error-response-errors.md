
# Subscription Group Create Error Response Errors

## Class Name

`SubscriptionGroupCreateErrorResponseErrors`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.SubscriptionGroupMembersArrayError`](../../../doc/models/subscription-group-members-array-error.md) | models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromSubscriptionGroupMembersArrayError(models.SubscriptionGroupMembersArrayError subscriptionGroupMembersArrayError) |
| [`models.SubscriptionGroupSingleError`](../../../doc/models/subscription-group-single-error.md) | models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromSubscriptionGroupSingleError(models.SubscriptionGroupSingleError subscriptionGroupSingleError) |
| `string` | models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromString(string mString) |

## models.SubscriptionGroupMembersArrayError

### Initialization Code

#### Example

```go
value := models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromSubscriptionGroupMembersArrayError(models.SubscriptionGroupMembersArrayError{
    Members:              []string{
        "members6",
    },
})
```

## models.SubscriptionGroupSingleError

### Initialization Code

#### Example

```go
value := models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromSubscriptionGroupSingleError(models.SubscriptionGroupSingleError{
    SubscriptionGroup:    "subscription_group2",
})
```

## string

### Initialization Code

#### Example

```go
value := models.SubscriptionGroupCreateErrorResponseErrorsContainer.FromString("String0")
```

