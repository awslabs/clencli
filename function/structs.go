package function

// ReadMe struct for readme.yaml config
type ReadMe struct {
	Logo struct {
		Label string `yaml:"label"`
		Theme string `yaml:"theme"`
		URL   string `yaml:"url"`
	} `yaml:"logo,omitempty"`
	Shields struct {
		Badges []struct {
			Description string `yaml:"description"`
			Image       string `yaml:"image"`
			URL         string `yaml:"url"`
		} `yaml:"badges"`
	} `yaml:"shields,omitempty"`
	App struct {
		Name     string `yaml:"name"`
		Function string `yaml:"function"`
		ID       string `yaml:"id"`
	} `yaml:"app,omitempty"`
	Screenshots []struct {
		Caption string `yaml:"caption"`
		Label   string `yaml:"label"`
		URL     string `yaml:"url"`
	} `yaml:"screenshots,omitempty"`
	Usage         string `yaml:"usage"`
	Prerequisites []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"prerequisites,omitempty"`
	Installing   string   `yaml:"installing,omitempty"`
	Testing      string   `yaml:"testing,omitempty"`
	Deployment   string   `yaml:"deployment,omitempty"`
	Include      []string `yaml:"include,omitempty"`
	Contributors []struct {
		Name  string `yaml:"name"`
		Role  string `yaml:"role"`
		Email string `yaml:"email"`
	} `yaml:"contributors,omitempty"`
	Acknowledgments []struct {
		Name string `yaml:"name"`
		Role string `yaml:"role"`
	} `yaml:"acknowledgments,omitempty"`
	References []struct {
		Description string `yaml:"description"`
		Name        string `yaml:"name"`
		URL         string `yaml:"url"`
	} `yaml:"references,omitempty"`
	License   string `yaml:"license,omitempty"`
	Copyright string `yaml:"copyright,omitempty"`
}

// GlobalConfig struct for the glogal config (~/.clencli.yaml)
type GlobalConfig struct {
	Unsplash struct {
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
	} `yaml:"unsplash,omitempty"`
	Init struct {
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
	} `yaml:"init"`
	ReadMe
}