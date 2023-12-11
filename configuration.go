package ab_golang_sdk

import (
	"os"
)

// ConfigurationOptions represents a function type that can be used to apply options to the Configuration struct.
type ConfigurationOptions func(*Configuration)

// Configuration holds configuration settings.
type Configuration struct {
	environment       Environment
	subdomain         string
	domain            string
	httpConfiguration HttpConfiguration
	basicAuthUserName string
	basicAuthPassword string
}

// newConfiguration creates a new Configuration with the provided options.
func newConfiguration(options ...ConfigurationOptions) Configuration {
	config := Configuration{}

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

// WithSubdomain is an option that sets the Subdomain in the Configuration.
func WithSubdomain(subdomain string) ConfigurationOptions {
	return func(c *Configuration) {
		c.subdomain = subdomain
	}
}

// WithDomain is an option that sets the Domain in the Configuration.
func WithDomain(domain string) ConfigurationOptions {
	return func(c *Configuration) {
		c.domain = domain
	}
}

// WithHttpConfiguration is an option that sets the HttpConfiguration in the Configuration.
func WithHttpConfiguration(httpConfiguration HttpConfiguration) ConfigurationOptions {
	return func(c *Configuration) {
		c.httpConfiguration = httpConfiguration
	}
}

// WithBasicAuthUserName is an option that sets the BasicAuthUserName in the Configuration.
func WithBasicAuthUserName(basicAuthUserName string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthUserName = basicAuthUserName
	}
}

// WithBasicAuthPassword is an option that sets the BasicAuthPassword in the Configuration.
func WithBasicAuthPassword(basicAuthPassword string) ConfigurationOptions {
	return func(c *Configuration) {
		c.basicAuthPassword = basicAuthPassword
	}
}

// Environment returns the Environment from the Configuration.
func (c *Configuration) Environment() Environment {
	return c.environment
}

// Subdomain returns the Subdomain from the Configuration.
func (c *Configuration) Subdomain() string {
	return c.subdomain
}

// Domain returns the Domain from the Configuration.
func (c *Configuration) Domain() string {
	return c.domain
}

// HttpConfiguration returns the HttpConfiguration from the Configuration.
func (c *Configuration) HttpConfiguration() HttpConfiguration {
	return c.httpConfiguration
}

// BasicAuthUserName returns the BasicAuthUserName from the Configuration.
func (c *Configuration) BasicAuthUserName() string {
	return c.basicAuthUserName
}

// BasicAuthPassword returns the BasicAuthPassword from the Configuration.
func (c *Configuration) BasicAuthPassword() string {
	return c.basicAuthPassword
}

// CreateConfigurationFromEnvironment creates a new Configuration with default settings.
// It also configures various Configuration options.
func CreateConfigurationFromEnvironment(options ...ConfigurationOptions) Configuration {
	config := DefaultConfiguration()

	environment := os.Getenv("MAXIOADVANCEDBILLING_ENVIRONMENT")
	if environment != "" {
		config.environment = Environment(environment)
	}
	subdomain := os.Getenv("MAXIOADVANCEDBILLING_SUBDOMAIN")
	if subdomain != "" {
		config.subdomain = subdomain
	}
	domain := os.Getenv("MAXIOADVANCEDBILLING_DOMAIN")
	if domain != "" {
		config.domain = domain
	}
	basicAuthUserName := os.Getenv("MAXIOADVANCEDBILLING_BASIC_AUTH_USER_NAME")
	if basicAuthUserName != "" {
		config.basicAuthUserName = basicAuthUserName
	}
	basicAuthPassword := os.Getenv("MAXIOADVANCEDBILLING_BASIC_AUTH_PASSWORD")
	if basicAuthPassword != "" {
		config.basicAuthPassword = basicAuthPassword
	}
	for _, option := range options {
		option(&config)
	}
	return config
}

// Server represents available servers.
type Server string

const (
	ENUMDEFAULT Server = "default"
)

// Environment represents available environments.
type Environment string

const (
	PRODUCTION   Environment = "production"
	ENVIRONMENT2 Environment = "environment2"
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
