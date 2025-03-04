package main

import (
	"flag"
	"log/slog"
	"os"
)

type application struct {
	logger *slog.Logger
	addr   *string
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// create an instance of the application struct
	app := &application{
		logger: logger,
		addr:   addr,
	}

	err := app.serve()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
