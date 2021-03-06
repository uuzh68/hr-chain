package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"github.com/uuzh68/hr-chain/config"
	"github.com/uuzh68/hr-chain/dbschema"
	"github.com/uuzh68/hr-chain/info"
	"github.com/uuzh68/hr-chain/workers"
)

var serveCommand = cli.Command{
	Name:   "serve",
	Usage:  "starts " + config.ServiceName + " workers",
	Action: serveAction,
}

func serveAction(c *cli.Context) error {
	config.Init(c.GlobalString(FlagConfig))
	cfg := config.Config()

	if cfg.AutoMigrate {
		dbschema.SetAssets()
		count, err := db.Migrate(config.Config().DB, "up")
		if err != nil {
			log.Default.WithError(err).Error("Migrations failed")
		}
		log.Default.Info(fmt.Sprintf("Applied %d %s migration", count, "up"))
	}

	workers.GetChief().RunAll(cfg.Log.AppName, cfg.Workers...)
	return nil
}
