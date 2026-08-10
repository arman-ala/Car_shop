package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Cors     CorsConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Validator Validator
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
}

type CorsConfig struct {
	AllowOrigins []string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 string
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

type Validator struct {
	Password Password
}

type Password struct {
	Min int
	Max int
	IncludeChars bool
	IncludeDigits bool
	IncludeUpper bool
	IncludeLower bool
	IncludeSpace bool
}

func SetConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	cfgFile, err := getConfigFilePath(env)
	if err != nil {
		panic(err)
	}
	v, err := LoadConfig(cfgFile)
	if err != nil {
		panic(err)
	}
	cfg, err := ParseConfig(v)
	if err != nil {
		panic(err)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(configFile string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func getConfigFileName(env string) string {
	switch env {
	case "docker":
		return "config-docker.yml"
	default:
		return "config-development.yml"
	}
}

func getConfigFilePath(env string) (string, error) {
	fileName := getConfigFileName(env)

	var candidates []string

	if wd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(wd, "config", fileName),
			filepath.Join(wd, "src", "config", fileName),
			filepath.Join(wd, "..", "config", fileName),
			filepath.Join(wd, "..", "src", "config", fileName),
		)
	}

	if _, thisFile, _, ok := runtime.Caller(0); ok {
		candidates = append(candidates, filepath.Join(filepath.Dir(thisFile), fileName))
	}

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return filepath.Abs(p)
		}
	}

	return "", fmt.Errorf("config file %q not found; searched: %v", fileName, candidates)
}
