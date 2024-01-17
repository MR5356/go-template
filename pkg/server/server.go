package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/MR5356/go-template/config"
	"github.com/MR5356/go-template/docs"
	"github.com/MR5356/go-template/pkg/controller"
	"github.com/MR5356/go-template/pkg/domain/demo"
	_ "github.com/MR5356/go-template/pkg/log"
	"github.com/MR5356/go-template/pkg/middleware/database"
	"github.com/MR5356/go-template/pkg/response"
	"github.com/MR5356/go-template/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	engine *gin.Engine
	config *config.Config
}

func New(config *config.Config) (server *Server, err error) {
	if config.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	database.GetDB()

	engine := gin.Default()
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(
		Record(),
	)

	engine.NoRoute(func(c *gin.Context) {
		response.New(c, http.StatusNotFound, response.CodeNotFound, response.MessageNotFound, nil)
	})

	api := engine.Group(config.Server.Prefix)

	// Prometheus
	api.GET("/metrics", func(handler http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		}
	}(promhttp.Handler()))

	// swagger
	docs.SwaggerInfo.Title = "API Docs"
	docs.SwaggerInfo.Description = "This is an auto-generated API Docs."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = config.Server.Prefix
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	demoSvc := demo.NewService()

	services := []service.Service{
		demoSvc,
	}

	for _, srv := range services {
		err := srv.Initialize()
		if err != nil {
			return nil, err
		}
	}

	controllers := []controller.Controller{
		demo.NewController(demoSvc),
	}
	for _, ctrl := range controllers {
		ctrl.RegisterRoute(api)
	}

	server = &Server{
		engine: engine,
		config: config,
	}
	return
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.Server.Port),
		Handler: s.engine,
	}

	go func() {
		logrus.Infof("Listening on port %d", s.config.Server.Port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.Server.GracePeriod)*time.Second)
	defer cancel()

	ch := <-sig
	logrus.Infof("Received signal %s", ch.String())
	return server.Shutdown(ctx)
}
