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
	RunMode string  `yaml:"runMode"`
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
	CertFile        string `yaml:"certFile"`
	KeyFile         string `yaml:"keyFile"`
	ShutdownTimeout int    `yaml:"shutdownTimeout"`
}

type Monitor struct {
	Enable    bool   `yaml:"enable"`
	Addr      string `yaml:"addr"`
	ConfigDir string `yaml:"configDir"`
}

type CORS struct {
	Enable           bool     `yaml:"enable"`
	AllowOrigins     []string `yaml:"allowOrigins"`
	AllowMethods     []string `yaml:"allowMethods"`
	AllowHeaders     []string `yaml:"allowHeaders"`
	AllowCredentials bool     `yaml:"allowCredentials"`
	MaxAge           int      `yaml:"maxAge"`
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
