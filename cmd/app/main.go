package main

import (
	"context"
	"dev/internal/app"
	"dev/internal/config"
	"dev/internal/persistence"
	"dev/internal/service"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"
)

func main() {

	app := cli.NewApp()
	app.Name = "Test app"
	app.Usage = "Serve an api"
	app.Version = "Pre_Beta"

	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   "./configs/.env",
			Usage:   "set path to enviroment file",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "serve",
			Usage:  "Serve server",
			Action: Serve,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "addr",
					Aliases: []string{"address"},
					Value:   "0.0.0.0:8080",
				},
			},
		},
		{
			Name:  "migration",
			Usage: "Migration",
			Subcommands: []*cli.Command{
				{
					Name:   "up",
					Usage:  "migration up",
					Action: MigrationUp,
				},
				{
					Name:   "down",
					Usage:  "migration down",
					Action: MigrationDown,
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		return godotenv.Load(c.String("env"))
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	ctx, cancel := context.WithCancel(context.Background())

	endSignal := make(chan os.Signal, 1)
	signal.Notify(endSignal, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}

	errChan := make(chan error, 1)

	wg.Add(1)

	go func(ctx context.Context, errChan chan error) {
		defer wg.Done()
		err := app.RunContext(ctx, os.Args)
		errChan <- err
	}(ctx, errChan)

	select {
	case sign := <-endSignal:
		log.Printf("Shutting down. reason:", sign)
	case err := <-errChan:
		if err == nil {
			break
		}
		log.Println("encountered error", err)
		return
	}

	cancel()
	wg.Wait()
}

func MigrationUp(c *cli.Context) error {
	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	ctx := c.Context

	err := persistence.LoadPayloadRespositoryMock(ctx)
	if err != nil {
		return err
	}

	workerService := service.InitWorkerService()

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		errChan <- workerService.Start(ctx)
	}()

	go func() {
		errChan <- app.Serve(ctx, c.String("addr"))
	}()

	for {
		select {
		case err := <-errChan:
			return err
		}
	}
}

func MigrationDown(c *cli.Context) error {
	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	ctx := c.Context

	err := persistence.LoadPayloadRespositoryMock(ctx)
	if err != nil {
		return err
	}

	workerService := service.InitWorkerService()

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		errChan <- workerService.Start(ctx)
	}()

	go func() {
		errChan <- app.Serve(ctx, c.String("addr"))
	}()

	for {
		select {
		case err := <-errChan:
			return err
		}
	}
}

func Serve(c *cli.Context) error {
	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	ctx := c.Context

	err := persistence.LoadPayloadRespositoryS3(ctx)
	if err != nil {
		return err
	}

	workerService := service.InitWorkerService()

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		errChan <- workerService.Start(ctx)
	}()

	go func() {
		errChan <- app.Serve(ctx, c.String("addr"))
	}()

	for {
		select {
		case err := <-errChan:
			return err
		}
	}
}
