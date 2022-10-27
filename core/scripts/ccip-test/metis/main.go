package main

import (
	"log"
	"os"

	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/recovery"
	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/metis/cmd"
)

// NAME:
//
//	Metis - CCIP sanity checker
//
// USAGE:
//
//	[global options] command [command options] [arguments...]
//
// COMMANDS:
//
//	state, s  prints CCIP config state
//	txs, t    prints recent txs
//	help, h   Shows a list of commands or help for one command
//
// GLOBAL OPTIONS:
//
//	--help, -h  show help
func main() {
	recovery.ReportPanics(func() {
		client := NewClient()
		app := cmd.NewMetisApp(client)

		client.Logger.ErrorIf(app.Run(os.Args), "Error running app")
		if err := client.CloseLogger(); err != nil {
			log.Fatal(err)
		}
	})
}

func NewClient() cmd.MetisClient {
	lggr, closeLggr := logger.NewLogger()
	return cmd.MetisClient{
		Logger:      lggr,
		CloseLogger: closeLggr,
	}
}
