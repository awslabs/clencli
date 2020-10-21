package model

// Config does TODO...
type Config struct {
	Profiles []Profile
}

// Profile does TODO...
type Profile struct {
	Name     string `yaml:"name"`
	Enabled  bool   `yaml:"enabled"`
	Unsplash `yaml:"unsplash"`
}
