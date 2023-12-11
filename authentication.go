package ab_golang_sdk

import (
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
)

// BasicAuthentication returns an authenticator function that provides authentication for API calls.
func BasicAuthentication(config Configuration) https.Authenticator {
	return func(requiresAuth bool) https.HttpInterceptor {
		if !requiresAuth {
			return https.PassThroughInterceptor
		}
		return func(req *http.Request,
			next https.HttpCallExecutor,
		) https.HttpContext {
			req.SetBasicAuth(config.BasicAuthUserName(), config.BasicAuthPassword())
			return next(req)
		}
	}
}
