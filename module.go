package storagessl

import (
	"os"
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/certmagic"
	"go.uber.org/zap"
)

// StorageParam Informacion para guardar los datos en el servidor
type StorageParam struct {
	Logger *zap.Logger

	Host         string `json:"host"`
	Port         string `json:"port"`
	DataBaseName string `json:"database"`
	Timeout      int    `json:"timeout"`
	AccessKey    string `json:"access_key"`
}

const (
	EnvNameHost      = "host"
	EnvNamePort      = "port"
	EnvNameDatabase  = "database"
	EnvNameTimeout   = "timeout"
	EnvNameAccessKey = "accesskey"

	DefaultHost      = "localhost"
	DefaultPort      = "80"
	DefaultDatabase  = "MultiComDatabase"
	DefaultTimeout   = 5
	DefaultAccessKey = "w9AFS267JWPTQkjfYXQMUTEg22BmmbfQ"
)

func init() {
	caddy.RegisterModule(StorageParam{})
}

func (rd *StorageParam) Provision(ctx caddy.Context) error {
	rd.Logger = ctx.Logger(rd)
	rd.GetConfigValue()
	return nil
}

// GetConfigValue get Config value from env, if already been set by Caddyfile, don't overwrite
func (rd *StorageParam) GetConfigValue() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	rd.Logger = logger
	rd.Logger.Info("Estoy en la funcion GetConfigValue. Esta funcion se tiene que recodificar")

	rd.Host = configureString(rd.Host, EnvNameHost, DefaultHost)
	rd.Port = configureString(rd.Port, EnvNamePort, DefaultPort)
	rd.DataBaseName = configureString(rd.DataBaseName, EnvNameDatabase, DefaultDatabase)

	rd.Timeout = configureInt(rd.Timeout, EnvNameTimeout, DefaultTimeout)
	rd.AccessKey = configureString(rd.AccessKey, EnvNameAccessKey, DefaultAccessKey)
}

func (rd *StorageParam) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		var value string
		key := d.Val()

		if !d.Args(&value) {
			continue
		}

		switch key {
		case "host":
			rd.Host = value
		case "port":
			rd.Port = value
		case "database":
			rd.DataBaseName = value
		case "timeout":
			if value != "" {
				timeParse, err := strconv.Atoi(value)
				if err == nil {
					rd.Timeout = timeParse
				} else {
					rd.Timeout = DefaultTimeout
				}
			} else {
				rd.Timeout = DefaultTimeout
			}
		case "access_key":
			rd.AccessKey = value
		}
	}

	return nil
}

// register caddy module with ID caddy.storage.MongoDb
func (StorageParam) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "caddy.storage.storagessl",
		New: func() caddy.Module {
			return new(StorageParam)
		},
	}
}

// CertMagicStorage converts s to a certmagic.Storage instance.
func (rd *StorageParam) CertMagicStorage() (certmagic.Storage, error) {
	return rd, nil
}

func configureInt(value int, envVariableName string, valueDefault int) int {
	if value != 0 {
		return value
	}
	if envVariableName != "" {
		valueEnvStr := os.Getenv(envVariableName)
		if valueEnvStr != "" {
			valueEnv, err := strconv.Atoi(os.Getenv(envVariableName))
			if err == nil {
				return valueEnv
			}
		}
	}
	return valueDefault
}

func configureString(value string, envVariableName string, valueDefault string) string {
	if value != "" {
		return value
	}
	if envVariableName != "" {
		valueEnvStr := os.Getenv(envVariableName)
		if valueEnvStr != "" {
			return valueEnvStr
		}
	}
	return valueDefault
}
