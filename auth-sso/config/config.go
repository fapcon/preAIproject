package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	migrateConfig "gitlab.com/golight/migrator/db/config"
)

type Config struct {
	Env           string        `yaml:"env" env:"ENV" env-default:"LOCAL"`
	GRPC          GRPCConfig    `yaml:"grpc"`
	TokenTTL      time.Duration `yaml:"token_ttl" env:"TOKEN_TTL" env-default:"1h"`
	SigningMethod string        `yaml:"signing_method" env:"JWT_SIGNING_METHOD"`
	SecretKey     string        `yaml:"secret_key" env:"JWT_SECRET_KEY"`
	Logger        Logger        `yaml:"logger"`
	DB            DB            `yaml:"db"`
	Server        HTTPServer    `yaml:"gateway"`
	Client        Client        `yaml:"clients"`
	Google        OauthGoogle   `yaml:"google"`
	Facebook      OauthFacebook `yaml:"facebook"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"GRPC_PORT" env-default:"8080"`
	Timeout time.Duration `yaml:"timeout" env:"GRPC_TIMEOUT" env-default:"999h"`
	Host    string        `yaml:"host" env:"GRPC_HOST" env-default:"0.0.0.0"`
}

type HTTPServer struct {
	Host string `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port int    `env:"HTTP_PORT" env-default:"8081"`
}

type Client struct {
	Timeout      time.Duration `env:"CLIENT_TIMEOUT" env-default:"5s"`
	RetriesCount int           `env:"CLIENT_RETRIES_COUNT" env-default:"3"`
}

type Logger struct {
	Level string `env:"LOG_LEVEL" env-default:"0"`
}

type DB struct {
	Net      string        `env:"DB_NET" env-default:"postgres"`
	Driver   string        `env:"DB_DRIVER" env-default:"postgres"`
	Name     string        `env:"DB_NAME" env-default:"auth-sso"`
	User     string        `env:"DB_USER" env-default:""`
	Password string        `env:"DB_PASSWORD" env-default:""`
	Host     string        `env:"DB_HOST" env-default:"localhost"`
	MaxConn  int           `env:"DB_MAX_CONN" env-default:"10"`
	Port     string        `env:"DB_PORT" env-default:"5432"`
	Timeout  time.Duration `env:"DB_TIMEOUT" env-default:"10s"`
}

type OauthGoogle struct {
	ClientID     string `env:"GOOGLE_OAUTH_CLIENT_ID" yaml:"client_id_google"`
	RedirectURI  string `env:"OAUTH_REDIRECT_URI" yaml:"redirect_uri_google"`
	ClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET" yaml:"client_secret_google"`
}

type OauthFacebook struct {
	ClientID     string `env:"FACEBOOK_OAUTH_CLIENT_ID" yaml:"client_id_facebook"`
	RedirectURI  string `env:"OAUTH_REDIRECT_URI" yaml:"redirect_uri_facebook"`
	ClientSecret string `env:"FACEBOOK_OAUTH_CLIENT_SECRET" yaml:"client_secret_facebook"`
}

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" && os.Getenv("CONFIG_PATH") == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func ConvertConfigs(cfg *Config, mCfg *migrateConfig.DB) migrateConfig.DB {
	mCfg.Net = cfg.DB.Net
	mCfg.Driver = cfg.DB.Driver
	mCfg.Name = cfg.DB.Name
	mCfg.User = cfg.DB.User
	mCfg.Password = cfg.DB.Password
	mCfg.Host = cfg.DB.Host
	mCfg.MaxConn = cfg.DB.MaxConn
	mCfg.Port = cfg.DB.Port
	mCfg.Timeout = int(cfg.DB.Timeout)

	return *mCfg
}
