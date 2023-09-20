package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/adamelfsborg-code/orders-api-go/order"
	"github.com/go-pg/pg/v10"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router   http.Handler
	database order.Repo
	config   Config
}

func New(config Config) *App {
	app := &App{
		config: config,
	}

	var repo order.Repo

	switch config.RepoAdapter {
	case "PSQL":
		// Initialize the PostgreSQL database connection
		pgConn := pg.Connect(&pg.Options{
			Addr:     config.PostgresAddr,
			User:     config.PostgresUser,
			Password: config.PostgresPassword,
			Database: config.PostgresDatabase,
		})

		repo = &order.PostgresRepo{
			Client: pgConn,
		}
	case "REDIS":
		redisConn := redis.NewClient(&redis.Options{
			Addr: config.RedisAddress,
		})

		repo = &order.RedisRepo{
			Client: redisConn,
		}
	}

	app.database = repo

	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.ServerPort),
		Handler: a.router,
	}

	err := a.database.Ping(ctx)
	if err != nil {
		return fmt.Errorf("Failed to connect to repo: %w", err)
	}
	defer func() {
		err := a.database.Close(ctx)
		if err != nil {
			fmt.Println("Failed to close Repo", err)
		}
	}()

	fmt.Printf("Starting Server! Connected to Repo: %q", a.config.RepoAdapter)

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("Failed to start server: %w", err)
		}

		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
