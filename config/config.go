package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	EnvProduction          = "production"
	EnvStaging             = "staging"
	EnvDevelopment         = "development"
	EnvLocalhost           = "localhost"
	AdminUserId            = int64(-1)
	RedisDefaultExpireTime = time.Second * 60 * 60 * 24 * 30 // 預設一個月
)

var EnvShortName = map[string]string{
	EnvProduction:  "prod",
	EnvStaging:     "stag",
	EnvDevelopment: "dev",
	EnvLocalhost:   "local",
}

var HtmlBasePath = map[string]string{
	EnvLocalhost: "http://localhost:8080/",
}

// Environment
func GetEnvironment() string {
	return os.Getenv("ENVIRONMENT")
}

func GetShortEnv() string {
	return EnvShortName[GetEnvironment()]
}

// Jwt Salt
func GetJwtSalt() string {
	return os.Getenv("JWT_SALT")
}

func GetOauthSalt() string {
	return os.Getenv("OAUTH_SALT")
}

// Base path
var (
	_, b, _, _ = runtime.Caller(0)
	basePath   = filepath.Dir(b)
)

func GetBasePath() string {
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetHtmlBasePath() string {
	return HtmlBasePath[GetEnvironment()]
}

// Cors
func GetCorsRule(origin string) bool {
	switch GetEnvironment() {
	case EnvLocalhost:
		return true
	case EnvDevelopment:
		return origin == "https://sample-development.website.com" || strings.Contains(origin, "http://localhost")
	case EnvStaging:
		return origin == "https://sample-staging.website.com"
	case EnvProduction:
		return origin == "https://sample.website.com"
	default:
		return true
	}
}

func InitEnv() {
	remoteBranch := os.Getenv("REMOTE_BRANCH")
	if remoteBranch == "" {
		// load env from .env file
		// load env from .env file
		path := GetBasePath() + "/.env"
		err := godotenv.Load(path)
		if err != nil {
			log.Panicln(err)
		}
	}
}
