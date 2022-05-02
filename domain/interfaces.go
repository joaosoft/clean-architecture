package domain

import (
	controller "clean-architecture/controllers/http"
	"clean-architecture/infrastructure/config"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IConfig interface {
	Load() (_ *config.Config, err error)
}

type IApp interface {
	WithConfigLoader(configLoader IConfig) IApp
	WithConfig(*config.Config) IApp
	Config() *config.Config
	WithLogger(logger ILogger) IApp
	Logger() ILogger
	WithDb(db *sql.DB) IApp
	Db() *sql.DB
	WithHttp(http *http.Server) IApp
	Http() *http.Server
	WithRouter(router *gin.Engine) IApp
	Router() *gin.Engine
	WithController(controller ...controller.IController) IApp
	Start() error
	Stop() error
}

type ILogger interface {
	Printf(format string, v ...any)
	Print(v ...any)
	Println(v ...any)
	Fatal(v ...any)
	Fatalf(format string, v ...any)
	Fatalln(v ...any)
	Panic(v ...any)
	Panicf(format string, v ...any)
	Panicln(v ...any)
}
