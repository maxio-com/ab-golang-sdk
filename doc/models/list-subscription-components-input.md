
# List Subscription Components Input

Input structure for the method ListSubscriptionComponents

## Structure

`ListSubscriptionComponentsInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `int` | Required | The Chargify id of the subscription. |
| `DateField` | [`*models.SubscriptionListDateField`](../../doc/models/subscription-list-date-field.md) | Optional | The type of filter you'd like to apply to your search. Use in query `date_field=updated_at`. |
| `Direction` | [`*models.SortingDirection`](../../doc/models/sorting-direction.md) | Optional | Controls the order in which results are returned.<br>Use in query `direction=asc`. |
| `Filter` | [`*models.ListSubscriptionComponentsFilter`](../../doc/models/list-subscription-components-filter.md) | Optional | Filter to use for List Subscription Components operation |
| `EndDate` | `*string` | Optional | The end date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp up to and including 11:59:59PM in your site’s time zone on the date specified. |
| `EndDatetime` | `*string` | Optional | The end date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or before exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of end_date. |
| `PricePointIds` | [`*models.IncludeNotNull`](../../doc/models/include-not-null.md) | Optional | Allows fetching components allocation only if price point id is present. Use in query `price_point_ids=not_null`. |
| `ProductFamilyIds` | `[]int` | Optional | Allows fetching components allocation with matching product family id based on provided ids. Use in query `product_family_ids=1,2,3`. |
| `Sort` | [`*models.ListSubscriptionComponentsSort`](../../doc/models/list-subscription-components-sort.md) | Optional | The attribute by which to sort. Use in query `sort=updated_at`. |
| `StartDate` | `*string` | Optional | The start date (format YYYY-MM-DD) with which to filter the date_field. Returns components with a timestamp at or after midnight (12:00:00 AM) in your site’s time zone on the date specified. |
| `StartDatetime` | `*string` | Optional | The start date and time (format YYYY-MM-DD HH:MM:SS) with which to filter the date_field. Returns components with a timestamp at or after exact time provided in query. You can specify timezone in query - otherwise your site''s time zone will be used. If provided, this parameter will be used instead of start_date. |
| `Include` | [`[]models.ListSubscriptionComponentsInclude`](../../doc/models/list-subscription-components-include.md) | Optional | Allows including additional data in the response. Use in query `include=subscription,historic_usages`. |
| `InUse` | `*bool` | Optional | If in_use is set to true, it returns only components that are currently in use. However, if it's set to false or not provided, it returns all components connected with the subscription. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listSubscriptionComponentsInput := models.ListSubscriptionComponentsInput{
        SubscriptionId:       222,
        DateField:            models.ToPointer(models.SubscriptionListDateField_UPDATEDAT),
        Direction:            models.ToPointer(models.SortingDirection_ASC),
        Filter:               models.ToPointer(models.ListSubscriptionComponentsFilter{
            Currencies:           []string{
                "EUR",
                "USD",
            },
            UseSiteExchangeRate:  models.ToPointer(false),
        }),
        EndDate:              models.ToPointer("end_date0"),
        EndDatetime:          models.ToPointer("end_datetime8"),
        PricePointIds:        models.ToPointer(models.IncludeNotNull_NOTNULL),
        ProductFamilyIds:     []int{
            1,
            2,
            3,
        },
        Sort:                 models.ToPointer(models.ListSubscriptionComponentsSort_UPDATEDAT),
        StartDate:            models.ToPointer("start_date6"),
        StartDatetime:        models.ToPointer("start_datetime0"),
        Include:              []models.ListSubscriptionComponentsInclude{
            models.ListSubscriptionComponentsInclude_SUBSCRIPTION,
            models.ListSubscriptionComponentsInclude_HISTORICUSAGES,
        },
        InUse:                models.ToPointer(true),
    }

}
```

