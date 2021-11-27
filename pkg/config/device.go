package config

// Device is a notification broadcast target.
type Device struct {
	Addr  string `yaml:"addr"`
	Token string `yaml:"token"`
}
