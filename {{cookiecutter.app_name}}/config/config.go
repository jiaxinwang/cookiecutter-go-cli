package config

import (
	"github.com/BurntSushi/toml"
	"{% if cookiecutter.use_github == "y" -%}github.com/{{cookiecutter.github_username}}/{%- endif %}{{cookiecutter.app_name}}/util/logger"

)

var log *zap.SugaredLogger

func init() {
	log = logger.S.Named("config")
}


// Config ...
var Config struct {
	Server struct {
		Listen             string `toml:"Listen"`
		Env                string `toml:"Env"`
		MaxHTTPRequestBody int    `toml:"MaxHttpRequestBody"`
		UseTLS             bool   `toml:"UseTLS"`
		TLSPem             string `toml:"TLSPem"`
		TLSKey             string `toml:"TLSKey"`
	} `toml:"Server"`
	Database struct {
		DSN string `toml:"DSN"`
	} `toml:"Database"`
	Permission struct {
		DSN string `toml:"DSN"`
		DB  string `toml:"DB"`
	} `toml:"Permission"`
}

// Load ...
func Load(filename string) {
	log.Info(filename)
	if _, err := toml.DecodeFile(filename, &Config); err != nil {
		log.Panic(err)
		return
	}
}
