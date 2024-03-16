package config

import (
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
)

type Cfg struct {
	Server Server `json:"server" yaml:"server"`
	Db     DB     `json:"db" yaml:"db"`
}

var cfgInstance *Cfg = nil

func (conf *Cfg) ParseConfig(cfgFileName string) {
	ext := path.Ext(cfgFileName)
	ext = strings.ToLower(ext)
	file, err := os.Open(cfgFileName)
	if err != nil {
		panic(err)
	}

	if ext == "" || ext == ".json" {
		err = json.NewDecoder(file).Decode(&conf)
	} else if ext == ".yaml" || ext == ".yml" {
		err = yaml.NewDecoder(file).Decode(&conf)
	} else {
		err = errors.New("unknown config provider")
	}
	if err != nil {
		panic(err)
	}
	log.Debug().Msgf("Readed config %+v", conf)
}

func Get() *Cfg {
	return cfgInstance
}

func set(cfg *Cfg) {
	cfgInstance = cfg
}

func ReadConfig(cfgFileName string) *Cfg {
	cfg := &Cfg{}
	cfg.ParseConfig(cfgFileName)
	set(cfg)
	return Get()
}
