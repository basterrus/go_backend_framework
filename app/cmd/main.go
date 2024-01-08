package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/basterrus/go_backend_framework/internal/config"
	"github.com/basterrus/go_backend_framework/internal/user"
	"github.com/basterrus/go_backend_framework/pkg/client"
	"github.com/basterrus/go_backend_framework/pkg/logging"
	"github.com/basterrus/go_backend_framework/pkg/metric"
	"github.com/basterrus/go_backend_framework/pkg/shutdown"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")

	logger.Println("config initializing")
	cfg := config.GetConfig(logger)

	logger.Println("router initializing")
	router := httprouter.New()

	logger.Println("metrics initializing")
	metricHandler := metric.Handler{Logger: logger}
	metricHandler.Register(router)

	logger.Println("postgresql initializing")
	connectionString := generateConnectionString(cfg)
	pgClient, err := client.NewPostgresClient(context.Background(), cfg, connectionString)
	if err != nil {
		logger.Fatal(err)
	}

	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	userStorage := user.NewStorage(pgClient, logger)
	userService, err := user.NewService(userStorage, logger)
	if err != nil {
		logger.Fatal(err)
	}

	usersHandler := user.Handler{
		Logger:      logger,
		UserService: userService,
	}

	usersHandler.Register(router)

	logger.Println("start application")
	start(router, logger, cfg)
}

func start(router http.Handler, logger logging.Logger, cfg *config.Config) {
	var server *http.Server
	var listener net.Listener

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		socketPath := path.Join(appDir, "app.sock")
		logger.Infof("socket path: %s", socketPath)

		logger.Info("create and listen unix socket")
		listener, err = net.Listen("unix", socketPath)
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		logger.Infof("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)

		var err error

		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		if err != nil {
			logger.Fatal(err)
		}
	}

	server = &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go shutdown.Graceful([]os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM},
		server)

	logger.Println("application initialized and started")

	if err := server.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
}

func generateConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database)
}
