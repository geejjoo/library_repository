package config

import (
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

const (
	AppName                      = "APP_NAME"
	serverPort                   = "SERVER_PORT"
	envShutdownTimeout           = "SHUTDOWN_TIMEOUT"
	parseShutdownTimeoutError    = "config: parse server shutdown timeout error"
	parseRpcShutdownTimeoutError = "config: parse rpc server shutdown timeout error"
)

type AppConf struct {
	AppName     string   `yaml:"app_name"`
	Environment string   `yaml:"environment"`
	Domain      string   `yaml:"domain"`
	APIUrl      string   `yaml:"api_url"`
	Server      Server   `yaml:"server"`
	Cors        Cors     `yaml:"cors"`
	Token       Token    `yaml:"token"`
	Provider    Provider `yaml:"provider"`
	Logger      Logger   `yaml:"logger"`
	DB          DB       `yaml:"db"`
	Cache       Cache    `yaml:"cache"`
}

type DB struct {
	Net      string `yaml:"net"`
	Driver   string `yaml:"driver"`
	Name     string `yaml:"name"`
	User     string `json:"-" yaml:"user"`
	Password string `json:"-" yaml:"password"`
	Host     string `yaml:"host"`
	MaxConn  int    `yaml:"max_conn"`
	Port     string `yaml:"port"`
	Timeout  int    `yaml:"timeout"`
}

type Cache struct {
	Address  string `yaml:"address"`
	Password string `json:"-" yaml:"password"`
	Port     string `yaml:"port"`
}

type Logger struct {
	Level string `yaml:"level"`
}

type Email struct {
	VerifyLinkTTL time.Duration `yaml:"verify_link_ttl"`
	From          string        `yaml:"from"`
	Port          string        `yaml:"port"`
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
	Host        string `json:"-" yaml:"host"`
	Login       string `json:"-" yaml:"login"`
	Password    string `json:"-" yaml:"password"`
	AccessToken string `json:"-" yaml:"access_token"`
	Secret      string `json:"-" yaml:"secret"`
	Key         string `json:"-" yaml:"key"`
	FilePath    string `json:"-" yaml:"file_path"`
}

type Token struct {
	AccessTTL     time.Duration `yaml:"access_ttl"`
	RefreshTTL    time.Duration `yaml:"refresh_ttl"`
	AccessSecret  string        `yaml:"access_secret"`
	RefreshSecret string        `yaml:"refresh_secret"`
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
	port := os.Getenv(serverPort)

	return AppConf{
		AppName: os.Getenv(AppName),
		Server: Server{
			Port: port,
		},
		DB: DB{
			Net:      os.Getenv("DB_NET"),
			Driver:   os.Getenv("DB_DRIVER"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},

		Cache: Cache{
			Address:  os.Getenv("CACHE_ADDRESS"),
			Password: os.Getenv("CACHE_PASSWORD"),
			Port:     os.Getenv("CACHE_PORT"),
		},
		Cors: *newCors(),
	}
}

func (a *AppConf) Init(logger *zap.Logger) {
	shutDownTimeOut, err := strconv.Atoi(os.Getenv(envShutdownTimeout))
	if err != nil {
		logger.Fatal(parseShutdownTimeoutError)
	}
	shutDownTimeout := time.Duration(shutDownTimeOut) * time.Second
	if err != nil {
		logger.Fatal(parseRpcShutdownTimeoutError)
	}
	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		logger.Fatal("config: parse db timeout err", zap.Error(err))
	}
	dbMaxConn, err := strconv.Atoi(os.Getenv("MAX_CONN"))
	if err != nil {
		logger.Fatal("config: parse db max connection err", zap.Error(err))
	}
	a.DB.Timeout = dbTimeout
	a.DB.MaxConn = dbMaxConn

	//var accessTTL int
	//accessTTL, err = strconv.Atoi(os.Getenv(envAccessTTL))
	//if err != nil {
	//	logger.Fatal(parseTokenTTlError)
	//}
	//a.Token.AccessTTL = time.Duration(accessTTL) * time.Minute
	//var refreshTTL int
	//refreshTTL, err = strconv.Atoi(os.Getenv(envRefreshTTL))
	//if err != nil {
	//	logger.Fatal(parseTokenTTlError)
	//}
	//var verifyLinkTTL int
	//verifyLinkTTL, err = strconv.Atoi(os.Getenv(envVerifyLinkTTL))
	//if err != nil {
	//	logger.Fatal(parseTokenTTlError)
	//}

	//a.Provider.Email.From = os.Getenv("EMAIL_FROM")
	//a.Provider.Email.Port = os.Getenv("EMAIL_PORT")
	//a.Provider.Email.Credentials.Host = os.Getenv("EMAIL_HOST")
	//a.Provider.Email.Credentials.Login = os.Getenv("EMAIL_LOGIN")
	//a.Provider.Email.Credentials.Password = os.Getenv("EMAIL_PASSWORD")
	//a.Token.AccessSecret = os.Getenv("ACCESS_SECRET")
	//a.Token.RefreshSecret = os.Getenv("REFRESH_SECRET")
	//a.Domain = os.Getenv("DOMAIN")
	//a.APIUrl = os.Getenv("API_URL")

	//a.Provider.Email.VerifyLinkTTL = time.Duration(verifyLinkTTL) * time.Minute

	//a.Token.RefreshTTL = time.Duration(refreshTTL) * time.Hour * 24

	a.Server.ShutdownTimeout = shutDownTimeout
}

type Server struct {
	Port            string        `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}
