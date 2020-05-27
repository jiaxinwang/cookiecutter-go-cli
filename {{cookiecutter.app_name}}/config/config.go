package config

// Config ...
type Config struct {
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
