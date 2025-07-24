// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "github.com/apimatic/go-core-runtime/https"
)

// createAuthenticationFromConfig returns a map of auth interfaces that provides authentications for API calls.
func createAuthenticationFromConfig(config Configuration) map[string]https.AuthInterface {
    return map[string]https.AuthInterface {
        "BasicAuth": config.BasicAuthCredentials(),
    }
}
