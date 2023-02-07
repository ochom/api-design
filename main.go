package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/ochom/api-design/pkg/database"
	"github.com/ochom/api-design/pkg/handlers"
	"github.com/ochom/api-design/pkg/utils"
	"github.com/ochom/grm"
)

func init() {
	databaseURL := utils.MustGetEnv("DATABASE_URL")

	// initialize database connection
	grm.Init(grm.Postgres, databaseURL, grm.Silent)

	// migrate database
	if err := grm.Migrate(database.GetSchema()...); err != nil {
		utils.NewLogger().Fatal(err)
	}
}

func main() {
	rht := flag.Int("rht", 5, "ReadHeaderTimeout in seconds")
	rt := flag.Int("rt", 5, "ReadTimeout in seconds")
	flag.Parse()

	handler := handlers.NewHandler()

	svr := http.Server{
		Addr:              ":8080",
		Handler:           handler.Routes(),
		ReadHeaderTimeout: time.Duration(*rht) * time.Second,
		ReadTimeout:       time.Duration(*rt) * time.Second,
	}

	utils.NewLogger().Info("Starting server on port 8080")
	go func() {
		if err := svr.ListenAndServe(); err != nil {
			utils.NewLogger().Fatal(err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		utils.NewLogger().Fatal(err)
	}

	utils.NewLogger().Info("Server gracefully stopped")

}
