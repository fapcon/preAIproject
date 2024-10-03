package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"
)

//go:generate easytags $GOFILE yaml
type AppConf struct {
	AppName     string   `env:"APP_NAME" yaml:"app_name"`
	Environment string   `yaml:"environment"`
	Domain      string   `yaml:"domain"`
	APIUrl      string   `yaml:"api_url"`
	Server      Server   `yaml:"server"`
	Cors        Cors     `yaml:"cors"`
	Token       Token    `yaml:"token"`
	Provider    Provider `yaml:"provider"`
	Logger      Logger   `yaml:"logger"`
	DB          DB       `yaml:"db"`
	NoSQL       NoSqlDB  `yaml:"no_sql"`
	Cache       Cache    `yaml:"cache"`
	GRPC        GRPC     `yaml:"rpc"`
	Google      OauthG   `yaml:"google"`
	Facebook    OauthF   `yaml:"facebook"`
	TgBot       TgBot    `yaml:"tg_bot"`
}

type GRPC struct {
	Port            string        `env:"GRPC_PORT" yaml:"port"`
	Host            string        `env:"GRPC_HOST" yaml:"address"`
	Type            string        `env:"USER_SERVICE_TYPE" yaml:"type"`
	ShutDownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" yaml:"shut_down_timeout"`
}

type OauthG struct {
	ClientID     string `env:"GOOGLE_OAUTH_CLIENT_ID" yaml:"client_id"`
	ClientSecret string `env:"GOOGLE_OAUTH_CLIENT_SECRET" yaml:"client"`
}

type OauthF struct {
	ClientID     string `env:"FACEBOOK_OAUTH_CLIENT_ID" yaml:"client_id"`
	ClientSecret string `env:"FACEBOOK_OAUTH_CLIENT_SECRET" yaml:"client"`
}

type TgBot struct {
	Token         string `env:"TELEGRAM_BOT_TOKEN"`
	TrustedChatID string `env:"TRUSTED_CHAN_ID"`
}

type DB struct {
	Net      string `env:"DB_NET"`
	Driver   string `env:"DB_DRIVER"`
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	MaxConn  int    `env:"MAX_CONN"`
	Port     string `env:"DB_PORT"`
	Timeout  int    `env:"DB_TIMEOUT"`
}

type NoSqlDB struct {
	Net      string `env:"MONGO_NET"`
	Driver   string `env:"MONGO_DRIVER"`
	Name     string `env:"MONGO_NAME"`
	User     string `env:"MONGO_USER"`
	Password string `env:"MONGO_PASSWORD"`
	Host     string `env:"MONGO_HOST"`
	MaxConn  int    `env:"MAX_CONN"`
	Port     string `env:"MONGO_PORT"`
	Timeout  int    `env:"DB_TIMEOUT"`
}

type Cache struct {
	Address  string `env:"CACHE_ADDRESS"`
	Password string `env:"CACHE_PASSWORD"`
	Port     string `env:"CACHE_PORT"`
}

type Logger struct {
	Level           string `env:"LEVEL"`
	StackTraceLevel string `env:"STACK_TRACE_LEVEL"`
}

type Email struct {
	VerifyLinkTTL time.Duration `env:"VERIFY_LINK_TTL"`
	From          string        `env:"EMAIL_FROM"`
	Port          string        `env:"EMAIL_PORT"`
	Credentials   Credentials   `json:"-" yaml:"credentials"`
}

type Provider struct {
	Email Email `yaml:"email"`
	Phone Phone `yaml:"phone"`
}

type Phone struct {
	VerifyCodeTTL time.Duration `yaml:"verify_code_ttl"`
	Credentials   Credentials   `json:"-" yaml:"credentials"`
}

type Credentials struct {
	Host        string `env:"EMAIL_HOST"`
	Login       string `env:"EMAIL_LOGIN"`
	Password    string `env:"EMAIL_PASSWORD"`
	AccessToken string `env:"ACCESS_TOKEN"`
	Secret      string `json:"-" yaml:"secret"`
	Key         string `json:"-" yaml:"key"`
	FilePath    string `json:"-" yaml:"file_path"`
}

type Token struct {
	AccessTTL     time.Duration `env:"ACCESS_TTL"`
	RefreshTTL    time.Duration `env:"REFRESH_TTL"`
	AccessSecret  string        `env:"ACCESS_SECRET"`
	RefreshSecret string        `env:"REFRESH_SECRET"`
}

type Server struct {
	Port            string        `env:"SERVER_PORT"`
	PrometheusPort  string        `env:"PROMETHEUS_PORT"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT"`
}

type Cors struct {
	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	AllowedOrigins   []string `yaml:"allowed_origins"`
	AllowedMethods   []string `yaml:"allowed_methods"`
	AllowedHeaders   []string `yaml:"allowed_headers"`
	ExposedHeaders   []string `yaml:"exposed_headers"`
	AllowCredentials bool     `yaml:"allow_credentials"`
	MaxAge           int      `yaml:"max_age"` // Maximum value not ignored by any of major browsers
}

func newCors() *Cors {
	return &Cors{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}
}

func NewAppConf() AppConf {
	var config AppConf
	err := Init(&config)
	if err != nil {
		fmt.Printf("Error mapping env to struct: %v\n", err)
		return AppConf{}
	}
	config.Cors = *newCors()

	config.Token.RefreshTTL *= time.Hour * 24
	config.Token.AccessTTL *= time.Minute

	return config
}

func Init(config interface{}) error {
	configValue := reflect.ValueOf(config)

	if configValue.Kind() == reflect.Ptr {
		configValue = configValue.Elem()
	}

	if configValue.Kind() != reflect.Struct {
		return fmt.Errorf("Input is not a struct")
	}

	configType := configValue.Type()

	for i := 0; i < configValue.NumField(); i++ {
		fieldValue := configValue.Field(i)
		fieldType := configType.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			err := Init(fieldValue.Addr().Interface())
			if err != nil {
				return err
			}
		} else {
			tagValue := fieldType.Tag.Get("env")
			if tagValue == "" {
				continue
			}
			envValue := os.Getenv(tagValue)
			if envValue == "" {
				continue
			}

			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(envValue)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if intValue, err := strconv.ParseInt(envValue, 10, 64); err == nil {
					fieldValue.SetInt(intValue)
				} else {
					return fmt.Errorf("error parsing int value for field %s: %v", fieldType.Name, err)
				}
			default:
				return fmt.Errorf("unsupported field type for field %s", fieldType.Name)
			}
		}
	}
	return nil
}
