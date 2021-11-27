package main

import (
	"context"
	"flag"
	"os"

	"github.com/blend/go-sdk/ansi/slant"
	"github.com/blend/go-sdk/configutil"
	"github.com/blend/go-sdk/logger"

	"github.com/wcharczuk/notifier/pkg/config"
	"github.com/wcharczuk/notifier/pkg/lametric"
)

var (
	flagConfig = flag.String("config", "_config/config.yml", "The configuration path")
)

func init() {
	flag.Parse()
}

func main() {
	var cfg config.Config
	configutil.MustRead(&cfg,
		configutil.OptAddPreferredPaths(*flagConfig),
	)

	log := logger.All()
	client := lametric.New(cfg.Devices[0].Addr, cfg.Devices[0].Token)
	client.Client.Log = log

	// START.
	slant.Print(os.Stdout, "LaMETRIC")

	_, err := client.CreateNotification(context.Background(), lametric.CreateNotificationInput{
		Model: lametric.DeviceModel{
			Frames: []lametric.Frame{
				{
					Icon: 2867,
					Text: "Hello!",
				},
			},
		},
	})
	logger.MaybeError(log, err)
	// FINISH.
}
