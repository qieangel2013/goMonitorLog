package handle

import (
	"github.com/BurntSushi/toml"
	"github.com/juju/errors"
	"io/ioutil"
	"time"
)

type Config struct {
	LogList        []string `toml:"log_list"`
	FillterList    []string `toml:"fillter_list"`
	FindList       []string `toml:"find_list"`
	DingWebhookUrl string   `toml:"ding_webhook_url"`
	Env            string   `toml:"env"`
	ServerLog      string   `toml:"server_log"`
	TailLine       string   `toml:"tail_line"`
}

// NewConfigWithFile creates a Config from file.
func NewConfigWithFile(name string) (*Config, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return NewConfig(string(data))
}

// NewConfig creates a Config from data.
func NewConfig(data string) (*Config, error) {
	var c Config

	_, err := toml.Decode(data, &c)
	if err != nil {
		return nil, errors.Trace(err)
	}

	return &c, nil
}

// TomlDuration supports time codec for TOML format.
type TomlDuration struct {
	time.Duration
}

// UnmarshalText implementes TOML UnmarshalText
func (d *TomlDuration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}
