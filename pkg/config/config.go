package config

// Config is a root config struct.
type Config struct {
	Devices []Device `yaml:"devices"`
}
