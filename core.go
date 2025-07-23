// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
	"github.com/apimatic/go-core-runtime/https"
)

// HttpConfiguration holds the configuration options for the client.
type HttpConfiguration = https.HttpConfiguration

// RetryConfiguration holds the configuration options for the client.
type RetryConfiguration = https.RetryConfiguration

// RetryConfigurationOptions is a function type that modifies the RetryConfiguration.
type RetryConfigurationOptions = https.RetryConfigurationOptions

// HttpConfigurationOptions is a function type that modifies the HttpConfiguration.
type HttpConfigurationOptions = https.HttpConfigurationOptions

// NewRetryConfiguration creates a new RetryConfiguration.
var NewRetryConfiguration = https.NewRetryConfiguration

// NewHttpConfiguration creates a new HttpConfiguration.
var NewHttpConfiguration = https.NewHttpConfiguration

// NewAndAuth creates a new AndAuth.
var NewAndAuth = https.NewAndAuth

// NewOrAuth creates a new OrAuth.
var NewOrAuth = https.NewOrAuth

// NewAuth creates a new Auth.
var NewAuth = https.NewAuth

// WithMaxRetryAttempts sets the MaxRetryAttempts.
var WithMaxRetryAttempts = https.WithMaxRetryAttempts

// WithRetryOnTimeout sets the RetryOnTimeout.
var WithRetryOnTimeout = https.WithRetryOnTimeout

// WithRetryInterval sets the RetryInterval.
var WithRetryInterval = https.WithRetryInterval

// WithMaximumRetryWaitTime sets the MaximumRetryWaitTime.
var WithMaximumRetryWaitTime = https.WithMaximumRetryWaitTime

// WithBackoffFactor sets the BackoffFactor.
var WithBackoffFactor = https.WithBackoffFactor

// WithHttpStatusCodesToRetry sets the HttpStatusCodesToRetry.
var WithHttpStatusCodesToRetry = https.WithHttpStatusCodesToRetry

// WithHttpMethodsToRetry sets the HttpMethodsToRetry.
var WithHttpMethodsToRetry = https.WithHttpMethodsToRetry

// WithTimeout sets the Timeout.
var WithTimeout = https.WithTimeout

// WithTransport sets the Transport.
var WithTransport = https.WithTransport

// WithRetryConfiguration sets the RetryConfiguration.
var WithRetryConfiguration = https.WithRetryConfiguration
