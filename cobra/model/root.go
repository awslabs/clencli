package model

// GlobalConfig struct for the glogal config (~/.clencli.yaml)
type GlobalConfig struct {
	Credentials []struct {
		Credential
	}
	Configs []struct {
		Config
	}
}

// Credential does ...
type Credential struct {
	Profile   string `yaml:"profile"`
	Provider  string `yaml:"provider"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

// ReadMe struct of the readme.yaml
type ReadMe struct {
	Logo struct {
		Provider string `yaml:"provider"`
		Label    string `yaml:"label"`
		URL      string `yaml:"url"`
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
