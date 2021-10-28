package config

import (
	"fmt"

	"github.com/iph0/conf"
	"github.com/iph0/conf/envconf"
	"github.com/iph0/conf/fileconf"
	"github.com/rs/zerolog"
)

type Config struct {
	Log        string `conf:"Log"`
	ListenAddr string `conf:"ListenAddr"`
	PGConnInfo string `conf:"PGConnInfo"`
}

type ConfigParsed struct {
	LogLevel zerolog.Level
	//ListenAddr  string
	//PGConnInfo  string
}

func (conf Config) Parse() (*ConfigParsed, error) {

	level, err := zerolog.ParseLevel(conf.Log)
	if err != nil {
		return &ConfigParsed{}, nil
	}

	return &ConfigParsed{
		LogLevel: level,
	}, nil
}

func LoadConfig() (Config, error) {
	fileLdr := fileconf.NewLoader("./etc/", "/etc/")
	envLdr := envconf.NewLoader()

	configProc := conf.NewProcessor(
		conf.ProcessorConfig{
			Loaders: map[string]conf.Loader{
				"file": fileLdr,
				"env":  envLdr,
			},
		},
	)
	configRaw, err := configProc.Load("file:config.yaml", "env:^HW_")

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = conf.Decode(configRaw["config"], &config)
	if err != nil {
		fmt.Println(err)
		return Config{}, err
	}
	return config, nil

}
