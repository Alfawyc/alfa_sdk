package alfa_sdk

import (
	"os"
	"time"
)

type Option func(*Options) error

type Options struct {
	Debug            bool
	Credentials      *Credentials
	UserSandBox      bool
	Timeout          time.Duration
	AutoRequestRetry bool
	AutoRetryCount   int
	RetryWaitTime    time.Duration
	RetryMaxWaitTime time.Duration
}

type Credentials struct {
	ClientId     string
	ClientSecret string
}

func WithAutoRequestRetry(v bool) Option {
	return func(opts *Options) error {
		opts.AutoRequestRetry = v
		return nil
	}
}

func WithAutoRetryCount(v int) Option {
	return func(opts *Options) error {
		opts.AutoRetryCount = v
		return nil
	}
}

func WithDebug(v bool) Option {
	return func(opts *Options) error {
		opts.Debug = v
		return nil
	}
}

func WithRetryWaitTime(v time.Duration) Option {
	return func(opts *Options) error {
		opts.RetryWaitTime = v
		return nil
	}
}

func WithRetryMaxWaitTime(v time.Duration) Option {
	return func(opts *Options) error {
		opts.RetryMaxWaitTime = v
		return nil
	}
}

func WithTimeout(v time.Duration) Option {
	return func(opts *Options) error {
		opts.Timeout = v
		return nil
	}
}

func WithUseSandbox(v bool) Option {
	return func(opts *Options) error {
		opts.UserSandBox = v
		return nil
	}
}

func WithCredentials(v *Credentials) Option {
	return func(opts *Options) error {
		opts.Credentials = v
		return nil
	}
}

func WithCredentialsFromEnv() Option {
	return func(opts *Options) error {
		cred, err := LoadCredentialsFromEvn()
		if err != nil {
			return err
		}
		opts.Credentials = cred
		return nil
	}
}

func LoadCredentialsFromEvn() (*Credentials, error) {
	keys := []string{
		"AMAZON_ADVERTISING_API_CLIENT_ID",
		"AMAZON_ADVERTISING_API_SECRET",
	}
	config := &Credentials{}
	for _, v := range keys {
		if val, ok := os.LookupEnv(v); ok && val != "" {
			switch v {
			case "AMAZON_ADVERTISING_API_CLIENT_ID":
				config.ClientId = val
			case "AMAZON_ADVERTISING_API_SECRET":
				config.ClientSecret = val
			}
		} else {
			//return nil, sdkerr.NewCustomErrorf(sdkerr.LoadEnvCredentialsMissing, "Could not find the following credentials %s", v)
		}
	}
	config.ClientId = "amzn1.application-oa2-client.632a4d1db3f44fb0bb445f5cb86faa61"
	config.ClientSecret = "c93e275bc3c8686e584b6532d9a529fefa3743221bda3229e71c7de9ca0bc2e1"

	return config, nil
}

type Seller struct {
	Region       string
	RefreshToken string
	ScopeId      string
}
