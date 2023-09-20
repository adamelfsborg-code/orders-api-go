package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/adamelfsborg-code/orders-api-go/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}

}
