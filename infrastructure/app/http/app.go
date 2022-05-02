package http

import (
	routes "clean-architecture/api/http"
	controller "clean-architecture/controllers/http"
	"clean-architecture/domain"
	"clean-architecture/infrastructure/config"
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	http         *http.Server
	router       *gin.Engine
	logger       domain.ILogger
	db           *sql.DB
	configLoader domain.IConfig
	config       *config.Config
}

func New() *App {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	app := &App{
		http: &http.Server{
			Handler: router,
		},
		router: router,
	}

	router.Use(gin.Recovery())

	return app
}

func (s *App) Start() (err error) {
	if s.configLoader != nil {
		if s.config, err = s.configLoader.Load(); err != nil {
			return err
		}
	}
	s.http.Addr = fmt.Sprintf(":%d", s.config.Http.Port)
	return s.http.ListenAndServe()
}

func (s *App) Stop() (err error) {
	return s.http.Shutdown(context.Background())
}

func (s *App) WithController(controller ...controller.IController) domain.IApp {
	routes.RegisterRoutes(s, controller...)
	return s
}

func (s *App) WithLogger(logger domain.ILogger) domain.IApp {
	s.logger = logger
	return s
}

func (s *App) Logger() domain.ILogger {
	return s.logger
}

func (s *App) WithDb(db *sql.DB) domain.IApp {
	s.db = db
	return s
}

func (s *App) Db() *sql.DB {
	return s.db
}

func (s *App) WithHttp(http *http.Server) domain.IApp {
	s.http = http
	return s
}

func (s *App) Http() *http.Server {
	return s.http
}

func (s *App) WithRouter(router *gin.Engine) domain.IApp {
	s.router = router
	return s
}

func (s *App) Router() *gin.Engine {
	return s.router
}

func (s *App) WithConfig(config *config.Config) domain.IApp {
	s.config = config
	return s
}

func (s *App) WithConfigLoader(configLoader domain.IConfig) domain.IApp {
	s.configLoader = configLoader
	return s
}

func (s *App) Config() *config.Config {
	return s.config
}
