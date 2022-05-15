package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/swaggo/swag"

	"github.com/ew0s/trade-bot/cmd/api/handler"
	_ "github.com/ew0s/trade-bot/cmd/api/swagger" // read swagger doc
	apiservice "github.com/ew0s/trade-bot/internal/api/service"
	"github.com/ew0s/trade-bot/internal/repos/postgres"
	"github.com/ew0s/trade-bot/internal/repos/redis"
	"github.com/ew0s/trade-bot/internal/service"
	"github.com/ew0s/trade-bot/pkg/api"
	"github.com/ew0s/trade-bot/pkg/httputils"
	logsetup "github.com/ew0s/trade-bot/pkg/log"
	"github.com/ew0s/trade-bot/pkg/openapi"
	"github.com/ew0s/trade-bot/pkg/resource"
)

//go:generate configer -app-name api -env local

// @title        Trade-bot API
// @version      1.0
// @description  API Server for Trade-bot Application

// @BasePath  /trade-bot/api/v1

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	config, err := mustParseAppConfig()
	if err != nil {
		log.WithError(err).Fatalf("can't parse config")
	}

	logger := logsetup.Setup(config.Log)

	ctx, cancel := context.WithCancel(context.Background())

	db, err := postgres.NewPostgresDB(config.Postgres)
	if err != nil {
		logger.WithError(err).Fatalf("can't create postgres db")
	}
	defer resource.Close(logger, db)

	redisClient, err := redis.NewRedisClient(ctx, config.Redis)
	if err != nil {
		logger.WithError(err).Fatalf("can't create redis client")
	}
	defer resource.Close(logger, redisClient)

	jwtService := service.NewJWTService(config.JWT.SigningKey, config.JWT.ExpirationDuration)

	userIdentityService := apiservice.NewUserIdentity(jwtService)
	userIdentity := handler.NewUserIdentity(userIdentityService)

	authRepo := postgres.NewAuth(db)
	identityRepo := redis.NewJWTRedis(redisClient)
	authService := apiservice.NewAuth(authRepo, identityRepo, jwtService)
	authHandler := handler.NewAuth(authService, userIdentity)

	r := api.MakeRoutes(config.BasePath, []chi.Router{
		authHandler.Routes(),
	})

	openapiHandler, err := setupOpenapiHandler(config.DocsPath)
	if err != nil {
		log.WithError(err).Fatalf("can't setup openapi handler")
	}

	setupDocsRoutes(r, openapiHandler, config.DocsPath)

	servers := []*httputils.Server{
		httputils.NewServer(config.ListenAddr, r),
	}

	for i := range servers {
		go func(srv *httputils.Server) {
			if err = srv.Run(); err != http.ErrServerClosed {
				logger.WithError(err).Fatalf("server cant't listen requests")
			}
		}(servers[i])

		logger.Info(servers[i].Info())
	}

	interrupt := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-interrupt

		logger.Info("interrupt signal caught")
		logger.Info("Trade bot server shutting down")

		cancel()

		for _, server := range servers {
			server := server

			go func() {
				if err = server.Shutdown(ctx); err != nil {
					logger.WithError(err).Fatalf("can't gracefully shotdown server")
				}
			}()
		}
	}()

	logger.Info("Trade bot server started")

	<-ctx.Done()

	logger.Info("Trade bot has been terminated")
}

func setupOpenapiHandler(docsPath string) (*openapi.Handler, error) {
	doc, err := swag.ReadDoc()
	if err != nil {
		return nil, fmt.Errorf("reading swagger (make sure doc import is presented): %w", err)
	}

	openapiHandler, err := openapi.NewHandler(docsPath, doc)
	if err != nil {
		return nil, fmt.Errorf("initializing openapi handler: %w", err)
	}

	return openapiHandler, nil
}

func setupDocsRoutes(r chi.Router, openapiHandler *openapi.Handler, docsPath string) {
	r.Route(docsPath, func(r chi.Router) {
		r.Get(openapi.DocsJSONPath, openapiHandler.DocJSON)
		r.Get(openapi.DocsIndexPath, openapiHandler.Index)
		r.Get("/*", openapiHandler.RedirectToIndex)
	})
}
