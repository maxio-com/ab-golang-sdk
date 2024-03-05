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
