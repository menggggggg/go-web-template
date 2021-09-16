package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	C    = new(Config)
	once sync.Once
)

type Config struct {
	RunMode string  `yaml:"run_mode"`
	Swagger bool    `yaml:"swagger"`
	Log     Log     `yaml:"log"`
	HTTP    HTTP    `yaml:"http"`
	Monitor Monitor `yaml:"monitor"`
	CORS    CORS    `yaml:"cores"`
	GZIP    GZIP    `yaml:"gzip"`
}

type Log struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type HTTP struct {
	Addr            string `yaml:"addr"`
	CertFile        string `yaml:"cert_file"`
	KeyFile         string `yaml:"key_file"`
	ShutdownTimeout int    `yaml:"shutdown_timeout"`
}

type Monitor struct {
	Enable    bool   `yaml:"enable"`
	Addr      string `yaml:"addr"`
	ConfigDir string `yaml:"config_dir"`
}

type CORS struct {
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allow_origins"`
	AllowMethods     []string `yaml:"allow_methods"`
	AllowHeaders     []string `yaml:"allow_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	MaxAge           int      `yaml:"max_age"`
}

type GZIP struct {
	Enable bool `yaml:"enable"`
}

func LoadConfig() {
	once.Do(func() {
		configPath := os.Getenv("CONF_FILE_PATH")
		if len(configPath) == 0 {
			panic("not found env variable: CONF_FILE_PATH ")
		}
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			panic("read config file error: " + err.Error())
		}

		if err := viper.Unmarshal(&C); err != nil {
			panic("unable to decode into stuct: " + err.Error())
		}
	})
}
