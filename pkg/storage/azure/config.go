package azure

type Config struct {
	Container              string
	TenantID               string
	ClientID               string
	ClientSecret           string
	Account                string
	DefaultAzureCredential bool
}

func (c *Config) SetDefaults() {}

func (c *Config) Validate() error {
	return nil
}
