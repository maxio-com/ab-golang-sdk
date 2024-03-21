# Webhooks

```go
webhooksController := client.WebhooksController()
```

## Class Name

`WebhooksController`

## Methods

* [List Webhooks](../../doc/controllers/webhooks.md#list-webhooks)
* [Enable Webhooks](../../doc/controllers/webhooks.md#enable-webhooks)
* [Replay Webhooks](../../doc/controllers/webhooks.md#replay-webhooks)
* [Create Endpoint](../../doc/controllers/webhooks.md#create-endpoint)
* [List Endpoints](../../doc/controllers/webhooks.md#list-endpoints)
* [Update Endpoint](../../doc/controllers/webhooks.md#update-endpoint)


# List Webhooks

## Webhooks Intro

The Webhooks API allows you to view a list of all webhooks and to selectively resend individual or groups of webhooks. Webhooks will be sent on endpoints specified by you. Endpoints can be added via API or Web UI. There is also an option to enable / disable webhooks via API request.

We recommend that you review Chargify's webhook documentation located in our help site. The following resources will help guide you on how to use webhooks in Chargify, in addition to these webhook endpoints:

+ [Adding/editing new webhooks](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404448450317#configure-webhook-url)
+ [Webhooks introduction and delivery information](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405568068365#webhooks-introduction-0-0)
+ [Main webhook overview](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#webhooks-reference-0-0)
+ [Available webhooks and payloads](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#events)

## List Webhooks for a Site

This method allows you to fetch data about webhooks. You can pass query parameters if you want to filter webhooks.

```go
ListWebhooks(
    ctx context.Context,
    input ListWebhooksInput) (
    models.ApiResponse[[]models.WebhookResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `status` | [`*models.WebhookStatus`](../../doc/models/webhook-status.md) | Query, Optional | Webhooks with matching status would be returned. |
| `sinceDate` | `*string` | Query, Optional | Format YYYY-MM-DD. Returns Webhooks with the created_at date greater than or equal to the one specified. |
| `untilDate` | `*string` | Query, Optional | Format YYYY-MM-DD. Returns Webhooks with the created_at date less than or equal to the one specified. |
| `page` | `*int` | Query, Optional | Result records are organized in pages. By default, the first page of results is displayed. The page parameter specifies a page number of results to fetch. You can start navigating through the pages to consume the results. You do this by passing in a page parameter. Retrieve the next page by adding ?page=2 to the query string. If there are no results to return, then an empty result set will be returned.<br>Use in query `page=1`. |
| `perPage` | `*int` | Query, Optional | This parameter indicates how many records to fetch in each request. Default value is 20. The maximum allowed values is 200; any per_page value over 200 will be changed to 200.<br>Use in query `per_page=200`. |
| `order` | [`*models.WebhookOrder`](../../doc/models/webhook-order.md) | Query, Optional | The order in which the Webhooks are returned. |
| `subscription` | `*int` | Query, Optional | The Chargify id of a subscription you'd like to filter for |

## Response Type

[`[]models.WebhookResponse`](../../doc/models/webhook-response.md)

## Example Usage

```go
ctx := context.Background()

collectedInput := advancedbilling.ListWebhooksInput{
    Page:         models.ToPointer(2),
    PerPage:      models.ToPointer(50),
}

apiResponse, err := webhooksController.ListWebhooks(ctx, collectedInput)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "webhook": {
      "event": "statement_settled",
      "id": 141765032,
      "created_at": "2016-11-08T16:22:26-05:00",
      "last_error": "404 Resource Not Found (retry 5 of 5)",
      "last_error_at": "2016-11-08T16:43:54-05:00",
      "accepted_at": null,
      "last_sent_at": "2016-11-08T16:43:54-05:00",
      "last_sent_url": "http://requestb.in/11u45x71",
      "successful": false,
      "body": "id=141765032&event=statement_settled&payload[site][id]=31615&payload[site][subdomain]=general-goods&payload[subscription][id]=15100141&payload[subscription][state]=active&payload[subscription][balance_in_cents]=0&payload[customer][id]=14585695&payload[customer][first_name]=Pookie&payload[customer][last_name]=Test&payload[customer][reference]=&payload[customer][organization]=&payload[customer][address]=&payload[customer][address_2]=&payload[customer][city]=&payload[customer][state]=&payload[customer][zip]=&payload[customer][country]=&payload[customer][email]=pookie999%40example.com&payload[customer][phone]=&payload[statement][closed_at]=2016-11-08%2016%3A22%3A20%20-0500&payload[statement][created_at]=2016-11-08%2016%3A22%3A18%20-0500&payload[statement][id]=80168049&payload[statement][opened_at]=2016-11-07%2016%3A22%3A15%20-0500&payload[statement][settled_at]=2016-11-08%2016%3A22%3A20%20-0500&payload[statement][subscription_id]=15100141&payload[statement][updated_at]=2016-11-08%2016%3A22%3A20%20-0500&payload[statement][starting_balance_in_cents]=0&payload[statement][ending_balance_in_cents]=0&payload[statement][total_in_cents]=6400&payload[statement][memo]=We%20thank%20you%20for%20your%20continued%20business!&payload[statement][events][0][id]=346956565&payload[statement][events][0][key]=renewal_success&payload[statement][events][0][message]=Successful%20renewal%20for%20Pookie%20Test's%20subscription%20to%20%2410%20Basic%20Plan&payload[statement][events][1][id]=346956579&payload[statement][events][1][key]=payment_success&payload[statement][events][1][message]=Successful%20payment%20of%20%2464.00%20for%20Pookie%20Test's%20subscription%20to%20%2410%20Basic%20Plan&payload[statement][events][2][id]=347299359&payload[statement][events][2][key]=renewal_success&payload[statement][events][2][message]=Successful%20renewal%20for%20Pookie%20Test's%20subscription%20to%20%2410%20Basic%20Plan&payload[statement][transactions][0][id]=161537343&payload[statement][transactions][0][subscription_id]=15100141&payload[statement][transactions][0][type]=Charge&payload[statement][transactions][0][kind]=baseline&payload[statement][transactions][0][transaction_type]=charge&payload[statement][transactions][0][success]=true&payload[statement][transactions][0][amount_in_cents]=1000&payload[statement][transactions][0][memo]=%2410%20Basic%20Plan%20(11%2F08%2F2016%20-%2011%2F09%2F2016)&payload[statement][transactions][0][created_at]=2016-11-08%2016%3A22%3A18%20-0500&payload[statement][transactions][0][starting_balance_in_cents]=0&payload[statement][transactions][0][ending_balance_in_cents]=1000&payload[statement][transactions][0][gateway_used]=&payload[statement][transactions][0][gateway_transaction_id]=&payload[statement][transactions][0][gateway_order_id]=&payload[statement][transactions][0][payment_id]=161537369&payload[statement][transactions][0][product_id]=3792003&payload[statement][transactions][0][tax_id]=&payload[statement][transactions][0][component_id]=&payload[statement][transactions][0][statement_id]=80168049&payload[statement][transactions][0][customer_id]=14585695&payload[statement][transactions][0][original_amount_in_cents]=&payload[statement][transactions][0][discount_amount_in_cents]=&payload[statement][transactions][0][taxable_amount_in_cents]=&payload[statement][transactions][1][id]=161537344&payload[statement][transactions][1][subscription_id]=15100141&payload[statement][transactions][1][type]=Charge&payload[statement][transactions][1][kind]=quantity_based_component&payload[statement][transactions][1][transaction_type]=charge&payload[statement][transactions][1][success]=true&payload[statement][transactions][1][amount_in_cents]=5400&payload[statement][transactions][1][memo]=Timesheet%20Users%3A%2018%20Timesheet%20Users&payload[statement][transactions][1][created_at]=2016-11-08%2016%3A22%3A18%20-0500&payload[statement][transactions][1][starting_balance_in_cents]=1000&payload[statement][transactions][1][ending_balance_in_cents]=6400&payload[statement][transactions][1][gateway_used]=&payload[statement][transactions][1][gateway_transaction_id]=&payload[statement][transactions][1][gateway_order_id]=&payload[statement][transactions][1][payment_id]=161537369&payload[statement][transactions][1][product_id]=3792003&payload[statement][transactions][1][tax_id]=&payload[statement][transactions][1][component_id]=277221&payload[statement][transactions][1][statement_id]=80168049&payload[statement][transactions][1][customer_id]=14585695&payload[statement][transactions][1][original_amount_in_cents]=&payload[statement][transactions][1][discount_amount_in_cents]=&payload[statement][transactions][1][taxable_amount_in_cents]=&payload[statement][transactions][2][id]=161537369&payload[statement][transactions][2][subscription_id]=15100141&payload[statement][transactions][2][type]=Payment&payload[statement][transactions][2][kind]=&payload[statement][transactions][2][transaction_type]=payment&payload[statement][transactions][2][success]=true&payload[statement][transactions][2][amount_in_cents]=6400&payload[statement][transactions][2][memo]=Pookie%20Test%20-%20%2410%20Basic%20Plan%3A%20Renewal%20payment&payload[statement][transactions][2][created_at]=2016-11-08%2016%3A22%3A20%20-0500&payload[statement][transactions][2][starting_balance_in_cents]=6400&payload[statement][transactions][2][ending_balance_in_cents]=0&payload[statement][transactions][2][gateway_used]=bogus&payload[statement][transactions][2][gateway_transaction_id]=53433&payload[statement][transactions][2][gateway_order_id]=&payload[statement][transactions][2][payment_id]=&payload[statement][transactions][2][product_id]=3792003&payload[statement][transactions][2][tax_id]=&payload[statement][transactions][2][component_id]=&payload[statement][transactions][2][statement_id]=80168049&payload[statement][transactions][2][customer_id]=14585695&payload[statement][transactions][2][card_number]=XXXX-XXXX-XXXX-1&payload[statement][transactions][2][card_expiration]=10%2F2020&payload[statement][transactions][2][card_type]=bogus&payload[statement][transactions][2][refunded_amount_in_cents]=0&payload[product][id]=3792003&payload[product][name]=%2410%20Basic%20Plan&payload[product_family][id]=527890&payload[product_family][name]=Acme%20Projects&payload[payment_profile][id]=10102821&payload[payment_profile][first_name]=Pookie&payload[payment_profile][last_name]=Test&payload[payment_profile][billing_address]=&payload[payment_profile][billing_address_2]=&payload[payment_profile][billing_city]=&payload[payment_profile][billing_country]=&payload[payment_profile][billing_state]=&payload[payment_profile][billing_zip]=&payload[event_id]=347299384",
      "signature": "7c606ec4628ce75ec46e284097ce163a",
      "signature_hmac_sha_256": "40f25e83dd324508bb2149e3e525821922fb210535ebfbfa81e7ab951996b41d"
    }
  },
  {
    "webhook": {
      "event": "payment_success",
      "id": 141765008,
      "created_at": "2016-11-08T16:22:25-05:00",
      "last_error": "404 Resource Not Found (retry 5 of 5)",
      "last_error_at": "2016-11-08T16:43:54-05:00",
      "accepted_at": null,
      "last_sent_at": "2016-11-08T16:43:54-05:00",
      "last_sent_url": "http://requestb.in/11u45x71",
      "successful": false,
      "body": "id=141765008&event=payment_success&payload[site][id]=31615&payload[site][subdomain]=general-goods&payload[subscription][id]=15100141&payload[subscription][state]=active&payload[subscription][trial_started_at]=&payload[subscription][trial_ended_at]=&payload[subscription][activated_at]=2016-11-04%2017%3A06%3A43%20-0400&payload[subscription][created_at]=2016-11-04%2017%3A06%3A42%20-0400&payload[subscription][updated_at]=2016-11-08%2016%3A22%3A22%20-0500&payload[subscription][expires_at]=&payload[subscription][balance_in_cents]=0&payload[subscription][current_period_ends_at]=2016-11-09%2016%3A06%3A42%20-0500&payload[subscription][next_assessment_at]=2016-11-09%2016%3A06%3A42%20-0500&payload[subscription][canceled_at]=&payload[subscription][cancellation_message]=&payload[subscription][next_product_id]=&payload[subscription][cancel_at_end_of_period]=false&payload[subscription][payment_collection_method]=automatic&payload[subscription][snap_day]=&payload[subscription][cancellation_method]=&payload[subscription][current_period_started_at]=2016-11-08%2016%3A06%3A42%20-0500&payload[subscription][previous_state]=active&payload[subscription][signup_payment_id]=161034048&payload[subscription][signup_revenue]=64.00&payload[subscription][delayed_cancel_at]=&payload[subscription][coupon_code]=&payload[subscription][total_revenue_in_cents]=32000&payload[subscription][product_price_in_cents]=1000&payload[subscription][product_version_number]=7&payload[subscription][payment_type]=credit_card&payload[subscription][referral_code]=pggn84&payload[subscription][coupon_use_count]=&payload[subscription][coupon_uses_allowed]=&payload[subscription][customer][id]=14585695&payload[subscription][customer][first_name]=Test&payload[subscription][customer][last_name]=Test&payload[subscription][customer][organization]=&payload[subscription][customer][email]=pookie999%40example.com&payload[subscription][customer][created_at]=2016-11-04%2017%3A06%3A42%20-0400&payload[subscription][customer][updated_at]=2016-11-04%2017%3A06%3A45%20-0400&payload[subscription][customer][reference]=&payload[subscription][customer][address]=&payload[subscription][customer][address_2]=&payload[subscription][customer][city]=&payload[subscription][customer][state]=&payload[subscription][customer][zip]=&payload[subscription][customer][country]=&payload[subscription][customer][phone]=&payload[subscription][customer][portal_invite_last_sent_at]=2016-11-04%2017%3A06%3A45%20-0400&payload[subscription][customer][portal_invite_last_accepted_at]=&payload[subscription][customer][verified]=false&payload[subscription][customer][portal_customer_created_at]=2016-11-04%2017%3A06%3A45%20-0400&payload[subscription][customer][cc_emails]=&payload[subscription][product][id]=3792003&payload[subscription][product][name]=%2410%20Basic%20Plan&payload[subscription][product][handle]=basic&payload[subscription][product][description]=lorem%20ipsum&payload[subscription][product][accounting_code]=basic&payload[subscription][product][request_credit_card]=false&payload[subscription][product][expiration_interval]=&payload[subscription][product][expiration_interval_unit]=never&payload[subscription][product][created_at]=2016-03-24%2013%3A38%3A39%20-0400&payload[subscription][product][updated_at]=2016-11-03%2013%3A03%3A05%20-0400&payload[subscription][product][price_in_cents]=1000&payload[subscription][product][interval]=1&payload[subscription][product][interval_unit]=day&payload[subscription][product][initial_charge_in_cents]=&payload[subscription][product][trial_price_in_cents]=&payload[subscription][product][trial_interval]=&payload[subscription][product][trial_interval_unit]=month&payload[subscription][product][archived_at]=&payload[subscription][product][require_credit_card]=false&payload[subscription][product][return_params]=&payload[subscription][product][taxable]=false&payload[subscription][product][update_return_url]=&payload[subscription][product][initial_charge_after_trial]=false&payload[subscription][product][version_number]=7&payload[subscription][product][update_return_params]=&payload[subscription][product][product_family][id]=527890&payload[subscription][product][product_family][name]=Acme%20Projects&payload[subscription][product][product_family][description]=&payload[subscription][product][product_family][handle]=billing-plans&payload[subscription][product][product_family][accounting_code]=&payload[subscription][product][public_signup_pages][id]=281054&payload[subscription][product][public_signup_pages][return_url]=http%3A%2F%2Fwww.example.com%3Fsuccessfulsignup&payload[subscription][product][public_signup_pages][return_params]=&payload[subscription][product][public_signup_pages][url]=https%3A%2F%2Fgeneral-goods.chargify.com%2Fsubscribe%2Fkqvmfrbgd89q%2Fbasic&payload[subscription][product][public_signup_pages][id]=281240&payload[subscription][product][public_signup_pages][return_url]=&payload[subscription][product][public_signup_pages][return_params]=&payload[subscription][product][public_signup_pages][url]=https%3A%2F%2Fgeneral-goods.chargify.com%2Fsubscribe%2Fdkffht5dxfd8%2Fbasic&payload[subscription][product][public_signup_pages][id]=282694&payload[subscription][product][public_signup_pages][return_url]=&payload[subscription][product][public_signup_pages][return_params]=&payload[subscription][product][public_signup_pages][url]=https%3A%2F%2Fgeneral-goods.chargify.com%2Fsubscribe%2Fjwffwgdd95s8%2Fbasic&payload[subscription][credit_card][id]=10102821&payload[subscription][credit_card][first_name]=Pookie&payload[subscription][credit_card][last_name]=Test&payload[subscription][credit_card][masked_card_number]=XXXX-XXXX-XXXX-1&payload[subscription][credit_card][card_type]=bogus&payload[subscription][credit_card][expiration_month]=10&payload[subscription][credit_card][expiration_year]=2020&payload[subscription][credit_card][customer_id]=14585695&payload[subscription][credit_card][current_vault]=bogus&payload[subscription][credit_card][vault_token]=1&payload[subscription][credit_card][billing_address]=&payload[subscription][credit_card][billing_city]=&payload[subscription][credit_card][billing_state]=&payload[subscription][credit_card][billing_zip]=&payload[subscription][credit_card][billing_country]=&payload[subscription][credit_card][customer_vault_token]=&payload[subscription][credit_card][billing_address_2]=&payload[subscription][credit_card][payment_type]=credit_card&payload[subscription][credit_card][site_gateway_setting_id]=&payload[subscription][credit_card][gateway_handle]=&payload[transaction][id]=161537369&payload[transaction][subscription_id]=15100141&payload[transaction][type]=Payment&payload[transaction][kind]=&payload[transaction][transaction_type]=payment&payload[transaction][success]=true&payload[transaction][amount_in_cents]=6400&payload[transaction][memo]=Pookie%20Test%20-%20%2410%20Basic%20Plan%3A%20Renewal%20payment&payload[transaction][created_at]=2016-11-08%2016%3A22%3A20%20-0500&payload[transaction][starting_balance_in_cents]=6400&payload[transaction][ending_balance_in_cents]=0&payload[transaction][gateway_used]=bogus&payload[transaction][gateway_transaction_id]=53433&payload[transaction][gateway_response_code]=&payload[transaction][gateway_order_id]=&payload[transaction][payment_id]=&payload[transaction][product_id]=3792003&payload[transaction][tax_id]=&payload[transaction][component_id]=&payload[transaction][statement_id]=80168049&payload[transaction][customer_id]=14585695&payload[transaction][card_number]=XXXX-XXXX-XXXX-1&payload[transaction][card_expiration]=10%2F2020&payload[transaction][card_type]=bogus&payload[transaction][refunded_amount_in_cents]=0&payload[transaction][invoice_id]=&payload[event_id]=347299364",
      "signature": "fbcf2f6be579f9658cff90c4373e0ca2",
      "signature_hmac_sha_256": "db96654f5456c5460062feb944ac8bb1418f9d181ae04a8ed982fe9ffdca8de1"
    }
  }
]
```


# Enable Webhooks

This method allows you to enable webhooks via API for your site

```go
EnableWebhooks(
    ctx context.Context,
    body *models.EnableWebhooksRequest) (
    models.ApiResponse[models.EnableWebhooksResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.EnableWebhooksRequest`](../../doc/models/enable-webhooks-request.md) | Body, Optional | - |

## Response Type

[`models.EnableWebhooksResponse`](../../doc/models/enable-webhooks-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.EnableWebhooksRequest{
    WebhooksEnabled: true,
}

apiResponse, err := webhooksController.EnableWebhooks(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "webhooks_enabled": true
}
```


# Replay Webhooks

Posting to the replay endpoint does not immediately resend the webhooks. They are added to a queue and will be sent as soon as possible, depending on available system resources.

You may submit an array of up to 1000 webhook IDs to replay in the request.

```go
ReplayWebhooks(
    ctx context.Context,
    body *models.ReplayWebhooksRequest) (
    models.ApiResponse[models.ReplayWebhooksResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.ReplayWebhooksRequest`](../../doc/models/replay-webhooks-request.md) | Body, Optional | - |

## Response Type

[`models.ReplayWebhooksResponse`](../../doc/models/replay-webhooks-response.md)

## Example Usage

```go
ctx := context.Background()

body := models.ReplayWebhooksRequest{
    Ids: []int64{int64(123456789), int64(123456788)},
}

apiResponse, err := webhooksController.ReplayWebhooks(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "status": "ok"
}
```


# Create Endpoint

The Chargify API allows you to create an endpoint and assign a list of webhooks subscriptions (events) to it.

You can check available events here.
[Event keys](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405357509645-Webhooks-Reference#example-payloads)

```go
CreateEndpoint(
    ctx context.Context,
    body *models.CreateOrUpdateEndpointRequest) (
    models.ApiResponse[models.EndpointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`*models.CreateOrUpdateEndpointRequest`](../../doc/models/create-or-update-endpoint-request.md) | Body, Optional | Used to Create or Update Endpoint |

## Response Type

[`models.EndpointResponse`](../../doc/models/endpoint-response.md)

## Example Usage

```go
ctx := context.Background()

bodyEndpoint := models.CreateOrUpdateEndpoint{
    Url:                  "https://your.site/webhooks",
    WebhookSubscriptions: []models.WebhookSubscription{models.WebhookSubscription("payment_success"), models.WebhookSubscription("payment_failure")},
}

body := models.CreateOrUpdateEndpointRequest{
    Endpoint: bodyEndpoint,
}

apiResponse, err := webhooksController.CreateEndpoint(ctx, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
{
  "endpoint": {
    "id": 1,
    "url": "https://your.site/webhooks",
    "site_id": 1,
    "status": "enabled",
    "webhook_subscriptions": [
      "payment_success",
      "payment_failure"
    ]
  }
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |


# List Endpoints

This method returns created endpoints for site.

```go
ListEndpoints(
    ctx context.Context) (
    models.ApiResponse[[]models.Endpoint],
    error)
```

## Response Type

[`[]models.Endpoint`](../../doc/models/endpoint.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := webhooksController.ListEndpoints(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Example Response *(as JSON)*

```json
[
  {
    "id": 11,
    "url": "https://foobar.com/webhooks",
    "site_id": 1,
    "status": "enabled",
    "webhook_subscriptions": [
      "payment_success",
      "payment_failure"
    ]
  },
  {
    "id": 12,
    "url": "https:/example.com/webhooks",
    "site_id": 1,
    "status": "enabled",
    "webhook_subscriptions": [
      "payment_success",
      "payment_failure",
      "refund_failure"
    ]
  }
]
```


# Update Endpoint

You can update an Endpoint via the API with a PUT request to the resource endpoint.

You can change the `url` of your endpoint which consumes webhooks or list of `webhook_subscriptions`.
Check available [Event keys](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404448450317-Webhooks#configure-webhook-url).

Always send a complete list of events which you want subscribe/watch.
Sending an PUT request for existing endpoint with empty list of `webhook_subscriptions` will end with unsubscribe from all events.

If you want unsubscribe from specific event, just send a list of `webhook_subscriptions` without the specific event key.

```go
UpdateEndpoint(
    ctx context.Context,
    endpointId int,
    body *models.CreateOrUpdateEndpointRequest) (
    models.ApiResponse[models.EndpointResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `endpointId` | `int` | Template, Required | The Chargify id for the endpoint that should be updated |
| `body` | [`*models.CreateOrUpdateEndpointRequest`](../../doc/models/create-or-update-endpoint-request.md) | Body, Optional | Used to Create or Update Endpoint |

## Response Type

[`models.EndpointResponse`](../../doc/models/endpoint-response.md)

## Example Usage

```go
ctx := context.Background()
endpointId := 42

bodyEndpoint := models.CreateOrUpdateEndpoint{
    Url:                  "https://yout.site/webhooks/1/json.",
    WebhookSubscriptions: []models.WebhookSubscription{models.WebhookSubscription("payment_failure"), models.WebhookSubscription("payment_success"), models.WebhookSubscription("refund_failure")},
}

body := models.CreateOrUpdateEndpointRequest{
    Endpoint: bodyEndpoint,
}

apiResponse, err := webhooksController.UpdateEndpoint(ctx, endpointId, &body)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

## Errors

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 404 | Not Found | `ApiError` |
| 422 | Unprocessable Entity (WebDAV) | [`ErrorListResponseException`](../../doc/models/error-list-response-exception.md) |

