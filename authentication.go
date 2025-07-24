// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "errors"
    "github.com/apimatic/go-core-runtime/https"
    "net/http"
)

// BasicAuthCredentials represents the credentials required for `BasicAuth` authentication.
type BasicAuthCredentials struct {
    username string
    password string
}

// NewBasicAuthCredentials creates a new instance of BasicAuthCredentials with provided parameters.
func NewBasicAuthCredentials(
    username string,
    password string) BasicAuthCredentials {
    return BasicAuthCredentials {
        username: username,
        password: password,
    }
}

// WithUsername sets username in BasicAuthCredentials.
func (b BasicAuthCredentials) WithUsername(username string) BasicAuthCredentials {
    b.username = username
    return b
}

// WithPassword sets password in BasicAuthCredentials.
func (b BasicAuthCredentials) WithPassword(password string) BasicAuthCredentials {
    b.password = password
    return b
}

// Username returns the username associated with the BasicAuthCredentials.
func (b BasicAuthCredentials) Username() string {
    return b.username
}

// Password returns the password associated with the BasicAuthCredentials.
func (b BasicAuthCredentials) Password() string {
    return b.password
}

// Validate function returns validation error associated with the BasicAuthCredentials.
func (b BasicAuthCredentials) Validate() error {
    if b.username == "" || b.password == "" {
        return errors.New("BasicAuth : missing auth credentials -> Username && Password.")
    }
    return nil
}

// Authenticator function returns HttpInterceptor function that provides authentication for API calls.
func (b BasicAuthCredentials) Authenticator() https.HttpInterceptor {
    return func(req *http.Request,
        next https.HttpCallExecutor,
    ) https.HttpContext {
            req.SetBasicAuth(b.Username(), b.Password())
        return next(req)
    }
}
