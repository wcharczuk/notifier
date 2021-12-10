package main

import (
	"context"
	"flag"
	"log"
	"sync"

	"github.com/wcharczuk/lametric/pkg/async"
	"github.com/wcharczuk/lametric/pkg/config"
	"github.com/wcharczuk/lametric/pkg/lametric"
)

var (
	flagConfig = flag.String("config", "_config/config.yml", "The configuration path")
)

func init() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.LUTC | log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {
	var cfg config.Config
	config.MustRead(&cfg,
		*flagConfig,
	)

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
	if len(errs) > 0 {
		maybeFatalExit(errs.All())
	}
	log.Printf("%d notifications sent", len(cfg.Devices))
}

func maybeFatalExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func send(ctx context.Context, device config.Device, notification lametric.Notification) error {
	client := lametric.New(device.Addr, device.Token)
	_, err := client.CreateNotification(context.Background(), notification)
	return err
}
