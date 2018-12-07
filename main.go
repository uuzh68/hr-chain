package main

import (
	"os"

	"github.com/urfave/cli"
	"github.com/uuzh68/hr-chain/cmd"
	"github.com/uuzh68/hr-chain/config"
	"github.com/uuzh68/hr-chain/info"
)

func main() {
	app := cli.NewApp()
	app.Usage = "A " + config.ServiceName + " service"
	app.Version = info.App.Version
	app.Flags = cmd.GetFlags()
	app.Commands = cmd.GetCommands()
	app.Run(os.Args)
}
