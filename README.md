
# Getting Started with Maxio Advanced Billing

## Introduction

Maxio Advanced Billing (formerly Chargify) provides an HTTP-based API that conforms to the principles of REST.
One of the many reasons to use Advanced Billing is the immense feature set and [client libraries](http://localhost:8080/go/development-tools/development-tools/client-libraries).
The Maxio API returns JSON responses as the primary and recommended format, but XML is also provided as a backwards compatible option for merchants who require it.

### Steps to make your first Maxio Advanced Billing API call

1. [Sign-up](https://app.chargify.com/signup/maxio-billing-sandbox) or [log-in](https://app.chargify.com/login.html) to your [test site](https://maxio.zendesk.com/hc/en-us/articles/24250712113165-Testing-Overview) account.
2. [Setup authentication](https://maxio.zendesk.com/hc/en-us/articles/24294819360525-API-Keys) credentials.
3. [Submit an API request and verify the response](http://localhost:8080/go/development-tools/development-tools/client-libraries#make-your-first-maxio-advanced-billing-api-request).
4. Test the Advanced Billing [integrations](https://www.maxio.com/integrations).

Next, you can explore [authentication methods](http://localhost:8080/go/getting-started/introduction/authentication), [basic concepts](http://localhost:8080/go) for interacting with Advanced Billing via the API, and the entire set of [application-based documentation](https://docs.maxio.com/hc/en-us) to aid in your discovery of the product.

#### Request Example

The following example uses the curl command-line tool to make an API request.

**Request**

curl -u <api_key>:x -H Accept:application/json -H Content-Type:application/json https://acme.chargify.com/subscriptions.json

### Requirements

The SDK requires **Go version 1.18 or above**.

## Installation

The following section explains how to use the advancedbilling library in a new project.

### 1. Install the Package

To use the package in your application, you can install the package from [pkg.go.dev](https://pkg.go.dev/) using the following command:

```bash
$ go get github.com/maxio-com/ab-golang-sdk@v0.7.1
```

You can also view the package at: https://pkg.go.dev/github.com/maxio-com/ab-golang-sdk@v0.7.1

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| site | `string` | The subdomain for your Advanced Billing site.<br>*Default*: `"subdomain"` |
| environment | `Environment` | The API environment. <br> **Default: `Environment.US`** |
| httpConfiguration | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries. |
| basicAuthCredentials | [`BasicAuthCredentials`](doc/auth/basic-authentication.md) | The Credentials Setter for Basic Authentication |

The API client can be initialized as follows:

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk"
)

func main() {
    client := advancedbilling.NewClient(
    advancedbilling.CreateConfiguration(
            advancedbilling.WithHttpConfiguration(
                advancedbilling.CreateHttpConfiguration(
                    advancedbilling.WithTimeout(120),
                ),
            ),
            advancedbilling.WithEnvironment(advancedbilling.US),
            advancedbilling.WithBasicAuthCredentials(
                advancedbilling.NewBasicAuthCredentials(
                    "BasicAuthUserName",
                    "BasicAuthPassword",
                ),
            ),
            advancedbilling.WithSite("subdomain"),
        ),
    )
}
```

## Environments

The SDK can be configured to use a different environment for making API calls. Available environments are:

### Fields

| Name | Description |
|  --- | --- |
| US | **Default** Default Advanced Billing environment hosted in US. Valid for the majority of our customers. |
| EU | Advanced Billing environment hosted in EU. Use only when you requested EU hosting for your AB account. |

## Authorization

This API uses the following authentication schemes.

* [`BasicAuth (Basic Authentication)`](doc/auth/basic-authentication.md)

## List of APIs

* [API Exports](doc/controllers/api-exports.md)
* [Advance Invoice](doc/controllers/advance-invoice.md)
* [Billing Portal](doc/controllers/billing-portal.md)
* [Component Price Points](doc/controllers/component-price-points.md)
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

## SDK Infrastructure

### Configuration

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

### Utilities

* [ApiResponse](doc/api-response.md)

