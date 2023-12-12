package app

import (
	"context"
	"fmt"
	"github.com/basterrus/go_backend_framework/internal/config"
	"github.com/basterrus/go_backend_framework/internal/user"
	"github.com/julienschmidt/httprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Application struct {
}

func NewApplication(ctx context.Context, cfg *config.Config) (*Application, error) {

	//connString := generateConnectionString(cfg)
	//pgclient, err := client.NewPostgresClient(ctx, cfg, connString)
	//if err != nil {
	//	panic(err)
	//}

	router := httprouter.New()

	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	userHandler := user.NewUserHandler()
	userHandler.Register(router)

	return &Application{}, nil
}

func (app *Application) Run() error {

	return nil
}

func generateConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database)
}
