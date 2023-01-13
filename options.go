package increase

import "increase/core"

type Environment string

const (
	EnvironmentProduction Environment = "production"
	EnvironmentSandbox    Environment = "sandbox"
)

func (e Environment) URL() (url string) {
	switch e {
	case EnvironmentProduction:
		url = "https://api.increase.com"
	case EnvironmentSandbox:
		url = "https://sandbox.increase.com"
	}
	return
}

type ClientOptions struct {
	APIKey      string      `env:"INCREASE_API_KEY"`
	BaseURL     string      `env:"INCREASE_BASE_URL"`
	Environment Environment `env:"INCREASE_ENVIRONMENT" default:"production"`
	Requester   *core.CoreClient
}

func NewClientOptionsWithDefaults() *ClientOptions {
	options := ClientOptions{}
	core.LoadEnvironmentVariables(&options)
	return &options
}

func (r *ClientOptions) LoadDefaults() {
	core.LoadEnvironmentVariables(r)
	if len(r.BaseURL) == 0 {
		r.BaseURL = r.Environment.URL()
	}
}

func (r *ClientOptions) Copy() *ClientOptions {
	copy := *r
	return &copy
}

func (r *ClientOptions) WithAPIKey(api_key string) *ClientOptions {
	opts := r.Copy()
	opts.APIKey = api_key
	return opts
}

func (r *ClientOptions) WithBaseURL(base_url string) *ClientOptions {
	opts := r.Copy()
	opts.BaseURL = base_url
	return opts
}

func (r *ClientOptions) WithEnvironment(environment Environment) *ClientOptions {
	opts := r.Copy()
	opts.Environment = environment
	return opts
}

func (r *ClientOptions) WithRequester(requester *core.CoreClient) *ClientOptions {
	opts := r.Copy()
	opts.Requester = requester
	return opts
}
