
# Chargify EBB

## Structure

`ChargifyEBB`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Timestamp` | `*time.Time` | Optional | This timestamp determines what billing period the event will be billed in. If your request payload does not include it, Chargify will add `chargify.timestamp` to the event payload and set the value to `now`. |
| `Id` | `*string` | Optional | A unique ID set by Chargify. This field is reserved. If `chargify.id` is present in the request payload, it will be overwritten. |
| `CreatedAt` | `*time.Time` | Optional | An ISO-8601 timestamp, set by Chargify at the time each event is recorded. This field is reserved. If `chargify.created_at` is present in the request payload, it will be overwritten. |
| `UniquenessToken` | `*string` | Optional | User-defined string scoped per-stream. Duplicate events within a stream will be silently ignored. Tokens expire after 31 days.<br><br>**Constraints**: *Maximum Length*: `64` |
| `SubscriptionId` | `*int` | Optional | Id of Maxio Advanced Billing Subscription which is connected to this event.<br>Provide `subscription_id` if you configured `chargify.subscription_id` as Subscription Identifier in your Event Stream. |
| `SubscriptionReference` | `*string` | Optional | Reference of Maxio Advanced Billing Subscription which is connected to this event.<br>Provide `subscription_reference` if you configured `chargify.subscription_reference` as Subscription Identifier in your Event Stream. |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    chargifyEBB := models.ChargifyEBB{
        Timestamp:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        Id:                    models.ToPointer("id2"),
        CreatedAt:             models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        UniquenessToken:       models.ToPointer("uniqueness_token8"),
        SubscriptionId:        models.ToPointer(224),
    }

}
```

