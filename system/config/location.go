package config

// DirLocations represent DirLocations
type DirLocations struct {
	Conf  string `json:"conf" yaml:"conf"`
	WWW   string `json:"www" yaml:"www"`
	Temp  string `json:"temp" yaml:"temp"`
	Log   string `json:"log" yaml:"log"`
	Cache string `json:"cache" yaml:"cache"`
}
