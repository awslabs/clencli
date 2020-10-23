package model

// Credentials does TODO
type Credentials struct {
	Profiles []CredentialProfile `yaml:"profiles"`
}

// CredentialProfile does TODO...
type CredentialProfile struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description,omitempty"`
	Enabled     bool         `yaml:"enabled"`
	CreatedAt   string       `yaml:"createdAt"`
	UpdatedAt   string       `yaml:"updatedAt"`
	Credentials []Credential `yaml:"credentials"`
}

// Credential does TODO
type Credential struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Enabled     bool   `yaml:"enabled"`
	CreatedAt   string `yaml:"createdAt"`
	UpdatedAt   string `yaml:"updatedAt"`
	Provider    string `yaml:"provider"`
	AccessKey   string `yaml:"accessKey"`
	SecretKey   string `yaml:"secretkey"`
}

// Configurations does TODO
type Configurations struct {
	Profiles []ConfigurationProfile `yaml:"profiles"`
}

// ConfigurationProfile does TODO...
type ConfigurationProfile struct {
	Name           string          `yaml:"name"`
	Description    string          `yaml:"description,omitempty"`
	Enabled        bool            `yaml:"enabled"`
	CreatedAt      string          `yaml:"createdAt"`
	UpdatedAt      string          `yaml:"updatedAt"`
	Configurations []Configuration `yaml:"credentials"`
}

// Configuration does TODO
type Configuration struct {
	Name        string `yaml:"name,omitempty"`
	Description string `yaml:"description,omitempty"`
	Enabled     bool   `yaml:"enabled"`
	CreatedAt   string `yaml:"createdAt"`
	UpdatedAt   string `yaml:"updatedAt"`
	Unsplash    `yaml:"unsplash,omitempty"`
}
