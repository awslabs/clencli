package model

// Credentials does TODO...
type Credentials struct {
	Profiles []CredentialProfile `yaml:"profiles"`
}

// CredentialProfile does TODO...
type CredentialProfile struct {
	Name       string `yaml:"name"`
	Enabled    bool   `yaml:"enabled"`
	Credential `yaml:"credential"`
}

// Credential does ...
type Credential struct {
	Provider  string `yaml:"provider"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretkey"`
}

// Configurations does TODO...
type Configurations struct {
	Profiles []ConfigurationProfile `yaml:"profiles"`
}

// ConfigurationProfile does TODO...
type ConfigurationProfile struct {
	Name     string `yaml:"name"`
	Enabled  bool   `yaml:"enabled"`
	Unsplash `yaml:"unsplash"`
}
