package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// MustRead reads a given path into a given config reference.
func MustRead(cfg interface{}, preferredPath string) {
	f, err := os.Open(preferredPath)
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		panic(err)
	}
}
