package config

import (
	"fmt"
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
	MySQL   MySQL   `yaml:"mysql"`
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

// MySQL mysql配置参数
type MySQL struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"dbName"`
	Parameters string `yaml:"parameters"`

	MaxLifetime  int `yaml:"maxLifetime"`
	MaxOpenConns int `yaml:"maxOpenConns"`
	MaxIdleConns int `yaml:"maxIdleConns"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
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
