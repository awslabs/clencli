package model

// Init struct to initalize things: projects, etc
type Init struct {
	Types []struct {
		Type    string `yaml:"type"`
		Name    string `yaml:"name"`
		Enabled bool   `yaml:"enabled"`
		Files   []struct {
			File struct {
				Path  string `yaml:"path"`
				Src   string `yaml:"src"`
				Dest  string `yaml:"dest"`
				State string `yaml:"state"`
			} `yaml:"file,omitempty"`
		} `yaml:"files"`
	} `yaml:"types"`
}
