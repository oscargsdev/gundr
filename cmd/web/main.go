package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/oscargsdev/gundr/internal/models"
)

type application struct {
	logger   *slog.Logger
	projects *models.ProjectModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// dsn := flag.String("dsn", "", "Data source name")
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger:   logger,
		projects: &models.ProjectModel{},
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
