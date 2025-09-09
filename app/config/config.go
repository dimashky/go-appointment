package config

type (
	App struct {
		FiberInfo bool `mapstructure:"fiber_info" json:"fiber_info" yaml:"fiber_info"`
	}
	Logging struct {
		Level       string `mapstructure:"level" json:"level" yaml:"level"`
		PrettyPrint bool   `mapstructure:"pretty_print" json:"pretty_print" yaml:"pretty_print"`
	}

	Database struct {
		Host     string `mapstructure:"host" json:"host" yaml:"host"`
		Port     int    `mapstructure:"port" json:"port" yaml:"port"`
		User     string `mapstructure:"user" json:"user" yaml:"user"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		Name     string `mapstructure:"name" json:"name" yaml:"name"`
		Charset  string `mapstructure:"charset" json:"charset" yaml:"charset"`
	}

	Config struct {
		Logging  Logging  `mapstructure:"logging" json:"logging" yaml:"logging"`
		HTTP     HTTP     `mapstructure:"http" json:"http" yaml:"http"`
		App      App      `mapstructure:"app" json:"app" yaml:"app"`
		Database Database `mapstructure:"database" json:"database" yaml:"database"`
	}
)
