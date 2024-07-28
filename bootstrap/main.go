package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/kuroshibaz/bootstrap/command"
	"github.com/kuroshibaz/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	cfg := loadConfig()

	app := cli.NewApp()
	app.Name = cfg.Server.ApplicationName
	app.Usage = "command line interface"

	app.Commands = []*cli.Command{
		{
			Name:    "command",
			Aliases: []string{"start", "serve", "s"},
			Usage:   fmt.Sprintf("start %s service", cfg.Server.ApplicationName),
			Action: func(context *cli.Context) error {
				command.Server(cfg)
				fmt.Printf("[%s] Server Shutdown", color.YellowString("INFO"))
				return nil
			},
		},
		{
			Name:  "seed",
			Usage: fmt.Sprintf("run seed"),
			Action: func(context *cli.Context) error {
				fmt.Printf("[%s] Seed data proceeded", color.YellowString("INFO"))
				command.Seed(cfg)
				fmt.Printf("[%s] Seed data has successfully", color.GreenString("COMPLETED"))
				return nil
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   fmt.Sprintf("generate..."),
			Action: func(context *cli.Context) error {
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func runSeed() {

}

/* loadConfig read/map to environment config */
func loadConfig() *config.Env {
	cfg, err := config.ReadConfig("config")
	if err != nil {
		log.Fatalf("error config: %v", err)
	}

	return cfg
}
