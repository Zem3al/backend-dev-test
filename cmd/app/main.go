package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dev/internal/app"
	"dev/internal/config"
	"dev/internal/persistence"
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
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "dir",
					Aliases: []string{"d"},
					Value:   "./migration",
				},
			},
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
	log.Println("Migration Up")
	mirgrationDir := c.String("dir")

	migration := &migrate.FileMigrationSource{
		Dir: mirgrationDir,
	}

	migrate.SetIgnoreUnknown(true)

	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	conn, err := db.DB()

	if err != nil {
		return err
	}

	_, err = migrate.Exec(conn, "postgres", migration, migrate.Up)
	if err != nil {
		return err
	}

	return nil

}

func MigrationDown(c *cli.Context) error {
	log.Println("Migration Down")
	mirgrationDir := c.String("dir")

	migration := &migrate.FileMigrationSource{
		Dir: mirgrationDir,
	}

	migrate.SetIgnoreUnknown(true)

	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	conn, err := db.DB()

	if err != nil {
		return err
	}

	_, err = migrate.Exec(conn, "postgres", migration, migrate.Down)
	if err != nil {
		return err
	}

	return nil
}

func Serve(c *cli.Context) error {
	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	ctx := c.Context

	err := persistence.LoadHubGroupRepository(ctx)
	if err != nil {
		return err
	}

	err = persistence.LoadTeamGroupRepository(ctx)
	if err != nil {
		return err
	}

	err = persistence.LoadUserGroupRepository(ctx)
	if err != nil {
		return err
	}

	errChan := make(chan error)
	defer close(errChan)

	err = app.Serve(ctx, c.String("addr"))

	log.Println("Shutting down server")

	if err != nil {
		return err
	}

	return nil
}
