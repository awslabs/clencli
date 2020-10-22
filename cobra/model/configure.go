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

// Config does TODO...
type Config struct {
	Profiles []ConfigProfile `yaml:"profiles"`
}

// ConfigProfile does TODO...
type ConfigProfile struct {
	Name     string `yaml:"name"`
	Enabled  bool   `yaml:"enabled"`
	Unsplash `yaml:"unsplash"`
}
