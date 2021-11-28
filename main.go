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

	slant.Print(os.Stdout, "LaMETRIC")

	log := logger.All()
	client := lametric.New(cfg.Devices[0].Addr, cfg.Devices[0].Token)
	client.Client.Log = log
	_, err := client.CreateNotification(context.Background(), lametric.Notification{
		Model: lametric.NotificationModel{
			Frames: []lametric.Frame{
				{
					Icon: lametric.IconAttention,
					Text: "ALERT",
				},
				{
					Icon: lametric.IconAppleLogo,
					Text: "Adam has been fired.",
				},
			},
			Sound: &lametric.Sound{
				Category: lametric.SoundCategoryAlarms,
				ID:       lametric.SoundAlarm10,
			},
		},
	})
	logger.MaybeError(log, err)
	// FINISH.
}
