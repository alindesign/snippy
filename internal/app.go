package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alindesign/snippy/web"
)

func NewApp(
	config Config,
	database DatabaseConnection,
	homeController HomeController,
	snippetController SnippetController,
) *App {
	return &App{
		database,
		config,
		homeController,
		snippetController,
	}
}

type App struct {
	database          DatabaseConnection
	config            Config
	homeController    HomeController
	snippetController SnippetController
}

func (a *App) Close() error {
	return nil
}

func (a *App) Run() error {
	switch a.config.ApplicationCommand {
	case "serve":
		return a.serve()
	case "migrate":
		return a.migrate()
	default:
		return fmt.Errorf("unknown command: '%s'. Available: (serve, migrate)", a.config.ApplicationCommand)
	}
}

func (a *App) migrate() error {
	migrationStatements := CreateMigrationString()
	if _, err := a.database.Exec(migrationStatements); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func (a *App) serve() error {
	mux := http.NewServeMux()
	a.homeController.Handler(mux)
	a.snippetController.Handler(mux)
	web.Handler(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.config.ServerHost, a.config.ServerPort),
		Handler: mux,
	}

	log.Printf("migrating database")
	if err := a.migrate(); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Printf("server listening on http://%s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
