package main

import (
	"context"
	"flag"
	"os"
	"sync"

	"github.com/blend/go-sdk/ansi/slant"
	"github.com/blend/go-sdk/async"
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

	notification := lametric.Notification{
		Model: lametric.NotificationModel{
			Frames: []lametric.Frame{
				{
					Icon: lametric.IconAttention,
					Text: "ALERT",
				},
			},
			/*
				Sound: &lametric.Sound{
					Category: lametric.SoundCategoryAlarms,
					ID:       lametric.SoundAlarm10,
				},
			*/
		},
	}

	errs := make(async.Errors, len(cfg.Devices))
	wg := sync.WaitGroup{}
	wg.Add(len(cfg.Devices))
	for x := 0; x < len(cfg.Devices); x++ {
		go func(index int) {
			defer wg.Done()
			if err := send(context.Background(), cfg.Devices[index], notification); err != nil {
				errs <- err
			}
		}(x)
	}
	wg.Wait()
	if err := errs.All(); err != nil {
		logger.MaybeFatalExit(log, errs.All())
	}
	logger.MaybeInfof(log, "%d notifications sent", len(cfg.Devices))
}

func send(ctx context.Context, device config.Device, notification lametric.Notification) error {
	client := lametric.New(device.Addr, device.Token)
	_, err := client.CreateNotification(context.Background(), notification)
	return err
}
