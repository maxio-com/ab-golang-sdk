
# Getting Started with Maxio Advanced Billing

## Introduction

### Introduction

#### API Integration

Maxio Advanced Billing (formerly Chargify API) can be integrated with many environments and programming languages via our REST API. Some of our users have contributed their API wrappers in various programming languages. Check out the [API Code Overview](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDI2-api-code-samples) doc for an introduction to the wrappers and available code samples.

#### Testing Guide

Maxio Advanced Billing has compiled a [testing guide](https://chargify.zendesk.com/hc/en-us/articles/4407904658587) that covers a list of important factors to consider while in the testing phase. Here's a high-level overiew of what's covered in our testing guide:

+ Test credit card basics
+ Test site limits
+ Live mode versus test mode

We strongly suggest reading over the testing guide, as well as the entire set of application-based documentation to aid in your discovery of the product.

#### Engage Support

We always enjoy (and appreciate) hearing about larger integrations ahead of time. If you’re planning on importing a large amount of data into Maxio via our API, we suggest sending a “heads up” to “support@chargify.com” so we can coordinate with you to ensure your import process goes smoothly.

Our API, while considered stable, is continually being improved and polished. Please feel free to contact support if you experience issues when integrating with the Maxio Advanced Billing API.

If you have any additional questions regarding our documentation please don't hesitate in reaching out.

#### Support Access

Access to our Technical Specialist team for API support is currently limited to purchasers of our larger Maxio support plans.

But don’t worry! There are quite a few options to help you get the answers you need:

- [Read our documentation for developers](https://developers.chargify.com/docs/developer-docs/ZG9jOjM0NjA3MQ-overview)
- Explore the endpoints of our API Documentation
- [Watch our videos and tutorials](https://chargify.com/tutorials)
- [Check out the Chargify tag on Stack Overflow](http://stackoverflow.com/questions/tagged/chargify)

### API Overview

The Chargify API allows you to interact with our system programmatically from your own application. Using the API you interact with Resources such as:

- Products
- Subscriptions
- Customers
- etc.

The API attempts to conform to the [RESTful](http://en.wikipedia.org/wiki/Representational_State_Transfer) design principles.
You interact with the resources exposed via the API by accessing resource collection and element URIs using the HTTP verbs (GET, POST, PUT, and DELETE).
Chargify accepts and returns both JSON and XML data via the API.

You’ll likely need access to a web developer or programmer (if you’re not one) to get the most use out of the API.

#### Available Formats: JSON and XML

JSON is the primary and recommended format for use with the Chargify API.  XML is also provided as a backwards compatible option for Merchants who require it.

#### Authentication

Authentication is implemented as HTTP Basic Authentication over TLS >= 1.2 (HTTPS), as described in [API Authentication](https://developers.chargify.com/docs/developer-docs/ZG9jOjE1NTUxNQ-authentication)

#### URL

The URL for API requests includes the subdomain of the Site you are working with:

`https://<subdomain>.chargify.com/<resource URI>`

#### Response Data

Response data is sent as either XML or JSON, depending on the type of data requested (`HTTP Content-Type` header) or the type specified as being accepted (HTTP `Accept` header).

GETs for individual statements & invoices may also be requested as PDF using `application/pdf` or appending `.pdf` to the resource URI.

Response codes are sent via the normal HTTP Response Code, and are documented separately for each resource.

For boolean fields, please note that a value of `null` may be considered as false. However, this is not true across all cases. Please excercise good judgement here, or contact support with any questions.

For example:

+ `null` can define that there's no data available for that attribute

#### Pagination

When an endpoint returns a list of items, it will be paginated.  Usually, 20 items will be returned by default, and you may request up to a maximum of 200 at a time.  Pagination is done with query string parameters, for example: `?page=5&per_page=200`

#### Response Time Zones

API responses from Chargify are sent with the timezone of current Chargify site.

Alternately, webhooks sent from Chargify globally utilize EST as the timezone for all content in the body of the payload.

#### Request Data

POST and PUT request data may be formatted as either XML (`application/xml`) or JSON (`application/json`). For best results, you should set your HTTP `Content-Type` request header accordingly, although you may also specify your format by appending `.xml` or `.json` extensions on to the resource URI.

Note that Chargify does not accept PUT or POST data sent as query params or form encoded data – data must be sent as either XML or JSON. If you fail to set your `Content-Type` to either `application/xml` or `application/json`, your request may fail due to triggering of forgery protection mechanisms.

##### About Decimal Numbers

In order to prevent losing precision, we serialize decimal numbers as strings instead of as JSON numbers.

We recommend parsing these strings into their decimal equivalent using a decimal number library in your programming language (i.e. `BigDecimal` in Ruby) instead of relying on floating point values or arithmetic.

##### About Amount Fields

Fields holding amount values are given as a string representing a decimal whole currency amount.

For example, `"1.23"` in currency `"USD"` would equate to `$1.23`.

Not all fields will be rounded to the smallest currency denomination.  Intermediate results, such as those that derive from line-level tax calculations, may hold precision up to 8 decimal places.  However, the top-level totals we provide (e.g. `total_amount`) will be rounded to the smallest currency denomination.

It is up to API consumers to parse the string into a decimal number representation and do any rounding necessary for your application.

#### Debugging

If you’re having difficulty executing a request via our API, try the simplest thing and attempt your request via the curl command-line tool, as shown in the below example. Add the `--verbose` flag to your request to receive even more debugging information.

Another handy tool is [Beeceptor](https://beeceptor.com/). You can use this to intercept your request to see exactly what is being sent.

If you are unable to connect at all, check that you are using TLS 1.2 or better.

If you see a "Could not resolve host" error, double check that the url is correct, including your subdomain. For example: `mysite.chargify.com`.  This error means your DNS server could not find an IP address for the domain you are trying to connect to.

#### Backwards Compatibility

We consider the following changes to be backwards compatible and may make them without advance notice:

+ Adding new API endpoints, or adding new attributes in the responses of existing endpoints
+ Adding new optional parameters to be sent to existing API endpoints
+ Adding new fields to exported data
+ Changing the type or length of any of the ID attributes
  + For example, most IDs are currently integers, but you should not assume that this will always be the case.

In addition, you should not depend on the order of attributes within the API response as this may change.

Chargify does not provide notifications for additions that are clearly defined as backwards compatible.

#### Examples

The following examples use the curl command-line tool to execute API requests.

##### Subscription Listing

**Request**

curl -u <api_key>:x -H Accept:application/json -H Content-Type:application/json https://acme.chargify.com/subscriptions.json

### API Access Limitations

There are a few scenarios that may end up in causing an API request to be blocked even with correct credentials.
**Please note:** All relevant API requests will be blocked if any of the below conditions are true. These limitations also apply to [Chargify Direct](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDE3-introduction).

Those scenarios are as follows:

- Your Chargify subscription is canceled.
- Your Chargify trial has reached an end.
- The site you're making a request for is in the process of ["clearing site data"](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405428327309)
  - _Note: any API request for another site that is in a good state will NOT be blocked_
- The site you're making a request for has been deleted.
  - _Note: any API request for another site that is in a good state will NOT be blocked_

Read more about your Chargify subscription [here](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405430043149-Advanced-Billing-Subscription#advanced-billing-subscription-0-0)

#### What happens when an API request is blocked

The request will fail with a `422` http status code. The response will also include a message explaining the reason for the request being blocked. For example:

- If your Chargify subscription is canceled:

```json
{
  "errors" => [
    [0] "Your Chargify account has been canceled. Please contact support@chargify.com to reactivate."
  ]
}
```

- If your Chargify trial has reached and end and you attempted to make an API request, the response body will look like:

```json
{
  "errors" => [
    [0] "Your trial has ended, please contact sales."
  ]
}
```

- If the site you're making a request for is in the process of ["clearing site data"](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405428327309):

```json
{
  "errors" => [
    [0] "Site data clearing is in progress. Please try later."
  ]
}
```

- If the site you're making a request for has been deleted:

```json
{
  "errors" => [
    [0] "This site has been deleted."
  ]
}
```

### Secure Applications

Please note that it is NOT possible to make API requests directly from the customer's browser or device.  Doing so would expose your API key on the client side, and anyone who has that key has full access to all of your Chargify data.

Instead you will need to take care to tokenize sensitive information by using [Chargify.js](https://developers.chargify.com/docs/developer-docs/ZG9jOjE0NjAzNDI0-overview) or a similar JavaScript library provided by your gateway, and then post the token and other information to your own server, from which you can make the API call to Chargify.

#### Troubleshooting

If you attempt to make a Chargify API request directly from the customer's browser, you may see an error such as:

```
Response to preflight request doesn't pass access control check: No 'Access-Control-Allow-Origin' header is present on the requested resource.
```

or

```
Origin 'https://example.com' is therefore not allowed access.` `The response had HTTP status code 404.
```

This is an error message indicating that Cross-Origin Resource Sharing (CORS) is not enabled on the Chargify server.

### Relationship Invoicing

#### API Compatibility for Relationship Invoicing

This section describes the API for the new, [Relationship Invoicing](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405078794253) style of invoices introduced in January 2018.

If you are an existing customer from prior to January 2018 or have not otherwise explicitly opted into this new style of invoices, you are probably looking for the legacy "Invoices" section that describes [invoice-billing legacy-style invoices](./b3A6MTQxMDgzNjQ-read-invoice).

These new invoices provide a single representation of all of your Chargify billing, whether you collect automatically or via remittance.

#### About Decimal Numbers

In order to prevent losing precision, we serialize decimal numbers as strings instead of as JSON numbers.

We recommend parsing these strings into their decimal equivalent using a decimal number library in your programming language (i.e. `BigDecimal` in Ruby) instead of relying on floating point values or arithmetic.

#### About Amount Fields

Fields holding amount values are given as a string representing a decimal whole currency amount.

For example, `"1.23"` in currency `"USD"` would equate to `$1.23`.

Not all fields will be rounded to the smallest currency denomination. Intermediate results, such as those that derive from line-level tax calculations, may hold precision up to 8 decimal places.  However, the top-level totals we provide (e.g. `total_amount`) will be rounded to the smallest currency denomination.

It is up to API consumers to parse the string into a decimal number representation and do any rounding necessary for your application.

##### Relationship Invoicing Summary

+ If your site **is** using relationship invoicing, you may only use the methods described in this section for working with invoices.

+ If your site is **not** using relationship invoicing, please use the legacy invoice methods:
  
  + [Invoices](./b3A6MTQxMTA0MTA-read-invoice)
  + [Invoices: Payments](./b3A6MTQxMTA0MTI-create-invoice-payment)
  + [Invoices: Charges](./b3A6MTQxMTA0MTM-create-charge)
  + [Invoices: Credits](./b3A6MTQxMTA0MTQ-create-invoice-credit)

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the maxioadvancedbilling library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace maxioadvancedbilling => ".\\Maxio Advanced Billing" // local path to the SDK

require maxioadvancedbilling v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Test the SDK

`Go Testing` is used as the testing framework. To run the unit tests of the SDK, navigate to the `tests` directory of the SDK and run the following command in the terminal:

```bash
$ go test
```

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `subdomain` | `string` | The subdomain for your Chargify site.<br>*Default*: `"subdomain"` |
| `domain` | `string` | The Chargify server domain.<br>*Default*: `"chargify.com"` |
| `environment` | Environment | The API environment. <br> **Default: `Environment.PRODUCTION`** |
| `httpConfiguration` | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := maxioadvancedbilling.CreateConfiguration(
    maxioadvancedbilling.WithHttpConfiguration(
        maxioadvancedbilling.CreateHttpConfiguration(
            maxioadvancedbilling.WithTimeout(0),
            maxioadvancedbilling.WithTransport(http.DefaultTransport),
            maxioadvancedbilling.WithRetryConfiguration(
                maxioadvancedbilling.CreateRetryConfiguration(
                    maxioadvancedbilling.WithMaxRetryAttempts(0),
                    maxioadvancedbilling.WithRetryOnTimeout(true),
                    maxioadvancedbilling.WithRetryInterval(1),
                    maxioadvancedbilling.WithMaximumRetryWaitTime(0),
                    maxioadvancedbilling.WithBackoffFactor(2),
                    maxioadvancedbilling.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    maxioadvancedbilling.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
                ),
            ),
        ),
    ),
    maxioadvancedbilling.WithEnvironment(maxioadvancedbilling.PRODUCTION),
    maxioadvancedbilling.WithBasicAuthUserName("BasicAuthUserName"),
    maxioadvancedbilling.WithBasicAuthPassword("BasicAuthPassword"),
)
client := maxioadvancedbilling.NewClient(config)
```

## Environments

The SDK can be configured to use a different environment for making API calls. Available environments are:

### Fields

| Name | Description |
|  --- | --- |
| production | **Default** Production server |
| environment2 | Production server |

## Authorization

This API uses `Basic Authentication`.

## List of APIs

* [API Exports](doc/controllers/api-exports.md)
* [Advance Invoice](doc/controllers/advance-invoice.md)
* [Billing Portal](doc/controllers/billing-portal.md)
* [Custom Fields](doc/controllers/custom-fields.md)
* [Events-Based Billing Segments](doc/controllers/events-based-billing-segments.md)
* [Payment Profiles](doc/controllers/payment-profiles.md)
* [Product Families](doc/controllers/product-families.md)
* [Product Price Points](doc/controllers/product-price-points.md)
* [Proforma Invoices](doc/controllers/proforma-invoices.md)
* [Reason Codes](doc/controllers/reason-codes.md)
* [Referral Codes](doc/controllers/referral-codes.md)
* [Sales Commissions](doc/controllers/sales-commissions.md)
* [Subscription Components](doc/controllers/subscription-components.md)
* [Subscription Groups](doc/controllers/subscription-groups.md)
* [Subscription Group Invoice Account](doc/controllers/subscription-group-invoice-account.md)
* [Subscription Group Status](doc/controllers/subscription-group-status.md)
* [Subscription Invoice Account](doc/controllers/subscription-invoice-account.md)
* [Subscription Notes](doc/controllers/subscription-notes.md)
* [Subscription Products](doc/controllers/subscription-products.md)
* [Subscription Status](doc/controllers/subscription-status.md)
* [Coupons](doc/controllers/coupons.md)
* [Components](doc/controllers/components.md)
* [Customers](doc/controllers/customers.md)
* [Events](doc/controllers/events.md)
* [Insights](doc/controllers/insights.md)
* [Invoices](doc/controllers/invoices.md)
* [Offers](doc/controllers/offers.md)
* [Products](doc/controllers/products.md)
* [Sites](doc/controllers/sites.md)
* [Subscriptions](doc/controllers/subscriptions.md)
* [Webhooks](doc/controllers/webhooks.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

