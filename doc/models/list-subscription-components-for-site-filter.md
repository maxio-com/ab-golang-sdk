
# List Subscription Components for Site Filter

## Structure

`ListSubscriptionComponentsForSiteFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currencies` | `[]string` | Optional | Allows fetching components allocation with matching currency based on provided values. Use in query `filter[currencies]=USD,EUR`.<br><br>**Constraints**: *Minimum Items*: `1` |
| `UseSiteExchangeRate` | `*bool` | Optional | Allows fetching components allocation with matching use_site_exchange_rate based on provided value. Use in query `filter[use_site_exchange_rate]=true`. |
| `Subscription` | [`*models.SubscriptionFilter`](../../doc/models/subscription-filter.md) | Optional | Nested filter used for List Subscription Components For Site Filter |

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
    listSubscriptionComponentsForSiteFilter := models.ListSubscriptionComponentsForSiteFilter{
        Currencies:           []string{
            "EUR",
            "USD",
        },
        UseSiteExchangeRate:  models.ToPointer(false),
        Subscription:         models.ToPointer(models.SubscriptionFilter{
            States:               []models.SubscriptionStateFilter{
                models.SubscriptionStateFilter_TRIALING,
                models.SubscriptionStateFilter_UNPAID,
                models.SubscriptionStateFilter_ACTIVE,
            },
            DateField:            models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
            StartDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            EndDate:              models.ToPointer(parseTime(models.DEFAULT_DATE, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
            StartDatetime:        models.ToPointer(parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) })),
        }),
    }

}
```

