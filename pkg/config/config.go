package config

import (
	"context"

	"github.com/blend/go-sdk/configmeta"
	"github.com/blend/go-sdk/configutil"
)

// Config is a root config struct.
type Config struct {
	configmeta.Meta `yaml:",inline"`
	Devices         []Device `yaml:"devices"`
}

// Resolve resolves the config.
func (c *Config) Resolve(ctx context.Context) error {
	return configutil.Resolve(ctx,
		(&c.Meta).Resolve,
	)
}
