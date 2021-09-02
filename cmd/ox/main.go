package main

import (
	"context"
	"fmt"
	"os"

	"cotizador_sounio_health"
	_ "cotizador_sounio_health/app/models"
	_ "cotizador_sounio_health/app/tasks"

	"github.com/wawandco/oxpecker/cli"
	"github.com/wawandco/oxpecker/tools/soda"
)

// main function for the tooling cli, will be invoked by Oxpecker
// when found in the source code. In here you can add/remove plugins that
// your app will use as part of its lifecycle.
func main() {
	cli.Use(soda.Plugins(cotizador_sounio_health.Migrations)...)
	err := cli.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Printf("[error] %v \n", err.Error())

		os.Exit(1)
	}
}
