package config

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var mu = &sync.Mutex{}

type ConfigProvider interface {
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetFloat(key string) float64
	GetBoolean(key string) bool
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetStringArray(key string) []string
	GetIntArray(key string) []int
	Get(key string) interface{}
	GetStringMap(key string) map[string]interface{}
	GetCustomConfigMap(config interface{}) interface{}
	GetStringMapString(key string) map[string]string
	GetStruct(key string, str any) error
}

type Option struct {
	Path       string
	ConfigType string
	ConfigName string
}

type config struct {
	configReader *viper.Viper
	path         string
	configType   string
	configName   string
}

func New(options ...func(*Option)) ConfigProvider {

	opts := &Option{}
	for _, o := range options {
		o(opts)
	}

	configReader := viper.New()
	configReader.AddConfigPath(opts.Path)
	configReader.SetConfigName(opts.ConfigName)
	configReader.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if opts.ConfigType != "" {
		configReader.SetConfigType(opts.ConfigType)
	} else {
		configReader.SetConfigType("yaml")
	}

	if err := configReader.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("config read error: %s", err.Error()))
	}

	configReader.WatchConfig()
	configReader.OnConfigChange(func(in fsnotify.Event) {
		mu.Lock()
		fmt.Println("config watch: updated")
		mu.Unlock()
	})

	return &config{
		configReader: configReader,
		path:         opts.Path,
		configType:   opts.ConfigType,
		configName:   opts.ConfigName,
	}
}

func WithPath(path string) func(*Option) {
	return func(o *Option) {
		o.Path = path
	}
}

func WithConfigType(configType string) func(*Option) {
	return func(o *Option) {
		o.ConfigType = configType
	}
}

func WithConfigName(cfgName string) func(*Option) {
	return func(o *Option) {
		o.ConfigName = cfgName
	}
}

func (c *config) GetInt(key string) int {
	return c.configReader.GetInt(key)
}

func (c *config) GetString(key string) string {
	return c.configReader.GetString(key)
}

func (c *config) GetInt64(key string) int64 {
	return c.configReader.GetInt64(key)
}

func (c *config) GetStringArray(key string) []string {
	return c.configReader.GetStringSlice(key)
}

func (c *config) GetIntArray(key string) []int {
	return c.configReader.GetIntSlice(key)
}

func (c *config) GetBoolean(key string) bool {
	return c.configReader.GetBool(key)
}

func (c *config) GetFloat(key string) float64 {
	return c.configReader.GetFloat64(key)
}

func (c *config) GetTime(key string) time.Time {
	return c.configReader.GetTime(key)
}

func (c *config) GetDuration(key string) time.Duration {
	return c.configReader.GetDuration(key)
}

func (c *config) Get(key string) interface{} {
	return c.configReader.Get(key)
}

func (c *config) GetStringMap(key string) map[string]interface{} {
	return c.configReader.GetStringMap(key)
}

func (c *config) GetCustomConfigMap(config interface{}) interface{} {
	c.configReader.Unmarshal(config)
	return config
}

func (c *config) GetStringMapString(key string) map[string]string {
	return c.configReader.GetStringMapString(key)
}

func (c *config) GetStruct(key string, vStruct any) error {
	val := c.configReader.Get(key)

	jsonData, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonData, &vStruct)
}

type Config struct {
	App    AppConfig    `mapstructure:"app"`
	Server ServerConfig `mapstructure:"server"`
	DB     DBConfig     `mapstructure:"db"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	Logger LoggerConfig `mapstructure:"logger"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	Timezone string `mapstructure:"timezone"`
}

type JWTConfig struct {
	SecretKey          string `mapstructure:"secret_key"`
	AccessTokenExpiry  string `mapstructure:"access_token_expiry"`
	RefreshTokenExpiry string `mapstructure:"refresh_token_expiry"`
}

type LoggerConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}
