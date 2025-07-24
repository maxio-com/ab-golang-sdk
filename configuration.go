// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package advancedbilling

import (
    "os"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
    environment          Environment
    site                 string
    httpConfiguration    HttpConfiguration
    basicAuthCredentials BasicAuthCredentials
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
    config := Configuration{}
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// cloneWithOptions provides configuration with the provided options.
func (config Configuration) cloneWithOptions(options ...ConfigurationOptions) Configuration {
    for _, option := range options {
        option(&config)
    }
    return config
}

// WithEnvironment is an option that sets the Environment in the Configuration.
func WithEnvironment(environment Environment) ConfigurationOptions {
    return func(c *Configuration) {
        c.environment = environment
    }
}

// WithSite is an option that sets the Site in the Configuration.
func WithSite(site string) ConfigurationOptions {
    return func(c *Configuration) {
        c.site = site
    }
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
    return func(c *Configuration) {
        c.httpConfiguration = httpConfiguration
    }
}

// WithBasicAuthCredentials is an option that sets the BasicAuthCredentials in the Configuration.
func WithBasicAuthCredentials(basicAuthCredentials BasicAuthCredentials) ConfigurationOptions {
    return func(c *Configuration) {
        c.basicAuthCredentials = basicAuthCredentials
    }
}

// Environment returns the environment from the Configuration.
func (c Configuration) Environment() Environment {
    return c.environment
}

// Site returns the site from the Configuration.
func (c Configuration) Site() string {
    return c.site
}

// HttpConfiguration returns the httpConfiguration from the Configuration.
func (c Configuration) HttpConfiguration() HttpConfiguration {
    return c.httpConfiguration
}

// BasicAuthCredentials returns the basicAuthCredentials from the Configuration.
func (c Configuration) BasicAuthCredentials() BasicAuthCredentials {
    return c.basicAuthCredentials
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    environment := os.Getenv("ADVANCEDBILLING_ENVIRONMENT")
    if environment != "" {
        config.environment = Environment(environment)
    }
    site := os.Getenv("ADVANCEDBILLING_SITE")
    if site != "" {
        config.site = site
    }
    basicAuthUserName := os.Getenv("ADVANCEDBILLING_BASIC_AUTH_USER_NAME")
    if basicAuthUserName != "" {
        config.basicAuthCredentials.username = basicAuthUserName
    }
    basicAuthPassword := os.Getenv("ADVANCEDBILLING_BASIC_AUTH_PASSWORD")
    if basicAuthPassword != "" {
        config.basicAuthCredentials.password = basicAuthPassword
    }
    for _, option := range options {
        option(&config)
    }
    return config
}

// Server represents available servers.
type Server string

const (
    PRODUCTION Server = "production"
    EBB        Server = "ebb"
)

// Environment represents available environments.
type Environment string

const (
    US Environment = "US"
    EU Environment = "EU"
)

// CreateRetryConfiguration creates a new RetryConfiguration with the provided options.
func CreateRetryConfiguration(options ...RetryConfigurationOptions) RetryConfiguration {
    config := DefaultRetryConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// CreateHttpConfiguration creates a new HttpConfiguration with the provided options.
func CreateHttpConfiguration(options ...HttpConfigurationOptions) HttpConfiguration {
    config := DefaultHttpConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}

// CreateConfiguration creates a new Configuration with the provided options.
func CreateConfiguration(options ...ConfigurationOptions) Configuration {
    config := DefaultConfiguration()
    
    for _, option := range options {
        option(&config)
    }
    return config
}
